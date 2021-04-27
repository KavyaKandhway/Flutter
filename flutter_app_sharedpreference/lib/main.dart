import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  _saveData() async {
    // we make sure we have an instance
    final prefs = await SharedPreferences.getInstance();
    prefs.setInt('yourKey', 100);
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: SafeArea(
          child: Center(
            child: Column(
              children: [
                FlatButton(
                  child: Text("SET"),
                  color: Colors.pink,
                  onPressed: () async {
                    final prefs = await SharedPreferences.getInstance();

                    prefs.setStringList('user1', ["abcd", "pqer"]);
                  },
                ),
                FlatButton(
                  child: Text("SET"),
                  color: Colors.pink,
                  onPressed: () async {
                    final prefs = await SharedPreferences.getInstance();
                    print("before");
                    List<String> ss = prefs.getStringList('user1') ?? [];
                    print(prefs.getString('user') ?? "");
                    print("abssssss");
                    for (int i = 0; i < ss.length; i++) {
                      print(ss[i]);
                    }
                    ss.add("pqrs");
                    prefs.setStringList("user1", ss);
                  },
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
