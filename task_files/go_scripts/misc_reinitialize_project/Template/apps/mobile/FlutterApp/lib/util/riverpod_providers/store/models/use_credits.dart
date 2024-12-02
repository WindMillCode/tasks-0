import 'package:json_annotation/json_annotation.dart';

part 'use_credits.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class UseCreditsRequestBody {

  // snake case keyof  ListUsersCredits
  String creditType;
  int amount;
  String accessToken;

  UseCreditsRequestBody({
    required this.creditType,
    required this.amount,
    required this.accessToken,
  });

  factory UseCreditsRequestBody.fromJson(Map<String, dynamic> json) => _$UseCreditsRequestBodyFromJson(json);
  Map<String, dynamic> toJson() => _$UseCreditsRequestBodyToJson(this);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class UseCreditsResponseBody {
  UseCreditsResponseBody();

  factory UseCreditsResponseBody.fromJson(Map<String, dynamic> json) => _$UseCreditsResponseBodyFromJson(json);
  Map<String, dynamic> toJson() => _$UseCreditsResponseBodyToJson(this);
}
