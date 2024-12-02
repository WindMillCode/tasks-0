import 'package:json_annotation/json_annotation.dart';

part 'socketio_client_response_model.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class SocketioClientResponseModel {
  SocketioClientResponseModel({
    required this.code,
    required this.data,
  });

  int code;
  SocketioClientResponseModelData data;


	Map<String, dynamic> toJson() => _$SocketioClientResponseModelToJson(this);

	factory SocketioClientResponseModel.fromJson(Map<String, dynamic> json) => _$SocketioClientResponseModelFromJson(json);

}

@JsonSerializable(fieldRename: FieldRename.snake)
class SocketioClientResponseModelData {
  SocketioClientResponseModelData({
    this.taskId,
    required this.result,
  });

  String? taskId;
  Map result;

	Map<String, dynamic> toJson() => _$SocketioClientResponseModelDataToJson(this);

	factory SocketioClientResponseModelData.fromJson(Map<String, dynamic> json) => _$SocketioClientResponseModelDataFromJson(json);

}
