import signal
from utils.local_deps import local_deps
local_deps()
from utils.env_vars import ENV_VARS, flask_backend_env, is_deployed_environment, is_dev_environment, is_preview_environment, is_test_environment, should_debug_app
if  is_dev_environment():
  from gevent import monkey
  monkey.patch_all()
  from geventwebsocket.handler import WebSocketHandler
  from gevent.pywsgi import WSGIServer
# if issues use gevent.ssl
from ssl import SSLContext, PROTOCOL_TLS_SERVER
from endpoints.scratchpad_endpoint import scratchpad_endpoint
from endpoints.healthcheck_endpoint import healthcheck_endpoint

import werkzeug
import socketio
import sqlalchemy
from flask import Flask, request, redirect
import sys
import argparse

# import utils.my_util
from utils.api_msg_format import APIMsgFormat
from utils.exceptions.api_exceptions import APIError, APIServerError, APINotFoundError
from utils.exceptions.general_exceptions import get_traceback
from utils.iterable_utils import list_get
from utils.print_if_dev import print_if_dev
from configs import CONFIGS
from flask_cors import CORS


CONFIGS.sentry_manager.init_sentry()
app = Flask(__name__)
CORS(app, supports_credentials=True, origins=[origin for origin in CONFIGS.app['access_control_allow_origin']] + ['vscode-webview*'] if flask_backend_env in ["DOCKER_DEV", "DEV"] else [])

app.config.update(
  FLASK_ENV=CONFIGS.app['flask_env'],
  DEBUG=False,
)
if is_dev_environment():
  app.config.update(
      DEBUG=True
  )
CONFIGS.limiter_manager.limiter.init_app(app)
CONFIGS.cache_manager.cache.init_app(app)


app.register_blueprint(healthcheck_endpoint)
if flask_backend_env in [
  "DEV", "TEST", "DOCKER_DEV"
  #  TODO comment out when ready
  # ,"PREVIEW","PROD"
  #  TODO
]:
  app.register_blueprint(scratchpad_endpoint)


@app.before_request
def upgrade_to_https():
  if not request.is_secure and flask_backend_env in ["DEV", "DOCKER_DEV"]:
    if request.method != 'OPTIONS' or 'Access-Control-Request-Method' not in request.headers:
      url = request.url.replace('http://', 'https://', 1)
      code = 301
      return redirect(url, code=code)


@app.after_request
def after_request(response):
  del response.headers["Server"]

  return response

@app.errorhandler(sqlalchemy.exc.PendingRollbackError)
def restart_because_of_socket_error(err):
  print_if_dev(err.__class__)
  error_string = f"DB Error: {get_traceback(err)}"
  CONFIGS.logger_manager.app_logger.error(
      error_string)
  CONFIGS.sentry_manager.debug_with_sentry(error_string)
  return APIServerError("There was an issue with the DB connection").return_flask_response()


@app.errorhandler(APIError)
def handle_api_exception(err):
  """Return custom JSON when APIError or its children are raised"""
  CONFIGS.logger_manager.app_logger.error(
      err.description, list_get(err.args, 0, "N/A")
  )
  CONFIGS.sentry_manager.debug_with_sentry(err.description)
  # TODO undo once work is finsihed remove is_prod_environment
  if not (is_dev_environment() or is_test_environment() or is_preview_environment()):
  # TODO
    if err.include_desc == False:
      err.description = None
  if err.return_desc_as_string:
    return err.description, err.code
  else:
    return err.return_flask_response()


@app.errorhandler(werkzeug.exceptions.NotFound)
def handle_not_found(err):
  CONFIGS.logger_manager.app_logger.error(
      f"Not found: {get_traceback(err)}")
  return APINotFoundError().return_flask_response()


@app.errorhandler(Exception)
def handle_unknown_exception(err):
  error_string = "Unknown Exception: {} \n {}".format(str(err), get_traceback(err))
  CONFIGS.logger_manager.app_logger.error(
      error_string)
  CONFIGS.sentry_manager.debug_with_sentry(error_string)
  if (isinstance(err, APIError)):
    return APIError().return_flask_response()
  else:
    # TODO  once work is done comment the following lines
    # return APIServerError(desc=get_traceback(err)).return_flask_response()
    # TODO
    if  (is_dev_environment() or is_test_environment() or is_preview_environment()) :
      return APIServerError(desc=get_traceback(err)).return_flask_response()
    else:
      return APIServerError().return_flask_response()


@app.teardown_appcontext
def app_shutdown(event):
  None



app.add_url_rule(
  '/', 'index', (lambda: APIMsgFormat("Hello World").return_flask_response()))

if is_dev_environment():
  app.add_url_rule(
      '/favicon.ico','index2', (lambda: APIMsgFormat("Hello World").return_flask_response()))

def get_ssl_context():
  try:
    ssl_cert = CONFIGS.app["ssl_cert"]
    ssl_key = CONFIGS.app["ssl_key"]

    ssl_context = SSLContext(PROTOCOL_TLS_SERVER)
    ssl_context.load_cert_chain(ssl_cert, ssl_key)
    return ssl_context
  except Exception as e:
    return None

def init_socketio(async_mode="threading"):

  server=socketio.Server(
    logger=False,
    async_mode=async_mode,
    cors_allowed_origins=CONFIGS.app['access_control_allow_origin'],
  )
  CONFIGS.socketio_manager.init_server(server)
  CONFIGS.socketio_manager.add_listeners_via_python_socketio()
  CONFIGS.socketio_manager.app = app
  app.wsgi_app =  socketio.WSGIApp(CONFIGS.socketio_manager.server, app.wsgi_app)

parser =None
args = None
if is_dev_environment():
  parser = argparse.ArgumentParser(
    prog='Flask App',
    description='options for running the flask app',
  )
  parser.add_argument('-r','--reloader_type',default="auto")
  args = parser.parse_args()


def run_app():
  if is_dev_environment():
    app.run(
      reloader_type=args.reloader_type,
      exclude_patterns="site-packages",
      debug=True,
      ssl_context=get_ssl_context(),
      host="0.0.0.0" if flask_backend_env in ["DOCKER_DEV"] else None,
      port=CONFIGS.app["backend_port"]
    )
  else:
    app.run()

def run_app_via_socketio():
  init_socketio()
  if is_dev_environment() or should_debug_app():
    run_app()

def run_app_via_gevent_base(listener,app,ssl_context=None):

  if  is_test_environment() or is_dev_environment() or should_debug_app():
    from geventwebsocket.handler import WebSocketHandler
    from gevent.pywsgi import WSGIServer

  http_server = WSGIServer(
    listener,
    app,
    ssl_context=ssl_context,
    handler_class=WebSocketHandler,
    log=CONFIGS.logger_manager.gevent_logger
  )

  def shutdown_server(signum, frame):
    if is_test_environment():
      return
    print("Shutting down server gracefully...")
    print(signum)
    sys.exit(0)

  signal.signal(signal.SIGTERM, shutdown_server)
  signal.signal(signal.SIGINT, shutdown_server)

  print("""Flask Gevent Server Running on Port: {}.
        You can connect to this server on {}""".format(CONFIGS.app["backend_port"], CONFIGS.app["domain_name"]))
  try:
    http_server.serve_forever()
  except OSError as e:
    pass # silence


def run_app_via_gevent():
  if is_dev_environment() or is_test_environment() or should_debug_app():
    run_app_via_gevent_base(
      CONFIGS.app['server_name'],
      app,
      ssl_context = get_ssl_context()
    )

def run_app_via_gevent_and_socketio():

  init_socketio("gevent")
  run_app_via_gevent_base(
    CONFIGS.app['server_name'],
    app.wsgi_app,
    ssl_context = get_ssl_context()
  )



if __name__ == "__main__":
  # run_app()
  # run_app_via_socketio()
  # run_app_via_gevent()
  run_app_via_gevent_and_socketio()


if  is_deployed_environment() and should_debug_app() == False:
  run_app_via_socketio()

