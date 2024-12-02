// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';

import 'current_page_control_zero_widget_riverpod_provider.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class CurrentPageControlZeroWidget extends ConsumerStatefulWidget {
  int currentPage;
  int totalPages;
  final void Function(int currentPageNum) onPageNumberChange;

  CurrentPageControlZeroWidget({
    super.key,
    this.currentPage = 1,
    required this.totalPages,
    void Function(int currentPageNum)? onPageNumberChange,
  }) : onPageNumberChange = onPageNumberChange ?? ((_) {});

  @override
  ConsumerState<CurrentPageControlZeroWidget> createState() => _CurrentPageControlZeroState();
}

class _CurrentPageControlZeroState extends ConsumerState<CurrentPageControlZeroWidget> {
  late int previousPageNum;
  late FocusNode focusNode;
  late TextEditingController ctrlr;

  @override
  void initState() {
    super.initState();
    setState(() {
      previousPageNum = widget.currentPage;
    });
    focusNode = FocusNode();
    ctrlr = TextEditingController(text: widget.currentPage.toString());
    focusNode.addListener(onPageNumberChange);
  }

  @override
  void didUpdateWidget(covariant CurrentPageControlZeroWidget oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (widget.currentPage != previousPageNum) {
      setState(() {
        previousPageNum = widget.currentPage;
        ctrlr.text = widget.currentPage.toString(); // Update the controller's text
      });
    }
  }

  onPageNumberChange() {
    if (focusNode.hasFocus) {
      ctrlr.clear();
      return;
    }
    // Convert the current text field value to an integer.
    String newPage = ctrlr.text;
    int? currentPageNum = int.tryParse(newPage);
    if (currentPageNum == null) {
      // If the conversion fails, you might want to reset the value or alert the user.
      ctrlr.text = previousPageNum.toString();
      return;
    }

    // Constrain the value between 1 and totalPages.
    if (currentPageNum < 1) {
      currentPageNum = 1;
      ctrlr.text = '1'; // Reset the text field to show the constrained value.
      ctrlr.selection =
          TextSelection.fromPosition(TextPosition(offset: ctrlr.text.length)); // Move cursor to the end of the input
    } else if (currentPageNum > widget.totalPages) {
      currentPageNum = widget.totalPages;
      ctrlr.text = widget.totalPages.toString(); // Reset the text field to show the constrained value.
      // ctrlr.selection = TextSelection.fromPosition(TextPosition(
      //     offset: ctrlr.text.length)); // Move cursor to the end of the input
    }

    setState(() {
      previousPageNum = currentPageNum as int;
    });
    widget.onPageNumberChange(currentPageNum);
  }

  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(CurrentPageControlZeroRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final contextHeight = MediaQuery.of(context).size.height;
    final contextWidth = MediaQuery.of(context).size.width;

    return Row(
      children: [
        Text(translate("CurrentPageZero.text")),
        SizedBox(width: wmlSpacing.medium),
        SizedBox(
          height: 35,
          child: IntrinsicWidth(
            child: ConstrainedBox(
              constraints: BoxConstraints(minWidth: 35), // Set the minimum width to 35
              child: Focus(
                focusNode: focusNode,
                child: TextFormField(
                  controller: ctrlr,
                  keyboardType: TextInputType.numberWithOptions(signed: true),
                  textInputAction: TextInputAction.done,
                  style: TextStyle(color: wmlColors.white),
                  textAlign: TextAlign.center,
                  decoration: InputDecoration(
                    isDense: true, // Reduces additional space within the input field
                    contentPadding: EdgeInsets.symmetric(
                        horizontal: wmlSpacing.medium, vertical: wmlSpacing.regular), // Adjusts the padding
                    border: OutlineInputBorder(
                      borderSide: BorderSide(color: wmlColors.white, width: wmlSpacing.medium),
                      borderRadius: BorderRadius.circular(wmlSpacing.cornerRadius),
                    ),
                  ),
                ),
              ),
            ),
          ),
        ),
        SizedBox(width: wmlSpacing.small),
        Text(
          ' / ', // Total pages
          style: TextStyle(color: wmlColors.white),
        ),
        Text(
          widget.totalPages.toString(), // Total pages
          style: TextStyle(color: wmlColors.white),
        ),
      ],
    );
  }
}
