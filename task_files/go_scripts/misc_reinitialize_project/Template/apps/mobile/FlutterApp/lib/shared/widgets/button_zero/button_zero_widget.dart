// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:flutter/material.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:[PROJECT_NAME]/shared/widgets/overlay_zero/overlay_zero_riverpod_provider.dart';

enum ButtonZeroWidgetType { primary, secondary, tertiary, quaternary, quinary }

class ButtonZeroWidget extends ConsumerStatefulWidget {
  ButtonZeroWidget(
      {super.key,
      this.child = const Text('Click Me'),
      this.onPressed,
      this.type = ButtonZeroWidgetType.primary,
      this.extraButtonStyle = const ButtonStyle(),
      this.extraTextStyle = const TextStyle()});

  dynamic child;
  final void Function()? onPressed;
  ButtonZeroWidgetType type;
  ButtonStyle extraButtonStyle;
  TextStyle extraTextStyle;

  @override
  ConsumerState<ButtonZeroWidget> createState() => _ButtonZeroState();
}

class _ButtonZeroState extends ConsumerState<ButtonZeroWidget> {
  @override
  Widget build(BuildContext context) {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    ButtonStyle buttonStyle = ButtonStyle();
    TextStyle textStyle = TextStyle();
    if (widget.type == ButtonZeroWidgetType.primary) {
      buttonStyle = ButtonStyle();
    }
    if (widget.type == ButtonZeroWidgetType.secondary) {
      buttonStyle = ButtonStyle(
        backgroundColor: WidgetStateProperty.all(wmlColors.secondaryColor),
      );
    }
    if (widget.type == ButtonZeroWidgetType.tertiary) {
      textStyle = TextStyle(color: wmlColors.black);
      buttonStyle = ButtonStyle(
        side: WidgetStateProperty.resolveWith<BorderSide>((Set<WidgetState> states) {
          return BorderSide(color: wmlColors.black);
        }),
        backgroundColor: WidgetStateProperty.all(wmlColors.tertiaryColor),
      );
    }
    if (widget.type == ButtonZeroWidgetType.quaternary) {
      buttonStyle = ButtonStyle(
        side: WidgetStateProperty.all(BorderSide(color: wmlColors.tertiaryColor)),
        backgroundColor: WidgetStateProperty.all(wmlColors.primaryColor),
      );
      // textStyle = TextStyle(color: wmlColors.tertiaryColor);
    }
    if (widget.child is String) {
      textStyle = textStyle.merge(widget.extraTextStyle);
      widget.child = Text(
        widget.child,
        style: textStyle,
        softWrap: true,
        overflow: TextOverflow.visible,
      );
    }
    buttonStyle = buttonStyle.merge(widget.extraButtonStyle);

    return ElevatedButton(
        style: buttonStyle,
        onPressed: widget.onPressed ??
            () {
              showSnackBarMsg(translate("global.comingSoon"), snackBarColor: wmlColors.snackBarNoticeColor);
            },
        child: widget.child);
  }
}
