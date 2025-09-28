import 'package:shared_preferences/shared_preferences.dart';

class DevModeService {
  static const String _devModeKey = 'dev_mode_enabled';
  static const String _mockUserKey = 'mock_user_data';

  static Future<bool> isDevModeEnabled() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getBool(_devModeKey) ?? false;
  }

  static Future<void> setDevModeEnabled(bool enabled) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setBool(_devModeKey, enabled);
  }

  static Future<Map<String, dynamic>?> getMockUser() async {
    final prefs = await SharedPreferences.getInstance();
    final userJson = prefs.getString(_mockUserKey);
    if (userJson != null) {
      // In a real app, you would parse JSON here
      return {
        'uid': 'dev-user-123',
        'email': 'dev-user@ecopoint.com',
        'displayName': 'Dev User',
        'role': 'USER',
        'photoURL': null,
      };
    }
    return null;
  }

  static Future<void> setMockUser(Map<String, dynamic> user) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString(_mockUserKey, user.toString());
  }

  static Future<void> clearMockUser() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove(_mockUserKey);
  }

  // Mock data for development
  static List<Map<String, dynamic>> getMockBookings() {
    return [
      {
        'id': 'booking-1',
        'wasteType': 'PLASTIC',
        'quantity': 5,
        'description': 'Plastic bottles and containers',
        'pickupAddress': {
          'street': '123 Main St',
          'city': 'Ho Chi Minh City',
          'state': 'HCMC',
          'zipCode': '70000',
          'latitude': 10.7769,
          'longitude': 106.7009,
        },
        'scheduledDate': '2024-01-15',
        'scheduledTime': '10:00',
        'status': 'PENDING',
        'estimatedPrice': 25.00,
        'createdAt': DateTime.now().toIso8601String(),
      },
      {
        'id': 'booking-2',
        'wasteType': 'PAPER',
        'quantity': 3,
        'description': 'Old newspapers and magazines',
        'pickupAddress': {
          'street': '456 Oak Ave',
          'city': 'Ho Chi Minh City',
          'state': 'HCMC',
          'zipCode': '70000',
          'latitude': 10.7769,
          'longitude': 106.7009,
        },
        'scheduledDate': '2024-01-16',
        'scheduledTime': '14:00',
        'status': 'CONFIRMED',
        'estimatedPrice': 15.00,
        'createdAt': DateTime.now().toIso8601String(),
      },
    ];
  }

  static List<Map<String, dynamic>> getMockAddresses() {
    return [
      {
        'id': 'addr-1',
        'street': '123 Main St',
        'city': 'Ho Chi Minh City',
        'state': 'HCMC',
        'zipCode': '70000',
        'latitude': 10.7769,
        'longitude': 106.7009,
        'isDefault': true,
        'createdAt': DateTime.now().toIso8601String(),
      },
      {
        'id': 'addr-2',
        'street': '456 Oak Ave',
        'city': 'Ho Chi Minh City',
        'state': 'HCMC',
        'zipCode': '70000',
        'latitude': 10.7769,
        'longitude': 106.7009,
        'isDefault': false,
        'createdAt': DateTime.now().toIso8601String(),
      },
    ];
  }
}