import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:gohub/providers/auth.dart';

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
          final authProvider =
              Provider.of<AuthProvider>(context, listen: false);
          authProvider.navigateToHome(context);

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
}
