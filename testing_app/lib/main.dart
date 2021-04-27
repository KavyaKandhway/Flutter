import 'package:flutter/material.dart';

void main() {
  runApp(
    MyApp(),
  );
}
class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
    home: Scaffold(
    appBar: AppBar(
      title: Text('Baby(Netflix)'),
      backgroundColor: Colors.pink[600],
    ),
    backgroundColor: Colors.pink[100],
    body: Center(
      child: ClipRRect(
        child: Image(
          image: AssetImage('images/baby_netflix.jpg'),
          fit: BoxFit.cover,
          height: double.infinity,
          width: double.infinity,
        ),
      ),
    ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Colors.pink[600],
        onPressed: () {print('Clicked');},
      ),
  ),
);
  }
}

