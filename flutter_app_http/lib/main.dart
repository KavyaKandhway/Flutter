import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(MyApp());
}

Future<Response> postRequest() async {
  final http.Response response = await http.get(
    'http://api.furrble.com',
  );
  print(response.statusCode);
  print(response.body);
  if (response.statusCode == 200) {
    return Response.fromJson(jsonDecode(response.body));
  } else {
    throw Exception('Failed to create album.');
  }
}

class Response {
  final String userID;
  final String message;
  Response({this.message, this.userID});

  factory Response.fromJson(Map<String, dynamic> json) {
    return Response(userID: json['userID'], message: json['message']);
  }
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: FlatButton(
          height: 100,
          minWidth: 100,
          color: Colors.blueAccent,
          onPressed: () {
            postRequest();
          },
        ),
      ),
    );
  }
}
