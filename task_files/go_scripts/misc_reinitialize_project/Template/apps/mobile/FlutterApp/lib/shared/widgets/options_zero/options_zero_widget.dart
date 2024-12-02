// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:[PROJECT_NAME]/shared/widgets/options_zero/models/options_zero_item.dart';
import 'package:[PROJECT_NAME]/util/misc/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'options_zero_widget_riverpod_provider.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class OptionsZeroWidget extends ConsumerStatefulWidget {
  OptionsZeroWidget({
    super.key,
    List<OptionsZeroItem>? chosen,
    required this.options,
    int? limit,
    required this.toggleOptionEvent,
  }) : limit = limit ?? INT_MAX_VALUE;

  List<OptionsZeroItem> chosen = [];
  List<OptionsZeroItem> options;
  int limit;
  void Function(Map<String, List<OptionsZeroItem>>) toggleOptionEvent;

  @override
  ConsumerState<OptionsZeroWidget> createState() => _OptionsZeroState();
}

class _OptionsZeroState extends ConsumerState<OptionsZeroWidget> {
  @override
  void initState() {
    super.initState();
    widget.chosen = widget.options.where((item) => item.isChosen).toList();
  }

  @override
  void didUpdateWidget(OptionsZeroWidget oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.options != widget.options) {
      widget.chosen = widget.options.where((item) => item.isChosen).toList();
    }
  }

  void toggleOption(OptionsZeroItem item) {
    setState(() {
      if (item.isChosen) {
        item.isChosen = false;
        widget.chosen.remove(item);
      } else {
        if (widget.chosen.length >= widget.limit) {
          OptionsZeroItem firstItem = widget.chosen.removeAt(0);
          firstItem.isChosen = false;
        }
        item.isChosen = true;
        widget.chosen.add(item);
      }
    });

    widget.toggleOptionEvent({
      "latest": [item],
      "chosen": widget.chosen
    });
  }

  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(OptionsZeroRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final contextHeight = MediaQuery.of(context).size.height;
    final contextWidth = MediaQuery.of(context).size.width;

    return Wrap(
      children: List.generate(widget.options.length, (index) {
        var item = widget.options[index];
        return GestureDetector(
          onTap: () => toggleOption(item),
          child: Container(
              width: item.containerWidth == "contextWidth" ? contextWidth : item.containerWidth,
              child: AbsorbPointer(child: item.builder(context, item.isChosen))),
        );
      }),
    );
  }
}
