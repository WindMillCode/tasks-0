// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables, must_be_immutable, unused_local_variable, unused_catch_stack

import 'package:tuli/shared/widgets/button_zero/button_zero_widget.dart';
import 'package:tuli/shared/widgets/input_zero/input_zero_widget.dart';
import 'package:tuli/shared/widgets/overlay_zero/overlay_zero_riverpod_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:windmillcode_flutter_translate/flutter_translate.dart';
import 'package:form_builder_validators/form_builder_validators.dart';
import 'multi_input_one_widget_riverpod_provider.dart';
import 'package:tuli/util/riverpod_providers/wml_colors/wml_colors.dart';
import 'package:tuli/util/riverpod_providers/wml_fonts/wml_fonts.dart';
import 'package:tuli/util/riverpod_providers/wml_nav/wml_nav.dart';
import 'package:tuli/util/riverpod_providers/wml_spacing/wml_spacing.dart';

class MultiInputOneWidget extends ConsumerStatefulWidget {
  MultiInputOneWidget(
      {super.key,
      this.useListViewForEntries = true,
      this.startingAmount = 1,
      Future<OverlayZeroProcessRequestPredicateResult> Function(List<dynamic> formValues)? onSubmit,
      void Function()? onReady,
      this.limit = 51,
      this.minEntries = 1,
      String? limitText,
      Map<String, dynamic> Function(int index)? createControllerForEntry,
      Widget? Function(int index, Map<String, dynamic> controller)? addEntry})
      : onSubmit = onSubmit ?? ((List<dynamic> formValues) async => OverlayZeroProcessRequestPredicateResult()),
        onReady = onReady ?? (() {}),
        limitText = limitText ?? translate("MultiInputOne.WMLNotifyOne.limit"),
        createControllerForEntry = createControllerForEntry ??
            ((int index) {
              return {"default": TextEditingController()};
            }),
        addEntry =
            // ignore: body_might_complete_normally_nullable
            addEntry ?? ((int index, Map<String, dynamic> controller) {});

  final Future<OverlayZeroProcessRequestPredicateResult> Function(List<dynamic> formValues) onSubmit;
  final void Function() onReady;
  final Widget? Function(int index, Map<String, dynamic> controller) addEntry;
  final bool useListViewForEntries;
  final int startingAmount;
  int limit;
  int minEntries;
  String limitText;
  final Map<String, dynamic> Function(int index) createControllerForEntry;

  @override
  ConsumerState<MultiInputOneWidget> createState() => _MultiInputOneState();
}

class _MultiInputOneState extends ConsumerState<MultiInputOneWidget> {
  List<dynamic> entries = [];
  Widget container = Column();
  final formKey = GlobalKey<FormState>();

  bool initialEntryAdded = false;
  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    if (!initialEntryAdded) {
      initialEntryAdded = true;
      for (var i = 0; i < widget.startingAmount; i++) {
        addEntry();
      }
      widget.onReady();
    }
  }

  removeEntry(int index) {
    entries.removeAt(index);
    updateEntryContainer();
    if (entries.length < widget.minEntries) {
      addEntry();
    }
  }

  addEntry() {
    if (widget.limit < entries.length + 1) {
      showSnackBarMsg(widget.limitText);
      return;
    }
    var controller = widget.createControllerForEntry(entries.length);
    var myWidget = widget.addEntry(entries.length, controller) ??
        InputZeroWidget(
          label: Text(translate("MultiInputOne.mainForm.valueField.label")),
          textEditingCtrlr: controller["default"],
          validator: FormBuilderValidators.required(errorText: translate("global.errorRequired")),
        );
    var obj = {"controller": controller, "widget": myWidget};
    entries.add(obj);
    updateEntryContainer();
  }

  void updateEntryContainer() {
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);

    setState(() {
      if (widget.useListViewForEntries) {
        container = Flexible(
          flex: 5,
          fit: FlexFit.tight,
          child: Form(
              key: formKey,
              child: ListView.builder(
                shrinkWrap: true,
                itemCount: entries.length,
                itemBuilder: (context, index) {
                  return Column(
                    crossAxisAlignment: CrossAxisAlignment.end,
                    children: [
                      entries[index]["widget"],
                      SizedBox(
                        height: wmlSpacing.medium,
                      ),
                      ButtonZeroWidget(
                        child: translate("global.remove"),
                        onPressed: () {
                          removeEntry(index);
                        },
                      ),
                    ],
                  );
                },
              )),
        );
      } else {
        container = Form(
          key: formKey,
          child: Column(
            children: entries.map((entry) {
              int index = entries.indexOf(entry);
              return Column(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: [
                  entry["widget"],
                  SizedBox(
                    height: wmlSpacing.medium,
                  ),
                  ButtonZeroWidget(
                    child: translate("global.remove"),
                    onPressed: () {
                      removeEntry(index);
                    },
                  ),
                ],
              );
            }).toList(),
          ),
        );
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    final rp = ref.watch(MultiInputOneRiverpodProvider);
    final wmlColors = ref.watch(WMLColorsRiverpodProvider);
    final wmlSpacing = ref.watch(WMLSpacingRiverpodProvider);
    final wmlFonts = ref.watch(WMLFontsRiverpodProvider);
    final wmlNav = ref.watch(WMLNavRiverpodProvider);
    final contextHeight = MediaQuery.of(context).size.height;
    final contextWidth = MediaQuery.of(context).size.width;

    return Column(
      children: [
        container,
        SizedBox(
          height: wmlSpacing.large,
        ),
        Align(
          alignment: Alignment.bottomCenter,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              Container(
                width: contextWidth,
                child: ButtonZeroWidget(
                  onPressed: addEntry,
                  type: ButtonZeroWidgetType.quaternary,
                  child: translate("global.add"),
                ),
              ),
              SizedBox(
                height: wmlSpacing.medium,
              ),
              Container(
                width: contextWidth,
                child: ButtonZeroWidget(
                  onPressed: () async {
                    if (!formKey.currentState!.validate()) {
                    } else {
                      return await processRequest(context, ref, () {
                        var formValues = entries.map((e) {
                          return {"value": e};
                        }).toList();
                        return widget.onSubmit(formValues);
                      });
                    }
                  },
                  type: ButtonZeroWidgetType.quaternary,
                  child: translate("global.submitBtn"),
                ),
              ),
            ],
          ),
        )
      ],
    );
  }
}
