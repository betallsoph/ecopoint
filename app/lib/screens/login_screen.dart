import 'package:flutter/material.dart';
import 'home_screen.dart';

class LoginScreen extends StatefulWidget {
  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _phoneController = TextEditingController();

  void _loginWithGoogle(BuildContext context) {
    // fake login
    Navigator.pushReplacement(
        context, MaterialPageRoute(builder: (_) => HomeScreen()));
  }

  void _loginWithPhone(BuildContext context) {
    // fake login
    Navigator.pushReplacement(
        context, MaterialPageRoute(builder: (_) => HomeScreen()));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
                SizedBox(height: 40),
                
                // Welcome section
                Text(
                  'Chào mừng đến với',
                  style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontSize: 28,
                    fontWeight: FontWeight.w500,
                    color: Color(0xFF388E3C),
                  ),
                ),
                SizedBox(height: 8),
                Text(
                  'ecoPoint',
                  style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontSize: 48,
                    fontWeight: FontWeight.bold,
                    color: Color(0xFF388E3C),
                  ),
                ),
                SizedBox(height: 12),
                Text(
                  'Tái chế vật liệu thải của bạn một cách thông minh',
                  style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontSize: 16,
                    fontWeight: FontWeight.w400,
                    color: Color(0xFF388E3C).withOpacity(0.8),
                  ),
                ),
                
                Spacer(),
                
                // Phone field
                Text(
                  'Số điện thoại',
                  style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                    color: Color(0xFF388E3C),
                  ),
                ),
                SizedBox(height: 8),
                _buildInputTile(
                  icon: Icons.phone,
                  controller: _phoneController,
                ),
                
                SizedBox(height: 32),
                
                // Login button
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    onPressed: () => _loginWithPhone(context),
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
                
                SizedBox(height: 16),
                
                // Or divider
                Row(
                  children: [
                    Expanded(child: Divider(color: Color(0xFF388E3C).withOpacity(0.3))),
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 16),
                      child: Text(
                        'hoặc',
                        style: TextStyle(
                          fontFamily: 'Montserrat',
                          color: Color(0xFF388E3C).withOpacity(0.7),
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                    ),
                    Expanded(child: Divider(color: Color(0xFF388E3C).withOpacity(0.3))),
                  ],
                ),
                
                SizedBox(height: 16),
                
                // Google login button
                SizedBox(
                  width: double.infinity,
                  child: OutlinedButton.icon(
                    onPressed: () => _loginWithGoogle(context),
                    icon: Icon(Icons.login),
                    label: Text(
                      'Tiếp tục với Google',
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
        ),
      ),
    );
  }

  Widget _buildInputTile({
    required IconData icon,
    required TextEditingController controller,
  }) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(8),
        border: Border.all(
          color: Color(0xFF388E3C).withOpacity(0.3),
          width: 1,
        ),
      ),
      child: TextField(
        controller: controller,
        keyboardType: TextInputType.phone,
        style: TextStyle(
          fontFamily: 'Montserrat',
          fontSize: 16,
          fontWeight: FontWeight.w400,
        ),
        decoration: InputDecoration(
          hintText: '0123456789',
          hintStyle: TextStyle(
            fontFamily: 'Montserrat',
            color: Color(0xFF388E3C).withOpacity(0.5),
            fontWeight: FontWeight.w400,
          ),
          prefixIcon: Icon(icon, color: Color(0xFF388E3C)),
          border: InputBorder.none,
          contentPadding: EdgeInsets.all(16),
          focusedBorder: InputBorder.none,
          enabledBorder: InputBorder.none,
        ),
      ),
    );
  }

  @override
  void dispose() {
    _phoneController.dispose();
    super.dispose();
  }
}
