import 'package:flutter/material.dart';

class SelectWasteScreen extends StatefulWidget {
  @override
  _SelectWasteScreenState createState() => _SelectWasteScreenState();
}

class _SelectWasteScreenState extends State<SelectWasteScreen> {
  final _weightController = TextEditingController();
  final _noteController = TextEditingController();
  
  List<String> _selectedWasteTypes = [];
  
  final List<Map<String, dynamic>> _wasteTypes = [
    {'name': 'Giấy', 'icon': Icons.description, 'color': Color(0xFF2196F3)},
    {'name': 'Nhựa', 'icon': Icons.local_drink, 'color': Color(0xFFFF9800)},
    {'name': 'Kim loại', 'icon': Icons.build, 'color': Color(0xFF9E9E9E)},
    {'name': 'Thủy tinh', 'icon': Icons.wine_bar, 'color': Color(0xFF4CAF50)},
    {'name': 'Thiết bị điện tử', 'icon': Icons.phone_android, 'color': Color(0xFF9C27B0)},
    {'name': 'Quần áo', 'icon': Icons.checkroom, 'color': Color(0xFFE91E63)},
    {'name': 'Pin/Ắc quy', 'icon': Icons.battery_full, 'color': Color(0xFFFF5722)},
    {'name': 'Khác', 'icon': Icons.category, 'color': Color(0xFF607D8B)},
  ];

  void _toggleWasteType(String wasteType) {
    setState(() {
      if (_selectedWasteTypes.contains(wasteType)) {
        _selectedWasteTypes.remove(wasteType);
      } else {
        _selectedWasteTypes.add(wasteType);
      }
    });
  }

  void _proceedToConfirmation() {
    if (_selectedWasteTypes.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Vui lòng chọn ít nhất một loại rác thải')),
      );
      return;
    }
    
    if (_weightController.text.trim().isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Vui lòng nhập khối lượng ước tính')),
      );
      return;
    }

    Navigator.pushNamed(
      context,
      '/confirm',
      arguments: {
        'wasteTypes': _selectedWasteTypes,
        'weight': _weightController.text.trim(),
        'note': _noteController.text.trim(),
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: Text(
          'Chọn loại rác thải',
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
        child: Column(
          children: [
            Expanded(
              child: SingleChildScrollView(
                padding: EdgeInsets.all(24.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    // Header text
                    Text(
                      'Loại rác thải cần thu gom',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                        color: Color(0xFF388E3C),
                      ),
                    ),
                    SizedBox(height: 8),
                    Text(
                      'Chọn các loại rác thải bạn muốn thu gom',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 14,
                        fontWeight: FontWeight.w400,
                        color: Color(0xFF388E3C).withOpacity(0.7),
                      ),
                    ),
                    
                    SizedBox(height: 24),
                    
                    // Waste type grid
                    GridView.builder(
                      shrinkWrap: true,
                      physics: NeverScrollableScrollPhysics(),
                      gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                        crossAxisCount: 2,
                        childAspectRatio: 1.2,
                        crossAxisSpacing: 12,
                        mainAxisSpacing: 12,
                      ),
                      itemCount: _wasteTypes.length,
                      itemBuilder: (context, index) {
                        final wasteType = _wasteTypes[index];
                        final isSelected = _selectedWasteTypes.contains(wasteType['name']);
                        
                        return GestureDetector(
                          onTap: () => _toggleWasteType(wasteType['name']),
                          child: Container(
                            decoration: BoxDecoration(
                              color: isSelected 
                                  ? Color(0xFF388E3C).withOpacity(0.1)
                                  : Colors.white,
                              borderRadius: BorderRadius.circular(12),
                              border: Border.all(
                                color: isSelected 
                                    ? Color(0xFF388E3C)
                                    : Color(0xFF388E3C).withOpacity(0.3),
                                width: isSelected ? 2 : 1,
                              ),
                            ),
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                Container(
                                  width: 48,
                                  height: 48,
                                  decoration: BoxDecoration(
                                    color: wasteType['color'].withOpacity(0.1),
                                    shape: BoxShape.circle,
                                  ),
                                  child: Icon(
                                    wasteType['icon'],
                                    color: wasteType['color'],
                                    size: 28,
                                  ),
                                ),
                                SizedBox(height: 8),
                                Text(
                                  wasteType['name'],
                                  style: TextStyle(
                                    fontFamily: 'Montserrat',
                                    fontSize: 14,
                                    fontWeight: FontWeight.w600,
                                    color: isSelected 
                                        ? Color(0xFF388E3C)
                                        : Color(0xFF388E3C).withOpacity(0.8),
                                  ),
                                  textAlign: TextAlign.center,
                                ),
                              ],
                            ),
                          ),
                        );
                      },
                    ),
                    
                    SizedBox(height: 32),
                    
                    // Weight input
                    Text(
                      'Khối lượng ước tính',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 18,
                        fontWeight: FontWeight.bold,
                        color: Color(0xFF388E3C),
                      ),
                    ),
                    SizedBox(height: 12),
                    Container(
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(8),
                        border: Border.all(
                          color: Color(0xFF388E3C).withOpacity(0.3),
                          width: 1,
                        ),
                      ),
                      child: TextField(
                        controller: _weightController,
                        keyboardType: TextInputType.number,
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 16,
                          fontWeight: FontWeight.w400,
                        ),
                        decoration: InputDecoration(
                          labelText: 'Nhập khối lượng (kg)',
                          labelStyle: TextStyle(
                            fontFamily: 'Montserrat',
                            color: Color(0xFF388E3C),
                            fontWeight: FontWeight.w500,
                          ),
                          prefixIcon: Icon(Icons.scale, color: Color(0xFF388E3C)),
                          suffixText: 'kg',
                          suffixStyle: TextStyle(
                            fontFamily: 'Montserrat',
                            color: Color(0xFF388E3C),
                            fontWeight: FontWeight.w500,
                          ),
                          border: InputBorder.none,
                          contentPadding: EdgeInsets.all(16),
                          focusedBorder: InputBorder.none,
                          enabledBorder: InputBorder.none,
                        ),
                      ),
                    ),
                    
                    SizedBox(height: 24),
                    
                    // Note input
                    Text(
                      'Ghi chú (tùy chọn)',
                      style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 18,
                        fontWeight: FontWeight.bold,
                        color: Color(0xFF388E3C),
                      ),
                    ),
                    SizedBox(height: 12),
                    Container(
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(8),
                        border: Border.all(
                          color: Color(0xFF388E3C).withOpacity(0.3),
                          width: 1,
                        ),
                      ),
                      child: TextField(
                        controller: _noteController,
                        maxLines: 4,
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          fontSize: 16,
                          fontWeight: FontWeight.w400,
                        ),
                        decoration: InputDecoration(
                          labelText: 'Mô tả thêm về rác thải...',
                          labelStyle: TextStyle(
                            fontFamily: 'Montserrat',
                            color: Color(0xFF388E3C),
                            fontWeight: FontWeight.w500,
                          ),
                          prefixIcon: Padding(
                            padding: EdgeInsets.only(top: 12),
                            child: Icon(Icons.note, color: Color(0xFF388E3C)),
                          ),
                          border: InputBorder.none,
                          contentPadding: EdgeInsets.all(16),
                          focusedBorder: InputBorder.none,
                          enabledBorder: InputBorder.none,
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ),
            
            // Bottom button
            Padding(
              padding: EdgeInsets.all(24.0),
              child: SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: _proceedToConfirmation,
                  child: Text(
                    'Tiếp tục',
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
            ),
          ],
        ),
      ),
    );
  }

  @override
  void dispose() {
    _weightController.dispose();
    _noteController.dispose();
    super.dispose();
  }
} 