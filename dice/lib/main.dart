import 'package:flutter/material.dart';
import 'dart:math';

void main() {
  runApp(DiceApp());
}

class DiceApp extends StatefulWidget {
  @override
  _DiceAppState createState() => _DiceAppState();
}

class _DiceAppState extends State<DiceApp> {
  int left = 1, right = 1;
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        backgroundColor: Colors.indigo[900],
        appBar: AppBar(
          title: Text('Dice'),
          backgroundColor: Colors.deepOrange,
        ),
        body: Center(
          child: Row(
            children: [
              Expanded(
                child: FlatButton(
                  onPressed: () {
                    setState(() {
                      left = Random().nextInt(6) + 1;
                      right = Random().nextInt(6) + 1;
                    });
                  },
                  child: Image.asset('images/dice$left.png'),
                ),
              ),
              Expanded(
                child: FlatButton(
                  onPressed: () {
                    setState(() {
                      left = Random().nextInt(6) + 1;
                      right = Random().nextInt(6) + 1;
                    });
                  },
                  child: Image.asset('images/dice$right.png'),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
