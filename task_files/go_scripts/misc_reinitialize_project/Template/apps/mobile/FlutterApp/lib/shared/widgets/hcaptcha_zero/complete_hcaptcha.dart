// ignore_for_file: prefer_const_constructors

import 'package:[PROJECT_NAME]/shared/widgets/hcaptcha_zero/hcaptcha_zero_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_smart_dialog/flutter_smart_dialog.dart';

Future<T> completeHcaptcha<T>(BuildContext context, [Function(String)? pOnCaptchaResult]) async {
  return await SmartDialog.show(builder: (BuildContext context) {
    return HcaptchaZeroWidget(onCaptchaResult: (result) async {
      if (pOnCaptchaResult == null) {
        SmartDialog.dismiss(result: result);
        return;
      }
      SmartDialog.dismiss(result: await pOnCaptchaResult(result));
    });
  });
}
