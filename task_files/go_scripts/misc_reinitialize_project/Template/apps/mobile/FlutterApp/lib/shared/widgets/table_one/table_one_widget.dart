// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:[PROJECT_NAME]/shared/widgets/current_page_control_zero/current_page_control_zero_widget.dart';
import 'package:[PROJECT_NAME]/util/wml/pagination/pagination.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'table_one_widget_riverpod_provider.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class TableOneWidget extends ConsumerStatefulWidget {
  TableOneWidget({required super.key, required this.tablePredicate});

  dynamic Function(WMLAPIPageRequestModel) tablePredicate;

  @override
  ConsumerState<TableOneWidget> createState() => _TableOneWidgetState();
}

class _TableOneWidgetState extends ConsumerState<TableOneWidget> {
  onPrevPressed() {
    final rp = ref.read(TableOneWidgetRiverpodProvider(widget.key!));
    final rpNotifier = ref.read(TableOneWidgetRiverpodProvider(widget.key!).notifier);
    rp.reqBody.pageNum -= 1;
    widget.tablePredicate(rp.reqBody);
  }

  onNextPressed() {
    final rp = ref.read(TableOneWidgetRiverpodProvider(widget.key!));
    final rpNotifier = ref.read(TableOneWidgetRiverpodProvider(widget.key!).notifier);
    rp.reqBody.pageNum += 1;
    widget.tablePredicate(rp.reqBody);
  }

  onPageNumberChange(int page) async {
    final rp = ref.read(TableOneWidgetRiverpodProvider(widget.key!));
    final rpNotifier = ref.read(TableOneWidgetRiverpodProvider(widget.key!).notifier);
    rp.reqBody.pageNum = page - 1;
    widget.tablePredicate(rp.reqBody);
  }

  toggleAdditionalControls([bool? isPresent]) {
    final rp = ref.read(TableOneWidgetRiverpodProvider(widget.key!));
    final rpNotifier = ref.read(TableOneWidgetRiverpodProvider(widget.key!).notifier);
    rpNotifier.updateAdditionalControlsIsPresent(isPresent ?? !rp.additionalControlsIsPresent);
  }

  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(TableOneWidgetRiverpodProvider(widget.key!));
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final contextHeight = MediaQuery.of(context).size.height;
    final contextWidth = MediaQuery.of(context).size.width;

    return Column(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Container(
                child: GestureDetector(
                    onTap: toggleAdditionalControls,
                    child: Container(
                      padding: EdgeInsets.all(wmlSpacing.small),
                      decoration: BoxDecoration(border: Border.all(color: Color.fromRGBO(0, 0, 0, 0.0))),
                      child:
                          Padding(padding: EdgeInsets.all(wmlSpacing.medium), child: Icon(rp.additionalControlsIcon)),
                    ))),
            Padding(
              padding: EdgeInsets.fromLTRB(0, 0, wmlSpacing.small, 0),
              child: CurrentPageControlZeroWidget(
                  currentPage: rp.reqBody.pageNum + 1,
                  totalPages: rp.respBody.totalPages,
                  onPageNumberChange: onPageNumberChange),
            )
          ],
        ),
        SizedBox(
          height: wmlSpacing.short,
        ),
        Expanded(
          child: Stack(
            children: [
              Row(
                children: [
                  GestureDetector(
                      onTap: onPrevPressed,
                      child: Container(
                        padding: EdgeInsets.all(wmlSpacing.small),
                        // needed so the hit boundary can be the container and not the icon
                        decoration: BoxDecoration(border: Border.all(color: Color.fromRGBO(0, 0, 0, 0.0))),
                        height: contextHeight,
                        child: Icon(CupertinoIcons.chevron_left),
                      )),
                  Expanded(
                    child: rp.displayWidget,
                  ),
                  GestureDetector(
                      onTap: onNextPressed,
                      child: Container(
                        padding: EdgeInsets.all(wmlSpacing.small),
                        decoration: BoxDecoration(border: Border.all(color: Color.fromRGBO(0, 0, 0, 0.0))),
                        height: contextHeight,
                        child: Icon(CupertinoIcons.chevron_right),
                      )),
                ],
              ),
              (() {
                if (rp.additionalControlsIsPresent) {
                  return rp.additionalControlsWidget;
                } else {
                  return Container();
                }
              })(),
            ],
          ),
        ),
      ],
    );
  }
}
