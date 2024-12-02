// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class CurrentPageControlZeroRiverpodProviderValue {
  CurrentPageControlZeroRiverpodProviderValue copyWith() {
    return CurrentPageControlZeroRiverpodProviderValue();
  }
}

var CurrentPageControlZeroRiverpodProviderInstance =
    CurrentPageControlZeroRiverpodProviderValue();

class CurrentPageControlZeroRiverpodNotifier
    extends Notifier<CurrentPageControlZeroRiverpodProviderValue> {
  @override
  CurrentPageControlZeroRiverpodProviderValue build() {
    return CurrentPageControlZeroRiverpodProviderInstance;
  }
}

final CurrentPageControlZeroRiverpodProvider = NotifierProvider<
    CurrentPageControlZeroRiverpodNotifier,
    CurrentPageControlZeroRiverpodProviderValue>(() {
  return CurrentPageControlZeroRiverpodNotifier();
});
