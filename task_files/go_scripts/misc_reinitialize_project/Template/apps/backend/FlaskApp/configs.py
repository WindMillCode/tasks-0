from utils.env_vars import ENV_VARS, is_dev_environment, is_test_environment, should_debug_app

class DevConfigs:

  endpoint_msg_codes = {
    'success': 'OK',
    'error': 'ERROR',
    'resp_via_socketio': "RESPONSE_VIA_SOCKETIO"
  }

  app = {}
  _managers ={}

  @property
  def cache_manager(self):
    if not self._managers.get("cache_manager"):
      from managers.cache_manager.cache_manager import CacheManager
      self._managers["cache_manager"] = CacheManager()
    return self._managers["cache_manager"]

  @property
  def email_manager(self):
    if not self._managers.get("email_manager"):
      from managers.email_manager.email_manager import EmailManager
      self._managers["email_manager"] = EmailManager()
    return self._managers["email_manager"]

  @property
  def firebase_manager(self):
    if not self._managers.get("firebase_manager"):
      from managers.firebase_manager.firebase_manager import FirebaseManager
      self._managers["firebase_manager"] = FirebaseManager()
    return self._managers["firebase_manager"]

  @property
  def limiter_manager(self):
    if not self._managers.get("limiter_manager"):
      from managers.limiter_manager.limiter_manager import LimiterManager
      self._managers["limiter_manager"] = LimiterManager()
    return self._managers["limiter_manager"]

  @property
  def logger_manager(self):
    if not self._managers.get("logger_manager"):
      from managers.logger_manager.logger_manager import LoggerManager
      self._managers["logger_manager"] = LoggerManager()
    return self._managers["logger_manager"]

  @property
  def one_signal_manager(self):
    if not self._managers.get("one_signal_manager"):
      from managers.one_signal_manager.one_signal_manager import OneSignalManager
      self._managers["one_signal_manager"] = OneSignalManager()
    return self._managers["one_signal_manager"]

  @property
  def postgresql_manager(self):
    if not self._managers.get("postgresql_manager"):
      from managers.postgresql_manager.postgresql_manager import PostgreSQLManager
      self._managers["postgresql_manager"] = PostgreSQLManager(
        ENV_VARS.get("SQLALCHEMY_POSTGRESSQL_0_CONN_STRING"),
        [
          ENV_VARS.get("CRYPTOGRAPHY_ENCRYPTION_KEY_0")
        ]
      )
    return self._managers["postgresql_manager"]

  @property
  def reportlab_manager(self):
    if not self._managers.get("reportlab_manager"):
      from managers.reportlab_manager.reportlab_manager import ReportlabManager
      self._managers["reportlab_manager"] = ReportlabManager()
    return self._managers["reportlab_manager"]

  @property
  def sentry_manager(self):
    if not self._managers.get("sentry_manager"):
      from managers.sentry_manager.sentry_manager import SentryManager
      self._managers["sentry_manager"] = SentryManager(ENV_VARS.get("SENTRY_DSN"))
    return self._managers["sentry_manager"]

  @property
  def socketio_manager(self):
    if not self._managers.get("socketio_manager"):
      from managers.socketio_manager.socketio_manager import SocketIOManager
      self._managers["socketio_manager"] = SocketIOManager(
        ssl_cert=self.app["ssl_cert"] if is_dev_environment() or is_test_environment() else None,
        ssl_key=self.app["ssl_key"] if is_dev_environment() or is_test_environment() else None,
      )
    return self._managers["socketio_manager"]

  def _create_app_info_obj(self, backend_port=ENV_VARS.get("FLASK_BACKEND_PORT",[Flask_Run_0])):

    mobile_url = "[PROXY_URLS_0]"
    return {
      'mobile_dev_server_domain':mobile_url,
      'server_name': 'example.com:{}'.format(backend_port),
      'domain_name': 'https://example.com:{}'.format(backend_port),
      'server_name_for_mobile':'{}'.format(mobile_url),
      'domain_name_for_mobile':'https://{}'.format(mobile_url),
      'flask_env': 'development',
      'access_control_allow_origin': ['https://example.com:4200', 'https://example.com:[Angular_Run_0]',"http://localhost:5050","https://{}:[Angular_Run_0]".format(ENV_VARS.get("HOST_REMOTE_IPS")),"http://{}:[Angular_Run_0]".format(ENV_VARS.get("HOST_REMOTE_IPS"))],
      # 'access_control_allow_origin':["http://192.168.1.90:[Angular_Run_0]"],
      'frontend_angular_app_domain': 'example.com:[Angular_Run_0]',
      'frontend_angular_app_url': 'https://example.com:[Angular_Run_0]',
      'mobile_android_app_domain':'main_app',
      'mobile_android_app_url':'[PROJECT_NAME]://main_app',
      'mobile_ios_app_domain':'[PROJECT_NAME]',
      'mobile_ios_app_url':'[PROJECT_NAME]://',
      'backend_port': backend_port,
      'ssl_cert': ENV_VARS.get("WML_CERT0", "cert.pem"),
      'ssl_key': ENV_VARS.get("WML_CERT_KEY0", "key.pem"),
    }

  def __init__(self):

    self.app = self._create_app_info_obj()


class DockerDevConfigs(DevConfigs):
  def __init__(self):
    super().__init__()
    self.app['server_name'] = '0.0.0.0:{}'.format(self.app['backend_port'])
    self.app['domain_name'] = 'https://0.0.0.0:{}'.format(self.app['backend_port'])
    self.app['server_name'] = '{}'.format(self.app['mobile_dev_server_domain'])
    self.app['domain_name'] = 'https://{}'.format(self.app['mobile_dev_server_domain'])
    self.app['access_control_allow_origin'] = ["https://localhost:4200"]
    self.app['frontend_angular_app_url'] = "https://localhost:4200"
    self.app['frontend_angular_app_domain'] = "localhost:4200"


class TestConfigs(DevConfigs):
  def __init__(self):
    super().__init__()


class PreviewConfigs(DevConfigs):

  def __init__(self) -> None:
    # self.app = self._create_app_info_obj(5004)
    self.app = self._create_app_info_obj()
    self.app['server_name'] = 'api.preview.[PROJECT_NAME].com'
    self.app['domain_name'] = 'https://api.preview.[PROJECT_NAME].com'
    self.app['server_name_for_mobile'] = 'api.preview.[PROJECT_NAME].com'
    self.app['domain_name_for_mobile'] = 'https://api.preview.[PROJECT_NAME].com'
    self.app['flask_env'] = 'production'
    self.app['access_control_allow_origin'] = ["https://ui.preview.[PROJECT_NAME].com"]
    self.app['frontend_angular_app_url'] = "https://ui.preview.[PROJECT_NAME].com"
    self.app['frontend_angular_app_domain'] = "ui.preview.[PROJECT_NAME].com"

class ProdConfigs(DevConfigs):

  def __init__(self) -> None:
    self.app = self._create_app_info_obj()
    self.app['server_name'] = 'api.[PROJECT_NAME].com'
    self.app['domain_name'] = 'https://api.[PROJECT_NAME].com'
    self.app['server_name_for_mobile'] = 'api.[PROJECT_NAME].com'
    self.app['domain_name_for_mobile'] = 'https://api.[PROJECT_NAME].com'
    self.app['flask_env'] = 'production'
    self.app['access_control_allow_origin'] = ["https://[PROJECT_NAME].com"]
    self.app['frontend_angular_app_url'] = "https://[PROJECT_NAME].com"
    self.app['frontend_angular_app_domain'] = "[PROJECT_NAME].com"

CONFIGS: DevConfigs = {
  'PROD': lambda x: ProdConfigs(),
  'PREVIEW': lambda x: PreviewConfigs(),
  'DEV': lambda x: DevConfigs(),
  'TEST': lambda x: TestConfigs(),
  'DOCKER_DEV': lambda x: DockerDevConfigs()
}[ENV_VARS.get("FLASK_BACKEND_ENV")](None)

if should_debug_app():
  CONFIGS = DockerDevConfigs()
