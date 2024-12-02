// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'use_credits.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

UseCreditsRequestBody _$UseCreditsRequestBodyFromJson(
        Map<String, dynamic> json) =>
    UseCreditsRequestBody(
      creditType: json['credit_type'] as String,
      amount: json['amount'] as int,
      accessToken: json['access_token'] as String,
    );

Map<String, dynamic> _$UseCreditsRequestBodyToJson(
        UseCreditsRequestBody instance) =>
    <String, dynamic>{
      'credit_type': instance.creditType,
      'amount': instance.amount,
      'access_token': instance.accessToken,
    };

UseCreditsResponseBody _$UseCreditsResponseBodyFromJson(
        Map<String, dynamic> json) =>
    UseCreditsResponseBody();

Map<String, dynamic> _$UseCreditsResponseBodyToJson(
        UseCreditsResponseBody instance) =>
    <String, dynamic>{};
