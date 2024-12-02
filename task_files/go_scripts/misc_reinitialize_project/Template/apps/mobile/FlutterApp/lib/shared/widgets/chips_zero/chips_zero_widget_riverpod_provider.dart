// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class ChipsZeroRiverpodProviderValue {
  ChipsZeroRiverpodProviderValue copyWith() {
    return ChipsZeroRiverpodProviderValue();
  }
}

var ChipsZeroRiverpodProviderInstance = ChipsZeroRiverpodProviderValue();

class ChipsZeroRiverpodNotifier
    extends Notifier<ChipsZeroRiverpodProviderValue> {
  @override
  ChipsZeroRiverpodProviderValue build() {
    return ChipsZeroRiverpodProviderInstance;
  }
}

final ChipsZeroRiverpodProvider =
    NotifierProvider<ChipsZeroRiverpodNotifier, ChipsZeroRiverpodProviderValue>(
        () {
  return ChipsZeroRiverpodNotifier();
});
