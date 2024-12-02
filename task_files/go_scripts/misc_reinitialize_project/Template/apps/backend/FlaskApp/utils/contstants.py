from enum import Enum

from utils.env_vars import ENV_VARS


SUCCESS = "SUCCESS"
FAILED = "FAILED"
ISSUES ="ISSUES"
HOST_REMOTE_IPS = ENV_VARS.get("HOST_REMOTE_IPS", "").split(",")

# dev additions

