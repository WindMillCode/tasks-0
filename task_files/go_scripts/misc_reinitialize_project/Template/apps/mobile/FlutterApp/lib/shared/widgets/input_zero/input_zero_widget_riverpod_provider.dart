// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class InputZeroRiverpodProviderValue {}

var InputZeroRiverpodProviderInstance = InputZeroRiverpodProviderValue();

class InputZeroRiverpodNotifier
    extends Notifier<InputZeroRiverpodProviderValue> {
  @override
  InputZeroRiverpodProviderValue build() {
    return InputZeroRiverpodProviderInstance;
  }
}

final InputZeroRiverpodProvider =
    NotifierProvider<InputZeroRiverpodNotifier, InputZeroRiverpodProviderValue>(
        () {
  return InputZeroRiverpodNotifier();
});
