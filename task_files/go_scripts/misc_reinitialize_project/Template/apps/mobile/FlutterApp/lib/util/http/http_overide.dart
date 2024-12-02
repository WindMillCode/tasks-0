import 'dart:io';

import 'package:tuli/util/env/env.dart';

class MyHttpOverrides extends HttpOverrides {
  @override
  HttpClient createHttpClient(SecurityContext? context) {
    if (APPENV.type.isDeployedEnvironment == false) {
      return super.createHttpClient(context)
        ..badCertificateCallback = (X509Certificate cert, String host, int port) => true;
    }
    return super.createHttpClient(context);
  }
}
