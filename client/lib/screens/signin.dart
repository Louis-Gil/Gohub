import 'package:flutter/material.dart';

import 'package:google_sign_in/google_sign_in.dart';

class LoginScreen extends StatefulWidget {
  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final GoogleSignIn _googleSignIn = GoogleSignIn();

  Future<void> _handleSignIn() async {
    try {
      await _googleSignIn.signIn();
      // on success, you can get the user's details as follows:
      final GoogleSignInAccount? user = _googleSignIn.currentUser;
      print('User name: ${user?.displayName}');
      // You can now use this user information for your app's needs
    } catch (error) {
      print('Sign in failed: $error');
      // Handle the error scenario
    }
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: ElevatedButton(
        onPressed: _handleSignIn,
        child: Text('Sign in with Google'),
      ),
    );
  }
}