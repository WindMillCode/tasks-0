import 'package:[PROJECT_NAME]/util/specific/classes.dart';
import 'package:json_annotation/json_annotation.dart';

part 'cancel_subscription.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class CancelSubscriptionRequestBody extends AuthBase {
  CancelSubscriptionRequestBody({required String accessToken, this.subscriptionContractInfo = const []})
      : super(
          accessToken: accessToken,
        );

  List<CancelSubscriptionRequestBodySubscriptionContractInfo> subscriptionContractInfo = [];
  Map<String, dynamic> toJson() => _$CancelSubscriptionRequestBodyToJson(this);

  factory CancelSubscriptionRequestBody.fromJson(Map<String, dynamic> json) =>
      _$CancelSubscriptionRequestBodyFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class CancelSubscriptionRequestBodySubscriptionContractInfo {
  CancelSubscriptionRequestBodySubscriptionContractInfo({this.id = '', this.confirmationNumber = ''});

  String id;
  String confirmationNumber;
  Map<String, dynamic> toJson() => _$CancelSubscriptionRequestBodySubscriptionContractInfoToJson(this);

  factory CancelSubscriptionRequestBodySubscriptionContractInfo.fromJson(Map<String, dynamic> json) =>
      _$CancelSubscriptionRequestBodySubscriptionContractInfoFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class CancelSubscriptionResponseBody {
  CancelSubscriptionResponseBody();
  Map<String, dynamic> toJson() => _$CancelSubscriptionResponseBodyToJson(this);

  factory CancelSubscriptionResponseBody.fromJson(Map<String, dynamic> json) =>
      _$CancelSubscriptionResponseBodyFromJson(json);
}
