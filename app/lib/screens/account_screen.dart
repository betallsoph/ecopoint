import 'package:flutter/material.dart';
import 'edit_address_screen.dart';

enum TransportOption { van, truck, bike }

class AccountScreen extends StatefulWidget {
  const AccountScreen({super.key});

  @override
  State<AccountScreen> createState() => _AccountScreenState();
}

class _AccountScreenState extends State<AccountScreen> {
  TransportOption? _selected = TransportOption.van;

  String address = '123 Đường ABC, Quận 1, TP.HCM';
  String phone = '0909 123 456';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Column(
          children: [
            Expanded(
              child: SingleChildScrollView(
                child: Padding(
                  padding: const EdgeInsets.fromLTRB(16.0, 24, 16.0, 12),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    // Địa chỉ
                    _buildInfoTile(
                      icon: Icons.location_on,
                      title: 'Địa chỉ',
                      value: address,
                    ),

                    SizedBox(height: 16),

                    // Số điện thoại
                    _buildInfoTile(
                      icon: Icons.phone,
                      title: 'Số điện thoại',
                      value: phone,
                    ),


                  ],
                ),
              ),
            ),
          ),

          // Phương tiện giao hàng và nút ở dưới cùng
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              children: [
                // Phương tiện giao hàng
                Text(
                  'Phương tiện thu gom',
                  style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontSize: 21,
                    fontWeight: FontWeight.bold,
                    color: Color(0xFF388E3C),
                  ),
                ),

                SizedBox(height: 12),

                Column(
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        _buildTransportTile(
                          option: TransportOption.bike,
                          label: 'Bike',
                          icon: Icons.pedal_bike,
                        ),
                        _buildTransportTile(
                          option: TransportOption.van,
                          label: 'Van',
                          icon: Icons.local_shipping,
                        ),
                        _buildTransportTile(
                          option: TransportOption.truck,
                          label: 'Truck',
                          icon: Icons.fire_truck, // fallback: Icons.local_shipping
                        ),
                      ],
                    ),
                    SizedBox(height: 8),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Expanded(
                          child: Text(
                            'nhẹ',
                            textAlign: TextAlign.center,
                            style: TextStyle(
                              fontFamily: 'Montserrat',
                              fontSize: 12,
                              color: Color(0xFF388E3C).withOpacity(0.7),
                              fontWeight: FontWeight.w400,
                            ),
                          ),
                        ),
                        Expanded(
                          child: Text(
                            'trung bình',
                            textAlign: TextAlign.center,
                            style: TextStyle(
                              fontFamily: 'Montserrat',
                              fontSize: 12,
                              color: Color(0xFF388E3C).withOpacity(0.7),
                              fontWeight: FontWeight.w400,
                            ),
                          ),
                        ),
                        Expanded(
                          child: Text(
                            'nặng',
                            textAlign: TextAlign.center,
                            style: TextStyle(
                              fontFamily: 'Montserrat',
                              fontSize: 12,
                              color: Color(0xFF388E3C).withOpacity(0.7),
                              fontWeight: FontWeight.w400,
                            ),
                          ),
                        ),
                      ],
                    ),
                  ],
                ),

                SizedBox(height: 20),

                // Nút CHỈNH SỬA
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    onPressed: () {
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                          builder: (_) => EditAddressScreen(
                            currentAddress: address,
                            currentPhone: phone,
                            onSave: (newAddress, newPhone) {
                              setState(() {
                                address = newAddress;
                                phone = newPhone;
                              });
                            },
                          ),
                        ),
                      );
                    },
                    child: Text('Chỉnh sửa', style: TextStyle(
                      fontFamily: 'Montserrat',
                      fontWeight: FontWeight.bold,
                    )),
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
              ],
            ),
          ),
        ],
        ),
      ),
    );
  }

  Widget _buildInfoTile({
    required IconData icon,
    required String title,
    required String value,
  }) {
    return Container(
      decoration: BoxDecoration(
        color: Color(0xFFF5FBF2),
        borderRadius: BorderRadius.circular(8),
        border: Border.all(
          color: Color(0xFF388E3C).withOpacity(0.3),
          width: 1,
        ),
      ),
      padding: EdgeInsets.all(16),
      child: Row(
        children: [
          Icon(icon, color: Color(0xFF388E3C)),
          SizedBox(width: 12),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(title,
                    style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontWeight: FontWeight.bold, 
                        fontSize: 21,
                        color: Color(0xFF388E3C))),
                SizedBox(height: 4),
                Text(
                    title == 'Địa chỉ' ? value.replaceAll(', ', ',\n') : value,
                    style: TextStyle(
                        fontFamily: 'Montserrat',
                        fontSize: 16, 
                        fontWeight: FontWeight.w400),
                    maxLines: null,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildTransportTile({
    required TransportOption option,
    required String label,
    required IconData icon,
  }) {
    final isSelected = _selected == option;
    return Expanded(
      child: GestureDetector(
        onTap: () {
          setState(() {
            _selected = option;
          });
        },
        child: Container(
          height: 100,
          margin: EdgeInsets.symmetric(horizontal: 4),
          decoration: BoxDecoration(
            color: isSelected ? Color(0xFF388E3C).withOpacity(0.15) : Colors.white,
            borderRadius: BorderRadius.circular(8),
            border: Border.all(
              color: isSelected ? Color(0xFF388E3C) : Color(0xFF388E3C).withOpacity(0.3),
              width: isSelected ? 2 : 1,
            ),
          ),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(icon, color: Color(0xFF388E3C)),
              SizedBox(height: 8),
              Text(label, style: TextStyle(
                fontFamily: 'Montserrat',
                fontWeight: FontWeight.w500,
              )),
            ],
          ),
        ),
      ),
    );
  }
}
