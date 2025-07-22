import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';
import 'firebase_options.dart';
import 'screens/auth_wrapper.dart'; // Import the new wrapper
import 'screens/splash_screen.dart';
import 'screens/login_screen.dart';
import 'screens/home_screen.dart';
import 'screens/select_waste_screen.dart';
import 'screens/confirm_booking_screen.dart';
import 'screens/waiting_for_collector_screen.dart';
import 'screens/collector_on_way_screen.dart';
import 'screens/finish_collection_screen.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'ecoPoint',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primaryColor: Color(0xFF388E3C), // xanh lá tươi
        scaffoldBackgroundColor: Colors.white,
      ),
      home: AuthWrapper(), // Use AuthWrapper as the home screen
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
