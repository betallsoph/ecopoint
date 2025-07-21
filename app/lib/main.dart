import 'package:flutter/material.dart';
import 'screens/splash_screen.dart';
import 'screens/login_screen.dart';
import 'screens/home_screen.dart';
import 'screens/select_waste_screen.dart';
import 'screens/confirm_booking_screen.dart';
import 'screens/waiting_for_collector_screen.dart';
import 'screens/collector_on_way_screen.dart';
import 'screens/finish_collection_screen.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  final Color lightBlue = Color(0xFFB3E5FC);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'ecoPoint',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primaryColor: Color(0xFF388E3C), // xanh lá tươi
        scaffoldBackgroundColor: Colors.white,
      ),
      home: SplashScreen(),
      routes: {
        '/login': (context) => LoginScreen(),
        '/home': (context) => HomeScreen(),
        '/select-waste': (context) => SelectWasteScreen(),
        '/confirm': (context) => ConfirmBookingScreen(),
        '/waiting': (context) => WaitingForCollectorScreen(),
        '/onway': (context) => CollectorOnWayScreen(),
        '/finish': (context) => FinishCollectionScreen(),
      },
    );
  }
}
