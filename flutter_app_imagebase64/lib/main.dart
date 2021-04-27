import 'package:flutter/material.dart';

import 'dart:async';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:typed_data';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  State createState() => new MyHomePageState();
}

class MyHomePageState extends State<MyHomePage> {
  String _base64;

  @override
  void initState() {
    super.initState();
    (() async {
      http.Response response = await http.get(
        'https://hdqwalls.com/wallpapers/kda-akali-5k-og.jpg',
      );
      if (mounted) {
        setState(() {
          _base64 = Base64Encoder().convert(response.bodyBytes);
          print("fkhjffffffffffffffffffffffffffffffffffffffff");
          print(_base64.length);
        });
      }
    })();
  }

  @override
  Widget build(BuildContext context) {
    if (_base64 == null) return new Container();
    Uint8List bytes = Base64Codec().decode(_base64);
    return new Scaffold(
      appBar: new AppBar(title: new Text('Example App')),
      body: Image.memory(
        bytes,
        fit: BoxFit.cover,
      ),
    );
  }
}
