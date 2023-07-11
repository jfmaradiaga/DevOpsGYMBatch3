import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class HelloWorldScreen extends StatefulWidget {
  @override
  _HelloWorldScreenState createState() => _HelloWorldScreenState();
}

class _HelloWorldScreenState extends State<HelloWorldScreen> {
  TextEditingController nameController = TextEditingController();
  String message = '';

  Future<void> fetchData() async {
    final url = Uri.parse('http://192.168.122.129:8080/api/hello');
    final response = await http.post(url, body: {'name': nameController.text});
    final responseData = json.decode(response.body);

    setState(() {
      message = responseData['message'];
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Hello World'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextField(
              controller: nameController,
              decoration: InputDecoration(
                labelText: 'Enter your name',
              ),
            ),
            SizedBox(height: 16),
            ElevatedButton(
              onPressed: fetchData,
              child: Text('Submit'),
            ),
            SizedBox(height: 16),
            Text(message),
          ],
        ),
      ),
    );
  }
}
