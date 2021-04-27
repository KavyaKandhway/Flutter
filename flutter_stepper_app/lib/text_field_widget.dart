import 'package:flutter/material.dart';

class TextFieldWidget extends StatelessWidget {
  TextFieldWidget({@required this.label, this.initialize, this.editable});
  final String label;
  String initialize;

  bool editable;
  TextEditingController textEditingController = TextEditingController();
  void initState() {
    textEditingController = TextEditingController(text: initialize);
  }

  @override
  Widget build(BuildContext context) {
    textEditingController.text = initialize;
    return Padding(
      padding: const EdgeInsets.fromLTRB(10.0, 8.0, 10.0, 10.0),
      child: TextField(
        enabled: editable,
        controller: textEditingController,
        decoration: InputDecoration(
          border: OutlineInputBorder(
            borderRadius: BorderRadius.zero,
          ),
          labelText: label,
        ),
      ),
    );
  }
}
