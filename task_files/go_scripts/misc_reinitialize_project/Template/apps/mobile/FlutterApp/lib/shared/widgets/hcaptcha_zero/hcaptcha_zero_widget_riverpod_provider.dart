// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class HcaptchaZeroRiverpodProviderValue {
  var passedAuthHcaptcha = false;
  HcaptchaZeroRiverpodProviderValue copyWith({
        bool? passedAuthHcaptcha,
  }) {
    return HcaptchaZeroRiverpodProviderValue()..passedAuthHcaptcha = passedAuthHcaptcha ?? this.passedAuthHcaptcha;

  }
}

var HcaptchaZeroRiverpodProviderInstance = HcaptchaZeroRiverpodProviderValue();

class HcaptchaZeroRiverpodNotifier extends Notifier<HcaptchaZeroRiverpodProviderValue> {
  @override
  HcaptchaZeroRiverpodProviderValue build() {
    return HcaptchaZeroRiverpodProviderInstance;
  }

    void updatePassedAuthHcaptcha([bool? passedAuthHcaptcha]) {
    state = state.copyWith(
      passedAuthHcaptcha: passedAuthHcaptcha,
    );
  }
}

final HcaptchaZeroRiverpodProvider =
    NotifierProvider<HcaptchaZeroRiverpodNotifier, HcaptchaZeroRiverpodProviderValue>(() {
  return HcaptchaZeroRiverpodNotifier();
});
