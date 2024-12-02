// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'cancel_subscription.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

CancelSubscriptionRequestBody _$CancelSubscriptionRequestBodyFromJson(
        Map<String, dynamic> json) =>
    CancelSubscriptionRequestBody(
      accessToken: json['access_token'] as String,
      subscriptionContractInfo: (json['subscription_contract_info']
                  as List<dynamic>?)
              ?.map((e) => CancelSubscriptionRequestBodySubscriptionContractInfo
                  .fromJson(e as Map<String, dynamic>))
              .toList() ??
          const [],
    );

Map<String, dynamic> _$CancelSubscriptionRequestBodyToJson(
        CancelSubscriptionRequestBody instance) =>
    <String, dynamic>{
      'access_token': instance.accessToken,
      'subscription_contract_info': instance.subscriptionContractInfo,
    };

CancelSubscriptionRequestBodySubscriptionContractInfo
    _$CancelSubscriptionRequestBodySubscriptionContractInfoFromJson(
            Map<String, dynamic> json) =>
        CancelSubscriptionRequestBodySubscriptionContractInfo(
          id: json['id'] as String? ?? '',
          confirmationNumber: json['confirmation_number'] as String? ?? '',
        );

Map<String, dynamic>
    _$CancelSubscriptionRequestBodySubscriptionContractInfoToJson(
            CancelSubscriptionRequestBodySubscriptionContractInfo instance) =>
        <String, dynamic>{
          'id': instance.id,
          'confirmation_number': instance.confirmationNumber,
        };

CancelSubscriptionResponseBody _$CancelSubscriptionResponseBodyFromJson(
        Map<String, dynamic> json) =>
    CancelSubscriptionResponseBody();

Map<String, dynamic> _$CancelSubscriptionResponseBodyToJson(
        CancelSubscriptionResponseBody instance) =>
    <String, dynamic>{};
