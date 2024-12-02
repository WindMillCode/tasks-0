// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class CheckboxZeroRiverpodProviderValue {}

var CheckboxZeroRiverpodProviderInstance = CheckboxZeroRiverpodProviderValue();

class CheckboxZeroRiverpodNotifier
    extends Notifier<CheckboxZeroRiverpodProviderValue> {
  @override
  CheckboxZeroRiverpodProviderValue build() {
    return CheckboxZeroRiverpodProviderInstance;
  }
}

final CheckboxZeroRiverpodProvider = NotifierProvider<
    CheckboxZeroRiverpodNotifier, CheckboxZeroRiverpodProviderValue>(() {
  return CheckboxZeroRiverpodNotifier();
});
