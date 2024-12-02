// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, , unused_catch_stack, unused_local_variable

import 'package:[PROJECT_NAME]/util/env/env.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;
import 'package:[PROJECT_NAME]/util/http/http_utils.dart';

class SocketioRiverpodProviderValue {
  IO.Socket? socket;
  String? id;
  List<String> taskIds = [];

  SocketioRiverpodProviderValue copyWith({String? id, List<String>? taskIds, IO.Socket? socket}) {
    return SocketioRiverpodProviderValue()
      ..id ??= id
      ..socket ??= socket
      ..taskIds = taskIds ?? this.taskIds;
  }
}

var SocketioRiverpodProviderInstance = SocketioRiverpodProviderValue();

bool getMessageFromSocketTask(List<String> taskIds, Map<String, dynamic> finishedTask) {
  try {
    return taskIds.any((taskId) => taskId == finishedTask['data']['task_id']);
  } catch (e) {
    throw e;
  }
}

Map<String, dynamic> validateSocketIOResponse(Map res2, {bool throwError = true}) {
  if (HttpUtils.isServerError(res2["code"]) || HttpUtils.isClientError(res2["code"])) {
    if (throwError == true) {
      throw Exception(res2["data"]['msg']);
    }
  }
  return {'data': res2["data"]['result']};
}

void onConnect(WidgetRef ref) {
  var socket = SocketioRiverpodProviderInstance.socket;
  socket?.onConnect((data) {
    ref.read(SocketioRiverpodProvider.notifier).updateId(socket.id!);
  });
}

performSocketioTask(WidgetRef ref, String event) {}

onYourEvent(WidgetRef ref) {
  var socket = SocketioRiverpodProviderInstance.socket;

  var events = [
    "your_event",
  ];
  events.forEach((event) {
    socket!.on(event, performSocketioTask(ref, event));
  });
}

void disposeSocketIO(WidgetRef ref) {
  var socketio = ref.read(SocketioRiverpodProvider);
  socketio.socket?.disconnect();
  socketio.socket?.close();
  socketio.socket?.dispose();
  socketio.socket?.destroy();
  SocketioRiverpodProviderInstance.socket = null;
}

initSocketIOConnection(BuildContext context, WidgetRef ref) {
  if (SocketioRiverpodProviderInstance.socket != null) {
    return;
  }
  var socket = IO.io(
      APPENV.backendURI0.fqdn, IO.OptionBuilder().setTransports(['websocket']).enableForceNewConnection().build());
  SocketioRiverpodProviderInstance..socket = socket;

  onConnect(ref);
  onYourEvent(ref);
}

class SocketioRiverpodNotifier extends Notifier<SocketioRiverpodProviderValue> {
  @override
  SocketioRiverpodProviderValue build() {
    return SocketioRiverpodProviderInstance;
  }

  updateId(String id) {
    state = state.copyWith(id: id);
  }

  updateSocket(IO.Socket? socket) {
    state = state.copyWith(socket: socket);
  }

  updateTaskIds(List<String> taskIds) {
    state = state.copyWith(taskIds: taskIds);
  }
}

final SocketioRiverpodProvider = NotifierProvider<SocketioRiverpodNotifier, SocketioRiverpodProviderValue>(() {
  return SocketioRiverpodNotifier();
});


