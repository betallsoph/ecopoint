import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import 'firebase_options.dart';
import 'screens/auth_wrapper.dart';
import 'services/graphql_service.dart';
import 'screens/login_screen.dart';
import 'screens/home_screen.dart';
import 'screens/order_details_screen.dart';
import 'screens/navigation_screen.dart';
import 'screens/complete_order_screen.dart';
import 'screens/profile_screen.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  
  // Initialize GraphQL service
  GraphQLService.instance.initialize();
  
  runApp(CollectorApp());
}

class CollectorApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return GraphQLProvider(
      client: ValueNotifier(GraphQLService.instance.client),
      child: MaterialApp(
        title: 'ecoPoint Collector',
        debugShowCheckedModeBanner: false,
        theme: ThemeData(
          primaryColor: Color(0xFF388E3C), // same green as main app
          scaffoldBackgroundColor: Colors.white,
          fontFamily: 'Montserrat',
        ),
        home: AuthWrapper(),
        routes: {
          '/login': (context) => LoginScreen(),
          '/home': (context) => HomeScreen(),
          '/order-details': (context) => OrderDetailsScreen(),
          '/navigation': (context) => NavigationScreen(),
          '/complete': (context) => CompleteOrderScreen(),
          '/profile': (context) => ProfileScreen(),
        },
      ),
    );
  }
}
