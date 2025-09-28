import { auth } from './firebase';
import { mockApi } from './mockApi';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080/api/v1';

class ApiClient {
  private baseURL: string;
  private isDevMode: boolean = false;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
    // Check if we're in dev mode (you can set this via environment variable)
    this.isDevMode = process.env.NODE_ENV === 'development' && process.env.NEXT_PUBLIC_DEV_MODE === 'true';
  }

  setDevMode(enabled: boolean) {
    this.isDevMode = enabled;
  }

  private async getAuthHeaders(): Promise<HeadersInit> {
    if (this.isDevMode) {
      return {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer dev-token',
      };
    }
    
    const token = await auth.currentUser?.getIdToken();
    return {
      'Content-Type': 'application/json',
      'Authorization': token ? `Bearer ${token}` : '',
    };
  }

  async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    // Use mock API in dev mode
    if (this.isDevMode) {
      return this.mockRequest<T>(endpoint, options);
    }

    const url = `${this.baseURL}${endpoint}`;
    const headers = await this.getAuthHeaders();

    const config: RequestInit = {
      ...options,
      headers: {
        ...headers,
        ...options.headers,
      },
    };

    const response = await fetch(url, config);

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || `HTTP error! status: ${response.status}`);
    }

    return response.json();
  }

  private async mockRequest<T>(endpoint: string, options: RequestInit): Promise<T> {
    // Simulate network delay
    await new Promise(resolve => setTimeout(resolve, 300));

    const method = options.method || 'GET';
    const body = options.body ? JSON.parse(options.body as string) : {};

    // Mock responses based on endpoint
    if (endpoint === '/users/me' && method === 'GET') {
      return mockApi.getCurrentUser('dev-user-123') as T;
    }

    if (endpoint === '/users/me' && method === 'PUT') {
      return mockApi.updateUser('dev-user-123', body) as T;
    }

    if (endpoint === '/addresses' && method === 'GET') {
      return mockApi.getAddresses('dev-user-123') as T;
    }

    if (endpoint === '/addresses' && method === 'POST') {
      return mockApi.createAddress({ ...body, userId: 'dev-user-123' }) as T;
    }

    if (endpoint.startsWith('/addresses/') && method === 'PUT') {
      const id = endpoint.split('/')[2];
      return { id, ...body } as T;
    }

    if (endpoint.startsWith('/addresses/') && method === 'DELETE') {
      return { success: true } as T;
    }

    if (endpoint === '/bookings' && method === 'GET') {
      return mockApi.getBookings('dev-user-123') as T;
    }

    if (endpoint.startsWith('/bookings/') && method === 'GET') {
      const id = endpoint.split('/')[2];
      const bookings = await mockApi.getBookings('dev-user-123');
      return bookings.find((b: Record<string, unknown>) => b.id === id) as T;
    }

    if (endpoint === '/bookings' && method === 'POST') {
      return mockApi.createBooking({ ...body, userId: 'dev-user-123' }) as T;
    }

    if (endpoint.startsWith('/bookings/') && method === 'PUT') {
      const id = endpoint.split('/')[2];
      return { id, ...body } as T;
    }

    if (endpoint.startsWith('/bookings/') && method === 'DELETE') {
      return { success: true } as T;
    }

    // Default mock response
    return { success: true, message: 'Mock response' } as T;
  }

  // User endpoints
  async getCurrentUser() {
    return this.request('/users/me');
  }

  async updateUser(data: Record<string, unknown>) {
    return this.request('/users/me', {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  // Address endpoints
  async getAddresses() {
    return this.request('/addresses');
  }

  async createAddress(data: Record<string, unknown>) {
    return this.request('/addresses', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateAddress(id: string, data: Record<string, unknown>) {
    return this.request(`/addresses/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteAddress(id: string) {
    return this.request(`/addresses/${id}`, {
      method: 'DELETE',
    });
  }

  // Booking endpoints
  async getBookings(params?: { status?: string; page?: number; limit?: number }) {
    const searchParams = new URLSearchParams();
    if (params?.status) searchParams.append('status', params.status);
    if (params?.page) searchParams.append('page', params.page.toString());
    if (params?.limit) searchParams.append('limit', params.limit.toString());
    
    const queryString = searchParams.toString();
    return this.request(`/bookings${queryString ? `?${queryString}` : ''}`);
  }

  async getBooking(id: string) {
    return this.request(`/bookings/${id}`);
  }

  async createBooking(data: Record<string, unknown>) {
    return this.request('/bookings', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateBooking(id: string, data: Record<string, unknown>) {
    return this.request(`/bookings/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async cancelBooking(id: string) {
    return this.request(`/bookings/${id}`, {
      method: 'DELETE',
    });
  }

  // Collector endpoints
  async getNearbyCollectors(lat: number, lng: number, radius: number = 10) {
    return this.request(`/collectors/nearby?lat=${lat}&lng=${lng}&radius=${radius}`);
  }

  async getCollectorLocation(collectorId: string) {
    return this.request(`/collectors/${collectorId}/location`);
  }

  // Chat endpoints
  async getConversations() {
    return this.request('/chat/conversations');
  }

  async getMessages(conversationId: string) {
    return this.request(`/chat/conversations/${conversationId}/messages`);
  }

  async sendMessage(conversationId: string, message: string, type: string = 'text') {
    return this.request(`/chat/conversations/${conversationId}/messages`, {
      method: 'POST',
      body: JSON.stringify({ message, type }),
    });
  }

  // Tracking endpoints
  async getTrackingEvents(bookingId: string) {
    return this.request(`/tracking/${bookingId}/events`);
  }

  async updateLocation(bookingId: string, location: { lat: number; lng: number; accuracy?: number }) {
    return this.request(`/tracking/${bookingId}/location`, {
      method: 'POST',
      body: JSON.stringify(location),
    });
  }
}

export const apiClient = new ApiClient(API_BASE_URL);