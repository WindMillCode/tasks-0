// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'socketio_client_response_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

SocketioClientResponseModel _$SocketioClientResponseModelFromJson(
        Map<String, dynamic> json) =>
    SocketioClientResponseModel(
      code: json['code'] as int,
      data: SocketioClientResponseModelData.fromJson(
          json['data'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$SocketioClientResponseModelToJson(
        SocketioClientResponseModel instance) =>
    <String, dynamic>{
      'code': instance.code,
      'data': instance.data,
    };

SocketioClientResponseModelData _$SocketioClientResponseModelDataFromJson(
        Map<String, dynamic> json) =>
    SocketioClientResponseModelData(
      taskId: json['task_id'] as String?,
      result: json['result'] as Map<String, dynamic>,
    );

Map<String, dynamic> _$SocketioClientResponseModelDataToJson(
        SocketioClientResponseModelData instance) =>
    <String, dynamic>{
      'task_id': instance.taskId,
      'result': instance.result,
    };
