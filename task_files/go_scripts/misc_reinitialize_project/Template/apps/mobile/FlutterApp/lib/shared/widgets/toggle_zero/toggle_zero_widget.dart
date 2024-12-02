// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:tuli/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:tuli/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:tuli/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class ToggleZeroWidget extends ConsumerStatefulWidget {
  ToggleZeroWidget(
      {super.key,
      this.label = const Text('Provide a Label'),
      this.validator,
      this.value = false,
      this.onChanged,
      this.disabled = false});

  final dynamic label;
  bool value;
  final FormFieldValidator<bool>? validator;
  final void Function(bool)? onChanged;
  bool disabled;

  @override
  ConsumerState<ToggleZeroWidget> createState() => _ToggleZeroState();
}

class _ToggleZeroState extends ConsumerState<ToggleZeroWidget> {
  Widget optionalOverwriteLabel() {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    if (widget.label is Text || widget.label is String) {
      String labelString = "";
      if (widget.label is Text) {
        labelString = (widget.label as Text).data!;
      } else if (widget.label is String) {
        labelString = widget.label;
      }
      return Text(
        labelString,
        style: TextStyle(
          fontSize: wmlFonts.captionSmall,
          color: wmlColors.fieldLabelDark0,
        ),
      );
    }
    return widget.label;
  }

  dynamic onChanged = null;
  bool value = false;

  @override
  void initState() {
    super.initState();
    toggleDisabled();
  }

  @override
  void didUpdateWidget(covariant ToggleZeroWidget oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.disabled != widget.disabled) {
      if (widget.label == "Push Notifications") {
        print(widget.disabled);
        print(oldWidget.disabled);
      }
      toggleDisabled();
    }
  }

  void toggleDisabled() {
    setState(() {
      if (widget.disabled) {
        onChanged = null;
      } else {
        onChanged = (state) {
          return (bool value) {
            state.didChange(value);
            if (widget.onChanged != null) {
              widget.onChanged!(value);
            }
          };
        };
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final contextWidth = MediaQuery.of(context).size.width;

    return Builder(builder: (context) {
      return Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          optionalOverwriteLabel(),
          FormField<bool>(
            key: ValueKey(widget.value),
            initialValue: widget.value,
            validator: widget.validator,
            builder: (FormFieldState<bool> state) {
              return Switch(
                value: state.value ?? false,
                onChanged: onChanged?.call(state),
                thumbIcon: WidgetStateProperty.resolveWith<Icon?>((Set<WidgetState> states) {
                  if (states.contains(WidgetState.selected)) {
                    // Icon when the switch is on
                    return const Icon(Icons.check);
                  }
                  // Icon when the switch is off
                  return const Icon(Icons.close);
                }),
                activeColor: wmlColors.tertiaryColor,
              );
            },
          ),
        ],
      );
    });
  }
}
