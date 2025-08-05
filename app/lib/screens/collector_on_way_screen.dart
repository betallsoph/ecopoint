import 'package:flutter/material.dart';
import 'dart:async';

class CollectorOnWayScreen extends StatefulWidget {
  @override
  _CollectorOnWayScreenState createState() => _CollectorOnWayScreenState();
}

class _CollectorOnWayScreenState extends State<CollectorOnWayScreen> with TickerProviderStateMixin {
  late AnimationController _pulseController;
  late Animation<double> _pulseAnimation;
  Timer? _navigationTimer;
  int _estimatedTime = 15; // minutes
  Timer? _timeTimer;

  @override
  void initState() {
    super.initState();
    
    // Pulse animation for the car icon
    _pulseController = AnimationController(
      duration: Duration(milliseconds: 1500),
      vsync: this,
    );
    _pulseAnimation = Tween<double>(
      begin: 0.9,
      end: 1.1,
    ).animate(CurvedAnimation(
      parent: _pulseController,
      curve: Curves.easeInOut,
    ));
    _pulseController.repeat(reverse: true);
    
    // Auto navigate to finish screen after 10 seconds
    _navigationTimer = Timer(Duration(seconds: 10), () {
      if (mounted) {
        Navigator.pushReplacementNamed(context, '/finish');
      }
    });
    
    // Update estimated time every 30 seconds (simulate decreasing time)
    _timeTimer = Timer.periodic(Duration(seconds: 30), (timer) {
      if (mounted && _estimatedTime > 1) {
        setState(() {
          _estimatedTime--;
        });
      }
    });
  }

  @override
  void dispose() {
    _pulseController.dispose();
    _navigationTimer?.cancel();
    _timeTimer?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: Text(
          'Nhân viên đang đến',
          style: TextStyle(
            fontFamily: 'Montserrat',
            fontWeight: FontWeight.bold,
            color: Colors.white,
          ),
        ),
        backgroundColor: Color(0xFF388E3C),
        leading: IconButton(
          icon: Icon(Icons.arrow_back, color: Colors.white),
          onPressed: () => Navigator.pop(context),
        ),
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            children: [
              Expanded(
                child: SingleChildScrollView(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                    // Animated car icon
                    AnimatedBuilder(
                      animation: _pulseAnimation,
                      builder: (context, child) {
                        return Transform.scale(
                          scale: _pulseAnimation.value,
                          child: Container(
                            width: 120,
                            height: 120,
                            decoration: BoxDecoration(
                              color: Color(0xFF388E3C).withOpacity(0.1),
                              shape: BoxShape.circle,
                              border: Border.all(
                                color: Color(0xFF388E3C),
                                width: 3,
                              ),
                            ),
                            child: Icon(
                              Icons.directions_car,
                              size: 60,
                              color: Color(0xFF388E3C),
                            ),
                          ),
                        );
                      },
                    ),
                    
                    SizedBox(height: 32),
                    
                    // Status text
                    Text(
                      'Nhân viên đang trên đường đến',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 24,
                        fontWeight: FontWeight.bold,
                        color: Color(0xFF388E3C),
                      ),
                      textAlign: TextAlign.center,
                    ),
                    
                    SizedBox(height: 16),
                    
                    Text(
                      'Nhân viên thu gom đã được phân công và đang di chuyển đến địa chỉ của bạn.',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 16,
                        fontWeight: FontWeight.w400,
                        color: Color(0xFF388E3C).withOpacity(0.8),
                      ),
                      textAlign: TextAlign.center,
                    ),
                    
                    SizedBox(height: 32),
                    
                    // Estimated time
                    Container(
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(8),
                        border: Border.all(
                          color: Color(0xFF388E3C).withOpacity(0.3),
                          width: 1,
                        ),
                      ),
                      padding: EdgeInsets.symmetric(horizontal: 24, vertical: 16),
                      child: Row(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          Icon(Icons.access_time, color: Color(0xFF388E3C)),
                          SizedBox(width: 8),
                          Text(
                            'Dự kiến: $_estimatedTime phút',
                            style: TextStyle(
                              fontFamily: 'Montserrat',
                              fontSize: 16,
                              fontWeight: FontWeight.w600,
                              color: Color(0xFF388E3C),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ],
                  ),
                ),
              ),
              
              // Collector info card
              Container(
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(12),
                  border: Border.all(
                    color: Color(0xFF388E3C).withOpacity(0.3),
                    width: 1,
                  ),
                  boxShadow: [
                    BoxShadow(
                      color: Color(0xFF388E3C).withOpacity(0.1),
                      blurRadius: 8,
                      offset: Offset(0, 4),
                    ),
                  ],
                ),
                padding: EdgeInsets.all(20),
                child: Column(
                  children: [
                    Row(
                      children: [
                        // Avatar
                        Container(
                          width: 60,
                          height: 60,
                          decoration: BoxDecoration(
                            color: Color(0xFF388E3C).withOpacity(0.1),
                            shape: BoxShape.circle,
                            border: Border.all(
                              color: Color(0xFF388E3C).withOpacity(0.3),
                              width: 2,
                            ),
                          ),
                          child: Icon(
                            Icons.person,
                            size: 30,
                            color: Color(0xFF388E3C),
                          ),
                        ),
                        SizedBox(width: 16),
                        // Collector info
                        Expanded(
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text(
                                'Nguyễn Văn A',
                                style: TextStyle(
                                  fontFamily: 'Montserrat',
                                  fontWeight: FontWeight.bold,
                                  fontSize: 18,
                                  color: Color(0xFF388E3C),
                                ),
                              ),
                              SizedBox(height: 4),
                              Text(
                                'Nhân viên thu gom',
                                style: TextStyle(
                                  fontFamily: 'Montserrat',
                                  fontSize: 14,
                                  fontWeight: FontWeight.w400,
                                  color: Color(0xFF388E3C).withOpacity(0.7),
                                ),
                              ),
                              SizedBox(height: 8),
                              Row(
                                children: [
                                  Icon(Icons.star, color: Colors.amber, size: 16),
                                  SizedBox(width: 4),
                                  Text(
                                    '4.8 (125 đánh giá)',
                                    style: TextStyle(
                                      fontFamily: 'Montserrat',
                                      fontSize: 12,
                                      fontWeight: FontWeight.w500,
                                      color: Color(0xFF388E3C).withOpacity(0.7),
                                    ),
                                  ),
                                ],
                              ),
                            ],
                          ),
                        ),
                        // Call button
                        Container(
                          decoration: BoxDecoration(
                            color: Color(0xFF388E3C),
                            shape: BoxShape.circle,
                          ),
                          child: IconButton(
                            onPressed: () {
                              // Make a call
                            },
                            icon: Icon(Icons.phone, color: Colors.white),
                          ),
                        ),
                      ],
                    ),
                    
                    SizedBox(height: 20),
                    
                    // Vehicle info
                    Container(
                      decoration: BoxDecoration(
                        color: Color(0xFF388E3C).withOpacity(0.1),
                        borderRadius: BorderRadius.circular(8),
                        border: Border.all(
                          color: Color(0xFF388E3C).withOpacity(0.3),
                          width: 1,
                        ),
                      ),
                      padding: EdgeInsets.all(16),
                      child: Row(
                        children: [
                          Icon(Icons.local_shipping, color: Color(0xFF388E3C)),
                          SizedBox(width: 12),
                          Expanded(
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Text(
                                  'Phương tiện: Xe tải nhỏ',
                                  style: TextStyle(
                                    fontFamily: 'Montserrat',
                                    fontWeight: FontWeight.w600,
                                    fontSize: 14,
                                    color: Color(0xFF388E3C),
                                  ),
                                ),
                                Text(
                                  'Biển số: 59A-12345',
                                  style: TextStyle(
                                    fontFamily: 'Montserrat',
                                    fontSize: 12,
                                    fontWeight: FontWeight.w400,
                                    color: Color(0xFF388E3C).withOpacity(0.7),
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ],
                      ),
                    ),
                  ],
                ),
              ),
              
              SizedBox(height: 24),
              
              // Progress indicator
              Container(
                decoration: BoxDecoration(
                  color: Color(0xFF388E3C).withOpacity(0.1),
                  borderRadius: BorderRadius.circular(12),
                  border: Border.all(
                    color: Color(0xFF388E3C).withOpacity(0.3),
                    width: 1,
                  ),
                ),
                padding: EdgeInsets.all(20),
                child: Column(
                  children: [
                    _buildProgressRow(
                      icon: Icons.check_circle,
                      title: 'Xác nhận đơn hàng',
                      subtitle: 'Hoàn thành',
                      isCompleted: true,
                    ),
                    SizedBox(height: 16),
                    _buildProgressRow(
                      icon: Icons.check_circle,
                      title: 'Tìm nhân viên',
                      subtitle: 'Hoàn thành',
                      isCompleted: true,
                    ),
                    SizedBox(height: 16),
                    _buildProgressRow(
                      icon: Icons.directions_car,
                      title: 'Trên đường đến',
                      subtitle: 'Đang thực hiện',
                      isActive: true,
                    ),
                    SizedBox(height: 16),
                    _buildProgressRow(
                      icon: Icons.home,
                      title: 'Thu gom tại nhà',
                      subtitle: 'Sắp diễn ra',
                      isNext: true,
                    ),
                  ],
                ),
              ),
              
              SizedBox(height: 24),
              
              // Cancel button
              SizedBox(
                width: double.infinity,
                child: OutlinedButton(
                  onPressed: () {
                    _navigationTimer?.cancel();
                    Navigator.pushNamedAndRemoveUntil(context, '/home', (route) => false);
                  },
                  child: Text(
                    'Hủy đặt dịch vụ',
                    style: TextStyle(
                      fontFamily: 'Montserrat',
                      fontWeight: FontWeight.w600,
                      fontSize: 16,
                    ),
                  ),
                  style: OutlinedButton.styleFrom(
                    side: BorderSide(color: Colors.red, width: 1.5),
                    foregroundColor: Colors.red,
                    padding: EdgeInsets.symmetric(vertical: 16),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildProgressRow({
    required IconData icon,
    required String title,
    required String subtitle,
    bool isCompleted = false,
    bool isActive = false,
    bool isNext = false,
  }) {
    return Row(
      children: [
        Container(
          width: 40,
          height: 40,
          decoration: BoxDecoration(
            color: isCompleted || isActive
                ? Color(0xFF388E3C)
                : Color(0xFF388E3C).withOpacity(0.3),
            shape: BoxShape.circle,
          ),
          child: Icon(
            icon,
            color: isCompleted || isActive ? Colors.white : Color(0xFF388E3C),
            size: 20,
          ),
        ),
        SizedBox(width: 16),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                title,
                style: TextStyle(
                  fontFamily: 'Montserrat',
                  fontWeight: FontWeight.bold,
                  fontSize: 16,
                  color: isNext ? Color(0xFF388E3C).withOpacity(0.5) : Color(0xFF388E3C),
                ),
              ),
              Text(
                subtitle,
                style: TextStyle(
                  fontFamily: 'Montserrat',
                  fontSize: 14,
                  fontWeight: FontWeight.w400,
                  color: isNext ? Color(0xFF388E3C).withOpacity(0.4) : Color(0xFF388E3C).withOpacity(0.7),
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }
} 