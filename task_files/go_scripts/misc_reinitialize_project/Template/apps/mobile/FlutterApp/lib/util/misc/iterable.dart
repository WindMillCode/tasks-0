import 'package:tuli/util/misc/map.dart';

List<T> removeDuplicates<T>(List<T> items, dynamic Function(T) keySelector) {
  var seenKeys = Set<dynamic>();
  var uniqueItems = <T>[];

  for (var item in items) {
    var key = keySelector(item);

    if (!seenKeys.contains(key)) {
      seenKeys.add(key);
      uniqueItems.add(item);
    }
  }

  return uniqueItems;
}

T? safeAccessIndexOfList<T>(List<T> list, int index, [dynamic defaultVal]) {
  if (index >= 0 && index < list.length) {
    return list[index];
  }
  return defaultVal;
}

List<dynamic> deepCopyList(List<dynamic> original) {
  return List<dynamic>.from(original.map((item) => item is Map
      ? deepCopyMap(item)
      : item is List
          ? deepCopyList(item)
          : item));
}
