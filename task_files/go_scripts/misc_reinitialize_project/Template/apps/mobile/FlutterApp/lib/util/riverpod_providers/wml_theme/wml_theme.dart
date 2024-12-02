// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:tuli/util/riverpod_providers/wml_spacing/wml_spacing.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../wml_colors/wml_colors.dart';
import '../wml_fonts/wml_fonts.dart';

ThemeData _buildWMLTheme(ThemeData base) {
  return base.copyWith(
    scaffoldBackgroundColor: WMLColorsRiverpodProviderInstance.primaryColor,
    buttonTheme: ButtonThemeData(buttonColor: WMLColorsRiverpodProviderInstance.primaryColor),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: ElevatedButton.styleFrom(
          elevation: 0,
          padding: EdgeInsets.symmetric(
              vertical: WMLSpacingRiverpodProviderInstance.large,
              horizontal: WMLSpacingRiverpodProviderInstance.enormous),
          foregroundColor: WMLColorsRiverpodProviderInstance.white,
          backgroundColor: WMLColorsRiverpodProviderInstance.primaryColor,
          shape: RoundedRectangleBorder(
              side: BorderSide(color: WMLColorsRiverpodProviderInstance.white, width: 1),
              borderRadius: BorderRadius.all(Radius.circular(WMLSpacingRiverpodProviderInstance.containerRadius)))
          // Add other button styling as needed
          ),
    ),
    colorScheme: base.colorScheme.copyWith(
      primary: WMLColorsRiverpodProviderInstance.primaryColor,
      onPrimary: WMLColorsRiverpodProviderInstance.secondaryColor,
      secondary: WMLColorsRiverpodProviderInstance.secondaryColor,
      onSecondary: WMLColorsRiverpodProviderInstance.primaryColor,
      error: WMLColorsRiverpodProviderInstance.error,
    ),
    textTheme: _buildWMLTextTheme(base.textTheme),
    textSelectionTheme: TextSelectionThemeData(
      selectionColor: WMLColorsRiverpodProviderInstance.iconPrimaryColor,
    ),
    listTileTheme: ListTileThemeData(
      iconColor: WMLColorsRiverpodProviderInstance.iconPrimaryColor,
    ),
    iconTheme: IconThemeData(color: WMLColorsRiverpodProviderInstance.white),
    appBarTheme: base.appBarTheme.copyWith(
        iconTheme: IconThemeData(color: WMLColorsRiverpodProviderInstance.white),
        // toolbarHeight: theme.toolbarHeight,
        titleTextStyle: TextStyle(fontSize: WMLFontsRiverpodProviderInstance.extraLarge),
        backgroundColor: WMLColorsRiverpodProviderInstance.primaryColor),

    // for textbtn it would be textbtn theme

    inputDecorationTheme: InputDecorationTheme(
      floatingLabelStyle: TextStyle(
        color: WMLColorsRiverpodProviderInstance.iconPrimaryColor,
      ),
    ),
  );
}

TextTheme _buildWMLTextTheme(TextTheme base) {
  return base.copyWith(
    headlineSmall: base.headlineSmall!.copyWith(
      fontWeight: FontWeight.w500,
    ),
    titleLarge: base.titleLarge!.copyWith(
      fontSize: 18.0,
    ),
    bodySmall: base.bodySmall!.copyWith(
      fontWeight: FontWeight.w400,
      fontSize: 14.0,
    ),
    bodyLarge: base.bodyLarge!.copyWith(
      fontWeight: FontWeight.w500,
      fontSize: 16.0,
    ),
  );
}

class WMLThemeRiverpodProviderValue {
  ThemeMode currentThemeMode = ThemeMode.dark;
  final ThemeData lightTheme = _buildWMLLightTheme();
  final ThemeData darkTheme = _buildWMLDarkTheme();

  static ThemeData _buildWMLLightTheme() {
    final ThemeData base = ThemeData(useMaterial3: true, fontFamily: "Segoe UI", colorScheme: ColorScheme.light());
    return _buildWMLTheme(base);
  }

  static ThemeData _buildWMLDarkTheme() {
    final ThemeData base = ThemeData(useMaterial3: true, fontFamily: "Segoe UI", colorScheme: ColorScheme.dark());
    return _buildWMLTheme(base);
  }

  WMLThemeRiverpodProviderValue copyWith({
    ThemeMode? currentThemeMode,
    ThemeData? lightTheme,
    ThemeData? darkTheme,
  }) {
    return WMLThemeRiverpodProviderValue()..currentThemeMode = currentThemeMode ?? this.currentThemeMode;
  }
}

var WMLThemeRiverpodProviderInstance = WMLThemeRiverpodProviderValue();

class WMLThemeRiverpodNotifier extends Notifier<WMLThemeRiverpodProviderValue> {
  @override
  WMLThemeRiverpodProviderValue build() {
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    return WMLThemeRiverpodProviderInstance;
  }

  void updateCurrentThemeMode(ThemeMode? val) {
    state = state.copyWith(
        currentThemeMode: state.currentThemeMode =
            val ?? ((state.currentThemeMode == ThemeMode.dark) ? ThemeMode.light : ThemeMode.dark));
  }
}

final WMLThemeRiverpodProvider = NotifierProvider<WMLThemeRiverpodNotifier, WMLThemeRiverpodProviderValue>(() {
  return WMLThemeRiverpodNotifier();
});
