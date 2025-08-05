import 'package:graphql_flutter/graphql_flutter.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class GraphQLService {
  static GraphQLService? _instance;
  static GraphQLService get instance => _instance ??= GraphQLService._();
  
  GraphQLService._();
  
  final _storage = const FlutterSecureStorage();
  late GraphQLClient _client;
  
  // GraphQL endpoint - thay đổi IP này nếu cần
  static const String _endpoint = 'http://192.168.21.71:4000/query';
  
  void initialize() {
    final HttpLink httpLink = HttpLink(_endpoint);
    
    final AuthLink authLink = AuthLink(
      getToken: () async {
        final token = await _storage.read(key: 'access_token');
        return token != null ? 'Bearer $token' : null;
      },
    );
    
    final Link link = authLink.concat(httpLink);
    
    _client = GraphQLClient(
      link: link,
      cache: GraphQLCache(store: InMemoryStore()),
    );
  }
  
  GraphQLClient get client => _client;
  
  // Save tokens after authentication
  Future<void> saveTokens(String accessToken, String refreshToken) async {
    await _storage.write(key: 'access_token', value: accessToken);
    await _storage.write(key: 'refresh_token', value: refreshToken);
  }
  
  // Get access token
  Future<String?> getAccessToken() async {
    return await _storage.read(key: 'access_token');
  }
  
  // Clear tokens (logout)
  Future<void> clearTokens() async {
    await _storage.delete(key: 'access_token');
    await _storage.delete(key: 'refresh_token');
  }
  
  // Check if user is authenticated
  Future<bool> isAuthenticated() async {
    final token = await getAccessToken();
    return token != null && token.isNotEmpty;
  }
}

// GraphQL Mutations - Same as customer app
class GraphQLMutations {
  static const String googleSignIn = '''
    mutation GoogleSignIn(\$input: GoogleSignInInput!) {
      googleSignIn(input: \$input) {
        accessToken
        refreshToken
        user {
          id
          displayName
          email
          userType
          isActive
        }
      }
    }
  ''';
  
  static const String acceptOrder = '''
    mutation AcceptOrder(\$orderID: ID!) {
      acceptOrder(orderID: \$orderID) {
        id
        status
        collectorID
      }
    }
  ''';
  
  static const String completeOrder = '''
    mutation CompleteOrder(\$orderID: ID!, \$input: CompleteOrderInput!) {
      completeOrder(orderID: \$orderID, input: \$input) {
        id
        status
        actualWeight
        completedTime
      }
    }
  ''';
}

// GraphQL Queries - Focused on collector needs
class GraphQLQueries {
  static const String me = '''
    query Me {
      me {
        id
        displayName
        email
        userType
        isActive
        rating
        isOnline
      }
    }
  ''';
  
  static const String myOrders = '''
    query MyOrders {
      myOrders {
        id
        customerID
        status
        wasteTypes
        estimatedWeight
        actualWeight
        pickupAddress {
          street
          district
          city
          lat
          lng
        }
        scheduledTime
        completedTime
        notes
        payment {
          amount
          currency
          method
          isPaid
        }
        createdAt
        updatedAt
      }
    }
  ''';
  
  static const String availableOrders = '''
    query AvailableOrders {
      availableOrders {
        id
        customerID
        status
        wasteTypes
        estimatedWeight
        pickupAddress {
          street
          district
          city
          lat
          lng
        }
        scheduledTime
        notes
        payment {
          amount
          currency
          method
        }
        createdAt
      }
    }
  ''';
  
  static const String userStats = '''
    query UserStats {
      userStats {
        totalOrders
        completedOrders
        rating
        totalEarnings
        isOnline
      }
    }
  ''';
}