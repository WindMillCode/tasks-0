// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class I18NRiverpodProviderValue {
  // Locale locale = PlatformDispatcher.instance.locale;
  Locale locale = Locale("en");
  List<String> supportedLocales = [
    "en",
    "am",
    "bn",
    "de",
    "es",
    "fr",
    "hi",
    "ms",
    "sw",
    "uk",
    "zh"
  ];
  I18NRiverpodProviderValue copyWith({Locale? locale}) {
    return I18NRiverpodProviderValue()..locale = locale ?? this.locale;
  }
}

var I18NRiverpodProviderInstance = I18NRiverpodProviderValue();

class I18NRiverpodNotifier extends Notifier<I18NRiverpodProviderValue> {
  @override
  I18NRiverpodProviderValue build() {
    return I18NRiverpodProviderInstance;
  }

  void updateLocale(Locale locale) {
    state = state.copyWith(locale: locale);
  }
}

final I18NRiverpodProvider =
    NotifierProvider<I18NRiverpodNotifier, I18NRiverpodProviderValue>(() {
  return I18NRiverpodNotifier();
});
