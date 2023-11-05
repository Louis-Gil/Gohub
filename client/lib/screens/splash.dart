import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:go_hub/providers/auth.dart';
import 'package:go_hub/screens/home.dart';
import 'package:go_hub/screens/signin.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  SplashScreenState createState() => SplashScreenState();
}

class SplashScreenState extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      _checkLoginStatus();
    });
  }

  Future<void> _checkLoginStatus() async {
    await Future.delayed(const Duration(seconds: 2), () {});

    final isLoggedIn =
        Provider.of<AuthProvider>(context, listen: false).isLoggedIn;
    if (isLoggedIn) {
      _navigateToHome();
    } else {
      _navigateToLogin();
    }
  }

  void _navigateToHome() {
    Navigator.of(context)
        .pushReplacement(MaterialPageRoute(builder: (_) => const HomeScreen()));
  }

  void _navigateToLogin() {
    Navigator.of(context)
        .pushReplacement(MaterialPageRoute(builder: (_) => LoginScreen()));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        height: MediaQuery.of(context).size.height,
        width: MediaQuery.of(context).size.width,
        child: FittedBox(
          fit: BoxFit.cover,
          child: Image.asset('assets/images/splash.png'),
        ),
      ),
    );
  }
}
