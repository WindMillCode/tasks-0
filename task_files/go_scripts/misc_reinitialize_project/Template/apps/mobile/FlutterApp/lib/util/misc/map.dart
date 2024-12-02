import 'package:tuli/util/misc/iterable.dart';

Map<String, dynamic> deepCopyMap(Map<dynamic, dynamic> original) {
  Map<String, dynamic> copy = {};
  original.forEach((key, value) {
    if (value is Map) {
      // Recursively copy nested maps
      copy[key] = deepCopyMap(value);
    } else if (value is List) {
      // Recursively copy nested lists
      copy[key] = deepCopyList(value);
    } else {
      // Copy primitive and other types directly
      copy[key] = value;
    }
  });
  return copy;
}

Map<String, dynamic> deepCopyInclude(List<String> keysArray, Map oldObject) {
  Map<String, dynamic> newObject = {};

  for (String key in keysArray) {
    if (oldObject.containsKey(key)) {
      newObject[key] = oldObject[key];
    }
  }

  return newObject;
}

Map<String, dynamic> deepCopyExclude(List<String> keysArray, Map oldObject) {
  Map<String, dynamic> filteredObject = {};

  for (String key in oldObject.keys) {
    if (!keysArray.contains(key)) {
      filteredObject[key] = oldObject[key];
    }
  }

  return filteredObject;
}
