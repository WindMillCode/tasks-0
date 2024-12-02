// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class CheckboxZeroWidget extends ConsumerStatefulWidget {
  CheckboxZeroWidget({super.key, this.value = false, this.onChanged});
  bool value;
  void Function(bool?)? onChanged;

  @override
  ConsumerState<CheckboxZeroWidget> createState() => _CheckboxZeroState();
}

class _CheckboxZeroState extends ConsumerState<CheckboxZeroWidget> {
  @override
  Widget build(BuildContext context) {
    return Checkbox(
        materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
        visualDensity: VisualDensity(horizontal: -4, vertical: -4),
        value: widget.value,
        onChanged: widget.onChanged ??
            (bool? value) {
              setState(() {
                widget.value = value!;
              });
            });
  }
}
