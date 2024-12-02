// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:[PROJECT_NAME]/shared/widgets/overlay_zero/overlay_zero_riverpod_provider.dart';
import 'package:[PROJECT_NAME]/util/env/env.dart';
import 'package:[PROJECT_NAME]/util/wml/base/base_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:webview_flutter/webview_flutter.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';

import 'hcaptcha_zero_widget_riverpod_provider.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class HcaptchaZeroWidget extends ConsumerStatefulWidget {
  final Function(String)? onCaptchaResult;

  HcaptchaZeroWidget({super.key, this.onCaptchaResult});

  @override
  ConsumerState<HcaptchaZeroWidget> createState() => _HcaptchaZeroState();
}

class _HcaptchaZeroState extends ConsumerState<HcaptchaZeroWidget> with WMLBaseWidget {
  late WebViewController controller;

  @override
  void initState() {
    super.initState();

    initWebView();
  }

  void initWebView() {
    controller = WebViewController()
      ..setJavaScriptMode(JavaScriptMode.unrestricted)
      ..setBackgroundColor(WMLColorsRiverpodProviderInstance.primaryColor)
      ..loadHtmlString(hCaptchaHtml())
      ..addJavaScriptChannel('CAPTCHA', onMessageReceived: (message) {
        if ([
          "rate-limited",
          "network-error",
          "invalid-data",
          "challenge-error",
          "challenge-closed",
          "challenge-expired",
          "missing-captcha",
          "invalid-captcha-id",
          "internal-error"
        ].contains(message.message)) {
          showSnackBarMsg(translate("global.systemError"));
        } else if (["success"].contains(message.message)) {
          widget.onCaptchaResult?.call(message.message);
        } else {
          showSnackBarMsg(translate("HcaptchaZero.WMLNotifyOne.failed"));
        }
      });
  }

  String hCaptchaHtml() {
    String siteKey = APPENV.hcaptcha["siteKey"]!;
    String host = APPENV.frontendURI0.uri.host;
    if (APPENV.type.isDevEnvironment) {
      host = "[PROJECT_NAME].com";
    }
    return '''
<!DOCTYPE html>
<html lang="en">
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<script src="https://js.hcaptcha.com/1/api.js?onload=onLoadCallback&render=explicit&host=$host" async defer></script>
</head>
<body>
  <div id="hcaptcha" data-theme=dark ></div>
  <br/>
<script>
  var onLoadCallback = function() {
    const options = {
      sitekey: "$siteKey",
      callback: (res) => {
        window.CAPTCHA.postMessage('success');
      },
      'expired-callback': () => {
        window.CAPTCHA.postMessage('expired');
      },
      'error-callback': (err) => {
        window.CAPTCHA.postMessage(err);
      },
    };
    hcaptcha.render('hcaptcha', options);
  };
</script>
</body>
</html>
''';
  }

  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(HcaptchaZeroRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);

    return SizedBox(
      width: MediaQuery.of(context).size.width,
      height: MediaQuery.of(context).size.height, // Set the height as needed
      child: Scaffold(appBar: AppBar(), body: WebViewWidget(controller: controller)), // Display the WebView
    );
  }
}
