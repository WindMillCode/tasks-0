// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class BannerAdRiverpodProviderValue {}

var BannerAdRiverpodProviderInstance = BannerAdRiverpodProviderValue();

class BannerAdRiverpodNotifier extends Notifier<BannerAdRiverpodProviderValue> {
  @override
  BannerAdRiverpodProviderValue build() {
    return BannerAdRiverpodProviderInstance;
  }
}

final BannerAdRiverpodProvider =
    NotifierProvider<BannerAdRiverpodNotifier, BannerAdRiverpodProviderValue>(
        () {
  return BannerAdRiverpodNotifier();
});
