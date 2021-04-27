import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        backgroundColor: Colors.teal,
        body: SafeArea(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              CircleAvatar(
                backgroundImage: AssetImage('images/before_cookie.jpg'),
                radius: 75.0,
              ),
              Text(
                'Shaswat Kandhway',
                style: TextStyle(
                  fontFamily: 'Pacifico',
                  color: Colors.white,
                  fontSize: 30.0,
                  fontWeight: FontWeight.bold,
                ),

              ),
              Text(
                  'REAPER',
                  style: TextStyle(
                    letterSpacing: 2.5,
                    color: Colors.teal[100],
                    fontSize: 15.0,
                    fontFamily: 'Source Sans Pro',
                ),
              ),
              SizedBox(
                child: Divider(
                  height: 10.0,
                  color: Colors.teal[100],
                ),
              ),
              Card(
                color: Colors.white,
                margin: EdgeInsets.fromLTRB(20.0, 20.0, 20.0, 0.0),
                child: Padding(
                  padding: const EdgeInsets.all(1.0),
                  child: ListTile(
                    leading: Icon(
                      Icons.call,
                      color: Colors.teal,
                    ),
                    title: Text(
                      '+91-747-986-5304',
                      style: TextStyle(
                        color: Colors.teal[400],
                        fontFamily: 'Source Sans Pro',
                        fontSize: 18.0,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                  ),

                ),
              ),
              Card(
                color: Colors.white,
                margin: EdgeInsets.fromLTRB(20.0, 20.0, 20.0, 0.0),
                child: Padding(
                  padding: const EdgeInsets.all(1.0),
                  child: ListTile(
                    leading: Icon(
                      Icons.email,
                      color: Colors.teal,
                    ),
                    title: Text(
                      'shaswatkandhway@gmail.com',
                      style: TextStyle(
                        color: Colors.teal[400],
                        fontFamily: 'Source Sans Pro',
                        fontSize: 18.0,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                  ),

                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
