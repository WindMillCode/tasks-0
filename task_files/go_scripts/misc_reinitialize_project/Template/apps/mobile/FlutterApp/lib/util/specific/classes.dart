import 'package:[PROJECT_NAME]/util/env/env.dart';
import 'package:json_annotation/json_annotation.dart';

part 'classes.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class AuthBase {
  String accessToken;
  final AppEnvPlatformType appPlatformTypeEnumValue = APPENV.platformType;

  AuthBase({
    required this.accessToken,
  });
  Map<String, dynamic> toJson() {
    var reqBody = _$AuthBaseToJson(this);
    reqBody["app_platform_type_enum_value"] = appPlatformTypeEnumValue.index;
    return reqBody;
  }

  factory AuthBase.fromJson(Map<String, dynamic> json) => _$AuthBaseFromJson(json);
}



