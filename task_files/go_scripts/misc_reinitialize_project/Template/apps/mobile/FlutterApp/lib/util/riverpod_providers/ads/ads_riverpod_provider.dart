// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'dart:async';
import 'dart:io';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_mobile_ads/google_mobile_ads.dart';
import 'package:[PROJECT_NAME]/util/env/env.dart';

class AdsRiverpodProviderValue {
  int widgetCtnrSize = 18;
  int bannerAdCtrnSize = 1;

  late String mainBannerID;
  late String openAppID;

  AdsRiverpodProviderValue() {
    if (APPENV.type.isDeployedEnvironment) {
      mainBannerID = {
        AppEnvPlatformType.ANDROID: "[FLUTTER_ANDRIOID_GOOGLE_ADS_ID]",
        AppEnvPlatformType.IOS: "[FLUTTER_IOS_GOOGLE_ADS_ID]"
      }[APPENV.platformType] as dynamic;

      openAppID = {
        AppEnvPlatformType.ANDROID: "[FLUTTER_ANDRIOID_GOOGLE_ADS_ID]",
        AppEnvPlatformType.IOS: "[FLUTTER_IOS_GOOGLE_ADS_ID]"
      }[APPENV.platformType] as dynamic;
    } else {
      openAppID = mainBannerID = getFakeAdID();
    }
  }

  static String getFakeAdID([bool? isAndroid]) {
    isAndroid ??= Platform.isAndroid;
    return isAndroid ? 'ca-app-pub-3940256099942544/5224354917' : 'ca-app-pub-3940256099942544/1712485313';
  }

  Future<BannerAd> loadBannerAd() async {
    Future<BannerAd> customFuture() async {
      final Completer<BannerAd> completer = Completer<BannerAd>();

      await BannerAd(
        adUnitId: mainBannerID,
        request: const AdRequest(),
        size: AdSize.largeBanner,
        listener: BannerAdListener(
          onAdLoaded: (ad) {
            completer.complete(ad as BannerAd);
          },
          onAdFailedToLoad: (ad, err) {
            ad.dispose();
          },
          onAdOpened: (Ad ad) {},
          onAdClosed: (Ad ad) {},
          onAdImpression: (Ad ad) {},
        ),
      ).load();
      return completer.future;
    }

    return await customFuture();
  }

  Future<AppOpenAd> loadOpenAppAd() async {
    Future<AppOpenAd> customFuture() {
      final Completer<AppOpenAd> completer = Completer<AppOpenAd>();

      AppOpenAd.load(
          // TODO figure out equivalent iisue
          // orientation: AppOpenAd.orientationPortrait,
          adUnitId: openAppID,
          request: const AdRequest(),
          adLoadCallback: AppOpenAdLoadCallback(
              onAdLoaded: (ad) {
                ad.fullScreenContentCallback = FullScreenContentCallback(
                    onAdShowedFullScreenContent: (ad) {},
                    onAdImpression: (ad) {},
                    onAdFailedToShowFullScreenContent: (ad, err) {},
                    onAdDismissedFullScreenContent: (ad) {},
                    onAdClicked: (ad) {});

                completer.complete(ad);
              },
              onAdFailedToLoad: (LoadAdError error) {}));

      return completer.future; // Return the Future
    }

    return await customFuture();
  }

  AdsRiverpodProviderValue copyWith() {
    return AdsRiverpodProviderValue();
  }
}

var AdsRiverpodProviderInstance = AdsRiverpodProviderValue();

class AdsRiverpodNotifier extends Notifier<AdsRiverpodProviderValue> {
  @override
  AdsRiverpodProviderValue build() {
    return AdsRiverpodProviderInstance;
  }
}

final AdsRiverpodProvider = NotifierProvider<AdsRiverpodNotifier, AdsRiverpodProviderValue>(() {
  return AdsRiverpodNotifier();
});
