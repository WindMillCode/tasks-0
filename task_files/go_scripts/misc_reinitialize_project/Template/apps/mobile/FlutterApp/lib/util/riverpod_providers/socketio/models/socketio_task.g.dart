// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'socketio_task.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

SocketIOTask _$SocketIOTaskFromJson(Map<String, dynamic> json) => SocketIOTask(
      data: (json['data'] as List<dynamic>)
          .map((e) => SocketIOTaskData.fromJson(e as Map<String, dynamic>))
          .toList(),
      msg: json['msg'] as String,
    );

Map<String, dynamic> _$SocketIOTaskToJson(SocketIOTask instance) =>
    <String, dynamic>{
      'data': instance.data,
      'msg': instance.msg,
    };

SocketIOTaskData _$SocketIOTaskDataFromJson(Map<String, dynamic> json) =>
    SocketIOTaskData(
      taskId: json['task_id'] as String,
    );

Map<String, dynamic> _$SocketIOTaskDataToJson(SocketIOTaskData instance) =>
    <String, dynamic>{
      'task_id': instance.taskId,
    };
