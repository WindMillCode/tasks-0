// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class OptionsZeroRiverpodProviderValue {
  OptionsZeroRiverpodProviderValue copyWith() {
    return OptionsZeroRiverpodProviderValue();
  }
}

var OptionsZeroRiverpodProviderInstance = OptionsZeroRiverpodProviderValue();

class OptionsZeroRiverpodNotifier
    extends Notifier<OptionsZeroRiverpodProviderValue> {
  @override
  OptionsZeroRiverpodProviderValue build() {
    return OptionsZeroRiverpodProviderInstance;
  }
}

final OptionsZeroRiverpodProvider = NotifierProvider<
    OptionsZeroRiverpodNotifier, OptionsZeroRiverpodProviderValue>(() {
  return OptionsZeroRiverpodNotifier();
});
