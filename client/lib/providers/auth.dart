import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'package:gohub/screens/home.dart';
import 'package:gohub/screens/login.dart';

class AuthProvider with ChangeNotifier {
  bool _isLoggedIn = false;
  GoogleSignInAccount? _googleUser;

  bool get isLoggedIn => _isLoggedIn;
  GoogleSignInAccount? get googleUser => _googleUser;

  AuthProvider() {
    initialize();
  }

  Future<void> initialize() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    // _isLoggedIn = prefs.getBool('isLoggedIn') ?? false;
    _isLoggedIn = false;
    notifyListeners();
  }

  Future<void> setLoginStatus(
      bool status, GoogleSignInAccount? googleUser) async {
    _isLoggedIn = status;
    _googleUser = googleUser;
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setBool('isLoggedIn', status);
    notifyListeners();
  }

  void _navigateToHome(BuildContext context) {
    if (_isLoggedIn) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        Navigator.of(context).pushReplacement(
            MaterialPageRoute(builder: (_) => const HomeScreen()));
      });
    } else {
      _navigateToLogin(context);
    }
  }

  void _navigateToLogin(BuildContext context) {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      Navigator.of(context).pushReplacement(
          MaterialPageRoute(builder: (_) => const LoginScreen()));
    });
  }

  void navigateToHome(BuildContext context) {
    _navigateToHome(context);
  }

  void navigateToLogin(BuildContext context) {
    _navigateToLogin(context);
  }
}
