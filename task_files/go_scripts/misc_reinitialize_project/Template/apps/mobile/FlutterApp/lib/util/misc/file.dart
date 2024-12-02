import 'dart:convert';
import 'dart:math';
import 'dart:typed_data';

import 'package:awesome_notifications/awesome_notifications.dart';
import 'package:form_builder_file_picker/form_builder_file_picker.dart';
import 'package:path_provider/path_provider.dart';
import 'dart:io';
import 'package:open_filex/open_filex.dart';
import 'dart:math' as math;

Future<void> downloadFile(String filename, dynamic content,
    {bool? openFile,
    bool? shouldNotify,
    String? shouldNotifyTitle,
    String? shouldNotifyBody,
    Map<String, String> Function(String filePath)? createPayload}) async {
  Uint8List bytes = base64.decode("");
  if (content is String) {
    bytes = base64.decode(content);
  } else if (content is Uint8List) {
    bytes = content;
  } else if (content == null) {
  } else {
    throw Exception("type must be String or Uint8List or null ");
  }

  String filePath;
  if (content != null) {
    filePath = await getDownloadFilePath(filename);
    final file = File(filePath);

    // Write the file
    await file.writeAsBytes(bytes);
  } else {
    filePath = filename;
  }
  if (openFile == true) {
    await OpenFilex.open(filePath);
  }

  if (shouldNotify == true) {
    AwesomeNotifications().createNotification(
        content: NotificationContent(
            id: 10,
            channelKey: 'basic_channel',
            actionType: ActionType.KeepOnTop,
            title: shouldNotifyTitle,
            body: shouldNotifyBody,
            payload: createPayload?.call(filePath)));
  }
}

Future<String> getDownloadFilePath(String filename) async {
  // var directory = await getApplicationDocumentsDirectory();
  late final directory;
  if (Platform.isAndroid) {
    directory = "/storage/emulated/0/Download/";
    return '$directory/$filename';
  } else if (Platform.isIOS) {
    directory = await getApplicationDocumentsDirectory();
    return '${directory.path}/$filename';
  } else{
    directory = await getApplicationDocumentsDirectory();
    return '${directory.path}/$filename';
  }

}

String uint8ListToBase64(Uint8List myBinary) {
  return base64.encode(myBinary);
}

Uint8List base64ToUint8List(String target) {
  final binaryString = base64.decode(target);
  return Uint8List.fromList(binaryString);
}

List<Map<String, dynamic>> chunkFileContent<T>(T content, {int chunkSize = 16384}) {
  final chunks = <Map<String, dynamic>>[];

  if (content is Uint8List) {
    final totalChunks = (content.lengthInBytes / chunkSize).ceil();

    for (var i = 0; i < totalChunks; i++) {
      final start = i * chunkSize;
      final end = start + chunkSize > content.lengthInBytes ? content.lengthInBytes : start + chunkSize;
      final chunk = content.sublist(start, end);
      final serializedChunk = uint8ListToBase64(chunk);
      chunks.add({
        'chunk': serializedChunk,
        'order': i,
      });
    }
  } else if (content is String) {
    final totalChunks = (content.length / chunkSize).ceil();

    for (var i = 0; i < totalChunks; i++) {
      final start = i * chunkSize;
      final end = start + chunkSize > content.length ? content.length : start + chunkSize;
      final chunk = content.substring(start, end);
      chunks.add({
        'chunk': chunk,
        'order': i,
      });
    }
  } else {
    throw Exception("Unsupported content type: ${content.runtimeType}");
  }

  return chunks;
}

String formatBytes(int bytes, [int decimals = 2]) {
  if (bytes == 0) return '0 Bytes';

  const int k = 1024;
  int dm = decimals < 0 ? 0 : decimals;
  const List<String> sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

  int i = (math.log(bytes) / math.log(k)).floor();

  return '${(bytes / math.pow(k, i)).toStringAsFixed(dm)} ${sizes[i]}';
}

Future<List<Map<String, dynamic>>> loadFilesFromPlatformFiles(List<PlatformFile> files) async {
  return Future.wait(files.map((file) async {
    Uint8List fileBytes;

    if (file.bytes != null) {
      fileBytes = file.bytes!;
    } else if (file.readStream != null) {
      final content = await file.readStream!.expand((bytes) => bytes).toList();
      fileBytes = Uint8List.fromList(content);
    } else if (file.path != null) {
      fileBytes = await File(file.path!).readAsBytes();
    } else {
      throw Exception("No valid file content found");
    }

    return {
      'name': file.name,
      'type': file.extension ?? 'unknown',
      'url': file.path ?? '',
      'content': fileBytes,
    };
  }).toList());
}

Future<Map> readChunk(RandomAccessFile raf, int offset, int length, int fileSize,
    { // "BINARY" | "BASE64"
    chunkType = "BASE64"}) async {
  await raf.setPosition(offset);
  dynamic chunk = await raf.read(length);
  if (chunkType == "BASE64") {
    chunk = base64Encode(chunk);
  }
  return {"chunk": chunk, "progress": min(100, (offset + length) / fileSize * 100), "chunkProgress": 100};
}
