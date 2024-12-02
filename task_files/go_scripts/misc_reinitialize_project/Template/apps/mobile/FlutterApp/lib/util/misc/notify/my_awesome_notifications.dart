// ignore_for_file: constant_identifier_names

import 'dart:async';

import 'package:awesome_notifications/awesome_notifications.dart';
import 'package:tuli/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:open_filex/open_filex.dart';

enum AwesomeNotifyMsgType { VIEW_EXPORT_DATA }

initAwesomeNotifications() {
  AwesomeNotifications().initialize(
      'resource://mipmap/notify_icon',
      [
        NotificationChannel(
            channelGroupKey: 'basic_channel_group',
            channelKey: 'basic_channel',
            channelName: 'Basic notifications',
            channelDescription: 'Notification channel for basic tests',
            defaultColor: WMLColorsRiverpodProviderInstance.notifyColor0,
            ledColor: WMLColorsRiverpodProviderInstance.notifyLedColor0)
      ],
      channelGroups: [
        NotificationChannelGroup(channelGroupKey: 'basic_channel_group', channelGroupName: 'Basic group')
      ],
      debug: true);
}

Future<void> notifyHandlerOnActionReceivedMethod(ReceivedAction action) async {
  Map<String, String?>? payload = action.payload;
  String? type = payload!["type"] ?? "";
  if (type == AwesomeNotifyMsgType.VIEW_EXPORT_DATA.name) {
    await OpenFilex.open(payload["filePath"]!);
  }
}

setAwesomeNotificationListeners() {
  AwesomeNotifications().setListeners(
    onActionReceivedMethod: notifyHandlerOnActionReceivedMethod,
  );
}

requestPermissionToSendNotifications() {
  AwesomeNotifications().isNotificationAllowed().then((isAllowed) {
    if (!isAllowed) {
      AwesomeNotifications().requestPermissionToSendNotifications();
    }
  });
}
