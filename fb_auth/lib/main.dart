import 'package:flutter/material.dart';
import 'package:flutter_facebook_login/flutter_facebook_login.dart';
import 'package:http/http.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    return _MyAppState();
  }
}

class _MyAppState extends State<MyApp> {
  bool isLoggedin = false;
  Map userProfile;
  loginFb() {}
  logoutFb() {}
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text(
            "FB Authentication",
          ),
        ),
        body: isLoggedin
            ? OutlineButton(
                child: Text("Logout"),
                onPressed: () {
                  logoutFb();
                },
              )
            : OutlineButton(
                onPressed: () {
                  loginFb();
                },
                child: Text("Login with facebook"),
              ),
      ),
    );
  }
}
