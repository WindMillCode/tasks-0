// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter_riverpod/flutter_riverpod.dart';

class DocZeroPageRiverpodProviderValue {}

var DocZeroPageRiverpodProviderInstance = DocZeroPageRiverpodProviderValue();

class DocZeroPageRiverpodNotifier
    extends Notifier<DocZeroPageRiverpodProviderValue> {
  @override
  DocZeroPageRiverpodProviderValue build() {
    return DocZeroPageRiverpodProviderInstance;
  }
}

final DocZeroPageRiverpodProvider = NotifierProvider<
    DocZeroPageRiverpodNotifier, DocZeroPageRiverpodProviderValue>(() {
  return DocZeroPageRiverpodNotifier();
});
