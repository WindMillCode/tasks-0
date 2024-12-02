// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class SpecificRiverpodProviderValue {

}

var SpecificRiverpodProviderInstance = SpecificRiverpodProviderValue();

class SpecificRiverpodNotifier extends Notifier<SpecificRiverpodProviderValue> {
  @override
  SpecificRiverpodProviderValue build() {
    return SpecificRiverpodProviderInstance;
  }
}

final SpecificRiverpodProvider =
    NotifierProvider<SpecificRiverpodNotifier, SpecificRiverpodProviderValue>(
        () {
  return SpecificRiverpodNotifier();
});