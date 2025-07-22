import 'package:flutter/material.dart';
import 'home_screen.dart';
import '../services/auth_service.dart';
import 'otp_screen.dart'; // Will be created next

class LoginScreen extends StatefulWidget {
  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _phoneController = TextEditingController();
  final AuthService _authService = AuthService();

  void _loginWithGoogle(BuildContext context) async {
    final user = await _authService.signInWithGoogle();
    if (user != null && mounted) {
      Navigator.pushReplacement(
          context, MaterialPageRoute(builder: (_) => HomeScreen()));
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Đăng nhập Google thất bại.')),
      );
    }
  }

  void _loginWithPhone(BuildContext context) {
    String phoneNumber = "+84" + _phoneController.text.trim();
    if (_phoneController.text.trim().isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Vui lòng nhập số điện thoại.')),
      );
      return;
    }
    _authService.verifyPhoneNumber(
      phoneNumber: phoneNumber,
      verificationCompleted: (credential) async {
        // Auto-retrieval or instant verification
        await _authService.signInWithCredential(credential);
        if (mounted) {
           Navigator.pushReplacement(context, MaterialPageRoute(builder: (_) => HomeScreen()));
        }
      },
      verificationFailed: (e) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text(e.message ?? 'Lỗi không xác định')),
        );
      },
      codeSent: (verificationId, resendToken) {
        Navigator.push(
          context,
          MaterialPageRoute(
            builder: (_) => OtpScreen(verificationId: verificationId),
          ),
        );
      },
      codeAutoRetrievalTimeout: (verificationId) {},
    );
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
                
                // Welcome section... (unchanged)
                Text('Chào mừng', style: TextStyle(fontFamily: 'Montserrat', fontSize: 28, fontWeight: FontWeight.w500, color: Color(0xFF388E3C))),
                SizedBox(height: 8),
                Text('ecoPoint Collector', style: TextStyle(fontFamily: 'Montserrat', fontSize: 48, fontWeight: FontWeight.bold, color: Color(0xFF388E3C))),
                SizedBox(height: 12),
                Text('Nhận đơn hàng thu gom và tăng thu nhập', style: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400, color: Color(0xFF388E3C).withOpacity(0.8))),
                
                Spacer(),
                
                // Phone field
                Text('Số điện thoại', style: TextStyle(fontFamily: 'Montserrat', fontSize: 18, fontWeight: FontWeight.bold, color: Color(0xFF388E3C))),
                SizedBox(height: 8),
                _buildInputTile(icon: Icons.phone, controller: _phoneController),
                
                SizedBox(height: 32),
                
                // Continue with Phone button
                _buildPhoneLoginButton(text: 'Tiếp tục', onPressed: () => _loginWithPhone(context)),
                
                SizedBox(height: 16),
                
                // Or divider
                Row(
                  children: [
                    Expanded(child: Divider(color: Color(0xFF388E3C).withOpacity(0.3))),
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 16),
                      child: Text('hoặc', style: TextStyle(fontFamily: 'Montserrat', color: Color(0xFF388E3C).withOpacity(0.7), fontWeight: FontWeight.w500)),
                    ),
                    Expanded(child: Divider(color: Color(0xFF388E3C).withOpacity(0.3))),
                  ],
                ),
                
                SizedBox(height: 16),
                
                // Google Sign-In Button
                _buildGoogleLoginButton(text: 'Tiếp tục với Google', onPressed: () => _loginWithGoogle(context)),
                
                SizedBox(height: 40),
            ],
          ),
        ),
      ),
    );
  }

  // Input tile for Phone Number
  Widget _buildInputTile({required IconData icon, required TextEditingController controller}) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(8),
        border: Border.all(color: Color(0xFF388E3C).withOpacity(0.3), width: 1),
      ),
      child: TextField(
        controller: controller,
        keyboardType: TextInputType.phone,
        style: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400),
        decoration: InputDecoration(
          hintText: '912345678',
          prefixText: '(+84) ',
          prefixStyle: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400, color: Colors.black),
          hintStyle: TextStyle(fontFamily: 'Montserrat', color: Color(0xFF388E3C).withOpacity(0.5), fontWeight: FontWeight.w400),
          prefixIcon: Icon(icon, color: Color(0xFF388E3C)),
          border: InputBorder.none,
          contentPadding: EdgeInsets.all(16),
        ),
      ),
    );
  }

  // Button for Phone Login
  Widget _buildPhoneLoginButton({required String text, required VoidCallback onPressed}) {
    return SizedBox(
      width: double.infinity,
      child: ElevatedButton(
        onPressed: onPressed,
        child: Text(text, style: TextStyle(fontFamily: 'Montserrat', fontWeight: FontWeight.bold, fontSize: 16)),
        style: ElevatedButton.styleFrom(
          backgroundColor: Color(0xFF388E3C),
          foregroundColor: Colors.white,
          padding: EdgeInsets.symmetric(vertical: 16),
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
        ),
      ),
    );
  }

  // Button for Google Login
  Widget _buildGoogleLoginButton({required String text, required VoidCallback onPressed}) {
    return SizedBox(
      width: double.infinity,
      child: OutlinedButton.icon(
        onPressed: onPressed,
        icon: Icon(Icons.g_mobiledata_outlined), // Example Google icon
        label: Text(text, style: TextStyle(fontFamily: 'Montserrat', fontWeight: FontWeight.w600, fontSize: 16)),
        style: OutlinedButton.styleFrom(
          side: BorderSide(color: Color(0xFF388E3C), width: 1.5),
          foregroundColor: Color(0xFF388E3C),
          padding: EdgeInsets.symmetric(vertical: 16),
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
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