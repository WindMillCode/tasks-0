// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'bullet_point_zero_widget_riverpod_provider.dart';
import 'package:tuli/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:tuli/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:tuli/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:tuli/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class BulletPointZeroWidget extends ConsumerStatefulWidget {
  BulletPointZeroWidget({super.key, this.child = const Text("Bullet Point")});
  dynamic child;

  @override
  ConsumerState<BulletPointZeroWidget> createState() => _BulletPointZeroState();
}

class _BulletPointZeroState extends ConsumerState<BulletPointZeroWidget> {
  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(BulletPointZeroRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final contextHeight = MediaQuery.of(context).size.height;
    final contextWidth = MediaQuery.of(context).size.width;
    TextSpan displayChild;
    if (widget.child is String) {
      displayChild = TextSpan(text: widget.child);
    } else {
      displayChild = widget.child;
    }

    return RichText(
      text: TextSpan(
        children: <TextSpan>[
          TextSpan(
              text: '• ',
              style: TextStyle(fontSize: wmlFonts.heading4)), // Bullet symbol
          displayChild
        ],
      ),
    );
  }
}
