import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import '../services/graphql_service.dart';

class CreateOrderScreen extends StatefulWidget {
  @override
  _CreateOrderScreenState createState() => _CreateOrderScreenState();
}

class _CreateOrderScreenState extends State<CreateOrderScreen> {
  final _formKey = GlobalKey<FormState>();
  final _streetController = TextEditingController();
  final _districtController = TextEditingController();
  final _cityController = TextEditingController();
  final _notesController = TextEditingController();
  final _weightController = TextEditingController();
  
  List<String> _selectedWasteTypes = [];
  DateTime _scheduledTime = DateTime.now().add(Duration(hours: 1));
  
  final List<Map<String, dynamic>> _wasteTypeOptions = [
    {'id': 'PAPER', 'name': 'Giấy', 'icon': Icons.description},
    {'id': 'PLASTIC', 'name': 'Nhựa', 'icon': Icons.local_drink},
    {'id': 'METAL', 'name': 'Kim loại', 'icon': Icons.build},
    {'id': 'GLASS', 'name': 'Thủy tinh', 'icon': Icons.wine_bar},
    {'id': 'ELECTRONIC', 'name': 'Điện tử', 'icon': Icons.phone_android},
  ];

  @override
  void dispose() {
    _streetController.dispose();
    _districtController.dispose();
    _cityController.dispose();
    _notesController.dispose();
    _weightController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Tạo đơn hàng mới'),
        backgroundColor: Color(0xFF388E3C),
        foregroundColor: Colors.white,
      ),
      body: Form(
        key: _formKey,
        child: SingleChildScrollView(
          padding: EdgeInsets.all(16),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Waste Types Selection
              Text(
                'Loại rác cần thu gom',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 12),
              Wrap(
                spacing: 8,
                runSpacing: 8,
                children: _wasteTypeOptions.map((option) {
                  final isSelected = _selectedWasteTypes.contains(option['id']);
                  return FilterChip(
                    label: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(option['icon'], size: 16),
                        SizedBox(width: 4),
                        Text(option['name']),
                      ],
                    ),
                    selected: isSelected,
                    onSelected: (selected) {
                      setState(() {
                        if (selected) {
                          _selectedWasteTypes.add(option['id']);
                        } else {
                          _selectedWasteTypes.remove(option['id']);
                        }
                      });
                    },
                    selectedColor: Color(0xFF388E3C).withOpacity(0.3),
                    checkmarkColor: Color(0xFF388E3C),
                  );
                }).toList(),
              ),
              SizedBox(height: 24),
              
              // Weight
              Text(
                'Khối lượng ước tính (kg)',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 8),
              TextFormField(
                controller: _weightController,
                keyboardType: TextInputType.numberWithOptions(decimal: true),
                decoration: InputDecoration(
                  hintText: 'Ví dụ: 5.5',
                  border: OutlineInputBorder(),
                  prefixIcon: Icon(Icons.scale),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Vui lòng nhập khối lượng';
                  }
                  if (double.tryParse(value) == null) {
                    return 'Vui lòng nhập số hợp lệ';
                  }
                  return null;
                },
              ),
              SizedBox(height: 24),
              
              // Address
              Text(
                'Địa chỉ thu gom',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 12),
              TextFormField(
                controller: _streetController,
                decoration: InputDecoration(
                  labelText: 'Số nhà, tên đường',
                  border: OutlineInputBorder(),
                  prefixIcon: Icon(Icons.home),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Vui lòng nhập địa chỉ';
                  }
                  return null;
                },
              ),
              SizedBox(height: 12),
              Row(
                children: [
                  Expanded(
                    child: TextFormField(
                      controller: _districtController,
                      decoration: InputDecoration(
                        labelText: 'Quận/Huyện',
                        border: OutlineInputBorder(),
                        prefixIcon: Icon(Icons.location_city),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Vui lòng nhập quận/huyện';
                        }
                        return null;
                      },
                    ),
                  ),
                  SizedBox(width: 12),
                  Expanded(
                    child: TextFormField(
                      controller: _cityController,
                      decoration: InputDecoration(
                        labelText: 'Thành phố',
                        border: OutlineInputBorder(),
                        prefixIcon: Icon(Icons.business),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Vui lòng nhập thành phố';
                        }
                        return null;
                      },
                    ),
                  ),
                ],
              ),
              SizedBox(height: 24),
              
              // Scheduled Time
              Text(
                'Thời gian thu gom',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 8),
              ListTile(
                leading: Icon(Icons.access_time),
                title: Text(
                  '${_scheduledTime.day}/${_scheduledTime.month}/${_scheduledTime.year} ${_scheduledTime.hour}:${_scheduledTime.minute.toString().padLeft(2, '0')}',
                ),
                subtitle: Text('Nhấn để thay đổi'),
                onTap: _selectDateTime,
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(8),
                  side: BorderSide(color: Colors.grey),
                ),
              ),
              SizedBox(height: 24),
              
              // Notes
              Text(
                'Ghi chú (tùy chọn)',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF388E3C),
                ),
              ),
              SizedBox(height: 8),
              TextFormField(
                controller: _notesController,
                maxLines: 3,
                decoration: InputDecoration(
                  hintText: 'Thêm ghi chú cho người thu gom...',
                  border: OutlineInputBorder(),
                  prefixIcon: Icon(Icons.note),
                ),
              ),
              SizedBox(height: 32),
              
              // Create Order Button
              Mutation(
                options: MutationOptions(
                  document: gql(GraphQLMutations.createOrder),
                  onCompleted: (data) {
                    if (data != null && data['createOrder'] != null) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text('Đơn hàng đã được tạo thành công!'),
                          backgroundColor: Colors.green,
                        ),
                      );
                      Navigator.pop(context, true); // Return true to indicate success
                    }
                  },
                  onError: (error) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text('Lỗi: ${error.toString()}'),
                        backgroundColor: Colors.red,
                      ),
                    );
                  },
                ),
                builder: (runMutation, result) {
                  return SizedBox(
                    width: double.infinity,
                    child: ElevatedButton(
                      onPressed: result?.isLoading == true ? null : () => _createOrder(runMutation),
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Color(0xFF388E3C),
                        foregroundColor: Colors.white,
                        padding: EdgeInsets.symmetric(vertical: 16),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      child: result?.isLoading == true
                          ? Row(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                SizedBox(
                                  width: 20,
                                  height: 20,
                                  child: CircularProgressIndicator(
                                    color: Colors.white,
                                    strokeWidth: 2,
                                  ),
                                ),
                                SizedBox(width: 12),
                                Text('Đang tạo đơn hàng...'),
                              ],
                            )
                          : Text(
                              'Tạo đơn hàng',
                              style: TextStyle(
                                fontSize: 16,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                    ),
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
  
  void _createOrder(RunMutation runMutation) {
    if (!_formKey.currentState!.validate()) {
      return;
    }
    
    if (_selectedWasteTypes.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Vui lòng chọn ít nhất một loại rác'),
          backgroundColor: Colors.orange,
        ),
      );
      return;
    }
    
    final variables = {
      'input': {
        'wasteTypes': _selectedWasteTypes,
        'estimatedWeight': double.parse(_weightController.text),
        'pickupAddress': {
          'street': _streetController.text,
          'district': _districtController.text,
          'city': _cityController.text,
          'lat': 10.7769, // Default coordinates (Ho Chi Minh City)
          'lng': 106.7009,
        },
        'scheduledTime': _scheduledTime.toUtc().toIso8601String(),
        'notes': _notesController.text.isEmpty ? null : _notesController.text,
      },
    };
    
    runMutation(variables);
  }
  
  Future<void> _selectDateTime() async {
    final date = await showDatePicker(
      context: context,
      initialDate: _scheduledTime,
      firstDate: DateTime.now(),
      lastDate: DateTime.now().add(Duration(days: 30)),
    );
    
    if (date != null) {
      final time = await showTimePicker(
        context: context,
        initialTime: TimeOfDay.fromDateTime(_scheduledTime),
      );
      
      if (time != null) {
        setState(() {
          _scheduledTime = DateTime(
            date.year,
            date.month,
            date.day,
            time.hour,
            time.minute,
          );
        });
      }
    }
  }
}