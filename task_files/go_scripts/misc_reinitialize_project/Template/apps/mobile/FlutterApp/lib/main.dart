// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'dart:io';
import 'package:tuli/shared/widgets/overlay_zero/overlay_zero.dart';
import 'package:tuli/util/env/env.dart';
import 'package:tuli/util/env/deep_links.dart';
import 'package:tuli/util/http/http_overide.dart';
import 'package:tuli/util/misc/notify/my_awesome_notifications.dart';
import 'package:tuli/util/riverpod_providers/i18n/i18n_riverpod_provider.dart';
import 'package:tuli/util/riverpod_providers/socketio/socketio_riverpod_provider.dart';
import 'package:tuli/util/wml/base/base_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_smart_dialog/flutter_smart_dialog.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:media_kit/media_kit.dart';
import 'package:sentry_flutter/sentry_flutter.dart';
import 'util/go_router/my_router.dart';
import 'util/riverpod_providers/wml_theme/wml_theme.dart';
import 'package:intl/date_symbol_data_local.dart';


void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  initDeepLinks();
  MediaKit.ensureInitialized();
  // MobileAds.instance.initialize();
  // var advertisingId = await AdvertisingId.id(true);
  HttpOverrides.global = MyHttpOverrides();
  initializeDateFormatting();
  // initAwesomeNotifications();

  FlutterTranslate.initialize(
      FlutterTranslateOptions(fallback: 'en', supported: I18NRiverpodProviderInstance.supportedLocales));
  if (APPENV.type.isDeployedEnvironment) {
    await SentryFlutter.init(
      (options) {
        const tracesSampleRate = {
          AppEnvType.PROD: 0.8,
          AppEnvType.PREVIEW: 1.0,
          AppEnvType.TEST: 0.0,
          AppEnvType.DEV: 1.0,
          AppEnvType.DOCKER_DEV: 0.0,
        };
        options.dsn = const String.fromEnvironment("SENTRY_DSN");
        options.environment = "Flutter_${APPENV.type.value}";
        options.tracesSampleRate = tracesSampleRate[APPENV.type];
      },
      appRunner: () => runApp(ProviderScope(child: LocalizedApp(MainApp()))),
    );
  } else {
    runApp(ProviderScope(child: LocalizedApp(MainApp())));
    runApp(app)
  }
}

class MainApp extends ConsumerStatefulWidget {
  MainApp({
    super.key,
    @visibleForTesting this.mockChild,
  });

  Widget? mockChild;
  @override
  ConsumerState<MainApp> createState() => MainAppState();
}

class MainAppState extends ConsumerState<MainApp> with WidgetsBindingObserver, WMLBaseWidget, TickerProviderStateMixin {
  @override
  void initState() {
    setAwesomeNotificationListeners();
    super.initState();

    WidgetsBinding.instance.addObserver(this);
    goRouter = createGoRouter(ref);

    // TODO going good so far use in other projects
    // WidgetsBinding.instance.addTimingsCallback((timings) {
    //   for (FrameTiming frameTiming in timings) {
    //     final buildTime = frameTiming.buildDuration.inMilliseconds;
    //     final rasterTime = frameTiming.rasterDuration.inMilliseconds;

    //     CAN_HANDLE_ANIMATIONS = buildTime > 16.67 || rasterTime > 16.67;
    //   }
    // });
  }

  bool isInit = false;

  @override
  void didChangeAppLifecycleState(AppLifecycleState state) {
    super.didChangeAppLifecycleState(state);
    if (state == AppLifecycleState.resumed) {
      // The app is resumed (user returned to the app)
    }
  }

  @override
  void dispose() {
    WidgetsBinding.instance.removeObserver(this);
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // final ads = ref.watch(AdsRiverpodProvider);
    // ads.loadOpenAppAd();

    final wmlTheme = ref.watch(WMLThemeRiverpodProvider);
    final socketio = ref.watch(SocketioRiverpodProvider);

    return widget.mockChild ??
        MaterialApp(
          debugShowCheckedModeBanner: false,
          theme: wmlTheme.lightTheme,
          darkTheme: wmlTheme.darkTheme,
          themeMode: wmlTheme.currentThemeMode,
          navigatorObservers: [FlutterSmartDialog.observer],
          builder: FlutterSmartDialog.init(builder: (context, child) {
            if (!isInit) {
              isInit = true;
              initSocketIOConnection(context, ref);

            }
            return Stack(
              children: [
                child!,
                OverlayZeroWidget(),
              ],
            );
          }),
          home: Router(
            routerDelegate: goRouter.routerDelegate,
            routeInformationParser: goRouter.routeInformationParser,
            routeInformationProvider: goRouter.routeInformationProvider,
            backButtonDispatcher: goRouter.backButtonDispatcher,
          ),
        );
  }
}
