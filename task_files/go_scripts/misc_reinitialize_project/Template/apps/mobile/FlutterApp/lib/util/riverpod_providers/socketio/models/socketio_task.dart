

import 'package:json_annotation/json_annotation.dart';

part 'socketio_task.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class SocketIOTask {
  List<SocketIOTaskData> data;
  String msg;

  SocketIOTask({required this.data, required this.msg});

  // From JSON
  factory SocketIOTask.fromJson(Map<String, dynamic> json) =>
      _$SocketIOTaskFromJson(json);

  // To JSON
  Map<String, dynamic> toJson() => _$SocketIOTaskToJson(this);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class SocketIOTaskData {
  String taskId;

  SocketIOTaskData({required this.taskId});

	factory SocketIOTaskData.fromJson(Map<String, dynamic> json) => _$SocketIOTaskDataFromJson(json);

}
