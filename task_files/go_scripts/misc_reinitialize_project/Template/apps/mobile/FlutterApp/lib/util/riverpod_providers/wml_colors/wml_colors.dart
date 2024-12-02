// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class WMLColorsRiverpodProviderValue {
  final primaryColor = const Color.fromARGB(255, 24, 26, 27);
  final secondaryColor = const Color.fromARGB(255, 54, 58, 61);
  final tertiaryColor = const Color.fromARGB(255, 228, 216, 4);
  final black = Color.fromARGB(255, 0, 0, 0);
  final white = Color.fromARGB(255, 255, 255, 255);
  final starSelected = Color.fromARGB(255, 255, 215, 0);
  final alert = Color.fromARGB(255, 255, 0, 0);
  late Color googlePrimary;
  final facebookPrimary = Color.fromARGB(255, 58, 89, 152);
  final twitterPrimary = Color.fromARGB(255, 29, 161, 242);
  final yahooPrimary = Color.fromARGB(255, 67, 2, 151);
  final microsoftPrimary = Color.fromARGB(255, 127, 186, 0);
  final githubPrimary = Color.fromARGB(255, 24, 23, 23);
  late Color applePrimary;
  final popupBackground = Color.fromARGB(255, 112, 112, 112);
  final successDark = Color.fromARGB(255, 132, 225, 188);
  final successLight = Color.fromARGB(255, 1, 71, 55);
  late Color navColor;
  final navColorDark = Color.fromARGB(255, 24, 26, 27);
  final inputBackgroundColorLight0 = Color.fromARGB(255, 211, 211, 211);
  final inputBackgroundColorDark0 = Color.fromARGB(255, 59, 59, 59);
  final fieldLabelDark0 = Color.fromARGB(255, 173, 216, 230);
  final blockColorDark0 = Color.fromARGB(255, 31, 30, 36);
  final blockColorDark1 = Color.fromARGB(255, 14, 14, 16);
  final Color snackBarSuccessColor = Colors.green;
  final Color snackBarErrorColor = Colors.red;
  final Color snackBarWarningColor = Colors.orange;
  final Color snackBarInfoColor = Colors.blue;
  final Color snackBarNoticeColor = Colors.grey;

  final notifyColor0 = Color(0xFF9D50DD);
  final notifyLedColor0 = Color.fromARGB(255, 255, 255, 255);
  late Gradient wmlGradient0;
  late Gradient wmlGradient1;

  // TODO: to be removed they are form the old app
  final quaternaryColor = const Color(0xFF8B0000);
  final iconPrimaryColor = const Color(0xFF70B1F5);
  final backgroundPrimaryColor = const Color(0xFF7A7A7A);
  final backgroundSecondaryColor = const Color(0xFFD7D7D7);

  final error = Colors.red;
  final dark = Color(0xff121212);

  WMLColorsRiverpodProviderValue() {
    navColor = black;
    applePrimary = black;
    googlePrimary = white;
    wmlGradient0 = RadialGradient(
      center: Alignment(-1, -2),
      radius: 5.5,
      colors: [
        Colors.grey,
        black,
      ],
      stops: [0.0, 4.5],
    );
    wmlGradient1 = RadialGradient(
      center: Alignment(1.5, 1.5),
      radius: 1.5,
      colors: [
        primaryColor,
        secondaryColor,
      ],
      stops: [0.0, 1.5],
    );
  }
}

var WMLColorsRiverpodProviderInstance = WMLColorsRiverpodProviderValue();

class WMLColorsRiverpodNotifier
    extends Notifier<WMLColorsRiverpodProviderValue> {
  @override
  WMLColorsRiverpodProviderValue build() {
    return WMLColorsRiverpodProviderInstance;
  }
}

final WMLColorsRiverpodProvider =
    NotifierProvider<WMLColorsRiverpodNotifier, WMLColorsRiverpodProviderValue>(
        () {
  return WMLColorsRiverpodNotifier();
});
