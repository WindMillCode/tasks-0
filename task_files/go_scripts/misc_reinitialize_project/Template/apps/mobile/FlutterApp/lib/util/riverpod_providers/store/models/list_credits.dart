import 'package:json_annotation/json_annotation.dart';

part 'list_credits.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class ListCreditsRequestBody {
  ListCreditsRequestBody();
  Map<String, dynamic> toJson() => _$ListCreditsRequestBodyToJson(this);

  factory ListCreditsRequestBody.fromJson(Map<String, dynamic> json) => _$ListCreditsRequestBodyFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class ListCreditsResponseBody {
  ListCreditsResponseBody();

  List<ListCreditsResponseBodyCredits> credits = [];
  Map<String, dynamic> toJson() => _$ListCreditsResponseBodyToJson(this);

  factory ListCreditsResponseBody.fromJson(Map<String, dynamic> json) => _$ListCreditsResponseBodyFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class ListCreditsResponseBodyCredits {
  String title;
  String variantId;
  String? appStoreProductId;
  ListCreditsResponseBodyCreditsPrice price;
  int backupCreditsDelta;
  int bulkEditCreditsDelta;
  int channelCreditsDelta;
  int downloadCreditsDelta;
  int fileTransferCreditsDelta;

  ListCreditsResponseBodyCredits({
    required this.title,
    required this.variantId,
    required this.price,
    this.appStoreProductId,
    this.backupCreditsDelta = 0,
    this.bulkEditCreditsDelta = 0,
    this.channelCreditsDelta = 0,
    this.downloadCreditsDelta = 0,
    this.fileTransferCreditsDelta = 0,
  });
  Map<String, dynamic> toJson() => _$ListCreditsResponseBodyCreditsToJson(this);

  factory ListCreditsResponseBodyCredits.fromJson(Map<String, dynamic> json) =>
      _$ListCreditsResponseBodyCreditsFromJson(json);
}

@JsonSerializable(fieldRename: FieldRename.snake)
class ListCreditsResponseBodyCreditsPrice {
  String amount;
  String currencyIcon;

  ListCreditsResponseBodyCreditsPrice({
    required this.amount,
    required this.currencyIcon,
  });
  Map<String, dynamic> toJson() => _$ListCreditsResponseBodyCreditsPriceToJson(this);

  factory ListCreditsResponseBodyCreditsPrice.fromJson(Map<String, dynamic> json) =>
      _$ListCreditsResponseBodyCreditsPriceFromJson(json);
}
