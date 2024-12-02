// ignore_for_file: non_constant_identifier_names, constant_identifier_names

import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';


import '../riverpod_providers/wml_nav/wml_nav.dart';
import 'nav_observer.dart';

GoRouter createGoRouter(WidgetRef ref) {
  return GoRouter(
      observers: [MyNavigatorObserver(ref)],
      initialLocation: WMLNavRiverpodProviderInstance.route["path"],
      navigatorKey: GlobalKey<NavigatorState>(),
      routes: [
        // Example route you need a layout and a page
        GoRoute(
            path: WMLNavRiverpodProviderInstance.route["path"],
            builder: (context, state) {
              return WMLNavRiverpodProviderInstance.route["widget"];
            }),
        //
      ]);
}

late GoRouter goRouter;
