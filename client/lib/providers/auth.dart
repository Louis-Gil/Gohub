import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AuthProvider with ChangeNotifier {
  bool _isLoggedIn = false;
  bool _isInitialized = false;

  bool get isLoggedIn => _isLoggedIn;
  bool get isInitialized => _isInitialized;

  AuthProvider() {
    _loadFromPrefs();
  }

  Future<void> _loadFromPrefs() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    _isLoggedIn = prefs.getBool('isLoggedIn') ?? false;
    Future.microtask(() {
      _isInitialized = true;
      notifyListeners();
    });
  }

  Future<void> setLoginStatus(bool status) async {
    _isLoggedIn = status;
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setBool('isLoggedIn', status);
    notifyListeners();
  }
}