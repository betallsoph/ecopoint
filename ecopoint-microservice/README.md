# EcoPoint - Waste Collection Microservice System

## 🏗️ Architecture Overview

Hệ thống thu gom rác thải tái chế với kiến trúc microservice, tương tự GrabFood/UberEats.

### 🎯 Core Features
- **User Management**: Quản lý người dùng, collector, admin với Firebase Auth
- **Booking System**: Đặt lịch thu gom rác thải với real-time tracking
- **Real-time Tracking**: Theo dõi collector real-time với Server-Sent Events
- **Chat System**: Chat giữa user và collector
- **Maps Integration**: Tích hợp Google Maps cho location services
- **Multi-platform**: Web (Next.js) + Mobile (Flutter)

### 🏛️ Tech Stack

#### Backend Services
- **API Gateway**: Go + Gin + Firebase Auth
- **User Service**: Node.js + Express + GraphQL + MongoDB
- **Booking Service**: Go + gRPC + PostgreSQL
- **Database**: PostgreSQL (primary) + MongoDB (users) + Redis (cache)

#### Frontend
- **Web App**: Next.js + TypeScript + Tailwind CSS + Firebase Auth
- **Mobile App**: Flutter + Dart + Firebase Auth

#### Real-time Features
- **Tracking**: Server-Sent Events (SSE)
- **Chat**: WebSocket connections
- **Maps**: Google Maps API

### 📁 Project Structure

```
ecopoint-microservice/
├── services/
│   ├── api-gateway/          # API Gateway (Go + Gin)
│   ├── user-service/         # User Management (Node.js + GraphQL)
│   └── booking-service/      # Booking System (Go + gRPC)
├── frontend/
│   └── web-app/              # Next.js Web Application
├── mobile/
│   └── flutter-app/          # Flutter Mobile App
├── docs/
│   └── database-schema.sql   # PostgreSQL Schema
├── docker-compose.yml        # Docker Compose Configuration
└── README.md
```

### 🚀 Quick Start

#### Prerequisites
- Docker & Docker Compose
- Node.js 18+ (for local development)
- Go 1.21+ (for local development)
- Flutter 3.0+ (for mobile development)

#### 1. Clone and Setup
```bash
git clone <repository-url>
cd ecopoint-microservice
```

#### 2. Environment Configuration
```bash
# Copy environment files
cp services/user-service/.env.example services/user-service/.env
cp services/api-gateway/.env.example services/api-gateway/.env
cp frontend/web-app/.env.local.example frontend/web-app/.env.local

# Edit the environment files with your Firebase configuration
```

#### 3. Start with Docker Compose
```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

#### 4. Access Applications
- **Web App**: http://localhost:3000
- **API Gateway**: http://localhost:8080
- **User Service GraphQL**: http://localhost:4000/graphql
- **PostgreSQL**: localhost:5432
- **MongoDB**: localhost:27017
- **Redis**: localhost:6379

### 🔧 Development

#### Running Services Individually

**User Service**
```bash
cd services/user-service
npm install
npm run dev
```

**Booking Service**
```bash
cd services/booking-service
go mod tidy
go run cmd/main.go
```

**API Gateway**
```bash
cd services/api-gateway
go mod tidy
go run cmd/main.go
```

**Web App**
```bash
cd frontend/web-app
npm install
npm run dev
```

**Flutter App**
```bash
cd mobile/flutter-app
flutter pub get
flutter run
```

### 🔐 Authentication

- **Firebase Authentication** cho tất cả platforms
- **JWT tokens** cho API communication
- **Role-based access control** (USER, COLLECTOR, ADMIN)

### 📊 Database Schema

#### PostgreSQL Tables
- `users` - User information (Firebase UID based)
- `user_addresses` - User addresses
- `user_preferences` - User preferences
- `waste_types` - Available waste types
- `bookings` - Booking records
- `booking_images` - Booking images
- `booking_status_history` - Status change history
- `collector_availability` - Collector availability
- `collector_locations` - Real-time collector locations
- `chat_conversations` - Chat conversations
- `chat_messages` - Chat messages
- `system_settings` - System configuration

### 🗺️ Maps Integration

- **Google Maps API** cho location services
- **Real-time tracking** với Server-Sent Events
- **Geolocation** cho collector positioning
- **Address autocomplete** cho pickup locations

### 💬 Chat System

- **WebSocket connections** cho real-time messaging
- **Message types**: text, image, location
- **Conversation management** per booking
- **Message history** với pagination

### 📱 Mobile App Features

#### User App
- Firebase Authentication
- Booking creation và management
- Real-time tracking
- Chat với collector
- Address management

#### Collector App
- Task management
- Location tracking
- Chat với users
- Booking status updates

#### Admin App
- User management
- System monitoring
- Analytics dashboard

### 🔄 Real-time Features

#### Server-Sent Events (SSE)
- **Tracking updates**: Collector location changes
- **Status updates**: Booking status changes
- **Notifications**: Real-time notifications

#### WebSocket
- **Chat messaging**: Real-time chat
- **Live updates**: Booking updates

### 🧪 Testing

```bash
# Run all tests
docker-compose -f docker-compose.test.yml up --abort-on-container-exit

# Run specific service tests
cd services/user-service && npm test
cd services/booking-service && go test ./...
```

### 📈 Monitoring

- **Health checks** cho tất cả services
- **Logging** với structured logs
- **Metrics** với Prometheus (planned)
- **Tracing** với OpenTelemetry (planned)

### 🚀 Deployment

#### Production Environment
```bash
# Build production images
docker-compose -f docker-compose.prod.yml build

# Deploy to production
docker-compose -f docker-compose.prod.yml up -d
```

#### Kubernetes (Optional)
```bash
# Apply Kubernetes manifests
kubectl apply -f infrastructure/kubernetes/
```

### 🔧 Configuration

#### Environment Variables

**API Gateway**
- `PORT`: Server port (default: 8080)
- `USER_SERVICE_URL`: User service URL
- `BOOKING_SERVICE_URL`: Booking service URL
- `FIREBASE_PROJECT_ID`: Firebase project ID

**User Service**
- `PORT`: Server port (default: 4000)
- `MONGODB_URI`: MongoDB connection string
- `JWT_SECRET`: JWT secret key

**Booking Service**
- `DATABASE_URL`: PostgreSQL connection string
- `GRPC_PORT`: gRPC server port (default: 50051)

### 📋 TODO

- [x] Setup project structure
- [x] Design database schema
- [x] Implement User Service
- [x] Implement Booking Service
- [x] Setup API Gateway
- [x] Create Web Frontend
- [x] Develop Flutter Mobile App
- [x] Setup Docker environment
- [ ] Implement Firebase Auth Integration
- [ ] Add Real-time tracking (SSE)
- [ ] Setup Maps Integration
- [ ] Implement Chat System
- [ ] Add comprehensive tests
- [ ] Setup monitoring & logging
- [ ] Production deployment

### 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

### 📄 License

This project is licensed under the MIT License.

---

## 🆘 Support

Nếu có vấn đề gì, vui lòng tạo issue hoặc liên hệ team development.

**Happy Coding! 🚀**