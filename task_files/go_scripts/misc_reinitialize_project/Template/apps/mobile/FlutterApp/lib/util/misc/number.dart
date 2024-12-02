import 'dart:math';

import 'package:intl/intl.dart';

String roundToXDecimals(double num, int decimals, {String locale = 'en'}) {
  NumberFormat decimalFormat = NumberFormat.currency(
    locale: locale,
    symbol: '',
    decimalDigits: decimals,
  );
  return decimalFormat.format(num);
}

double log10(double value) {
  return (value == 0) ? 0 : (value > 0) ? log(value) / ln10 : log(-value) / ln10;
}

