// ignore_for_file: non_constant_identifier_names, constant_identifier_names

import 'dart:io';
import 'package:tuli/util/wml/url.dart';
import 'package:flutter/foundation.dart';

enum AppEnvType { DEV, DOCKER_DEV, PREVIEW, PROD, TEST }

extension AppEnvTypeExtension on AppEnvType {
  String get value {
    return toString().split('.').last;
  }

  static AppEnvType fromString(String value) {
    return AppEnvType.values.firstWhere(
      (type) => type.value == value,
      orElse: () => throw ArgumentError('Invalid enum value: $value'),
    );
  }

  bool get isDevEnvironment => [AppEnvType.DEV, AppEnvType.DOCKER_DEV].contains(this);

  bool get isTestEnvironment => this == AppEnvType.TEST;

  bool get isPreviewEnvironment => this == AppEnvType.PREVIEW;

  bool get isProdEnvironment => this == AppEnvType.PROD;

  bool get isDeployedEnvironment => [AppEnvType.PROD, AppEnvType.PREVIEW].contains(this);
}

enum AppEnvPlatformType { WEB, ANDROID, IOS, LINUX, MACOS, WINDOWS }

extension AppEnvPlatformTypeExtension on AppEnvPlatformType {
  static AppEnvPlatformType getPlatform() {
    if (kIsWeb) {
      return AppEnvPlatformType.WEB;
    } else if (Platform.isAndroid) {
      return AppEnvPlatformType.ANDROID;
    } else if (Platform.isIOS) {
      return AppEnvPlatformType.IOS;
    } else if (Platform.isLinux) {
      return AppEnvPlatformType.LINUX;
    } else if (Platform.isMacOS) {
      return AppEnvPlatformType.MACOS;
    } else if (Platform.isWindows) {
      return AppEnvPlatformType.WINDOWS;
    } else {
      // Handle an unexpected platform; throw an error or default to a specific case
      throw Exception("Platform not recognized");
    }
  }
}

class AppEnvEndpoints {
  late final Map<String, WMLEndpoint> accountsRiverpodProvider;
  late final Map<String, WMLEndpoint> blogRiverpodProvider;
  late final Map<String, WMLEndpoint> storeRiverpodProvider;
  late final Map<String, WMLEndpoint> eventRiverpodProvider;
  late final Map<String, WMLEndpoint> contactRiverpodProvider;
  late final Map<String, WMLEndpoint> labsRiverpodProvider;
  late final Map<String, WMLEndpoint> subscriptionsRiverpodProvider;
  late final Map<String, WMLEndpoint> platformsRiverpodProvider;
  late final Map<String, WMLEndpoint> jobsRiverpodProvider;
  late final Map<String, WMLEndpoint> storageRiverpodProvider;

  AppEnvEndpoints({required backendURI0}) {
    String httpsBackendFQDN0 = backendURI0.fqdn;
  }
}

class AppEnv {
  WMLUri backendURI0 = WMLUri(
    scheme: 'https',
    host: '10.0.2.2',
    port: 10073,
  );
  WMLUri frontendURI0 = WMLUri(
    scheme: 'https',
    host: 'example.com',
    port: 10003,
  );
  WMLUri firebaseURI0 = WMLUri(scheme: 'http', host: 'localhost');

  AppEnvPlatformType platformType = AppEnvPlatformTypeExtension.getPlatform();
  AppEnvType type = AppEnvType.DEV;
  final apiMsgs = {"SUCCESS": 0, "FAILURE": 1};
  late Map<String, dynamic> firebase;
  late AppEnvEndpoints endpoints;
  late String hostComputer;
  List<String> hostRemoteIPS = ["localhost"];

  AppEnv() {
    String typeString = const String.fromEnvironment("FLUTTER_MOBILE_ENV");
    hostComputer = const String.fromEnvironment("HOST_COMPUTER");
    type = AppEnvTypeExtension.fromString(typeString);
    // LOCAL | REMOTE

    if (platformType != AppEnvPlatformType.ANDROID) {
      backendURI0.uri = backendURI0.uri.replace(host: "example.com", port: 10073);
    }
    if (platformType == AppEnvPlatformType.IOS) {
      app["authContainerHeight"] = .88;
    }
    if (hostComputer == "REMOTE" && type.isDevEnvironment == true) {
      hostRemoteIPS = const String.fromEnvironment("HOST_REMOTE_IPS").split(",");
      firebaseURI0.uri = firebaseURI0.uri.replace(host: hostRemoteIPS[0]);
      frontendURI0.uri = frontendURI0.uri.replace(host: hostRemoteIPS[0]);
      backendURI0.uri = backendURI0.uri.replace(host: "0xdaf8738d7d6dca8774f2f2742cbd7be913e9c3eb.diode.link", port: 443);
    }
    // must be like this
    firebase = {
      'storageImageUrl': '${firebaseURI0.uri.scheme}://${firebaseURI0.uri.host}:9199/v0/b/default.appspot.com/o/'
    };
    if (type == AppEnvType.PREVIEW) {
      backendURI0.uri = backendURI0.uri.replace(host: "api.preview.tuli.com", port: 443);
      frontendURI0.uri = frontendURI0.uri.replace(host: "ui.preview.tuli.com", port: 443);
      firebase['storageImageUrl'] = 'https://firebasestorage.googleapis.com/v0/b/tuli-preview.appspot.com/o/';
    }
    if (type == AppEnvType.PROD) {
      backendURI0.uri = backendURI0.uri.replace(host: "api.tuli.com", port: 443);
      frontendURI0.uri = frontendURI0.uri.replace(host: "tuli.com", port: 443);
      firebase['storageImageUrl'] = 'https://firebasestorage.googleapis.com/v0/b/tuli.appspot.com/o/';
    }

    endpoints = AppEnvEndpoints(backendURI0: backendURI0.fqdn);
  }

  String getImageUrlFromFirebaseStorage(String? resourcePath) {
    var storageImgUrl = firebase['storageImageUrl'];
    return '$storageImgUrl${Uri.encodeComponent(resourcePath ?? '')}?alt=media';
  }

  createAPIMSG({code, data = ""}) {
    code ??= APPENV.apiMsgs["SUCCESS"];
    return {"data": data, "code": code};
  }

  createAPIServerError({data = "", msg = "There was an error while the system was processing your request"}) {
    return {"data": data, "msg": msg};
  }

  var app = {
    // dev additions
  };

  // dev additions
  var hcaptcha = {"siteKey": "7dee141b-7877-4153-bef2-b4e04ef037e1"};
}

var APPENV = AppEnv();
