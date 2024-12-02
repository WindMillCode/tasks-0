import 'package:tuli/shared/widgets/overlay_zero/overlay_zero_riverpod_provider.dart';
import 'package:tuli/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:tuli/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:tuli/util/riverpod_providers/wml_spacing/wml_spacing.dart';
import 'package:flutter/material.dart';

import 'package:flutter_riverpod/flutter_riverpod.dart';

mixin WMLBaseWidget<T extends ConsumerStatefulWidget> on ConsumerState<T> {
  Widget createStrikeThru({EdgeInsets? padding}) {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    return Padding(
      padding: padding ?? EdgeInsets.all(wmlSpacing.large),
      child: Container(
        height: 1,
        color: wmlColors.white,
      ),
    );
  }

  void toggleOverlayLoading(bool isPresent, {bool? forceClose}) async {
    await Future(() {
      if (!mounted) {
        return;
      }
      final overlayZeroNotifer = ref.read(OverlayZeroRiverpodProvider.notifier);
      overlayZeroNotifer.updateIsPresent(isPresent, forceClose: forceClose);
    });
  }

  Text createFieldLabel(String text) {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    return Text(
      text,
      style: TextStyle(
        fontSize: wmlFonts.captionSmall,
        color: wmlColors.fieldLabelDark0,
      ),
    );
  }

  Text createErrorLabel(FormFieldState<dynamic> state) {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    return Text(
      state.errorText ?? '',
      style: TextStyle(color: wmlColors.error, fontSize: wmlFonts.small),
    );
  }
}
