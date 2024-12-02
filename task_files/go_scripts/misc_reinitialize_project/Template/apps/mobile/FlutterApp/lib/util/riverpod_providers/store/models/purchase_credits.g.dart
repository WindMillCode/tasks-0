// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'purchase_credits.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

PurchaseCreditsRequestBody _$PurchaseCreditsRequestBodyFromJson(
        Map<String, dynamic> json) =>
    PurchaseCreditsRequestBody(
      accessToken: json['access_token'] as String,
      items: (json['items'] as List<dynamic>?)
              ?.map((e) => PurchaseCreditsRequestBodyItems.fromJson(
                  e as Map<String, dynamic>))
              .toList() ??
          const [],
    )..appPlatformTypeEnumValue = $enumDecode(
        _$AppEnvPlatformTypeEnumMap, json['app_platform_type_enum_value']);

Map<String, dynamic> _$PurchaseCreditsRequestBodyToJson(
        PurchaseCreditsRequestBody instance) =>
    <String, dynamic>{
      'access_token': instance.accessToken,
      'items': instance.items,
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

PurchaseCreditsRequestBodyItems _$PurchaseCreditsRequestBodyItemsFromJson(
        Map<String, dynamic> json) =>
    PurchaseCreditsRequestBodyItems(
      quantity: json['quantity'] as int? ?? 0,
      variantId: json['variant_id'] as String? ?? '',
      iosInfo: json['ios_info'] as Map<String, dynamic>?,
    );

Map<String, dynamic> _$PurchaseCreditsRequestBodyItemsToJson(
        PurchaseCreditsRequestBodyItems instance) =>
    <String, dynamic>{
      'quantity': instance.quantity,
      'variant_id': instance.variantId,
      'ios_info': instance.iosInfo,
    };

PurchaseCreditsResponseBody _$PurchaseCreditsResponseBodyFromJson(
        Map<String, dynamic> json) =>
    PurchaseCreditsResponseBody()..weburl = json['weburl'] as String;

Map<String, dynamic> _$PurchaseCreditsResponseBodyToJson(
        PurchaseCreditsResponseBody instance) =>
    <String, dynamic>{
      'weburl': instance.weburl,
    };
