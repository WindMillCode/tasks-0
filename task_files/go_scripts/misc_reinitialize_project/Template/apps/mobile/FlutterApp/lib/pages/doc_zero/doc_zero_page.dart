// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:[PROJECT_NAME]/shared/widgets/bullet_point_zero/bullet_point_zero_widget.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:[PROJECT_NAME]/util/riverpod_providers/wml_spacing/wml_spacing.dart';
import 'package:flutter/material.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import 'doc_zero_page_riverpod_provider.dart';

class DocZeroPageWidget extends ConsumerStatefulWidget {
  const DocZeroPageWidget({super.key});

  @override
  ConsumerState<DocZeroPageWidget> createState() => _DocZeroPageState();
}

class _DocZeroPageState extends ConsumerState<DocZeroPageWidget> {
  late Future<List<dynamic>> sectionsFuture;

  Future<List<dynamic>> getPriavcyPolicySections() async {
    final response = getValueAtKeyPath("PrivacyPolicy.sections");
    return response as List<dynamic>;
  }


  @override
  Widget build(BuildContext context) {
    String currentRoute = GoRouter.of(context).routeInformationProvider.value.uri.path;
    final rp = ref.watch(DocZeroPageRiverpodProvider);
    final contextWidth = MediaQuery.of(context).size.width;
    final contextHeight = MediaQuery.of(context).size.height;
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);

    Map<String, dynamic>? options;

    if (currentRoute == wmlNav.route["path"]) {
      options = {
        "title": translate("PrivacyPolicy.title"),
        "prefix": "PrivacyPolicy",
        "sections": [
          "bold_text",
          "text",
          "text",
          "text",
          "text",
          "text",
          "text",
          "section",
          "list",
          "text",
          "list",
          "text",
          "section",
          "list",
          "section",
          "list",
          "section",
          "list",
          "section",
          "list",
          "section",
          "list",
          "section",
          "list",
          "section",
          "list",
          "section",
          "list"
        ]
      };
      sectionsFuture = getPriavcyPolicySections();
    }

    return FutureBuilder(
        future: sectionsFuture,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<dynamic> sections = snapshot.data as List<dynamic>;

            return Container(
              width: contextWidth,
              height: MediaQuery.of(context).size.height,
              decoration: BoxDecoration(
                color: wmlColors.black,
              ),
              child: Padding(
                padding: EdgeInsets.all(wmlSpacing.large),
                child: SingleChildScrollView(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(options!["title"]!,
                          style: TextStyle(
                              fontWeight: FontWeight.bold, fontSize: wmlFonts.heading2, color: wmlColors.white)),
                      SizedBox(height: wmlSpacing.extraLarge),
                      ...List.generate(sections.length, (index) {
                        String type = options!["sections"][index];
                        if (type == "text") {
                          return Text(translate(
                            "${options["prefix"]}.sections.$index",
                          ));
                        } else if (type == "bold_text") {
                          return Text(
                            translate(
                              "${options["prefix"]}.sections.$index",
                            ),
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                            ),
                          );
                        } else if (type == "section") {
                          return Column(
                            children: [
                              SizedBox(height: wmlSpacing.gigantic),
                              Text(
                                  translate(
                                    "${options["prefix"]}.sections.$index",
                                  ),
                                  style: TextStyle(
                                      fontWeight: FontWeight.bold,
                                      fontSize: wmlFonts.heading4,
                                      color: wmlColors.white)),
                              SizedBox(height: wmlSpacing.medium),
                            ],
                          );
                        } else if (type == "list") {
                          List<dynamic> items = sections[index];

                          return Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              SizedBox(height: wmlSpacing.medium),
                              ...List.generate(items.length, (index1) {
                                return BulletPointZeroWidget(
                                    child: translate(
                                  "${options!["prefix"]}.sections.$index.$index1",
                                ));
                              }),
                              SizedBox(height: wmlSpacing.medium),
                            ],
                          );
                        } else {
                          return Text(translate(
                            "${options["prefix"]}.sections.$index",
                          ));
                        }
                      })
                    ],
                  ),
                ),
              ),
            );
          }
          return CircularProgressIndicator();
        });
  }
}
