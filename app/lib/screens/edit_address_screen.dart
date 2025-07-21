import 'package:flutter/material.dart';

class EditAddressScreen extends StatefulWidget {
  final String currentAddress;
  final String currentPhone;
  final Function(String, String) onSave;

  const EditAddressScreen({
    super.key,
    required this.currentAddress,
    required this.currentPhone,
    required this.onSave,
  });

  @override
  State<EditAddressScreen> createState() => _EditAddressScreenState();
}

class _EditAddressScreenState extends State<EditAddressScreen> {
  late TextEditingController _addressController;
  late TextEditingController _phoneController;
  bool _hasChanges = false;

  @override
  void initState() {
    super.initState();
    _addressController = TextEditingController(text: widget.currentAddress);
    _phoneController = TextEditingController(text: widget.currentPhone);
    
    // Lắng nghe thay đổi
    _addressController.addListener(_checkForChanges);
    _phoneController.addListener(_checkForChanges);
  }

  @override
  void dispose() {
    _addressController.dispose();
    _phoneController.dispose();
    super.dispose();
  }

  void _checkForChanges() {
    final hasChanges = _addressController.text != widget.currentAddress ||
        _phoneController.text != widget.currentPhone;
    
    if (hasChanges != _hasChanges) {
      setState(() {
        _hasChanges = hasChanges;
      });
    }
  }

  void _saveChanges() {
    if (_hasChanges) {
    widget.onSave(_addressController.text, _phoneController.text);
    Navigator.pop(context);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: true,
      body: SafeArea(
        child: Column(
          children: [
            Expanded(
              child: Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  children: [
                    // Khoảng trống để đẩy ô nhập xuống dưới
                    Expanded(child: SizedBox()),
                    
                    // Label và ô nhập địa chỉ
                    Align(
                      alignment: Alignment.centerLeft,
                      child: Text(
                        'Địa chỉ',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 18,
                          fontWeight: FontWeight.bold,
                          color: Color(0xFF388E3C),
                        ),
                      ),
                    ),
                    SizedBox(height: 8),
            TextField(
              controller: _addressController,
              decoration: InputDecoration(
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: BorderSide(color: Color(0xFF388E3C).withOpacity(0.3), width: 1.5),
                        ),
                        enabledBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: BorderSide(color: Color(0xFF388E3C).withOpacity(0.3), width: 1.5),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: BorderSide(color: Color(0xFF388E3C), width: 2),
            ),
                        contentPadding: EdgeInsets.symmetric(horizontal: 16, vertical: 18),
                      ),
                    ),
                    
            SizedBox(height: 20),
                    
                    // Label và ô nhập số điện thoại
                    Align(
                      alignment: Alignment.centerLeft,
                      child: Text(
                        'Số điện thoại',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 18,
                          fontWeight: FontWeight.bold,
                          color: Color(0xFF388E3C),
                        ),
                      ),
                    ),
                    SizedBox(height: 8),
            TextField(
              controller: _phoneController,
              decoration: InputDecoration(
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: BorderSide(color: Color(0xFF388E3C).withOpacity(0.3), width: 1.5),
                        ),
                        enabledBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: BorderSide(color: Color(0xFF388E3C).withOpacity(0.3), width: 1.5),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: BorderSide(color: Color(0xFF388E3C), width: 2),
                        ),
                        contentPadding: EdgeInsets.symmetric(horizontal: 16, vertical: 18),
              ),
              keyboardType: TextInputType.phone,
            ),
                    
                    SizedBox(height: 20),
                  ],
                ),
              ),
            ),
            
            // Các nút ở dưới cùng
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                children: [
                  // Nút lưu thay đổi
            SizedBox(
              width: double.infinity,
              child: ElevatedButton(
                      onPressed: _hasChanges ? _saveChanges : null,
                      child: Text('Lưu thay đổi', style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 16,
                        fontWeight: FontWeight.w600,
                      )),
                style: ElevatedButton.styleFrom(
                        backgroundColor: _hasChanges ? Color(0xFF388E3C) : Color(0xFF388E3C).withOpacity(0.25),
                        foregroundColor: _hasChanges ? Colors.white : Color(0xFF388E3C).withOpacity(0.7),
                        padding: EdgeInsets.symmetric(vertical: 16),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                        elevation: _hasChanges ? 2 : 0,
                      ),
                    ),
                  ),
                  
                  // Chỉ hiển thị nút trở lại khi chưa có thay đổi
                  if (!_hasChanges) ...[
                    SizedBox(height: 12),
                    
                    // Nút trở lại
                    SizedBox(
                      width: double.infinity,
                      child: OutlinedButton(
                        onPressed: () => Navigator.pop(context),
                        child: Text('Trở lại', style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 16,
                          fontWeight: FontWeight.w600,
                        )),
                        style: OutlinedButton.styleFrom(
                          backgroundColor: Colors.white,
                          foregroundColor: Color(0xFF388E3C),
                                                      side: BorderSide(color: Color(0xFF388E3C)),
                  padding: EdgeInsets.symmetric(vertical: 16),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                      ),
                ),
                  ],
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
