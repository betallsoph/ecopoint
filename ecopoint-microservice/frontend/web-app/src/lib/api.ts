import { auth } from './firebase';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080/api/v1';

class ApiClient {
  private baseURL: string;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  private async getAuthHeaders(): Promise<HeadersInit> {
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