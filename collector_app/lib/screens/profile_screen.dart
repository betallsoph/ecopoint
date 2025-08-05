import 'package:flutter/material.dart';

class ProfileScreen extends StatelessWidget {
  Widget _buildStatCard({
    required String title,
    required String value,
    required IconData icon,
    required Color color,
  }) {
    return Container(
      padding: EdgeInsets.all(20),
      decoration: BoxDecoration(
        color: color.withOpacity(0.1),
        borderRadius: BorderRadius.circular(16),
        border: Border.all(
          color: color.withOpacity(0.3),
        ),
      ),
      child: Column(
        children: [
          Icon(icon, color: color, size: 32),
          SizedBox(height: 12),
          Text(
            value,
            style: TextStyle(
              fontFamily: 'Montserrat',
              fontSize: 24,
              fontWeight: FontWeight.bold,
              color: color,
            ),
          ),
          SizedBox(height: 4),
          Text(
            title,
            style: TextStyle(
              fontFamily: 'Montserrat',
              fontSize: 14,
              color: Colors.grey[600],
            ),
            textAlign: TextAlign.center,
          ),
        ],
      ),
    );
  }

  Widget _buildInfoTile({
    required IconData icon,
    required String title,
    required String subtitle,
    VoidCallback? onTap,
  }) {
    return Container(
      margin: EdgeInsets.only(bottom: 12),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(
          color: Color(0xFF388E3C).withOpacity(0.2),
        ),
      ),
      child: ListTile(
        onTap: onTap,
        leading: Container(
          padding: EdgeInsets.all(8),
          decoration: BoxDecoration(
            color: Color(0xFF388E3C).withOpacity(0.1),
            borderRadius: BorderRadius.circular(8),
          ),
          child: Icon(icon, color: Color(0xFF388E3C)),
        ),
        title: Text(
          title,
          style: TextStyle(
            fontFamily: 'Montserrat',
            fontWeight: FontWeight.w600,
            color: Colors.black87,
          ),
        ),
        subtitle: Text(
          subtitle,
          style: TextStyle(
            fontFamily: 'Montserrat',
            color: Colors.grey[600],
          ),
        ),
        trailing: onTap != null 
            ? Icon(Icons.arrow_forward_ios, color: Colors.grey[400], size: 16)
            : null,
        contentPadding: EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[50],
      body: SafeArea(
        child: SingleChildScrollView(
          padding: EdgeInsets.all(24),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Profile header
              Container(
                padding: EdgeInsets.all(24),
                decoration: BoxDecoration(
                  color: Color(0xFF388E3C),
                  borderRadius: BorderRadius.circular(20),
                ),
                child: Row(
                  children: [
                    CircleAvatar(
                      radius: 40,
                      backgroundColor: Colors.white,
                      child: Icon(
                        Icons.person,
                        size: 40,
                        color: Color(0xFF388E3C),
                      ),
                    ),
                    SizedBox(width: 20),
                    Expanded(
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            'Nguyễn Văn Collector',
                            style: TextStyle(
                              fontFamily: 'Montserrat',
                              fontSize: 22,
                              fontWeight: FontWeight.bold,
                              color: Colors.white,
                            ),
                          ),
                          SizedBox(height: 4),
                          Text(
                            'Collector ID: COL001',
                            style: TextStyle(
                              fontFamily: 'Montserrat',
                              fontSize: 14,
                              color: Colors.white.withOpacity(0.8),
                            ),
                          ),
                          SizedBox(height: 8),
                          Container(
                            padding: EdgeInsets.symmetric(horizontal: 12, vertical: 4),
                            decoration: BoxDecoration(
                              color: Colors.white.withOpacity(0.2),
                              borderRadius: BorderRadius.circular(12),
                            ),
                            child: Text(
                              '⭐ 4.8 (156 đánh giá)',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                fontSize: 12,
                                color: Colors.white,
                                fontWeight: FontWeight.w600,
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ],
                ),
              ),

              SizedBox(height: 24),

              // Statistics
              Text(
                'Thống kê tháng này',
                style: TextStyle(
                  fontFamily: 'Montserrat',
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 16),

              Row(
                children: [
                  Expanded(
                    child: _buildStatCard(
                      title: 'Đơn hàng\nhoàn thành',
                      value: '47',
                      icon: Icons.check_circle,
                      color: Color(0xFF388E3C),
                    ),
                  ),
                  SizedBox(width: 16),
                  Expanded(
                    child: _buildStatCard(
                      title: 'Thu nhập\ntháng này',
                      value: '2.8M',
                      icon: Icons.account_balance_wallet,
                      color: Colors.orange,
                    ),
                  ),
                ],
              ),

              SizedBox(height: 16),

              Row(
                children: [
                  Expanded(
                    child: _buildStatCard(
                      title: 'Tổng khối lượng\nthu gom',
                      value: '238kg',
                      icon: Icons.scale,
                      color: Colors.blue,
                    ),
                  ),
                  SizedBox(width: 16),
                  Expanded(
                    child: _buildStatCard(
                      title: 'Thời gian\ntrực tuyến',
                      value: '156h',
                      icon: Icons.access_time,
                      color: Colors.purple,
                    ),
                  ),
                ],
              ),

              SizedBox(height: 32),

              // Menu items
              Text(
                'Cài đặt tài khoản',
                style: TextStyle(
                  fontFamily: 'Montserrat',
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 16),

              _buildInfoTile(
                icon: Icons.person_outline,
                title: 'Thông tin cá nhân',
                subtitle: 'Chỉnh sửa thông tin cá nhân',
                onTap: () {
                  // Navigate to edit profile
                },
              ),

              _buildInfoTile(
                icon: Icons.history,
                title: 'Lịch sử đơn hàng',
                subtitle: 'Xem các đơn hàng đã hoàn thành',
                onTap: () {
                  // Navigate to order history
                },
              ),

              _buildInfoTile(
                icon: Icons.account_balance_wallet,
                title: 'Thu nhập & Thanh toán',
                subtitle: 'Quản lý thu nhập và rút tiền',
                onTap: () {
                  // Navigate to earnings
                },
              ),

              _buildInfoTile(
                icon: Icons.notifications_outlined,
                title: 'Thông báo',
                subtitle: 'Cài đặt thông báo',
                onTap: () {
                  // Navigate to notifications settings
                },
              ),

              _buildInfoTile(
                icon: Icons.help_outline,
                title: 'Hỗ trợ',
                subtitle: 'Liên hệ hỗ trợ khách hàng',
                onTap: () {
                  // Navigate to support
                },
              ),

              _buildInfoTile(
                icon: Icons.logout,
                title: 'Đăng xuất',
                subtitle: 'Đăng xuất khỏi ứng dụng',
                onTap: () {
                  // Show logout confirmation
                  showDialog(
                    context: context,
                    builder: (BuildContext context) {
                      return AlertDialog(
                        backgroundColor: Colors.white,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(16),
                        ),
                        title: Text(
                          'Đăng xuất',
                          style: TextStyle(
                            fontFamily: 'Montserrat',
                            fontWeight: FontWeight.bold,
                            color: Color(0xFF388E3C),
                          ),
                        ),
                        content: Text(
                          'Bạn có chắc chắn muốn đăng xuất?',
                          style: TextStyle(
                            fontFamily: 'Montserrat',
                            color: Colors.grey[600],
                          ),
                        ),
                        actions: [
                          TextButton(
                            onPressed: () => Navigator.pop(context),
                            child: Text(
                              'Hủy',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                color: Colors.grey[600],
                              ),
                            ),
                          ),
                          ElevatedButton(
                            onPressed: () {
                              Navigator.pushNamedAndRemoveUntil(
                                context,
                                '/login',
                                (route) => false,
                              );
                            },
                            style: ElevatedButton.styleFrom(
                              backgroundColor: Color(0xFF388E3C),
                              foregroundColor: Colors.white,
                            ),
                            child: Text(
                              'Đăng xuất',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                fontWeight: FontWeight.w600,
                              ),
                            ),
                          ),
                        ],
                      );
                    },
                  );
                },
              ),

              SizedBox(height: 20),
            ],
          ),
        ),
      ),
    );
  }
} 