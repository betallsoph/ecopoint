import 'package:flutter/material.dart';

class OrderDetailsScreen extends StatefulWidget {
  @override
  _OrderDetailsScreenState createState() => _OrderDetailsScreenState();
}

class _OrderDetailsScreenState extends State<OrderDetailsScreen> {
  bool _isAccepted = false;

  Widget _buildInfoRow(IconData icon, String label, String value) {
    return Padding(
      padding: EdgeInsets.symmetric(vertical: 8),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Icon(icon, color: Color(0xFF388E3C), size: 20),
          SizedBox(width: 12),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  label,
                  style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontSize: 14,
                    color: Colors.grey[600],
                  ),
                ),
                SizedBox(height: 2),
                Text(
                  value,
                  style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontSize: 16,
                    fontWeight: FontWeight.w600,
                    color: Colors.black87,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildActionButton({
    required String text,
    required VoidCallback onPressed,
    required Color color,
    required Color textColor,
    IconData? icon,
  }) {
    return Container(
      width: double.infinity,
      height: 56,
      child: ElevatedButton(
        onPressed: onPressed,
        style: ElevatedButton.styleFrom(
          backgroundColor: color,
          foregroundColor: textColor,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(12),
          ),
          elevation: 2,
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            if (icon != null) ...[
              Icon(icon, size: 20),
              SizedBox(width: 8),
            ],
            Text(
              text,
              style: TextStyle(
                fontFamily: 'Montserrat',
                fontSize: 16,
                fontWeight: FontWeight.w600,
              ),
            ),
          ],
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    final order = ModalRoute.of(context)?.settings.arguments as Map<String, dynamic>?;
    
    if (order == null) {
      return Scaffold(
        appBar: AppBar(title: Text('Lỗi')),
        body: Center(child: Text('Không tìm thấy thông tin đơn hàng')),
      );
    }

    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        leading: IconButton(
          icon: Icon(Icons.arrow_back, color: Color(0xFF388E3C)),
          onPressed: () => Navigator.pop(context),
        ),
        title: Text(
          'Chi tiết đơn hàng',
          style: TextStyle(
            fontFamily: 'Montserrat',
            color: Color(0xFF388E3C),
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: SingleChildScrollView(
        padding: EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Order ID and status
            Container(
              padding: EdgeInsets.all(16),
              decoration: BoxDecoration(
                color: Color(0xFF388E3C).withOpacity(0.1),
                borderRadius: BorderRadius.circular(12),
                border: Border.all(
                  color: Color(0xFF388E3C).withOpacity(0.3),
                ),
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        'Mã đơn hàng',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 14,
                          color: Colors.grey[600],
                        ),
                      ),
                      Text(
                        order['id'],
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 18,
                          fontWeight: FontWeight.bold,
                          color: Color(0xFF388E3C),
                        ),
                      ),
                    ],
                  ),
                  Container(
                    padding: EdgeInsets.symmetric(horizontal: 12, vertical: 6),
                    decoration: BoxDecoration(
                      color: _isAccepted ? Colors.green : Colors.orange,
                      borderRadius: BorderRadius.circular(20),
                    ),
                    child: Text(
                      _isAccepted ? 'Đã nhận' : 'Chờ xác nhận',
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

            SizedBox(height: 24),

            // Customer information
            Text(
              'Thông tin khách hàng',
              style: TextStyle(
                fontFamily: 'Montserrat',
                fontSize: 20,
                fontWeight: FontWeight.bold,
                color: Color(0xFF388E3C),
              ),
            ),
            SizedBox(height: 16),
            
            _buildInfoRow(Icons.person, 'Tên khách hàng', order['customerName']),
            _buildInfoRow(Icons.location_on, 'Địa chỉ', order['address']),
            _buildInfoRow(Icons.access_time, 'Thời gian đặt', order['timeAgo']),

            SizedBox(height: 24),

            // Waste information
            Text(
              'Thông tin rác thải',
              style: TextStyle(
                fontFamily: 'Montserrat',
                fontSize: 20,
                fontWeight: FontWeight.bold,
                color: Color(0xFF388E3C),
              ),
            ),
            SizedBox(height: 16),

            // Waste types
            Wrap(
              spacing: 12,
              runSpacing: 8,
              children: (order['wasteTypes'] as List<String>).map((type) {
                return Container(
                  padding: EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                  decoration: BoxDecoration(
                    color: Color(0xFF388E3C).withOpacity(0.1),
                    borderRadius: BorderRadius.circular(20),
                    border: Border.all(
                      color: Color(0xFF388E3C).withOpacity(0.3),
                    ),
                  ),
                  child: Text(
                    type,
                    style: TextStyle(
                      fontFamily: 'Montserrat',
                      fontSize: 14,
                      fontWeight: FontWeight.w600,
                      color: Color(0xFF388E3C),
                    ),
                  ),
                );
              }).toList(),
            ),

            SizedBox(height: 16),

            _buildInfoRow(Icons.scale, 'Khối lượng ước tính', order['estimatedWeight']),
            _buildInfoRow(Icons.directions, 'Khoảng cách', order['distance']),
            _buildInfoRow(Icons.payment, 'Thu nhập dự kiến', order['payment']),

            SizedBox(height: 32),

            // Action buttons
            if (!_isAccepted) ...[
              Row(
                children: [
                  Expanded(
                    child: _buildActionButton(
                      text: 'Từ chối',
                      onPressed: () {
                        Navigator.pop(context);
                      },
                      color: Colors.grey[300]!,
                      textColor: Colors.grey[700]!,
                    ),
                  ),
                  SizedBox(width: 16),
                  Expanded(
                    child: _buildActionButton(
                      text: 'Nhận đơn',
                      onPressed: () {
                        setState(() {
                          _isAccepted = true;
                        });
                      },
                      color: Color(0xFF388E3C),
                      textColor: Colors.white,
                      icon: Icons.check,
                    ),
                  ),
                ],
              ),
            ] else ...[
              _buildActionButton(
                text: 'Bắt đầu thu gom',
                onPressed: () {
                  Navigator.pushNamed(
                    context,
                    '/navigation',
                    arguments: order,
                  );
                },
                color: Color(0xFF388E3C),
                textColor: Colors.white,
                icon: Icons.navigation,
              ),
            ],

            SizedBox(height: 16),
          ],
        ),
      ),
    );
  }
} 