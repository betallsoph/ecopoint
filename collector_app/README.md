# ecoPoint Collector App

Ứng dụng dành cho người thu gom rác tái chế trong hệ thống ecoPoint.

## Tính năng chính

### 🏠 **Màn hình chính**
- Xem danh sách đơn hàng có sẵn
- Thông tin khách hàng và địa chỉ
- Loại rác thải cần thu gom
- Khoảng cách và thu nhập dự kiến
- Thông báo đơn hàng mới

### 📋 **Chi tiết đơn hàng**
- Thông tin chi tiết khách hàng
- Danh sách loại rác thải
- Khối lượng ước tính
- Nút nhận/từ chối đơn hàng
- Chuyển đến màn hình điều hướng

### 🗺️ **Điều hướng**
- Placeholder cho Google Maps
- Thông tin khoảng cách và thời gian
- Nút "Đã đến địa điểm"
- Chuyển đến màn hình hoàn thành

### ✅ **Hoàn thành đơn hàng**
- Chụp ảnh xác nhận thu gom
- Nhập khối lượng thực tế
- Ghi chú bổ sung
- Dialog xác nhận hoàn thành

### 👤 **Hồ sơ cá nhân**
- Thông tin collector và rating
- Thống kê tháng (đơn hoàn thành, thu nhập, khối lượng)
- Cài đặt tài khoản
- Lịch sử đơn hàng
- Quản lý thu nhập

## Thiết kế

### Màu sắc
- **Màu chính**: `Color(0xFF388E3C)` (xanh lá tươi)
- **Màu nền**: `Colors.white`
- **Màu nền nhạt**: `Color(0xFFC8E6C9)`

### Font chữ
- **Font family**: Montserrat
- **Weights**: Regular (400), Medium (500), SemiBold (600), Bold (700)

### UI Components
- Cards với border radius 12-16px
- Shadows nhẹ với opacity 0.1
- Icons từ Material Design
- Consistent spacing và padding

## Cấu trúc thư mục

```
lib/
├── main.dart                 # Entry point và routing
└── screens/
    ├── splash_screen.dart    # Màn hình splash
    ├── login_screen.dart     # Đăng nhập
    ├── home_screen.dart      # Danh sách đơn hàng
    ├── order_details_screen.dart    # Chi tiết đơn
    ├── navigation_screen.dart       # Điều hướng
    ├── complete_order_screen.dart   # Hoàn thành
    └── profile_screen.dart          # Hồ sơ
```

## Flow ứng dụng

1. **Splash Screen** → **Login** → **Home**
2. **Home** → **Order Details** → **Navigation** → **Complete**
3. **Profile** có thể truy cập từ bottom navigation

## So sánh với app chính

| Tính năng | Customer App | Collector App |
|-----------|-------------|---------------|
| Màn hình chính | Đặt dịch vụ | Xem đơn hàng |
| Quy trình | Đặt → Chờ → Hoàn thành | Nhận → Đi → Thu gom |
| Focus | UX đặt hàng | UX thu gom |
| Thông tin | Chọn rác thải | Chi tiết thu gom |

## Chạy ứng dụng

```bash
cd collector_app
flutter pub get
flutter run
```

## Tính năng có thể mở rộng

- 🗺️ Tích hợp Google Maps thực tế
- 📸 Camera API cho chụp ảnh
- 💰 Hệ thống thanh toán
- 📊 Analytics và reporting
- 🔔 Push notifications
- 📱 Real-time tracking
