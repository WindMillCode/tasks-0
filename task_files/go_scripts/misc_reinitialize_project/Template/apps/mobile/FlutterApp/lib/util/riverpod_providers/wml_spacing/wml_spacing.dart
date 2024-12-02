// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class WMLSpacingRiverpodProviderValue {
  final double tiny = 2.0;
  final double small = 4.0;
  final double short = 6.0;
  final double regular = 8.0;
  final double medium = 12.0;
  final double large = 16.0;
  final double extraLarge = 20.0;
  final double huge = 24.0;
  final double enormous = 32.0;
  final double gigantic = 48.0;

  final double cornerRadius = 5.0;
  final double cardRadius = 20.0;
  final double containerRadius = 36.0;
  final double cylinnderRadius = 100.0;

  WMLSpacingRiverpodProviderValue copyWith() {
    return WMLSpacingRiverpodProviderValue();
  }
}

var WMLSpacingRiverpodProviderInstance = WMLSpacingRiverpodProviderValue();

class WMLSpacingRiverpodNotifier
    extends Notifier<WMLSpacingRiverpodProviderValue> {
  @override
  WMLSpacingRiverpodProviderValue build() {
    return WMLSpacingRiverpodProviderInstance;
  }
}

final WMLSpacingRiverpodProvider = NotifierProvider<WMLSpacingRiverpodNotifier,
    WMLSpacingRiverpodProviderValue>(() {
  return WMLSpacingRiverpodNotifier();
});
