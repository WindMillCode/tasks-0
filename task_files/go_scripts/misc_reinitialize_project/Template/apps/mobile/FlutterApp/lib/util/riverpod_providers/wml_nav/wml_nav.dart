// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:tuli/pages/doc_zero/doc_zero_page.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class WMLNavRiverpodProviderValue {
  Map<String, dynamic> route = {
    "path": "/legal/privacy-policy",
    "title": "Privacy Policy",
    "widget": DocZeroPageWidget(),
  };


  List<String> authenticatedRoutes = [];

  WMLNavRiverpodProviderValue() {
    authenticatedRoutes = [
      route["path"],
    ];
  }
}

var WMLNavRiverpodProviderInstance = WMLNavRiverpodProviderValue();

class WMLNavRiverpodNotifier extends Notifier<WMLNavRiverpodProviderValue> {
  @override
  WMLNavRiverpodProviderValue build() {
    return WMLNavRiverpodProviderInstance;
  }
}

final WMLNavRiverpodProvider = NotifierProvider<WMLNavRiverpodNotifier, WMLNavRiverpodProviderValue>(() {
  return WMLNavRiverpodNotifier();
});
