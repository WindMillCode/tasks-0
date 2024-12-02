// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'list_credits.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ListCreditsRequestBody _$ListCreditsRequestBodyFromJson(
        Map<String, dynamic> json) =>
    ListCreditsRequestBody();

Map<String, dynamic> _$ListCreditsRequestBodyToJson(
        ListCreditsRequestBody instance) =>
    <String, dynamic>{};

ListCreditsResponseBody _$ListCreditsResponseBodyFromJson(
        Map<String, dynamic> json) =>
    ListCreditsResponseBody()
      ..credits = (json['credits'] as List<dynamic>)
          .map((e) => ListCreditsResponseBodyCredits.fromJson(
              e as Map<String, dynamic>))
          .toList();

Map<String, dynamic> _$ListCreditsResponseBodyToJson(
        ListCreditsResponseBody instance) =>
    <String, dynamic>{
      'credits': instance.credits,
    };

ListCreditsResponseBodyCredits _$ListCreditsResponseBodyCreditsFromJson(
        Map<String, dynamic> json) =>
    ListCreditsResponseBodyCredits(
      title: json['title'] as String,
      variantId: json['variant_id'] as String,
      price: ListCreditsResponseBodyCreditsPrice.fromJson(
          json['price'] as Map<String, dynamic>),
      appStoreProductId: json['app_store_product_id'] as String?,
      backupCreditsDelta: json['backup_credits_delta'] as int? ?? 0,
      bulkEditCreditsDelta: json['bulk_edit_credits_delta'] as int? ?? 0,
      channelCreditsDelta: json['channel_credits_delta'] as int? ?? 0,
      downloadCreditsDelta: json['download_credits_delta'] as int? ?? 0,
      fileTransferCreditsDelta:
          json['file_transfer_credits_delta'] as int? ?? 0,
    );

Map<String, dynamic> _$ListCreditsResponseBodyCreditsToJson(
        ListCreditsResponseBodyCredits instance) =>
    <String, dynamic>{
      'title': instance.title,
      'variant_id': instance.variantId,
      'app_store_product_id': instance.appStoreProductId,
      'price': instance.price,
      'backup_credits_delta': instance.backupCreditsDelta,
      'bulk_edit_credits_delta': instance.bulkEditCreditsDelta,
      'channel_credits_delta': instance.channelCreditsDelta,
      'download_credits_delta': instance.downloadCreditsDelta,
      'file_transfer_credits_delta': instance.fileTransferCreditsDelta,
    };

ListCreditsResponseBodyCreditsPrice
    _$ListCreditsResponseBodyCreditsPriceFromJson(Map<String, dynamic> json) =>
        ListCreditsResponseBodyCreditsPrice(
          amount: json['amount'] as String,
          currencyIcon: json['currency_icon'] as String,
        );

Map<String, dynamic> _$ListCreditsResponseBodyCreditsPriceToJson(
        ListCreditsResponseBodyCreditsPrice instance) =>
    <String, dynamic>{
      'amount': instance.amount,
      'currency_icon': instance.currencyIcon,
    };
