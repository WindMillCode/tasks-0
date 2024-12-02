import pytest
import datetime
import json
from handlers.wml_template_handler import *
from unit_tests.test_utils.common_utils import ImportStrings
from unit_tests.test_utils.fake_classes import *
from utils.exceptions.api_exceptions import APIAuthorizationError, APIServerError
from utils.wml_libs.pagination import WMLAPIPageRequestModel, WMLAPIPageResponseModel

# Import strings for mocking


class TestImportStrings(ImportStrings):
  def wml_template_handler(self, x=""):
    return "handlers.wml_template_handler" + self.append_submodule(x)


import_strings = TestImportStrings()


def set_mocks(monkeypatch):
  monkeypatch.setattr(
      "handlers.wml_template_handler.APIAuthorizationError",
      FakeAPIAuthorizationError
  )
  monkeypatch.setattr(
      "handlers.wml_template_handler.APIServerError",
      FakeAPIServerError
  )
  monkeypatch.setattr(
      import_strings.wml_template_handler("CONFIGS"),
      FakeCONFIGS
  )


