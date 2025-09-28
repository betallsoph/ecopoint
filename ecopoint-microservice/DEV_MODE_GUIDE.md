# 🚀 Dev Mode Guide - Test App Ngay Lập Tức!

## ✅ **Dev Mode đã sẵn sàng!**

Bạn có thể test toàn bộ hệ thống EcoPoint ngay bây giờ mà không cần setup Firebase hay database!

## 🎯 **Cách sử dụng Dev Mode:**

### **📱 Web App (React/Next.js)**

#### **1. Start Web App:**
```bash
cd frontend/web-app
npm run dev
```

#### **2. Mở trình duyệt:**
```
http://localhost:3000
```

#### **3. Bật Dev Mode:**
- Nhìn góc trên bên phải màn hình
- Toggle **"🚀 Dev Mode"** switch
- Chọn role: **User** / **Collector** / **Admin**
- Dùng **bất kỳ email/password nào** để login

#### **4. Test các tính năng:**
- ✅ Login với bất kỳ email/password
- ✅ Switch giữa các roles
- ✅ Xem mock data (bookings, addresses)
- ✅ Test UI/UX cho từng role

---

### **📱 Flutter Apps**

#### **1. User App:**
```bash
cd mobile/user-app
flutter run
```

#### **2. Collector App:**
```bash
cd mobile/collector-app
flutter run
```

#### **3. Admin App:**
```bash
cd mobile/admin-app
flutter run
```

#### **4. Bật Dev Mode trong app:**
- Mở login screen
- Toggle **"🚀 Dev Mode"** switch
- Dùng **bất kỳ email/password nào** để login
- Mock data sẽ được load tự động

---

## 🎨 **Tính năng Dev Mode:**

### **✅ Web App Features:**
- **Dev Mode Toggle**: Switch on/off ở góc trên phải
- **Role Switching**: User/Collector/Admin với 1 click
- **Mock Authentication**: Bypass Firebase hoàn toàn
- **Mock API**: Tất cả API calls trả về mock data
- **Mock Data**: Bookings, addresses, users sẵn sàng

### **✅ Flutter Apps Features:**
- **Dev Mode Toggle**: Trong login screen
- **Mock Authentication**: Bypass Firebase
- **Mock Data**: Load tự động khi enable dev mode
- **Role-based UI**: Khác nhau cho từng app

### **✅ Mock Data Available:**
- **Users**: 3 roles (User, Collector, Admin)
- **Bookings**: 2 sample bookings với đầy đủ thông tin
- **Addresses**: 2 sample addresses
- **Realistic Data**: Giống như data thật

---

## 🔧 **Technical Details:**

### **Web App Dev Mode:**
```typescript
// DevModeContext.tsx - Quản lý dev mode state
// DevModeToggle.tsx - UI component để toggle
// mockApi.ts - Mock API service với realistic data
// apiClient.ts - API client với dev mode support
```

### **Flutter Apps Dev Mode:**
```dart
// dev_mode_service.dart - Quản lý dev mode state
// login_screen.dart - Login với dev mode toggle
// Mock data được load từ service
```

### **Bypass Features:**
- ✅ **Firebase Auth**: Hoàn toàn bypass
- ✅ **API Calls**: Trả về mock data
- ✅ **Database**: Không cần kết nối
- ✅ **Real-time**: Mock responses

---

## 🚀 **Quick Start Commands:**

### **Start Web App:**
```bash
cd ecopoint-microservice/frontend/web-app
npm run dev
# Mở http://localhost:3000
# Toggle Dev Mode ở góc trên phải
```

### **Start Flutter Apps:**
```bash
# User App
cd ecopoint-microservice/mobile/user-app
flutter run

# Collector App  
cd ecopoint-microservice/mobile/collector-app
flutter run

# Admin App
cd ecopoint-microservice/mobile/admin-app
flutter run
```

### **Start Backend (Optional):**
```bash
cd ecopoint-microservice
docker-compose up -d
# Chỉ cần nếu muốn test real API
```

---

## 🎯 **Test Scenarios:**

### **1. Test User Flow:**
1. Bật Dev Mode → Chọn "User"
2. Login với email: `test@example.com`, password: `123456`
3. Xem dashboard với mock bookings
4. Test tạo booking mới
5. Test quản lý addresses

### **2. Test Collector Flow:**
1. Bật Dev Mode → Chọn "Collector"  
2. Login với email: `collector@example.com`, password: `123456`
3. Xem task dashboard
4. Test online/offline status
5. Test location tracking

### **3. Test Admin Flow:**
1. Bật Dev Mode → Chọn "Admin"
2. Login với email: `admin@example.com`, password: `123456`
3. Xem admin dashboard
4. Test analytics
5. Test user management

---

## 🎉 **Kết quả:**

**Bạn có thể test toàn bộ hệ thống EcoPoint ngay bây giờ!**

- ✅ **3 Flutter Apps** riêng biệt cho 3 roles
- ✅ **Web App** với role switching
- ✅ **Mock data** realistic
- ✅ **Bypass Firebase** hoàn toàn
- ✅ **UI/UX testing** đầy đủ
- ✅ **Không cần setup** gì thêm

**Chỉ cần chạy `npm run dev` và `flutter run` là có thể test ngay!** 🚀

---

## 📝 **Notes:**

- Dev Mode chỉ hoạt động trong development
- Production sẽ sử dụng real Firebase và API
- Mock data có thể customize trong `mockApi.ts`
- Tất cả features UI đã sẵn sàng để test

**Happy testing! 🎉**