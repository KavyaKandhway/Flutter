import 'dart:convert';

import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter_facebook_auth/flutter_facebook_auth.dart';

void main() {
  runApp(LoginPage());
}

Future<UserCredential> signInWithFacebook() async {
  // Trigger the sign-in flow
  await Firebase.initializeApp();
  final AccessToken result = await FacebookAuth.instance.login();

  // Create a credential from the access token
  final FacebookAuthCredential facebookAuthCredential =
      FacebookAuthProvider.credential(result.token);
  print(result.applicationId);
  print(result.userId);
  print(result.token);
  // Once signed in, return the UserCredential
  return await FirebaseAuth.instance
      .signInWithCredential(facebookAuthCredential);
}

class LoginPage extends StatefulWidget {
  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  Widget build(BuildContext context) {
    return MaterialApp(
      home: SafeArea(
        child: Scaffold(
          body: Center(
            child: FlatButton(
              height: 100.0,
              minWidth: 100,
              color: Colors.purple,
              onPressed: () {
                signInWithFacebook().then((value) {
                  if (value != null) {
                    print(value.user.refreshToken);
                    print(value);
                  }
                });
              },
            ),
          ),
        ),
      ),
    );
  }
}
