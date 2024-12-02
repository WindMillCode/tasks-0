// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../util/riverpod_providers/wml_colors/wml_colors.dart';
import 'overlay_zero_riverpod_provider.dart';

class OverlayZeroWidget extends ConsumerStatefulWidget {
  OverlayZeroWidget({super.key, this.isPresent});

  bool? isPresent;

  @override
  ConsumerState<OverlayZeroWidget> createState() => OverlayZeroState();
}

class OverlayZeroState extends ConsumerState<OverlayZeroWidget> {
  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(OverlayZeroRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);

    if (widget.isPresent ?? rp.isPresent) {
      return Opacity(
        opacity: .8,
        child: Container(
          height: MediaQuery.of(context).size.height,
          width: MediaQuery.of(context).size.width,
          color: wmlColors.primaryColor,
          child: Center(
            child: CircularProgressIndicator(
              color: wmlColors.tertiaryColor,
            ),
          ),
        ),
      );
    } else {
      return Container();
    }

  }
}
