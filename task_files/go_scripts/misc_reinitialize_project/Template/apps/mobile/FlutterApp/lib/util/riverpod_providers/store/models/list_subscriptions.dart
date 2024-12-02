import 'package:tuli/util/wml/pagination/pagination.dart';
import 'package:json_annotation/json_annotation.dart';

part 'list_subscriptions.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class ListSubscriptionsRequestBody extends WMLAPIPageRequestModel {
  ListSubscriptionsRequestBody();

  factory ListSubscriptionsRequestBody.fromJson(Map<String, dynamic> json) =>
      _$ListSubscriptionsRequestBodyFromJson(json);

  Map<String, dynamic> toJson() => _$ListSubscriptionsRequestBodyToJson(this);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class ListSubscriptionsResponseBody extends WMLAPIPageResponseModel<ListSubscriptionsEntity> {
  ListSubscriptionsResponseBody();

  factory ListSubscriptionsResponseBody.fromJson(Map<String, dynamic> json) =>
      _$ListSubscriptionsResponseBodyFromJson(json);

  Map<String, dynamic> toJson() => _$ListSubscriptionsResponseBodyToJson(this);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class ListSubscriptionsEntity {
  String id;
  String name;
  String productId;
  String? appStoreProductId;
  ListSubscriptionsEntityProductVariant productVariant;
  List<ListSubscriptionsEntitySellingPlans> sellingPlans = [];

  ListSubscriptionsEntity({
    required this.id,
    required this.name,
    required this.productId,
    required this.productVariant,
    required this.sellingPlans,
    this.appStoreProductId,
  });

  factory ListSubscriptionsEntity.fromJson(Map<String, dynamic> json) => _$ListSubscriptionsEntityFromJson(json);

  Map<String, dynamic> toJson() => _$ListSubscriptionsEntityToJson(this);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class ListSubscriptionsEntityProductVariant {
  String id;

  ListSubscriptionsEntityProductVariant({
    required this.id,
  });

  factory ListSubscriptionsEntityProductVariant.fromJson(Map<String, dynamic> json) =>
      _$ListSubscriptionsEntityProductVariantFromJson(json);

  Map<String, dynamic> toJson() => _$ListSubscriptionsEntityProductVariantToJson(this);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class ListSubscriptionsEntitySellingPlans {
  String price;
  String currency;
  String id;
  String interval;
  List<String> description;

  ListSubscriptionsEntitySellingPlans({
    required this.price,
    required this.currency,
    required this.id,
    required this.interval,
    required this.description,
  });

  factory ListSubscriptionsEntitySellingPlans.fromJson(Map<String, dynamic> json) =>
      _$ListSubscriptionsEntitySellingPlansFromJson(json);

  Map<String, dynamic> toJson() => _$ListSubscriptionsEntitySellingPlansToJson(this);
}
