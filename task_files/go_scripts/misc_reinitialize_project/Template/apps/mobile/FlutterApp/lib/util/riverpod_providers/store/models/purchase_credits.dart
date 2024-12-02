import 'package:[PROJECT_NAME]/util/env/env.dart';
import 'package:[PROJECT_NAME]/util/specific/classes.dart';
import 'package:json_annotation/json_annotation.dart';

part 'purchase_credits.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class PurchaseCreditsRequestBody extends AuthBase {
  PurchaseCreditsRequestBody({required String accessToken, this.items = const []})
      : super(
          accessToken: accessToken,
        );

  List<PurchaseCreditsRequestBodyItems> items;
  AppEnvPlatformType appPlatformTypeEnumValue = APPENV.platformType;
  Map<String, dynamic> toJson() {
    var reqBody = _$PurchaseCreditsRequestBodyToJson(this);
    reqBody["app_platform_type_enum_value"] = appPlatformTypeEnumValue.index;
    return reqBody;
  }

  factory PurchaseCreditsRequestBody.fromJson(Map<String, dynamic> json) => _$PurchaseCreditsRequestBodyFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class PurchaseCreditsRequestBodyItems {
  PurchaseCreditsRequestBodyItems({this.quantity = 0, this.variantId = '', this.iosInfo});

  int quantity = 0;
  String? variantId = '';
  Map? iosInfo;
  Map<String, dynamic> toJson() => _$PurchaseCreditsRequestBodyItemsToJson(this);

  factory PurchaseCreditsRequestBodyItems.fromJson(Map<String, dynamic> json) =>
      _$PurchaseCreditsRequestBodyItemsFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class PurchaseCreditsResponseBody {
  PurchaseCreditsResponseBody();

  String weburl = '';
  Map<String, dynamic> toJson() => _$PurchaseCreditsResponseBodyToJson(this);

  factory PurchaseCreditsResponseBody.fromJson(Map<String, dynamic> json) =>
      _$PurchaseCreditsResponseBodyFromJson(json);
}
