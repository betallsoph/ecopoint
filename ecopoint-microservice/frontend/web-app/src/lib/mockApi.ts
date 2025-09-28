// Mock API service for dev mode
/* eslint-disable @typescript-eslint/no-explicit-any */
export const mockApi = {
  // Mock user data
  users: {
    'dev-user-123': {
      id: 'dev-user-123',
      email: 'dev-user@ecopoint.com',
      firstName: 'Dev',
      lastName: 'User',
      role: 'USER',
      isActive: true,
      profileImageUrl: null,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    },
    'dev-collector-123': {
      id: 'dev-collector-123',
      email: 'dev-collector@ecopoint.com',
      firstName: 'Dev',
      lastName: 'Collector',
      role: 'COLLECTOR',
      isActive: true,
      profileImageUrl: null,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    },
    'dev-admin-123': {
      id: 'dev-admin-123',
      email: 'dev-admin@ecopoint.com',
      firstName: 'Dev',
      lastName: 'Admin',
      role: 'ADMIN',
      isActive: true,
      profileImageUrl: null,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    },
  } as Record<string, any>,

  // Mock bookings data
  bookings: [
    {
      id: 'booking-1',
      userId: 'dev-user-123',
      wasteType: 'PLASTIC',
      quantity: 5,
      description: 'Plastic bottles and containers',
      pickupAddress: {
        street: '123 Main St',
        city: 'Ho Chi Minh City',
        state: 'HCMC',
        zipCode: '70000',
        latitude: 10.7769,
        longitude: 106.7009,
      },
      scheduledDate: '2024-01-15',
      scheduledTime: '10:00',
      status: 'PENDING',
      estimatedPrice: 25.00,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    },
    {
      id: 'booking-2',
      userId: 'dev-user-123',
      wasteType: 'PAPER',
      quantity: 3,
      description: 'Old newspapers and magazines',
      pickupAddress: {
        street: '456 Oak Ave',
        city: 'Ho Chi Minh City',
        state: 'HCMC',
        zipCode: '70000',
        latitude: 10.7769,
        longitude: 106.7009,
      },
      scheduledDate: '2024-01-16',
      scheduledTime: '14:00',
      status: 'CONFIRMED',
      estimatedPrice: 15.00,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    },
  ] as any[],

  // Mock addresses data
  addresses: [
    {
      id: 'addr-1',
      userId: 'dev-user-123',
      street: '123 Main St',
      city: 'Ho Chi Minh City',
      state: 'HCMC',
      zipCode: '70000',
      latitude: 10.7769,
      longitude: 106.7009,
      isDefault: true,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    },
    {
      id: 'addr-2',
      userId: 'dev-user-123',
      street: '456 Oak Ave',
      city: 'Ho Chi Minh City',
      state: 'HCMC',
      zipCode: '70000',
      latitude: 10.7769,
      longitude: 106.7009,
      isDefault: false,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    },
  ] as any[],

  // Mock API methods
  async getCurrentUser(userId: string) {
    await new Promise(resolve => setTimeout(resolve, 500)); // Simulate network delay
    return this.users[userId] || null;
  },

  async getBookings(userId: string) {
    await new Promise(resolve => setTimeout(resolve, 500));
    return this.bookings.filter(booking => booking.userId === userId);
  },

  async getAddresses(userId: string) {
    await new Promise(resolve => setTimeout(resolve, 500));
    return this.addresses.filter(addr => addr.userId === userId);
  },

  async createBooking(bookingData: Record<string, unknown>) {
    await new Promise(resolve => setTimeout(resolve, 1000));
    const newBooking = {
      id: `booking-${Date.now()}`,
      ...bookingData,
      status: 'PENDING',
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };
    this.bookings.push(newBooking);
    return newBooking;
  },

  async createAddress(addressData: Record<string, unknown>) {
    await new Promise(resolve => setTimeout(resolve, 500));
    const newAddress = {
      id: `addr-${Date.now()}`,
      ...addressData,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };
    this.addresses.push(newAddress);
    return newAddress;
  },

  async updateUser(userId: string, userData: Record<string, unknown>) {
    await new Promise(resolve => setTimeout(resolve, 500));
    if (this.users[userId]) {
      this.users[userId] = { ...this.users[userId], ...userData, updatedAt: new Date().toISOString() };
      return this.users[userId];
    }
    return null;
  },
};