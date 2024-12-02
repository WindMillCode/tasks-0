// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'purchase_subscription.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

PurchaseSubscriptionRequestBody _$PurchaseSubscriptionRequestBodyFromJson(
        Map<String, dynamic> json) =>
    PurchaseSubscriptionRequestBody(
      accessToken: json['access_token'] as String,
      subscriptionContractIds:
          (json['subscription_contract_ids'] as List<dynamic>?)
                  ?.map((e) => e as String)
                  .toList() ??
              const [],
      items: (json['items'] as List<dynamic>?)
              ?.map((e) => PurchaseSubscriptionRequestBodyItems.fromJson(
                  e as Map<String, dynamic>))
              .toList() ??
          const [],
      redirectUrl: json['redirect_url'] as String? ?? '',
    )..appPlatformTypeEnumValue = $enumDecode(
        _$AppEnvPlatformTypeEnumMap, json['app_platform_type_enum_value']);

Map<String, dynamic> _$PurchaseSubscriptionRequestBodyToJson(
        PurchaseSubscriptionRequestBody instance) =>
    <String, dynamic>{
      'access_token': instance.accessToken,
      'subscription_contract_ids': instance.subscriptionContractIds,
      'items': instance.items,
      'redirect_url': instance.redirectUrl,
      'app_platform_type_enum_value':
          _$AppEnvPlatformTypeEnumMap[instance.appPlatformTypeEnumValue]!,
    };

const _$AppEnvPlatformTypeEnumMap = {
  AppEnvPlatformType.WEB: 'WEB',
  AppEnvPlatformType.ANDROID: 'ANDROID',
  AppEnvPlatformType.IOS: 'IOS',
  AppEnvPlatformType.LINUX: 'LINUX',
  AppEnvPlatformType.MACOS: 'MACOS',
  AppEnvPlatformType.WINDOWS: 'WINDOWS',
};

PurchaseSubscriptionRequestBodyItems
    _$PurchaseSubscriptionRequestBodyItemsFromJson(Map<String, dynamic> json) =>
        PurchaseSubscriptionRequestBodyItems(
          quantity: json['quantity'] as int? ?? 1,
          variantId: json['variant_id'] as String,
          sellingPlanId: json['selling_plan_id'] as String,
          iosInfo: json['ios_info'] as Map<String, dynamic>?,
        );

Map<String, dynamic> _$PurchaseSubscriptionRequestBodyItemsToJson(
        PurchaseSubscriptionRequestBodyItems instance) =>
    <String, dynamic>{
      'quantity': instance.quantity,
      'variant_id': instance.variantId,
      'selling_plan_id': instance.sellingPlanId,
      'ios_info': instance.iosInfo,
    };

PurchaseSubscriptionResponseBody _$PurchaseSubscriptionResponseBodyFromJson(
        Map<String, dynamic> json) =>
    PurchaseSubscriptionResponseBody()
      ..id = json['id'] as String?
      ..weburl = json['weburl'] as String;

Map<String, dynamic> _$PurchaseSubscriptionResponseBodyToJson(
        PurchaseSubscriptionResponseBody instance) =>
    <String, dynamic>{
      'id': instance.id,
      'weburl': instance.weburl,
    };
