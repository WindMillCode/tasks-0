from configs import CONFIGS
from managers.firebase_manager.check_for_authentication import check_for_authentication
from flask import Blueprint, request
from utils.api_msg_format import APIMsgFormat
from handlers.scratchpad_handler import *


scratchpad_endpoint = Blueprint(
    "scratchpad", __name__, url_prefix="/scratchpad")

