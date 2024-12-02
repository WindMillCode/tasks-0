// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';


class MyNavigatorObserver extends NavigatorObserver {
  final WidgetRef ref;
  MyNavigatorObserver(this.ref);

  @override
  void didPush(Route<dynamic> route, Route<dynamic>? previousRoute) {
    super.didPush(route, previousRoute);

    // Handle route changes here
    final currentLocation = route.settings.name;
    Future.delayed(Duration.zero, () {
      if (previousRoute != null) {

      }
    });
  }



  @override
  void didPop(Route route, Route? previousRoute) {

  }
}
