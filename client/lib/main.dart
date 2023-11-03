import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:go_hub/screens/splash.dart';

void main() {
  runApp(
    Provider<String>(
      create: (context) => 'Hello World',
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const SplashScreen(),
    );
  }
}
