import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'authGoogle.dart';
import 'main.dart';

User userInfo;
String x;

class FirstScreen extends StatefulWidget {
  @override
  _FirstScreenState createState() => _FirstScreenState();
}

class _FirstScreenState extends State<FirstScreen> {
  //String x;
  TextEditingController fcontroler = TextEditingController();
  Future<User> user = signInWithGoogle().then((result) {
    if (result != null) {
      userInfo = result;
    }
    return null;
  });
  String x = "kkk";

  void initState() {
    fcontroler = TextEditingController(text: x);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
          child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          FlatButton(
            onPressed: () {
              signOutGoogle();
              Navigator.of(context).pop();
            },
            child: Text(
              'Logout',
            ),
          ),
          TextField(
            controller: fcontroler,
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.zero,
              ),
              labelText: 'TEST',
            ),
          )
        ],
      )),
    );
  }
}
