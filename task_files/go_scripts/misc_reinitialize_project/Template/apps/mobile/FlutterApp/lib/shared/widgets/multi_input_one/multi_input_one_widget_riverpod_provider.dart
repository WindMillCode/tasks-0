// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class MultiInputOneRiverpodProviderValue {
  MultiInputOneRiverpodProviderValue copyWith() {
    return MultiInputOneRiverpodProviderValue();
  }
}

var MultiInputOneRiverpodProviderInstance =
    MultiInputOneRiverpodProviderValue();

class MultiInputOneRiverpodNotifier
    extends Notifier<MultiInputOneRiverpodProviderValue> {
  @override
  MultiInputOneRiverpodProviderValue build() {
    return MultiInputOneRiverpodProviderInstance;
  }
}

final MultiInputOneRiverpodProvider = NotifierProvider<
    MultiInputOneRiverpodNotifier, MultiInputOneRiverpodProviderValue>(() {
  return MultiInputOneRiverpodNotifier();
});
