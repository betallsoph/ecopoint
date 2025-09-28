class Booking {
  final String id;
  final String userId;
  final String? collectorId;
  final WasteType wasteType;
  final double quantity;
  final String description;
  final Address pickupAddress;
  final DateTime scheduledDate;
  final String scheduledTime;
  final BookingStatus status;
  final double? estimatedPrice;
  final double? finalPrice;
  final String? notes;
  final List<String> images;
  final DateTime createdAt;
  final DateTime updatedAt;

  Booking({
    required this.id,
    required this.userId,
    this.collectorId,
    required this.wasteType,
    required this.quantity,
    required this.description,
    required this.pickupAddress,
    required this.scheduledDate,
    required this.scheduledTime,
    required this.status,
    this.estimatedPrice,
    this.finalPrice,
    this.notes,
    required this.images,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Booking.fromJson(Map<String, dynamic> json) {
    return Booking(
      id: json['id'] ?? '',
      userId: json['user_id'] ?? '',
      collectorId: json['collector_id'],
      wasteType: WasteType.fromString(json['waste_type'] ?? 'OTHER'),
      quantity: (json['quantity'] ?? 0).toDouble(),
      description: json['description'] ?? '',
      pickupAddress: Address.fromJson(json['pickup_address'] ?? {}),
      scheduledDate: DateTime.parse(json['scheduled_date'] ?? DateTime.now().toIso8601String()),
      scheduledTime: json['scheduled_time'] ?? '',
      status: BookingStatus.fromString(json['status'] ?? 'PENDING'),
      estimatedPrice: json['estimated_price']?.toDouble(),
      finalPrice: json['final_price']?.toDouble(),
      notes: json['notes'],
      images: List<String>.from(json['images'] ?? []),
      createdAt: DateTime.parse(json['created_at'] ?? DateTime.now().toIso8601String()),
      updatedAt: DateTime.parse(json['updated_at'] ?? DateTime.now().toIso8601String()),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_id': userId,
      'collector_id': collectorId,
      'waste_type': wasteType.value,
      'quantity': quantity,
      'description': description,
      'pickup_address': pickupAddress.toJson(),
      'scheduled_date': scheduledDate.toIso8601String().split('T')[0],
      'scheduled_time': scheduledTime,
      'status': status.value,
      'estimated_price': estimatedPrice,
      'final_price': finalPrice,
      'notes': notes,
      'images': images,
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt.toIso8601String(),
    };
  }
}

class Address {
  final String street;
  final String city;
  final String state;
  final String zipCode;
  final double latitude;
  final double longitude;

  Address({
    required this.street,
    required this.city,
    required this.state,
    required this.zipCode,
    required this.latitude,
    required this.longitude,
  });

  factory Address.fromJson(Map<String, dynamic> json) {
    return Address(
      street: json['street'] ?? '',
      city: json['city'] ?? '',
      state: json['state'] ?? '',
      zipCode: json['zip_code'] ?? '',
      latitude: (json['latitude'] ?? 0).toDouble(),
      longitude: (json['longitude'] ?? 0).toDouble(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'street': street,
      'city': city,
      'state': state,
      'zip_code': zipCode,
      'latitude': latitude,
      'longitude': longitude,
    };
  }

  String get fullAddress => '$street, $city, $state $zipCode';
}

enum WasteType {
  plastic('PLASTIC'),
  paper('PAPER'),
  glass('GLASS'),
  metal('METAL'),
  organic('ORGANIC'),
  electronic('ELECTRONIC'),
  other('OTHER');

  const WasteType(this.value);
  final String value;

  static WasteType fromString(String value) {
    switch (value.toUpperCase()) {
      case 'PLASTIC':
        return WasteType.plastic;
      case 'PAPER':
        return WasteType.paper;
      case 'GLASS':
        return WasteType.glass;
      case 'METAL':
        return WasteType.metal;
      case 'ORGANIC':
        return WasteType.organic;
      case 'ELECTRONIC':
        return WasteType.electronic;
      default:
        return WasteType.other;
    }
  }
}

enum BookingStatus {
  pending('PENDING'),
  confirmed('CONFIRMED'),
  assigned('ASSIGNED'),
  inProgress('IN_PROGRESS'),
  completed('COMPLETED'),
  cancelled('CANCELLED'),
  failed('FAILED');

  const BookingStatus(this.value);
  final String value;

  static BookingStatus fromString(String value) {
    switch (value.toUpperCase()) {
      case 'PENDING':
        return BookingStatus.pending;
      case 'CONFIRMED':
        return BookingStatus.confirmed;
      case 'ASSIGNED':
        return BookingStatus.assigned;
      case 'IN_PROGRESS':
        return BookingStatus.inProgress;
      case 'COMPLETED':
        return BookingStatus.completed;
      case 'CANCELLED':
        return BookingStatus.cancelled;
      case 'FAILED':
        return BookingStatus.failed;
      default:
        return BookingStatus.pending;
    }
  }
}