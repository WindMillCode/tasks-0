// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class RadioButtonZeroRiverpodProviderValue {
  RadioButtonZeroRiverpodProviderValue copyWith() {
    return RadioButtonZeroRiverpodProviderValue();
  }
}

var RadioButtonZeroRiverpodProviderInstance =
    RadioButtonZeroRiverpodProviderValue();

class RadioButtonZeroRiverpodNotifier
    extends Notifier<RadioButtonZeroRiverpodProviderValue> {
  @override
  RadioButtonZeroRiverpodProviderValue build() {
    return RadioButtonZeroRiverpodProviderInstance;
  }
}

final RadioButtonZeroRiverpodProvider = NotifierProvider<
    RadioButtonZeroRiverpodNotifier, RadioButtonZeroRiverpodProviderValue>(() {
  return RadioButtonZeroRiverpodNotifier();
});
