// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:tuli/util/misc/misc.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'radio_button_zero_widget_riverpod_provider.dart';
import 'package:tuli/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:tuli/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:tuli/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:tuli/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class RadioButtonZeroWidget extends ConsumerStatefulWidget {
  RadioButtonZeroWidget({
    super.key,
    dynamic text,
    required this.isChosen,
  }) : text = ensureTextWidget(text) ?? ensureTextWidget("Radio Option");

  dynamic text;
  final bool isChosen;

  @override
  ConsumerState<RadioButtonZeroWidget> createState() => _RadioButtonZeroState();
}

class _RadioButtonZeroState extends ConsumerState<RadioButtonZeroWidget> {
  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(RadioButtonZeroRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final contextHeight = MediaQuery.of(context).size.height;
    final contextWidth = MediaQuery.of(context).size.width;

    return Row(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Container(
          height: wmlSpacing.large,
          width: wmlSpacing.large,
          decoration: BoxDecoration(
              border: Border.all(color: wmlColors.secondaryColor),
              color: widget.isChosen ? wmlColors.tertiaryColor : Colors.transparent,
              shape: BoxShape.circle),
        ),
        SizedBox(width: wmlSpacing.medium),
        Expanded(child: widget.text)
      ],
    );
  }
}
