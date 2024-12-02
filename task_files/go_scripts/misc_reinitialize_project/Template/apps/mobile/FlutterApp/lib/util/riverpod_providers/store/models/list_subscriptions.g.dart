// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'list_subscriptions.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ListSubscriptionsRequestBody _$ListSubscriptionsRequestBodyFromJson(
        Map<String, dynamic> json) =>
    ListSubscriptionsRequestBody()
      ..filter = (json['filter'] as List<dynamic>)
          .map((e) => e as Map<String, dynamic>)
          .toList()
      ..sort = (json['sort'] as List<dynamic>)
          .map((e) => Map<String, String>.from(e as Map))
          .toList()
      ..fields = (json['fields'] as List<dynamic>)
          .map((e) => Map<String, String>.from(e as Map))
          .toList()
      ..cursor = Map<String, String>.from(json['cursor'] as Map)
      ..pageNum = json['page_num'] as int
      ..pageSize = json['page_size'] as int
      ..errorOccuredIsPresent = json['error_occured_is_present'] as bool;

Map<String, dynamic> _$ListSubscriptionsRequestBodyToJson(
        ListSubscriptionsRequestBody instance) =>
    <String, dynamic>{
      'filter': instance.filter,
      'sort': instance.sort,
      'fields': instance.fields,
      'cursor': instance.cursor,
      'page_num': instance.pageNum,
      'page_size': instance.pageSize,
      'error_occured_is_present': instance.errorOccuredIsPresent,
    };

ListSubscriptionsResponseBody _$ListSubscriptionsResponseBodyFromJson(
        Map<String, dynamic> json) =>
    ListSubscriptionsResponseBody()
      ..columns = (json['columns'] as List<dynamic>?)
          ?.map((e) => Map<String, String?>.from(e as Map))
          .toList()
      ..data = (json['data'] as List<dynamic>)
          .map((e) =>
              ListSubscriptionsEntity.fromJson(e as Map<String, dynamic>))
          .toList()
      ..pageNum = json['page_num'] as int
      ..pageSize = json['page_size'] as int
      ..totalPages = json['total_pages'] as int
      ..totalItems = json['total_items'] as int
      ..metadata = json['metadata'] as Map<String, dynamic>;

Map<String, dynamic> _$ListSubscriptionsResponseBodyToJson(
        ListSubscriptionsResponseBody instance) =>
    <String, dynamic>{
      'columns': instance.columns,
      'data': instance.data,
      'page_num': instance.pageNum,
      'page_size': instance.pageSize,
      'total_pages': instance.totalPages,
      'total_items': instance.totalItems,
      'metadata': instance.metadata,
    };

ListSubscriptionsEntity _$ListSubscriptionsEntityFromJson(
        Map<String, dynamic> json) =>
    ListSubscriptionsEntity(
      id: json['id'] as String,
      name: json['name'] as String,
      productId: json['product_id'] as String,
      productVariant: ListSubscriptionsEntityProductVariant.fromJson(
          json['product_variant'] as Map<String, dynamic>),
      sellingPlans: (json['selling_plans'] as List<dynamic>)
          .map((e) => ListSubscriptionsEntitySellingPlans.fromJson(
              e as Map<String, dynamic>))
          .toList(),
      appStoreProductId: json['app_store_product_id'] as String?,
    );

Map<String, dynamic> _$ListSubscriptionsEntityToJson(
        ListSubscriptionsEntity instance) =>
    <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'product_id': instance.productId,
      'app_store_product_id': instance.appStoreProductId,
      'product_variant': instance.productVariant,
      'selling_plans': instance.sellingPlans,
    };

ListSubscriptionsEntityProductVariant
    _$ListSubscriptionsEntityProductVariantFromJson(
            Map<String, dynamic> json) =>
        ListSubscriptionsEntityProductVariant(
          id: json['id'] as String,
        );

Map<String, dynamic> _$ListSubscriptionsEntityProductVariantToJson(
        ListSubscriptionsEntityProductVariant instance) =>
    <String, dynamic>{
      'id': instance.id,
    };

ListSubscriptionsEntitySellingPlans
    _$ListSubscriptionsEntitySellingPlansFromJson(Map<String, dynamic> json) =>
        ListSubscriptionsEntitySellingPlans(
          price: json['price'] as String,
          currency: json['currency'] as String,
          id: json['id'] as String,
          interval: json['interval'] as String,
          description: (json['description'] as List<dynamic>)
              .map((e) => e as String)
              .toList(),
        );

Map<String, dynamic> _$ListSubscriptionsEntitySellingPlansToJson(
        ListSubscriptionsEntitySellingPlans instance) =>
    <String, dynamic>{
      'price': instance.price,
      'currency': instance.currency,
      'id': instance.id,
      'interval': instance.interval,
      'description': instance.description,
    };
