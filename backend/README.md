# ecoPoint Backend - Microservices Architecture

Hệ thống backend cho ứng dụng ecoPoint sử dụng kiến trúc microservices với Go và MongoDB.

## 🏗️ Kiến trúc

```
┌─────────────────┐    ┌──────────────────┐
│   Flutter Apps  │────│   API Gateway    │
│ (Customer/      │    │   (Port 8080)    │
│  Collector)     │    │                  │
└─────────────────┘    └────────┬─────────┘
                               │
        ┌──────────────────────┼──────────────────────┐
        │                      │                      │
   ┌────▼───┐            ┌─────▼────┐           ┌─────▼────┐
   │  Auth  │            │   User   │           │  Order   │
   │Service │            │ Service  │           │ Service  │
   │:8081   │            │  :8082   │           │  :8083   │
   └────────┘            └──────────┘           └──────────┘
```

## 🚀 Services

### 1. API Gateway (Port 8080)
- Điểm vào duy nhất cho tất cả requests
- Routing requests đến các services
- Load balancing và health checking

### 2. Auth Service (Port 8081)
- Xác thực người dùng (Google OAuth, Phone)
- Quản lý JWT tokens
- Tích hợp Firebase Auth

### 3. User Service (Port 8082)
- Quản lý thông tin người dùng
- Profile management
- Thống kê user/collector

### 4. Order Service (Port 8083)
- Quản lý đơn hàng thu gom
- Workflow từ tạo đơn đến hoàn thành
- Tính toán thanh toán

## 🛠️ Tech Stack

- **Language:** Go 1.21
- **Framework:** Gin
- **Database:** MongoDB
- **Authentication:** JWT + Firebase Auth
- **Containerization:** Docker
- **Orchestration:** Docker Compose

## 📦 Quick Start

### Prerequisites
- Docker và Docker Compose
- Go 1.21+ (cho development)
- MongoDB (hoặc MongoDB Atlas)

### 1. Clone và Setup
```bash
git clone <repo>
cd backend
cp .env.example .env
# Chỉnh sửa .env với thông tin của bạn
```

### 2. Chạy với Docker Compose
```bash
# Build và start tất cả services
docker-compose up --build

# Chạy ở background
docker-compose up -d

# Xem logs
docker-compose logs -f

# Stop services
docker-compose down
```

### 3. Development Mode
```bash
# Install dependencies
go mod download

# Chạy từng service riêng lẻ
cd auth-service && go run main.go
cd user-service && go run main.go
cd order-service && go run main.go
cd gateway && go run main.go
```

## 🌐 API Endpoints

### Gateway (http://localhost:8080)
```
GET  /health                    # Health check tất cả services
POST /api/v1/auth/google        # Google sign-in
POST /api/v1/auth/phone/verify  # Phone verification
GET  /api/v1/users/profile      # User profile
POST /api/v1/orders             # Tạo đơn hàng
GET  /api/v1/orders/available   # Đơn hàng có sẵn (collectors)
```

### Auth Service (http://localhost:8081)
```
POST /api/v1/auth/google        # Google authentication
POST /api/v1/auth/phone/verify  # Phone number verification
POST /api/v1/auth/refresh       # Refresh JWT token
GET  /api/v1/auth/profile       # Get current user profile
```

### User Service (http://localhost:8082)
```
GET  /api/v1/users/profile      # Get user profile
PUT  /api/v1/users/profile      # Update user profile
GET  /api/v1/users/stats        # User statistics
```

### Order Service (http://localhost:8083)
```
POST /api/v1/orders             # Create new order
GET  /api/v1/orders             # Get user's orders
GET  /api/v1/orders/:id         # Get specific order
PUT  /api/v1/orders/:id/accept  # Accept order (collectors)
PUT  /api/v1/orders/:id/complete # Complete order
GET  /api/v1/orders/available   # Available orders (collectors)
```

## 🗄️ Database Schema

### Users Collection
```json
{
  "_id": "ObjectId",
  "firebase_uid": "string",
  "email": "string",
  "display_name": "string",
  "phone_number": "string",
  "user_type": "customer|collector",
  "is_active": true,
  "created_at": "ISODate",
  "updated_at": "ISODate"
}
```

### Orders Collection
```json
{
  "_id": "ObjectId",
  "customer_id": "ObjectId",
  "collector_id": "ObjectId",
  "status": "pending|accepted|in_progress|completed|cancelled",
  "waste_types": ["paper", "plastic", "metal", "glass", "electronic"],
  "estimated_weight": 5.0,
  "actual_weight": 4.8,
  "pickup_address": {
    "street": "string",
    "district": "string", 
    "city": "string",
    "lat": 10.762622,
    "lng": 106.660172
  },
  "scheduled_time": "ISODate",
  "completed_time": "ISODate",
  "payment": {
    "amount": 50000,
    "currency": "VND",
    "method": "cash",
    "is_paid": false
  },
  "created_at": "ISODate",
  "updated_at": "ISODate"
}
```

## 🔐 Authentication Flow

1. **Client** đăng nhập qua Google/Phone
2. **Auth Service** xác thực và trả về JWT
3. **Client** gửi JWT trong Authorization header cho các requests tiếp theo
4. **Gateway** forward request kèm JWT đến service tương ứng
5. **Service** validate JWT và thực hiện business logic

## 🐛 Development

### Testing APIs
```bash
# Health check
curl http://localhost:8080/health

# Google sign-in (cần ID token thật từ client)
curl -X POST http://localhost:8080/api/v1/auth/google \
  -H "Content-Type: application/json" \
  -d '{"id_token": "...", "user_type": "customer"}'

# Create order (cần JWT token)
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"waste_types": ["paper"], "estimated_weight": 2.0, ...}'
```

### Monitoring
```bash
# Xem logs của tất cả services
docker-compose logs -f

# Xem logs của một service cụ thể
docker-compose logs -f auth-service

# Monitor resource usage
docker stats
```

## 📈 Scaling

Để scale hệ thống:

1. **Horizontal scaling:** Tăng số instance của từng service
2. **Database:** Sử dụng MongoDB Atlas với auto-scaling
3. **Load balancer:** Nginx hoặc cloud load balancer trước Gateway
4. **Caching:** Redis cho session và frequently accessed data

## 🔧 Configuration

Tất cả cấu hình qua environment variables trong `.env`:

- `MONGO_URI`: Connection string MongoDB
- `JWT_SECRET`: Secret key cho JWT signing
- `FIREBASE_CONFIG_PATH`: Đường dẫn đến Firebase config file

## 🚀 Deployment

### Production với Docker
```bash
# Build production images
docker-compose -f docker-compose.prod.yml up --build -d

# Với external MongoDB
MONGO_URI=mongodb+srv://user:pass@cluster.mongodb.net/ecopoint docker-compose up -d
```

### Cloud Deployment
- **Containers:** AWS ECS, Google Cloud Run, Azure Container Instances
- **Database:** MongoDB Atlas
- **Monitoring:** CloudWatch, Stackdriver, Application Insights

---

**Made with ❤️ for a cleaner planet 🌱**