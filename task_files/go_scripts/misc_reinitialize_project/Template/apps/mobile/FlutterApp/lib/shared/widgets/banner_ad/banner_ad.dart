// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_mobile_ads/google_mobile_ads.dart';

import '../../../util/riverpod_providers/ads/ads_riverpod_provider.dart';
import '../../../util/riverpod_providers/wml_colors/wml_colors.dart';

class BannerAdWidget extends ConsumerStatefulWidget {
  const BannerAdWidget({super.key});

  @override
  ConsumerState<BannerAdWidget> createState() => BannerAdState();
}

class BannerAdState extends ConsumerState<BannerAdWidget> {
  BannerAd? bannerAd;

  @override
  void initState() {
    super.initState();
    Future.delayed(Duration.zero, () async {
      final ad = await ref.read(AdsRiverpodProvider).loadBannerAd();

      setState(() {
        bannerAd = ad;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    if (bannerAd != null) {
      return Row(
        children: [
          AdWidget(ad: bannerAd!),
          Align(
              alignment: Alignment.topRight,
              child: IconButton(
                icon: Icon(
                  CupertinoIcons.xmark,
                  weight: 1000,
                ),
                color: wmlColors.white,
                onPressed: () {
                  bannerAd!.dispose();
                  setState(() {
                    bannerAd = null;
                  });
                },
              ))
        ],
      );
    } else {
      return Container();
    }
  }
}
