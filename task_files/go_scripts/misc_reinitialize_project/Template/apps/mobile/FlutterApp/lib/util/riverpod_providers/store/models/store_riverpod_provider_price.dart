import 'package:tuli/util/misc/number.dart';

class StoreRiverpodProviderPrice {
  double business;
  String currency;
  String? serverDisplay;

  StoreRiverpodProviderPrice({
    this.business = 0.0,
    this.currency = "\$",
    this.serverDisplay,
  });

  String get display {
    return serverDisplay ?? currency + roundToXDecimals(business, 2).toString();
  }
}
