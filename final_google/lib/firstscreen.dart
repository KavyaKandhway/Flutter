import 'package:flutter/material.dart';
import 'authGoogle.dart';
import 'main.dart';

class FirstScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: FlatButton(
          onPressed: () {
            signOutGoogle();
            Navigator.of(context).pop();
          },
          child: Text(
            'Logout',
          ),
        ),
      ),
    );
  }
}
