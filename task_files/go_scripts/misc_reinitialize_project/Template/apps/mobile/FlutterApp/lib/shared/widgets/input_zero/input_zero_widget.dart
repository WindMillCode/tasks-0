// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:[PROJECT_NAME]/util/misc/misc.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class InputZeroWidget extends ConsumerStatefulWidget {
  InputZeroWidget(
      {super.key,
      dynamic label,
      this.validator,
      this.obscureText = false,
      TextEditingController? textEditingCtrlr,
      TextInputType? keyboardType,
      List<FilteringTextInputFormatter>? inputFormatters,
      bool? enabled,
      FocusNode? focusNode})
      : textEditingCtrlr = textEditingCtrlr ?? TextEditingController(),
        label = label != null ? ensureTextWidget(label) : const Text('Provide a Label'),
        keyboardType = keyboardType ?? TextInputType.text,
        inputFormatters = inputFormatters ?? [],
        enabled = enabled ?? true,
        focusNode = focusNode;

  bool enabled;
  final FocusNode? focusNode;
  final bool obscureText;
  final dynamic label;
  final TextEditingController textEditingCtrlr;
  final FormFieldValidator<String>? validator;
  final TextInputType keyboardType;
  final List<FilteringTextInputFormatter> inputFormatters;

  @override
  ConsumerState<InputZeroWidget> createState() => _InputZeroState();
}

class _InputZeroState extends ConsumerState<InputZeroWidget> {
  Widget optionalOverwriteLabelColor() {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    if (widget.label is Text) {
      return Text(
        (widget.label as Text).data!,
        style: TextStyle(
          fontSize: wmlFonts.captionSmall,
          color: wmlColors.fieldLabelDark0,
        ),
      );
    }
    return widget.label;
  }

  @override
  Widget build(BuildContext context) {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        optionalOverwriteLabelColor(),
        TextFormField(
          focusNode: widget.focusNode,
          enabled: widget.enabled,
          controller: widget.textEditingCtrlr,
          keyboardType: widget.keyboardType,
          inputFormatters: widget.inputFormatters,
          validator: widget.validator,
          obscureText: widget.obscureText,
          style: TextStyle(fontSize: wmlFonts.captionSmall, color: wmlColors.fieldLabelDark0),
          decoration: InputDecoration(
            filled: true,
            isDense: true,
            errorStyle: TextStyle(),
            contentPadding: EdgeInsets.symmetric(horizontal: wmlSpacing.medium, vertical: wmlSpacing.regular),
            fillColor: wmlColors.primaryColor,
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(wmlSpacing.cornerRadius),
              // borderSide: BorderSide.none,
            ),
          ),
        ),
      ],
    );
  }
}
