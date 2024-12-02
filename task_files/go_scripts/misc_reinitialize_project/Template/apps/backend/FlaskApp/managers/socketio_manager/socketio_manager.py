import http
from firebase_admin.auth import EmailIdentifier, UidIdentifier
import json
import ssl
import time
from managers.firebase_manager.check_for_authentication import  verify_access_token
from socketio.server import Server
from utils.exceptions.api_exceptions import APINotFoundError
from utils.api_msg_format import APIMsgFormat
from utils.env_vars import ENV_VARS, is_deployed_environment, is_dev_environment, is_test_environment
from utils.exceptions.general_exceptions import get_traceback
from utils.iterable_utils import is_iterable
from utils.print_if_dev import print_if_dev
from utils.exceptions.singleton_exception import SingletonException
from configs import CONFIGS

class SocketIOManager():
  init = False
  ssl_cert = None
  ssl_key = None
  server = None
  app = None
  user_session_id_map = {}
  listeners_were_added = False

  def __init__(self, ssl_cert=None, ssl_key=None):
    if (SocketIOManager.init):
      raise SingletonException
    else:
      SocketIOManager.init = True
      self.ssl_cert = ssl_cert
      self.ssl_key = ssl_key

  def init_server(self,server=Server(logger=False)):
    self.server = server

  def connect(self, sid, environ, auth):
    print_if_dev("Connnected to client with session Id {}".format(sid))

  def disconnect(self, sid):
    print_if_dev("Client disconnected {}".format(sid))
    self.update_user_session_id_map(sid)

  # TODO replace with socketio rooms
  def update_user_session_id_map(self, session_id, user_id=None):
    if user_id:
      if user_id in self.user_session_id_map:
        self.user_session_id_map[user_id].add(session_id)
      else:
        self.user_session_id_map[user_id] = set([session_id])
    else:
      for uid, sessions in self.user_session_id_map.items():
        if session_id in sessions:
          sessions.remove(session_id)
          if not sessions:
            del self.user_session_id_map[uid]
          break

  def get_session_ids_via_user_id(self, user_id):
    return self.user_session_id_map.get(user_id)

  def run_via_eventlet(self):
    self._add_listeners()
    self.server.run(
      app=self.app,
      host=CONFIGS.app['server_name'],
      certfile=self.ssl_cert, keyfile=self.ssl_key,
      use_reloader=True,
      debug=True,
      port=CONFIGS.app["backend_port"]
    )

  def run_via_gevent(self):
    context = ssl.create_default_context(ssl.Purpose.CLIENT_AUTH)
    context.load_cert_chain(certfile=self.ssl_cert, keyfile=self.ssl_key)
    self._add_listeners()
    self.server.run(
      app=self.app,
      async_mode="gevent",
      cors_allowed_origins=CONFIGS.app['access_control_allow_origin'],
      host=CONFIGS.app['server_name'],
      ssl_context=context,
      use_reloader=True,
      debug=True,
      port=CONFIGS.app["backend_port"]
    )

  def _add_listeners(self):
    if not self.listeners_were_added:
      self.listeners_were_added = True
      self.server.on_event("connect", self.connect)
      self.server.on_event("disconnect", self.disconnect)
      self.server.on_event("file_transfer_request",self.request_file_transfer)
      self.server.on_event("respond_to_file_transfer_request",self.respond_to_file_transfer_request)
      self.server.on_event("connect",self.peerjs_connect,namespace="/peerjs")
      self.server.on_event("message",self.peerjs_message,namespace="/peerjs")
      self.server.on_event("disconnect",self.peerjs_disconnect,namespace="/peerjs")

  def add_listeners_via_python_socketio(self):
    if not self.listeners_were_added:
      self.listeners_were_added = True
      # see the difference?
      self.server.on("connect", self.connect)
      self.server.on("disconnect", self.disconnect)
      self.server.on("file_transfer_request",self.request_file_transfer)
      self.server.on("respond_to_file_transfer_request",self.respond_to_file_transfer_request)
      self.server.on("connect",self.peerjs_connect,namespace="/peerjs")
      self.server.on("message",self.peerjs_message,namespace="/peerjs")
      self.server.on("disconnect",self.peerjs_disconnect,namespace="/peerjs")

  def emit_result(self,event, task_id = None, result=None,session_ids=[],code=http.client.OK.value):
    if code ==None:
      code = http.client.OK.value

    resp_body = APIMsgFormat(
      code=code,
      data={
        "task_id": task_id,
        "result": result
      },
    )

    print_if_dev("Clients of given user {}".format(session_ids))
    if is_iterable(session_ids):
      for room in session_ids:
        self.server.emit(event, resp_body.get_dict(),room=room )

  def run_task(self, task_id, task_name,room, fn_to_call):
    result = None
    code = http.client.OK.value
    if not is_test_environment():
        time.sleep(1)

    try:
        result,code = fn_to_call()
    except BaseException as e:
        code = http.client.INTERNAL_SERVER_ERROR.value
        print_if_dev("Error calling function: {}".format(e))

    self.emit_result(task_id, task_name, result,room,code)

  def request_file_transfer(self,sid,data):
    try:
      access_token = data["data"]["access_token"]
      recipients = data["data"]["recipients"]
      job_id = data["data"]["job_id"]
      firebase_uid, error = verify_access_token(access_token)
      if error != None:
        self.emit_result(
          "file_transfer_error",
          code=http.client.UNAUTHORIZED.value,
          session_ids=set([sid]),
          result=error if is_dev_environment() else None
        )
        return


      identifiers = [EmailIdentifier(email) for email in recipients]
      identifiers.append(UidIdentifier(firebase_uid))
      firebase_users = CONFIGS.firebase_manager.auth.get_users(identifiers).users
      sender_user = None
      sending_to_self =  len(identifiers) == len(firebase_users) +1
      for user in firebase_users:
        if user.uid == firebase_uid:
          if sending_to_self ==False:
            sender_user = firebase_users.pop(firebase_users.index(user))
          else:
            sender_user = user
          break

      for firebase_user in firebase_users:
        recipient_session_ids =self.get_session_ids_via_user_id(firebase_user.uid)
        recipient_session_ids = list(filter(lambda x:x != sid,recipient_session_ids))
        self.emit_result(
          "respond_to_file_transfer_request",
          session_ids=recipient_session_ids,
          result={
            "sender_email":sender_user.email,
            "sender_uid":sender_user.uid,
            "sender_sid":sid,
            "job_id":job_id
          }
        )

      self.emit_result(
        "file_transfer_request",
        session_ids=set([sid]),
        result={
          # hopefully we wont need to go to firebase_manager to get a server_job_id
          "recipients":recipients,
          "job_id":job_id
        }
      )
    except BaseException as e:
      errMsg =get_traceback(e) if is_dev_environment() else None
      if isinstance(e,APINotFoundError):
        if e.description == NEED_MORE_CREDITS:
          errMsg = NEED_MORE_CREDITS
      self.emit_result(
        "file_transfer_error",
        code=http.client.INTERNAL_SERVER_ERROR.value,
        session_ids=set([sid]),
        result=errMsg
      )
      if is_deployed_environment():
        CONFIGS.sentry_manager.debug_with_sentry(errMsg)

  def respond_to_file_transfer_request(self,sid,data):
    try:
      access_token = data["data"]["access_token"]
      sender_sid = data["data"]["sender_sid"]
      sender_peer_id = data["data"]["sender_peer_id"]
      sender_email = data["data"]["sender_email"]
      job_id = data["data"]["job_id"]
      firebase_uid, error = verify_access_token(access_token)
      if error != None:
        self.emit_result(
          "file_transfer_error",
          code=http.client.UNAUTHORIZED.value,
          session_ids=set([sid]),
          result=error if is_dev_environment() else None
        )
        return
      self.emit_result(
        "file_transfer_response",
        session_ids=[sender_sid],
        result={
          "sender_email":sender_email,
          "sender_peer_id":sender_peer_id,
          "job_id":job_id
        }
      )
    except BaseException as e:
      errMsg = get_traceback(e)
      self.emit_result(
        "file_transfer_error",
        code=http.client.INTERNAL_SERVER_ERROR.value,
        session_ids=set([sid]),
        result=errMsg if is_dev_environment() else None
      )
      if is_deployed_environment():
        CONFIGS.sentry_manager.debug_with_sentry(errMsg)

  # PEERJS SERVER IMPLEMENTATION
  def peerjs_connect(self, sid, environ, auth):
    print_if_dev(f'PEERJS: Client connected: {sid}')
    self.server.emit('open', {'id': sid},room=sid,namespace="/peerjs")

  def peerjs_message(self, sid, message):
    if isinstance(message,str):
      message = json.loads(message)
    dst = message.get('dst')
    if dst in self.server.rooms(dst,namespace="/peerjs"):
      message["src"] =sid
      self.server.emit('message', message, room=dst, namespace='/peerjs')
    else:
      self.server.emit('error', {'msg': 'Destination not found'}, room=sid, namespace='/peerjs')

  def peerjs_disconnect(self, sid):
    print_if_dev(f'PEERJS: Client disconnected: {sid}')




