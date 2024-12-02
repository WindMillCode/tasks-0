import 'package:app_links/app_links.dart';
import 'package:tuli/shared/widgets/overlay_zero/overlay_zero_riverpod_provider.dart';
import 'package:tuli/util/go_router/my_router.dart';
import 'package:tuli/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:sentry_flutter/sentry_flutter.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';

final appLinks = AppLinks();

void initDeepLinks() {
  appLinks.uriLinkStream.listen((link) {
    handleLink(link);
  }, onError: (err) {
    // Handle exception
  });
}

void handleLink(Uri uri) {
  Sentry.captureMessage("Deep link: ${uri.path}");
  if ([WMLNavRiverpodProviderInstance.route["path"]].contains(uri.path)) {
    // PERFORM ACTION
  }


}
