import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:gohub/providers/auth.dart';
import 'package:gohub/screens/home.dart';
import 'package:gohub/screens/login.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  SplashScreenState createState() => SplashScreenState();
}

class SplashScreenState extends State<SplashScreen> {
  late Future _initialization;

  @override
  void initState() {
    super.initState();
    _initialization = Future.wait([
      Provider.of<AuthProvider>(context, listen: false).initialize(),
      Future.delayed(const Duration(seconds: 1)),
    ]);
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _initialization,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
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
        } else if (snapshot.connectionState == ConnectionState.done) {
          if (Provider.of<AuthProvider>(context, listen: false).isLoggedIn) {
            _navigateToHome();
          } else {
            _navigateToLogin();
          }
          return Container();
        } else {
          return const Scaffold(
            body: Center(
              child: Text('failed to load'),
            ),
          );
        }
      },
    );
  }

  void _navigateToHome() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (mounted) {
        Navigator.of(context).pushReplacement(
            MaterialPageRoute(builder: (_) => const HomeScreen()));
      }
    });
  }

  void _navigateToLogin() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (mounted) {
        Navigator.of(context)
            .pushReplacement(MaterialPageRoute(builder: (_) => const LoginScreen()));
      }
    });
  }
}
