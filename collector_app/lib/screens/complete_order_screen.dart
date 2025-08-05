import 'package:flutter/material.dart';

class CompleteOrderScreen extends StatefulWidget {
  @override
  _CompleteOrderScreenState createState() => _CompleteOrderScreenState();
}

class _CompleteOrderScreenState extends State<CompleteOrderScreen> {
  final _weightController = TextEditingController();
  final _notesController = TextEditingController();
  bool _photoTaken = false;

  void _completeOrder() {
    // Show success dialog
    showDialog(
      context: context,
      barrierDismissible: false,
      builder: (BuildContext context) {
        return AlertDialog(
          backgroundColor: Colors.white,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(16),
          ),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Icon(
                Icons.check_circle,
                color: Color(0xFF388E3C),
                size: 64,
              ),
              SizedBox(height: 16),
              Text(
                'Hoàn thành!',
                style: TextStyle(
                  fontFamily: 'Montserrat',
                  fontSize: 24,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 8),
              Text(
                'Đơn hàng đã được hoàn thành thành công',
                textAlign: TextAlign.center,
                style: TextStyle(
                  fontFamily: 'Montserrat',
                  fontSize: 16,
                  color: Colors.grey[600],
                ),
              ),
              SizedBox(height: 24),
              Container(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: () {
                    Navigator.popUntil(context, (route) => route.isFirst);
                  },
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Color(0xFF388E3C),
                    foregroundColor: Colors.white,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(12),
                    ),
                  ),
                  child: Text(
                    'Về trang chủ',
                    style: TextStyle(
                      fontFamily: 'Montserrat',
                      fontWeight: FontWeight.w600,
                    ),
                  ),
                ),
              ),
            ],
          ),
        );
      },
    );
  }

  Widget _buildInputField({
    required String label,
    required TextEditingController controller,
    String? hint,
    int maxLines = 1,
  }) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          label,
          style: TextStyle(
            fontFamily: 'Montserrat',
            fontSize: 16,
            fontWeight: FontWeight.w600,
            color: Color(0xFF388E3C),
          ),
        ),
        SizedBox(height: 8),
        Container(
          decoration: BoxDecoration(
            color: Color(0xFF388E3C).withOpacity(0.1),
            borderRadius: BorderRadius.circular(12),
            border: Border.all(
              color: Color(0xFF388E3C).withOpacity(0.3),
            ),
          ),
          child: TextField(
            controller: controller,
            maxLines: maxLines,
            decoration: InputDecoration(
              hintText: hint,
              border: InputBorder.none,
              contentPadding: EdgeInsets.all(16),
              hintStyle: TextStyle(
                fontFamily: 'Montserrat',
                color: Colors.grey[500],
              ),
            ),
            style: TextStyle(
              fontFamily: 'Montserrat',
              color: Colors.black87,
            ),
          ),
        ),
      ],
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
          'Hoàn thành đơn hàng',
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
            // Order summary
            Container(
              padding: EdgeInsets.all(16),
              decoration: BoxDecoration(
                color: Color(0xFF388E3C).withOpacity(0.1),
                borderRadius: BorderRadius.circular(12),
                border: Border.all(
                  color: Color(0xFF388E3C).withOpacity(0.3),
                ),
              ),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    'Thông tin đơn hàng',
                    style: TextStyle(
                      fontFamily: 'Montserrat',
                      fontSize: 18,
                      fontWeight: FontWeight.bold,
                      color: Color(0xFF388E3C),
                    ),
                  ),
                  SizedBox(height: 12),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        'Mã đơn:',
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
                          fontSize: 14,
                          fontWeight: FontWeight.w600,
                          color: Colors.black87,
                        ),
                      ),
                    ],
                  ),
                  SizedBox(height: 8),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        'Thu nhập:',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 14,
                          color: Colors.grey[600],
                        ),
                      ),
                      Text(
                        order['payment'],
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 16,
                          fontWeight: FontWeight.bold,
                          color: Color(0xFF388E3C),
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),

            SizedBox(height: 24),

            // Photo section
            Text(
              'Chụp ảnh xác nhận',
              style: TextStyle(
                fontFamily: 'Montserrat',
                fontSize: 18,
                fontWeight: FontWeight.bold,
                color: Color(0xFF388E3C),
              ),
            ),
            SizedBox(height: 8),
            Text(
              'Chụp ảnh rác thải đã thu gom để xác nhận',
              style: TextStyle(
                fontFamily: 'Montserrat',
                fontSize: 14,
                color: Colors.grey[600],
              ),
            ),
            SizedBox(height: 16),

            Container(
              width: double.infinity,
              height: 200,
              decoration: BoxDecoration(
                color: Colors.grey[200],
                borderRadius: BorderRadius.circular(12),
                border: Border.all(
                  color: _photoTaken ? Color(0xFF388E3C) : Colors.grey[400]!,
                  width: 2,
                ),
              ),
              child: InkWell(
                onTap: () {
                  setState(() {
                    _photoTaken = true;
                  });
                },
                borderRadius: BorderRadius.circular(12),
                child: Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Icon(
                        _photoTaken ? Icons.check_circle : Icons.camera_alt,
                        size: 48,
                        color: _photoTaken ? Color(0xFF388E3C) : Colors.grey[600],
                      ),
                      SizedBox(height: 8),
                      Text(
                        _photoTaken ? 'Đã chụp ảnh' : 'Chụp ảnh',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 16,
                          fontWeight: FontWeight.w600,
                          color: _photoTaken ? Color(0xFF388E3C) : Colors.grey[600],
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ),

            SizedBox(height: 24),

            // Weight input
            _buildInputField(
              label: 'Khối lượng thực tế (kg)',
              controller: _weightController,
              hint: 'Nhập khối lượng rác đã thu gom',
            ),

            SizedBox(height: 20),

            // Notes input
            _buildInputField(
              label: 'Ghi chú (không bắt buộc)',
              controller: _notesController,
              hint: 'Ghi chú thêm về quá trình thu gom...',
              maxLines: 3,
            ),

            SizedBox(height: 32),

            // Complete button
            Container(
              width: double.infinity,
              height: 56,
              child: ElevatedButton(
                onPressed: _photoTaken && _weightController.text.isNotEmpty
                    ? _completeOrder
                    : null,
                style: ElevatedButton.styleFrom(
                  backgroundColor: Color(0xFF388E3C),
                  foregroundColor: Colors.white,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(12),
                  ),
                  elevation: 2,
                ),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Icon(Icons.done_all, size: 20),
                    SizedBox(width: 8),
                    Text(
                      'Hoàn thành đơn hàng',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 16,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                  ],
                ),
              ),
            ),

            SizedBox(height: 16),
          ],
        ),
      ),
    );
  }
} 