import 'package:flutter/material.dart';

class FinishCollectionScreen extends StatefulWidget {
  @override
  _FinishCollectionScreenState createState() => _FinishCollectionScreenState();
}

class _FinishCollectionScreenState extends State<FinishCollectionScreen> with TickerProviderStateMixin {
  late AnimationController _scaleController;
  late Animation<double> _scaleAnimation;
  int _rating = 0;

  @override
  void initState() {
    super.initState();
    
    // Scale animation for the success icon
    _scaleController = AnimationController(
      duration: Duration(milliseconds: 800),
      vsync: this,
    );
    _scaleAnimation = Tween<double>(
      begin: 0.0,
      end: 1.0,
    ).animate(CurvedAnimation(
      parent: _scaleController,
      curve: Curves.elasticOut,
    ));
    
    // Start animation when screen loads
    _scaleController.forward();
  }

  @override
  void dispose() {
    _scaleController.dispose();
    super.dispose();
  }

  void _submitRating() {
    // Handle rating submission
    Navigator.pushNamedAndRemoveUntil(context, '/home', (route) => false);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: Text(
          'Hoàn thành thu gom',
          style: TextStyle(
            fontFamily: 'Montserrat',
            fontWeight: FontWeight.bold,
            color: Colors.white,
          ),
        ),
        backgroundColor: Color(0xFF388E3C),
        automaticallyImplyLeading: false,
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            children: [
              Expanded(
                child: SingleChildScrollView(
                  child: Column(
                    children: [
                      SizedBox(height: 40),
                      
                      // Success animation
                      AnimatedBuilder(
                        animation: _scaleAnimation,
                        builder: (context, child) {
                          return Transform.scale(
                            scale: _scaleAnimation.value,
                            child: Container(
                              width: 120,
                              height: 120,
                              decoration: BoxDecoration(
                                color: Color(0xFF4CAF50),
                                shape: BoxShape.circle,
                                boxShadow: [
                                  BoxShadow(
                                    color: Color(0xFF4CAF50).withOpacity(0.3),
                                    blurRadius: 20,
                                    offset: Offset(0, 10),
                                  ),
                                ],
                              ),
                              child: Icon(
                                Icons.check,
                                size: 60,
                                color: Colors.white,
                              ),
                            ),
                          );
                        },
                      ),
                      
                      SizedBox(height: 32),
                      
                      // Success message
                      Text(
                        'Thu gom thành công!',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 28,
                          fontWeight: FontWeight.bold,
                          color: Color(0xFF388E3C),
                        ),
                        textAlign: TextAlign.center,
                      ),
                      
                      SizedBox(height: 16),
                      
                      Text(
                        'Cảm ơn bạn đã sử dụng dịch vụ ecoPoint. Vật liệu của bạn sẽ được tái chế một cách có trách nhiệm.',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 16,
                          fontWeight: FontWeight.w400,
                          color: Color(0xFF388E3C).withOpacity(0.8),
                        ),
                        textAlign: TextAlign.center,
                      ),
                      
                      SizedBox(height: 40),
                      
                      // Collection summary
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
                        padding: EdgeInsets.all(24),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              'Tóm tắt thu gom',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                fontSize: 20,
                                fontWeight: FontWeight.bold,
                                color: Color(0xFF388E3C),
                              ),
                            ),
                            SizedBox(height: 20),
                            
                            _buildSummaryRow(
                              icon: Icons.schedule,
                              label: 'Thời gian bắt đầu',
                              value: '14:30',
                            ),
                            SizedBox(height: 12),
                            _buildSummaryRow(
                              icon: Icons.check_circle,
                              label: 'Thời gian kết thúc',
                              value: '14:45',
                            ),
                            SizedBox(height: 12),
                            _buildSummaryRow(
                              icon: Icons.scale,
                              label: 'Khối lượng thu gom',
                              value: '15.5 kg',
                            ),
                            SizedBox(height: 12),
                            _buildSummaryRow(
                              icon: Icons.recycling,
                              label: 'Loại vật liệu',
                              value: 'Giấy, Nhựa, Kim loại',
                            ),
                            SizedBox(height: 12),
                            _buildSummaryRow(
                              icon: Icons.person,
                              label: 'Thu gom bởi',
                              value: 'Nguyễn Văn A',
                            ),
                          ],
                        ),
                      ),
                      
                      SizedBox(height: 32),
                      
                      // Rating section
                      Container(
                        decoration: BoxDecoration(
                          color: Color(0xFF388E3C).withOpacity(0.1),
                          borderRadius: BorderRadius.circular(12),
                          border: Border.all(
                            color: Color(0xFF388E3C).withOpacity(0.3),
                            width: 1,
                          ),
                        ),
                        padding: EdgeInsets.all(24),
                        child: Column(
                          children: [
                            Text(
                              'Đánh giá dịch vụ',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                fontSize: 18,
                                fontWeight: FontWeight.bold,
                                color: Color(0xFF388E3C),
                              ),
                            ),
                            SizedBox(height: 16),
                            Text(
                              'Bạn cảm thấy dịch vụ thế nào?',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                fontSize: 14,
                                fontWeight: FontWeight.w400,
                                color: Color(0xFF388E3C).withOpacity(0.8),
                              ),
                            ),
                            SizedBox(height: 20),
                            
                            // Star rating
                            Row(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: List.generate(5, (index) {
                                return GestureDetector(
                                  onTap: () {
                                    setState(() {
                                      _rating = index + 1;
                                    });
                                  },
                                  child: Padding(
                                    padding: EdgeInsets.symmetric(horizontal: 4),
                                    child: Icon(
                                      Icons.star,
                                      size: 40,
                                      color: index < _rating ? Colors.amber : Colors.grey[300],
                                    ),
                                  ),
                                );
                              }),
                            ),
                            
                            SizedBox(height: 20),
                            
                            // Rating text
                            if (_rating > 0)
                              Text(
                                _getRatingText(_rating),
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
                      
                      SizedBox(height: 32),
                      
                      // Environmental impact
                      Container(
                        decoration: BoxDecoration(
                          color: Colors.white,
                          borderRadius: BorderRadius.circular(12),
                          border: Border.all(
                            color: Color(0xFF388E3C).withOpacity(0.3),
                            width: 1,
                          ),
                        ),
                        padding: EdgeInsets.all(24),
                        child: Column(
                          children: [
                            Icon(
                              Icons.eco,
                              size: 48,
                              color: Color(0xFF4CAF50),
                            ),
                            SizedBox(height: 16),
                            Text(
                              'Tác động môi trường',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                fontSize: 18,
                                fontWeight: FontWeight.bold,
                                color: Color(0xFF388E3C),
                              ),
                            ),
                            SizedBox(height: 12),
                            Text(
                              'Bạn đã giúp tiết kiệm khoảng 12.3 kg CO₂ và bảo vệ môi trường. Cảm ơn bạn!',
                              style: TextStyle(
                                fontFamily: 'Montserrat',
                                fontSize: 14,
                                fontWeight: FontWeight.w400,
                                color: Color(0xFF388E3C).withOpacity(0.8),
                              ),
                              textAlign: TextAlign.center,
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              
              SizedBox(height: 24),
              
              // Action buttons
              Column(
                children: [
                  if (_rating > 0)
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton(
                        onPressed: _submitRating,
                        child: Text(
                          'Gửi đánh giá',
                          style: TextStyle(
                            fontFamily: 'Montserrat',
                            fontWeight: FontWeight.bold,
                            fontSize: 16,
                          ),
                        ),
                        style: ElevatedButton.styleFrom(
                          backgroundColor: Color(0xFF388E3C),
                          foregroundColor: Colors.white,
                          padding: EdgeInsets.symmetric(vertical: 16),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                      ),
                    ),
                  
                  if (_rating > 0) SizedBox(height: 12),
                  
                  SizedBox(
                    width: double.infinity,
                    child: OutlinedButton(
                      onPressed: () {
                        Navigator.pushNamedAndRemoveUntil(context, '/home', (route) => false);
                      },
                      child: Text(
                        'Về trang chủ',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontWeight: FontWeight.w600,
                          fontSize: 16,
                        ),
                      ),
                      style: OutlinedButton.styleFrom(
                        side: BorderSide(color: Color(0xFF388E3C), width: 1.5),
                        foregroundColor: Color(0xFF388E3C),
                        padding: EdgeInsets.symmetric(vertical: 16),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildSummaryRow({
    required IconData icon,
    required String label,
    required String value,
  }) {
    return Row(
      children: [
        Icon(icon, color: Color(0xFF388E3C), size: 20),
        SizedBox(width: 12),
        Expanded(
          child: Text(
            label,
            style: TextStyle(
              fontFamily: 'Montserrat',
              fontSize: 14,
              fontWeight: FontWeight.w500,
              color: Color(0xFF388E3C).withOpacity(0.8),
            ),
          ),
        ),
        Text(
          value,
          style: TextStyle(
            fontFamily: 'Montserrat',
            fontSize: 14,
            fontWeight: FontWeight.bold,
            color: Color(0xFF388E3C),
          ),
        ),
      ],
    );
  }

  String _getRatingText(int rating) {
    switch (rating) {
      case 1:
        return 'Cần cải thiện';
      case 2:
        return 'Tạm được';
      case 3:
        return 'Tốt';
      case 4:
        return 'Rất tốt';
      case 5:
        return 'Xuất sắc';
      default:
        return '';
    }
  }
} 