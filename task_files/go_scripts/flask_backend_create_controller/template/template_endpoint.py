from configs import CONFIGS
from flask import Blueprint,request
from urllib.parse import urlparse

from utils.api_msg_format import APIMsgFormat
from handlers.wml_template_handler  import  post_endpoint

wml_template_endpoint =Blueprint("wml_template", __name__, url_prefix="/wml_template_url_prefix")

@wml_template_endpoint.route('/get',methods=['GET'])
def wml_template_get_endpoint():
  res = APIMsgFormat(msg="A-OK",code=CONFIGS.endpointMsgCodes["success"])
  return res.return_flask_response(),200


@wml_template_endpoint.route('/post',methods=['POST'])
def wml_template_post_endpoint():
  data = request.json.get('data',{})
  new_data = post_endpoint(data)
  res = APIMsgFormat(data=data, msg=request.url,code=CONFIGS.endpointMsgCodes["success"])
  return res.return_flask_response(),200

