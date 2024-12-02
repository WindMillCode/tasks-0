// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class BulletPointZeroRiverpodProviderValue {}

var BulletPointZeroRiverpodProviderInstance =
    BulletPointZeroRiverpodProviderValue();

class BulletPointZeroRiverpodNotifier
    extends Notifier<BulletPointZeroRiverpodProviderValue> {
  @override
  BulletPointZeroRiverpodProviderValue build() {
    return BulletPointZeroRiverpodProviderInstance;
  }
}

final BulletPointZeroRiverpodProvider = NotifierProvider<
    BulletPointZeroRiverpodNotifier, BulletPointZeroRiverpodProviderValue>(() {
  return BulletPointZeroRiverpodNotifier();
});
