import 'package:flutter/material.dart';

class Label extends StatelessWidget {
  Label(
      {@required this.title,
      @required this.labelColor,
      @required this.labelSize,
      this.labelWeight});
  final String title;
  final int labelColor;
  final double labelSize;
  FontWeight labelWeight = null;
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Text(
        title,
        style: TextStyle(
          color: Color(labelColor),
          fontSize: labelSize,
          fontWeight: labelWeight,
          fontFamily: 'Jost',
        ),
      ),
      width: MediaQuery.of(context).size.width,
    );
  }
}
