import 'package:tuli/util/riverpod_providers/store/models/store_riverpod_provider_category.dart';
import 'package:tuli/util/riverpod_providers/store/models/store_riverpod_provider_price.dart';

class StoreRiverpodProviderItem {
  String? id;
  // maybe map
  List<dynamic> imgs = [];
  String? title;
  String? subtitle;
  int quantity = 1;
  List<StoreRiverpodProviderCategory> categories = [];
  StoreRiverpodProviderPrice price;
  // RatingStarsParams rating = RatingStarsParams(rating: 4);

  StoreRiverpodProviderItem(
      {this.id,
      this.title,
      this.subtitle,
      this.imgs = const [],
      this.quantity = 1,
      this.categories = const [],
      StoreRiverpodProviderPrice? price})
      : price = price ?? StoreRiverpodProviderPrice();

  get displayImg => imgs.isNotEmpty ? imgs[0] : null;

  set displayImg(value) {
    imgs.insert(0, value);
  }

  StoreRiverpodProviderPrice get totalPrice {
    StoreRiverpodProviderPrice result = StoreRiverpodProviderPrice();

    double value = price.business * quantity;
    result.business = value;
    result.currency = price.currency;

    return result;
  }
}
