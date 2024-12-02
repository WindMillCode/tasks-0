// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:tuli/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:flutter/material.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';


bool CAN_HANDLE_ANIMATIONS = true;

dynamic ensureTextWidget(dynamic child,
    {bool shouldTranslate = false, Widget Function(String)? customWidget, double? fontSize}) {
  if (fontSize == null) {
    fontSize = WMLFontsRiverpodProviderInstance.large;
  }
  if (child is String) {
    if (shouldTranslate) {
      child = translate(child);
    }
    if (customWidget != null) {
      return customWidget(child);
    }
    return Text(child, style: TextStyle(fontSize: fontSize));
  } else if (child is Widget) {
    return child;
  }
}

String bulletPointString = "\u2022";

Widget onFailedImageLoadError(BuildContext context, Object error, StackTrace? stackTrace) {
  return Image.asset('assets/media/app-logo-no-bg.png');
}

// Widget onFailedImageLoadErrorForCachedNetworkImage(
//     BuildContext context,String url, Object error) {
//   return Image.asset('assets/media/app-logo-no-bg.png');
// }

Color determineContrastColor(Color backgroundColor, {Color light = Colors.white, Color dark = Colors.black}) {
  return backgroundColor.computeLuminance() < 0.5 ? light : dark;
}

// dev additions

String? toJsonReturnOrginalVal(String? val) {
  return val;
}
