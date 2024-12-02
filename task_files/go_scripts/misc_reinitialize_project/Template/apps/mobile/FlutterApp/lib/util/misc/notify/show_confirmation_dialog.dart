// ignore_for_file: prefer_const_constructors

import 'package:[PROJECT_NAME]/shared/widgets/button_zero/button_zero_widget.dart';
import 'package:[PROJECT_NAME]/util/misc/misc.dart';
import 'package:flutter/material.dart';
import 'package:flutter_smart_dialog/flutter_smart_dialog.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';

Future<T> showConfirmationDialog<T>(BuildContext context, Function onConfirm, {dynamic title, dynamic subtitle}) async {
  return await SmartDialog.show(
    builder: (BuildContext context) {
      return AlertDialog(
        scrollable: true,
        title: ensureTextWidget(title ?? translate("ConfirmDialogZero.title")),
        content: ensureTextWidget(subtitle ?? translate("ConfirmDialogZero.subtitle")),
        actions: <Widget>[
          ButtonZeroWidget(
            onPressed: () async {
              var result = await onConfirm();
              SmartDialog.dismiss(result: result);
            },
            child: translate("ConfirmDialogZero.options.0"),
          ),
          ButtonZeroWidget(
            onPressed: () {
              SmartDialog.dismiss();
            },
            type: ButtonZeroWidgetType.tertiary,
            child: translate("ConfirmDialogZero.options.1"),
          ),
        ],
      );
    },
  );
}
