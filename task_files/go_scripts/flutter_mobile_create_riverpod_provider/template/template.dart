// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class TemplateRiverpodProviderValue {
  TemplateRiverpodProviderValue copyWith() {
    return TemplateRiverpodProviderValue();
  }
}

var TemplateRiverpodProviderInstance =
    TemplateRiverpodProviderValue();

class TemplateRiverpodNotifier  extends Notifier<TemplateRiverpodProviderValue>{


  @override
  TemplateRiverpodProviderValue build() {
    return TemplateRiverpodProviderInstance;
  }
}

final TemplateRiverpodProvider = NotifierProvider<
    TemplateRiverpodNotifier,
    TemplateRiverpodProviderValue>(() {
  return TemplateRiverpodNotifier();
});

