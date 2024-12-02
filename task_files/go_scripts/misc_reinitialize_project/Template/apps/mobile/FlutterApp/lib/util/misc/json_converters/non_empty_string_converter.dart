import 'package:json_annotation/json_annotation.dart';

class NonEmptyStringConverter implements JsonConverter<String?, String?> {
  const NonEmptyStringConverter();

  @override
  String? fromJson(String? json) => json;

  @override
  String? toJson(String? object) {
    return (object == null || object.isEmpty) ? null : object;
  }
}
