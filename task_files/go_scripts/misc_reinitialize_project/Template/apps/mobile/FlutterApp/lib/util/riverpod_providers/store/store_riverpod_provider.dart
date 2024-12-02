// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'dart:convert';
import 'dart:io';
import 'package:[PROJECT_NAME]/util/riverpod_providers/store/models/cancel_subscription.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/store/models/list_credits.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/store/models/list_subscriptions.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/store/models/purchase_credits.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/store/models/purchase_subscription.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/store/models/use_credits.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:[PROJECT_NAME]/util/env/env.dart';

class StoreRiverpodProviderValue {
  Future<ListSubscriptionsResponseBody> listSubscriptions({
    Uri? uri,
    AppEnv? env,
    http.Client? client,
    required ListSubscriptionsRequestBody reqBody,
  }) async {
    env ??= APPENV;
    uri ??= Uri.parse(APPENV.endpoints.storeRiverpodProvider['listSubscriptions']!.url());
    client ??= http.Client();

    var apiReqBody = jsonEncode({"data": reqBody.toJson()});

    final response = await client.post(
      uri,
      headers: {"Content-Type": "application/json"},
      body: apiReqBody,
    );

    if ([HttpStatus.ok].contains(response.statusCode)) {
      return ListSubscriptionsResponseBody.fromJson(jsonDecode(response.body)["data"]);
    } else {
      throw Exception('Failed to list subscriptions with status code: ${response.statusCode}');
    }
  }

  Future<PurchaseSubscriptionResponseBody> purchaseSubscription({
    Uri? uri,
    AppEnv? env,
    http.Client? client,
    required PurchaseSubscriptionRequestBody reqBody,
  }) async {
    env ??= APPENV;
    uri ??= Uri.parse(APPENV.endpoints.storeRiverpodProvider['purchaseSubscription']!.url());
    client ??= http.Client();

    var apiReqBody = jsonEncode({"data": reqBody.toJson()});

    final response = await client.post(
      uri,
      headers: {"Content-Type": "application/json"},
      body: apiReqBody,
    );

    if ([HttpStatus.ok].contains(response.statusCode)) {
      return PurchaseSubscriptionResponseBody.fromJson(jsonDecode(response.body)["data"]);
    } else {
      throw Exception('Failed to purchase subscription with status code: ${response.statusCode}');
    }
  }

  Future<CancelSubscriptionResponseBody> cancelSubscription({
    Uri? uri,
    AppEnv? env,
    http.Client? client,
    required CancelSubscriptionRequestBody reqBody,
  }) async {
    env ??= APPENV;
    uri ??= Uri.parse(APPENV.endpoints.storeRiverpodProvider['cancelSubscription']!.url());
    client ??= http.Client();

    var apiReqBody = jsonEncode({"data": reqBody.toJson()});

    final response = await client.delete(
      uri,
      headers: {"Content-Type": "application/json"},
      body: apiReqBody,
    );

    if ([HttpStatus.ok].contains(response.statusCode)) {
      return CancelSubscriptionResponseBody();
    } else {
      throw Exception('Failed to cancel subscription with status code: ${response.statusCode}');
    }
  }

  Future<PurchaseCreditsResponseBody> purchaseCredits({
    Uri? uri,
    AppEnv? env,
    http.Client? client,
    required PurchaseCreditsRequestBody reqBody,
  }) async {
    env ??= APPENV;
    uri ??= Uri.parse(APPENV.endpoints.storeRiverpodProvider['purchaseCredits']!.url());
    client ??= http.Client();

    var apiReqBody = jsonEncode({"data": reqBody.toJson()});

    final response = await client.post(
      uri,
      headers: {"Content-Type": "application/json"},
      body: apiReqBody,
    );

    if ([HttpStatus.ok].contains(response.statusCode)) {
      return PurchaseCreditsResponseBody.fromJson(jsonDecode(response.body)["data"]);
    } else {
      throw Exception('Failed to purchase credits with status code: ${response.statusCode}');
    }
  }

  Future<ListCreditsResponseBody> listCredits({
    Uri? uri,
    AppEnv? env,
    http.Client? client,
    ListCreditsRequestBody? reqBody,
  }) async {
    env ??= APPENV;
    uri ??= Uri.parse(APPENV.endpoints.storeRiverpodProvider['listCredits']!.url());
    client ??= http.Client();
    reqBody ??= ListCreditsRequestBody();

    var apiReqBody = jsonEncode({"data": reqBody.toJson()});

    final response = await client.post(
      uri,
      headers: {"Content-Type": "application/json"},
      body: apiReqBody,
    );

    if ([HttpStatus.ok].contains(response.statusCode)) {
      var resp = ListCreditsResponseBody.fromJson(jsonDecode(response.body)["data"]);
      return resp;
    } else {
      throw Exception('Failed to fetch list credits with status code: ${response.statusCode}');
    }
  }

  Future<void> useCredits({
    Uri? uri,
    AppEnv? env,
    http.Client? client,
    required UseCreditsRequestBody reqBody,
  }) async {
    env ??= APPENV;
    uri ??= Uri.parse(APPENV.endpoints.storeRiverpodProvider['useCredits']!.url());
    client ??= http.Client();

    var apiReqBody = jsonEncode({"data": reqBody.toJson()});

    final response = await client.post(
      uri,
      headers: {"Content-Type": "application/json"},
      body: apiReqBody,
    );

    if ([HttpStatus.ok].contains(response.statusCode)) {
      return;
    } else {
      throw Exception('Failed to use credits with status code: ${response.statusCode}');
    }
  }

  StoreRiverpodProviderValue copyWith() {
    return StoreRiverpodProviderValue();
  }
}

var StoreRiverpodProviderInstance = StoreRiverpodProviderValue();

class StoreRiverpodNotifier extends Notifier<StoreRiverpodProviderValue> {
  @override
  StoreRiverpodProviderValue build() {
    return StoreRiverpodProviderInstance;
  }
}

final StoreRiverpodProvider = NotifierProvider<StoreRiverpodNotifier, StoreRiverpodProviderValue>(() {
  return StoreRiverpodNotifier();
});
