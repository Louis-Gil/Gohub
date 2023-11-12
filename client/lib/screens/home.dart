import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:gohub/providers/auth.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final authProvider = Provider.of<AuthProvider>(context, listen: false);
    final userEmail = authProvider.googleUser?.email ?? 'No email found';

    return Scaffold(
      appBar: AppBar(
        title: const Text('Home Screen'),
      ),
      body: Center(
        child: Text('Welcome to the home screen!, your Email is: $userEmail'),
      ),
    );
  }
}