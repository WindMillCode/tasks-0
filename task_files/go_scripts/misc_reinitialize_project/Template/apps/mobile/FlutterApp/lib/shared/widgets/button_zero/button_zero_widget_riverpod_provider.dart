// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class ButtonZeroRiverpodProviderValue {}

var ButtonZeroRiverpodProviderInstance = ButtonZeroRiverpodProviderValue();

class ButtonZeroRiverpodNotifier
    extends Notifier<ButtonZeroRiverpodProviderValue> {
  @override
  ButtonZeroRiverpodProviderValue build() {
    return ButtonZeroRiverpodProviderInstance;
  }
}

final ButtonZeroRiverpodProvider = NotifierProvider<ButtonZeroRiverpodNotifier,
    ButtonZeroRiverpodProviderValue>(() {
  return ButtonZeroRiverpodNotifier();
});


