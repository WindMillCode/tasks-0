// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:tuli/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:tuli/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:tuli/util/wml/pagination/pagination.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';

class TableOneWidgetRiverpodProviderValue {
  Widget displayWidget = Align(
      alignment: Alignment.topCenter,
      child: Text(
        translate("TableOne.noData"),
        style: TextStyle(fontSize: WMLFontsRiverpodProviderInstance.heading2),
      ));
  WMLAPIPageRequestModel reqBody = WMLAPIPageRequestModel();
  bool additionalControlsIsPresent = false;
  Widget additionalControlsWidget = Container(
    color: WMLColorsRiverpodProviderInstance.popupBackground,
  );
  WMLAPIPageResponseModel respBody = WMLAPIPageResponseModel();
  IconData additionalControlsIcon = CupertinoIcons.chevron_down;

  TableOneWidgetRiverpodProviderValue copyWith(
      {Widget? displayWidget,
      WMLAPIPageRequestModel? reqBody,
      WMLAPIPageResponseModel? respBody,
      bool? additionalControlsIsPresent,
      Widget? additionalControlsWidget,
      IconData? additionalControlsIcon}) {
    return TableOneWidgetRiverpodProviderValue()
      ..displayWidget = displayWidget ?? this.displayWidget
      ..reqBody = reqBody ?? this.reqBody
      ..respBody = respBody ?? this.respBody
      ..additionalControlsWidget = additionalControlsWidget ?? this.additionalControlsWidget
      ..additionalControlsIsPresent = additionalControlsIsPresent ?? this.additionalControlsIsPresent
      ..additionalControlsIcon = additionalControlsIcon ?? this.additionalControlsIcon;
  }
}

var TableOneWidgetRiverpodProviderInstance = TableOneWidgetRiverpodProviderValue();

class TableOneWidgetRiverpodNotifier extends FamilyNotifier<TableOneWidgetRiverpodProviderValue, Key> {
  @override
  TableOneWidgetRiverpodProviderValue build(Key arg) {
    return TableOneWidgetRiverpodProviderInstance;
  }

  void updateDisplayWidget(Widget displayWidget) {
    state = state.copyWith(displayWidget: displayWidget);
  }

  void updateReqBody(WMLAPIPageRequestModel reqBody) {
    state = state.copyWith(reqBody: reqBody);
  }

  void updateRespBody(WMLAPIPageResponseModel respBody) {
    state = state.copyWith(respBody: respBody);
  }

  void updateAdditionalControlsIsPresent(bool additionalControlsIsPresent) {
    var additionalControlsIcon = additionalControlsIsPresent ? CupertinoIcons.chevron_up : CupertinoIcons.chevron_down;
    state = state.copyWith(
        additionalControlsIsPresent: additionalControlsIsPresent, additionalControlsIcon: additionalControlsIcon);
  }

  void updateAdditionalControlsWidget(Widget additionalControlsWidget) {
    state = state.copyWith(additionalControlsWidget: additionalControlsWidget);
  }
}

final TableOneWidgetRiverpodProvider =
    NotifierProvider.family<TableOneWidgetRiverpodNotifier, TableOneWidgetRiverpodProviderValue, Key>(() {
  return TableOneWidgetRiverpodNotifier();
});
