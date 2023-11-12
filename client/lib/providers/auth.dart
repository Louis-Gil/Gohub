import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'package:gohub/screens/home.dart';
import 'package:gohub/screens/login.dart';

class AuthProvider with ChangeNotifier {
  bool _isLoggedIn = false;

  bool get isLoggedIn => _isLoggedIn;

  AuthProvider() {
    initialize();
  }

  Future<void> initialize() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    _isLoggedIn = prefs.getBool('isLoggedIn') ?? false;
    notifyListeners();
  }

  Future<void> setLoginStatus(bool status) async {
    _isLoggedIn = status;
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
