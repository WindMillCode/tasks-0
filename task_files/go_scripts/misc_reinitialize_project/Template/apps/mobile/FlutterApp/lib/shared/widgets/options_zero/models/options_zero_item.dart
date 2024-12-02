import 'package:flutter/material.dart';

class OptionsZeroItem<T> {
  T value;
  bool isChosen;
  final Widget Function(BuildContext, bool) builder;
  // null contextWidth number
  dynamic containerWidth;

  OptionsZeroItem({
    required this.value,
    this.isChosen = false,
    required this.builder,
    this.containerWidth = null,
  });
}
