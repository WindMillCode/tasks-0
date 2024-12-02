// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class ToggleZeroRiverpodProviderValue {}

var ToggleZeroRiverpodProviderInstance = ToggleZeroRiverpodProviderValue();

class ToggleZeroRiverpodNotifier
    extends Notifier<ToggleZeroRiverpodProviderValue> {
  @override
  ToggleZeroRiverpodProviderValue build() {
    return ToggleZeroRiverpodProviderInstance;
  }
}

final ToggleZeroRiverpodProvider = NotifierProvider<ToggleZeroRiverpodNotifier,
    ToggleZeroRiverpodProviderValue>(() {
  return ToggleZeroRiverpodNotifier();
});
