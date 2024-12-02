import 'package:[PROJECT_NAME]/shared/widgets/hcaptcha_zero/complete_hcaptcha.dart';
import 'package:[PROJECT_NAME]/util/env/env.dart';
import 'package:[PROJECT_NAME]/util/misc/misc.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';
import 'package:flutter/material.dart';
import 'package:flutter_smart_dialog/flutter_smart_dialog.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:sentry_flutter/sentry_flutter.dart';

class OverlayZeroRiverpodProviderValue {
  bool isPresent;
  int counter;

  OverlayZeroRiverpodProviderValue({this.isPresent = false, this.counter = 0});

  OverlayZeroRiverpodProviderValue copyWith({
    bool? isPresent,
    int? counter,
  }) {
    return OverlayZeroRiverpodProviderValue(
      isPresent: isPresent ?? this.isPresent,
      counter: counter ?? this.counter,
    );
  }
}

void showSnackBarMsg(String msg, {Color? snackBarColor, Duration? displayTime = const Duration(seconds: 5)}) {
  snackBarColor ??= WMLColorsRiverpodProviderInstance.snackBarErrorColor;

  SmartDialog.showToast(msg, useAnimation: true, alignment: Alignment.bottomRight, displayTime: displayTime,
      builder: (context) {
    return Container(
      width: MediaQuery.of(context).size.width,
      padding: EdgeInsets.all(WMLSpacingRiverpodProviderInstance.medium),
      decoration: BoxDecoration(color: snackBarColor),
      child: Text('$msg', style: TextStyle(color: determineContrastColor(snackBarColor!))),
    );
  }, maskColor: snackBarColor);
}

class OverlayZeroProcessRequestPredicateResult {
  String? msg;
  Color? snackBarColor;

  OverlayZeroProcessRequestPredicateResult({this.msg, this.snackBarColor});
}

processRequest(
    BuildContext? context, WidgetRef ref, Future<OverlayZeroProcessRequestPredicateResult> Function() predicate,
    {Function customShowMsg = showSnackBarMsg, challengeWithHcaptcha = false, toggleOverlay = true}) async {
  if (challengeWithHcaptcha) {
    // if(false){
    var hcaptchaPass = await completeHcaptcha(context!);
    if (hcaptchaPass != "success") {
      return;
    }
  }

  final wmlColors = WMLColorsRiverpodProviderInstance;
  late final overlayZeroNotifier;
  if (toggleOverlay) {
    overlayZeroNotifier = ref.read(OverlayZeroRiverpodProvider.notifier);
    overlayZeroNotifier.updateIsPresent(true);
  }
  var msg = "";
  Color snackBarColor = wmlColors.snackBarErrorColor;
  var result;
  try {
    result = await predicate();
    msg = result.msg ?? msg;
  } catch (err, stack) {
    msg = APPENV.type.isProdEnvironment ? "${err.toString()}\n${stack.toString()}" : "global.systemError";
    Sentry.captureException(err, stackTrace: stack);
  }
  if (msg.isNotEmpty) {
    customShowMsg(translate(msg), snackBarColor: result?.snackBarColor ?? snackBarColor);
  }
  if (toggleOverlay) {
    overlayZeroNotifier.updateIsPresent(false);
  }
}

class OverlayZeroRiverpodNotifier extends StateNotifier<OverlayZeroRiverpodProviderValue> {
  OverlayZeroRiverpodNotifier() : super(OverlayZeroRiverpodProviderValue());

  void updateIsPresent(bool val, {bool? forceClose}) {
    if (forceClose == true) {
      state = state.copyWith(isPresent: false, counter: 0);
      return;
    }
    if (val) {
      state = state.copyWith(isPresent: true, counter: state.counter + 1);
    } else if (state.counter > 0) {
      int newCounter = state.counter - 1;
      state = state.copyWith(isPresent: newCounter > 0, counter: newCounter);
    }
  }
}

final OverlayZeroRiverpodProvider =
    StateNotifierProvider<OverlayZeroRiverpodNotifier, OverlayZeroRiverpodProviderValue>(
  (ref) => OverlayZeroRiverpodNotifier(),
);
