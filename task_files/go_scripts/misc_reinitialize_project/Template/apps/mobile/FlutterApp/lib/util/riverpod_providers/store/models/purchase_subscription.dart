import 'package:[PROJECT_NAME]/util/env/env.dart';
import 'package:[PROJECT_NAME]/util/specific/classes.dart';
import 'package:json_annotation/json_annotation.dart';

part 'purchase_subscription.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class PurchaseSubscriptionRequestBody extends AuthBase {
  PurchaseSubscriptionRequestBody(
      {required String accessToken,
      this.subscriptionContractIds = const [],
      this.items = const [],
      String redirectUrl = ''})
      : super(
          accessToken: accessToken,
        );

  List<String> subscriptionContractIds = [];
  List<PurchaseSubscriptionRequestBodyItems> items = [];
  String redirectUrl = '';
  AppEnvPlatformType appPlatformTypeEnumValue = APPENV.platformType;
  Map<String, dynamic> toJson() {
    var reqBody = _$PurchaseSubscriptionRequestBodyToJson(this);
    reqBody["app_platform_type_enum_value"] = appPlatformTypeEnumValue.index;
    return reqBody;
  }

  factory PurchaseSubscriptionRequestBody.fromJson(Map<String, dynamic> json) =>
      _$PurchaseSubscriptionRequestBodyFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class PurchaseSubscriptionRequestBodyItems {
  PurchaseSubscriptionRequestBodyItems(
      {this.quantity = 1, required this.variantId, required this.sellingPlanId, this.iosInfo});

  int quantity;
  String variantId;
  String sellingPlanId;
  Map? iosInfo;
  Map<String, dynamic> toJson() => _$PurchaseSubscriptionRequestBodyItemsToJson(this);

  factory PurchaseSubscriptionRequestBodyItems.fromJson(Map<String, dynamic> json) =>
      _$PurchaseSubscriptionRequestBodyItemsFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class PurchaseSubscriptionResponseBody {
  PurchaseSubscriptionResponseBody();

  String? id = '';
  String weburl = '';
  Map<String, dynamic> toJson() => _$PurchaseSubscriptionResponseBodyToJson(this);

  factory PurchaseSubscriptionResponseBody.fromJson(Map<String, dynamic> json) =>
      _$PurchaseSubscriptionResponseBodyFromJson(json);
}
