// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:[PROJECT_NAME]/shared/widgets/button_zero/button_zero_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';

import 'chips_zero_widget_riverpod_provider.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class ChipsZeroWidget extends ConsumerStatefulWidget {
  ChipsZeroWidget({super.key, List<String>? items, double? containerHeight, this.onChanged})
      : items = items ?? ["Default", "chips", "to", "be", "replaced", "by", "passing", "'items'", "property"],
        containerHeight = containerHeight ?? 300.0;

  List<String> items;
  double containerHeight;
  final void Function(Map<String, dynamic>)? onChanged;

  @override
  ConsumerState<ChipsZeroWidget> createState() => _ChipsZeroState();
}

class _ChipsZeroState extends ConsumerState<ChipsZeroWidget> {
  final FocusNode inputFocusNode = FocusNode();
  final inputCtrlr = TextEditingController();

  @override
  void initState() {
    super.initState();
    inputFocusNode.addListener(onInputFocusChange);
  }

  @override
  void dispose() {
    inputFocusNode.removeListener(onInputFocusChange);
    inputFocusNode.dispose();
    inputCtrlr.dispose();
    super.dispose();
  }

  addChip() {
    setState(() {
      widget.items.add(inputCtrlr.text);
    });
    widget.onChanged?.call({"action": "ADD", "value": inputCtrlr.text, "current": widget.items});

    setState(() {
      inputCtrlr.clear(); // Optionally clear the text field after adding the item
    });
  }

  // TODO focus blur event does not work anymore
  void onInputFocusChange() {
    if (!inputFocusNode.hasFocus) {}
  }

  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(ChipsZeroRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final contextHeight = MediaQuery.of(context).size.height;
    final contextWidth = MediaQuery.of(context).size.width;

    return Container(
      padding: EdgeInsets.all(wmlSpacing.short),
      decoration: BoxDecoration(
        border: Border.all(color: wmlColors.white),
      ),
      width: contextWidth,
      height: widget.containerHeight,
      child: Column(
        children: [
          Expanded(
            child: SingleChildScrollView(
              child: Wrap(
                spacing: wmlSpacing.short,
                children: List.generate(widget.items.length, (index) {
                  var item = widget.items[index];
                  return InputChip(
                    label: Text(item),
                    deleteIconColor: wmlColors.white,
                    onDeleted: () {
                      var item;
                      setState(() {
                        item = widget.items.removeAt(index);
                      });
                      widget.onChanged?.call({"action": "DELETE", "value": item, "current": widget.items});
                    },
                  );
                }),
              ),
            ),
          ),
          TextFormField(
            maxLines: 2,
            // focusNode: inputFocusNode,
            controller: inputCtrlr,
            decoration: InputDecoration(
              hintText: translate("global.wmlChipsZero.placeholder"),
            ),
          ),
          Align(
            alignment: Alignment.bottomRight,
            child: Row(
              children: [
                ButtonZeroWidget(
                  child: translate("global.wmlChipsZero.enterBtn"),
                  onPressed: () {
                    addChip();
                  },
                ),
                ButtonZeroWidget(
                  child: translate("global.wmlChipsZero.clearBtn"),
                  onPressed: () {
                    setState(() {
                      widget.items.clear();
                    });
                    widget.onChanged?.call({"action": "CLEAR", "value": "", "current": widget.items});
                  },
                ),
              ],
            ),
          )
        ],
      ),
    );
  }
}
