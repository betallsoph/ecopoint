import 'package:flutter/material.dart';
import 'dart:async';

class WaitingForCollectorScreen extends StatefulWidget {
  @override
  _WaitingForCollectorScreenState createState() => _WaitingForCollectorScreenState();
}

class _WaitingForCollectorScreenState extends State<WaitingForCollectorScreen> with TickerProviderStateMixin {
  late AnimationController _pulseController;
  late Animation<double> _pulseAnimation;
  Timer? _navigationTimer;
  int _waitingTime = 0;
  Timer? _timeTimer;

  @override
  void initState() {
    super.initState();
    
    // Pulse animation for the waiting icon
    _pulseController = AnimationController(
      duration: Duration(seconds: 2),
      vsync: this,
    );
    _pulseAnimation = Tween<double>(
      begin: 0.8,
      end: 1.2,
    ).animate(CurvedAnimation(
      parent: _pulseController,
      curve: Curves.easeInOut,
    ));
    _pulseController.repeat(reverse: true);
    
    // Auto navigate to next screen after 5 seconds
    _navigationTimer = Timer(Duration(seconds: 5), () {
      if (mounted) {
        Navigator.pushReplacementNamed(context, '/onway');
      }
    });
    
    // Update waiting time every second
    _timeTimer = Timer.periodic(Duration(seconds: 1), (timer) {
      if (mounted) {
        setState(() {
          _waitingTime++;
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

  String _formatTime(int seconds) {
    int minutes = seconds ~/ 60;
    int remainingSeconds = seconds % 60;
    return '${minutes.toString().padLeft(2, '0')}:${remainingSeconds.toString().padLeft(2, '0')}';
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: Text(
          'Đang chờ thu gom',
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
                    // Animated waiting icon
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
                              Icons.hourglass_empty,
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
                      'Đang tìm nhân viên thu gom',
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
                      'Chúng tôi đang kết nối bạn với nhân viên thu gom gần nhất. Vui lòng chờ trong giây lát.',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 16,
                        fontWeight: FontWeight.w400,
                        color: Color(0xFF388E3C).withOpacity(0.8),
                      ),
                      textAlign: TextAlign.center,
                    ),
                    
                    SizedBox(height: 32),
                    
                    // Waiting time
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
                          Icon(Icons.timer, color: Color(0xFF388E3C)),
                          SizedBox(width: 8),
                          Text(
                            'Thời gian chờ: ${_formatTime(_waitingTime)}',
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
              
              // Info cards
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
                    _buildInfoRow(
                      icon: Icons.check_circle_outline,
                      title: 'Xác nhận đơn hàng',
                      subtitle: 'Đơn hàng của bạn đã được tiếp nhận',
                    ),
                    SizedBox(height: 16),
                    _buildInfoRow(
                      icon: Icons.search,
                      title: 'Tìm nhân viên',
                      subtitle: 'Đang tìm nhân viên thu gom phù hợp',
                      isActive: true,
                    ),
                    SizedBox(height: 16),
                    _buildInfoRow(
                      icon: Icons.directions_car,
                      title: 'Trên đường đến',
                      subtitle: 'Nhân viên sẽ đến địa chỉ của bạn',
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
                    Navigator.pop(context);
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

  Widget _buildInfoRow({
    required IconData icon,
    required String title,
    required String subtitle,
    bool isActive = false,
    bool isNext = false,
  }) {
    return Row(
      children: [
        Container(
          width: 40,
          height: 40,
          decoration: BoxDecoration(
            color: isActive
                ? Color(0xFF388E3C)
                : isNext
                    ? Color(0xFF388E3C).withOpacity(0.3)
                    : Color(0xFF388E3C),
            shape: BoxShape.circle,
          ),
          child: Icon(
            icon,
            color: isActive || !isNext ? Colors.white : Color(0xFF388E3C),
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