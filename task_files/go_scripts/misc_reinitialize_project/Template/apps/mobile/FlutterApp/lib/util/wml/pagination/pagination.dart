import 'package:tuli/util/misc/sort.dart';

enum WMLAPIPageRequestModelFilterPredicateEnum {
  EQUALS,
  STARTSWITH,
  ENDSWITH,
  CONTAINS,
}

class WMLAPIPageRequestModel {
  List<Map<String, dynamic>> filter;
  List<Map<String, String>> sort;
  List<Map<String, String>> fields;
  Map<String, String> cursor;
  int pageNum;
  int pageSize;
  bool errorOccuredIsPresent;

  WMLAPIPageRequestModel(
      {this.fields = const [],
      this.filter = const [],
      this.sort = const [],
      this.pageNum = 0,
      this.pageSize = 0,
      this.errorOccuredIsPresent = false,
      this.cursor = const {}});

  factory WMLAPIPageRequestModel.fromJson(Map<String, dynamic> json) {
    return WMLAPIPageRequestModel(
      fields: (json['fields'] as List<dynamic>?)
              ?.map((filterMap) => (filterMap as Map<String, dynamic>).cast<String, String>())
              .toList() ??
          [],
      filter: (json['filter'] as List<dynamic>?)
              ?.map((filterMap) => (filterMap as Map<String, dynamic>).cast<String, String>())
              .toList() ??
          [],
      sort: (json['sort'] as List<dynamic>?)
              ?.map((sortMap) => (sortMap as Map<String, dynamic>).cast<String, String>())
              .toList() ??
          [],
      cursor: (json['cursor'] as Map<String, dynamic>?)?.cast<String, String>() ?? {},
      pageNum: json['page_num'] as int? ?? 0,
      pageSize: json['page_size'] as int? ?? 0,
      errorOccuredIsPresent: json['error_occured_is_present'] as bool? ?? false,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'fields': fields,
      'filter': filter,
      'sort': sort,
      'cursor': cursor,
      'page_num': pageNum,
      'page_size': pageSize,
      'error_occured_is_present': errorOccuredIsPresent,
    };
  }

  WMLAPIPageRequestModel copyWith({
    List<Map<String, dynamic>>? filter,
    List<Map<String, String>>? sort,
    List<Map<String, String>>? fields,
    Map<String, String>? cursor,
    int? pageNum,
    int? pageSize,
    bool? errorOccuredIsPresent,
  }) {
    return WMLAPIPageRequestModel()
      ..filter = filter ?? this.filter
      ..sort = sort ?? this.sort
      ..fields = fields ?? this.fields
      ..cursor = cursor ?? this.cursor
      ..pageNum = pageNum ?? this.pageNum
      ..pageSize = pageSize ?? this.pageSize
      ..errorOccuredIsPresent = errorOccuredIsPresent ?? this.errorOccuredIsPresent;
  }
}

class WMLAPIPageResponseModel<DT> {
  List<Map<String, String?>>? columns;
  List<DT> data;
  int pageNum;
  int pageSize;
  int totalPages;
  int totalItems;
  Map<dynamic, dynamic> metadata;

  WMLAPIPageResponseModel(
      {this.columns = const [],
      this.data = const [],
      this.pageNum = 0,
      this.pageSize = 0,
      this.totalPages = 0,
      this.totalItems = 0,
      this.metadata = const {}});

  factory WMLAPIPageResponseModel.fromJson(
      Map<String, dynamic> json,
      // TODO why is this here
      Function dataTransformPredicate) {
    return WMLAPIPageResponseModel<DT>(
        columns: (json['columns'] as List<dynamic>?)
                ?.map((columnMap) => (columnMap as Map<String, dynamic>).cast<String, String?>())
                .toList() ??
            [],
        data: json['data'],
        pageNum: json['page_num'] as int? ?? 0,
        pageSize: json['page_size'] as int? ?? 0,
        totalPages: json['total_pages'] as int? ?? 0,
        totalItems: json['total_items'] as int? ?? 0,
        metadata: json["metadata"] ?? {});
  }

  Map<String, dynamic> toJson() {
    return {
      'columns': columns,
      'data': data,
      'page_num': pageNum,
      'page_size': pageSize,
      'total_pages': totalPages,
      'total_items': totalItems,
      'metadata': metadata
    };
  }

  WMLAPIPageResponseModel<DT> copyWith({
    List<Map<String, String?>>? columns,
    List<DT>? data,
    int? pageNum,
    int? pageSize,
    int? totalPages,
    int? totalItems,
    Map<dynamic, dynamic>? metadata,
  }) {
    return WMLAPIPageResponseModel<DT>()
      ..columns = columns ?? this.columns
      ..data = data ?? this.data
      ..pageNum = pageNum ?? this.pageNum
      ..pageSize = pageSize ?? this.pageSize
      ..totalPages = totalPages ?? this.totalPages
      ..totalItems = totalItems ?? this.totalItems
      ..metadata = metadata ?? this.metadata;
  }

  WMLAPIPageResponseModel<DT> calculateCurrentState({int? totalPages, int? totalItems, int? pageSize}) {
    int displayPageNum = pageNum + 1;
    if (totalItems != null) {
      this.pageSize = data.length;
      this.totalItems = totalItems;
      this.totalPages = (totalItems / (pageSize ?? this.pageSize)).ceil();
    } else {
      totalPages = totalPages ?? displayPageNum;
      this.pageSize = data.length;
      this.totalPages = totalPages;
      this.totalItems = totalPages * data.length;
    }
    return this;
  }

  WMLAPIPageResponseModel<DT> calcSectionBasedOnPageDetails({
    required List<DT> data,
    required int pageSize,
    required int pageNum,
    int? totalItems,
  }) {
    int startIndex = pageNum * pageSize;
    int endIndex = (pageNum + 1) * pageSize;
    this.pageNum = pageNum;
    this.pageSize = pageSize;
    this.totalItems = totalItems ?? data.length;
    if (this.totalItems == 0) {
      this.totalPages = 0;
    } else {
      this.totalPages = (this.totalItems / this.pageSize).ceil();
    }
    this.data = data.sublist(startIndex.clamp(0, data.length), endIndex.clamp(0, data.length));
    return this;
  }
}

List<T> filterRows<T extends dynamic>(List<Map<String, dynamic>> filters, List<T> rows) {
  if (filters.isEmpty) {
    return rows;
  }

  // Apply filter based on predicate
  bool applyPredicate(String source, String value, WMLAPIPageRequestModelFilterPredicateEnum? predicate) {
    switch (predicate) {
      case WMLAPIPageRequestModelFilterPredicateEnum.STARTSWITH:
        return source.toLowerCase().startsWith(value.toLowerCase());
      case WMLAPIPageRequestModelFilterPredicateEnum.ENDSWITH:
        return source.toLowerCase().endsWith(value.toLowerCase());
      case WMLAPIPageRequestModelFilterPredicateEnum.CONTAINS:
        return source.toLowerCase().contains(value.toLowerCase());
      case WMLAPIPageRequestModelFilterPredicateEnum.EQUALS:
      default:
        return source == value;
    }
  }

  // Grouping filters by their 'key' values
  var groupedFilters = <String, List<Map<String, dynamic>>>{};
  for (var filter in filters) {
    var key = filter['key'];
    if (key != null) {
      groupedFilters.putIfAbsent(key, () => []).add(filter);
    }
  }

  // Apply filters
  return rows.where((row) {
    var myJson = row.toJson();

    // Apply AND logic across different keys
    return groupedFilters.entries.every((group) {
      // Apply OR logic for filters with the same key
      return group.value.any((filter) {
        var key = filter['key'];
        var value = filter['value'];
        var predicate = filter['predicate'] ?? WMLAPIPageRequestModelFilterPredicateEnum.EQUALS;
        return key != null && value != null && applyPredicate(myJson[key].toString(), value, predicate);
      });
    });
  }).toList();
}

abstract class SortRowsRow {
  Map<String, dynamic> toJson();
}

// TODO quite cant use extends SortRowsRow becuase the row needs to inheretnly extend it
List<T> sortRows<T extends dynamic>(List<Map<String, String>> sorts, List<T> rows) {
  if (sorts.isEmpty) {
    return rows;
  }

  // We need to create a custom comparator that leverages genericSort
  rows.sort((a, b) {
    for (var sortMap in sorts) {
      var sortColumn = sortMap['key'];
      var sortOrder = sortMap['direction'] ?? 'ASC';

      // Key extractor function for the current sort column
      Function(T) keyExtractor = (T obj) => obj.toJson()[sortColumn];

      // Use genericSort to compare the two objects based on the current sort column
      int result = genericSort(a, b, sortOrder, keyExtractor);

      if (result != 0) {
        return result; // Return the result if there's a difference
      }
      // If equal, proceed to the next sort criterion
    }
    return 0; // All sort criteria resulted in equality
  });

  return rows;
}

List<T> paginateRows<T extends dynamic>(List<T> rows, int pageNum, int pageSize) {
  int startIndex = pageNum * pageSize;
  int endIndex = startIndex + pageSize;
  if (startIndex > rows.length) return [];
  endIndex = endIndex > rows.length ? rows.length : endIndex;
  return rows.sublist(startIndex, endIndex);
}
