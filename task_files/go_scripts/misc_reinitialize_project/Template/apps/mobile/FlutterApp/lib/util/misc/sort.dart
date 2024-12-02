int genericSort<T>(T a, T b, [String order = 'ASC', Function(T)? aKeyExtractor, Function(T)? bKeyExtractor]) {
  // If only one extractor is provided, use it for both a and b
  aKeyExtractor ??= bKeyExtractor ?? (T value) => value;
  bKeyExtractor ??= aKeyExtractor;

  // Apply key extractors to get the comparison keys
  var keyA = aKeyExtractor(a);
  var keyB = bKeyExtractor(b);

  // Handle null values
  if (keyA == null) return 1;
  if (keyB == null) return -1;

  // Determine the result based on the type of the keys
  int result;
  if (keyA is num && keyB is num) {
    result = keyA.compareTo(keyB);
  } else if (keyA is DateTime && keyB is DateTime) {
    result = keyA.compareTo(keyB);
  } else {
    // Default to string comparison if keys are not numbers or DateTime
    result = keyA.toString().toLowerCase().compareTo(keyB.toString().toLowerCase());
  }

  // Adjust for order
  return order == 'DESC' ? -result : result;
}

List<T> sortListByProperty<T>(List<T> list, Function(T) keyExtractor, [String order = 'asc']) {
  var sortedList = List<T>.from(list);
  sortedList.sort((a, b) => genericSort(a, b, order, keyExtractor));
  return sortedList;
}
