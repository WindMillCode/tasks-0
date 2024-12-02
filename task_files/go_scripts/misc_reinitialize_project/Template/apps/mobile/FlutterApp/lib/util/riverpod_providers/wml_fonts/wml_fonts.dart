// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class WMLFontsRiverpodProviderValue {
  // Default sizes
  final double small = 12.0;

  final double regular = 16.0;
  final double medium = 18.0;
  final double large = 24.0;
  final double extraLarge = 32.0;

  // Headings
  final double heading0 = 48.0;
  final double heading1 = 36.0;
  final double heading2 = 28.0;
  final double heading3 = 24.0;
  final double heading3_5 = 22.0;
  final double heading4 = 20.0;
  final double heading5 = 18.0;

  // Subtitles
  final double subtitle1 = 18.0;
  final double subtitle2 = 16.0;

  // Buttons
  final double buttonSmall = 14.0;
  final double buttonMedium = 16.0;
  final double buttonLarge = 18.0;

  // Captions
  final double captionSmall = 12.0;
  final double captionMedium = 14.0;
  final double captionLarge = 16.0;

  // Labels
  final double labelSmall = 12.0;
  final double labelMedium = 14.0;
  final double labelLarge = 16.0;

  // Notes
  final double noteSmall = 12.0;
  final double noteMedium = 14.0;
  final double noteLarge = 16.0;

  // Icon sizes
  final double iconSmall = 16.0;
  final double iconMedium = 24.0;
  final double iconLarge = 32.0;
  final double iconXLarge = 48.0;
}

var WMLFontsRiverpodProviderInstance = WMLFontsRiverpodProviderValue();

class WMLFontsRiverpodNotifier extends Notifier<WMLFontsRiverpodProviderValue> {
  @override
  WMLFontsRiverpodProviderValue build() {
    return WMLFontsRiverpodProviderInstance;
  }
}

final WMLFontsRiverpodProvider =
    NotifierProvider<WMLFontsRiverpodNotifier, WMLFontsRiverpodProviderValue>(
        () {
  return WMLFontsRiverpodNotifier();
});
