import 'package:flutter/material.dart';
import 'login_screen.dart';

class SettingsScreen extends StatefulWidget {
  @override
  _SettingsScreenState createState() => _SettingsScreenState();
}

class _SettingsScreenState extends State<SettingsScreen> {
  bool _isDarkMode = false;
  String _language = 'Tiếng Việt';

  ThemeOption _themeOption = ThemeOption.system;
  LanguageOption _langOption = LanguageOption.vi;

  Widget _buildOptionTile({required String title, required Widget child}) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            title,
            style: TextStyle(
              fontFamily: 'Montserrat',
              fontWeight: FontWeight.bold,
              fontSize: 21,
              color: Color(0xFF388E3C),
            ),
          ),
          SizedBox(height: 8),
          child,
        ],
      ),
    );
  }

  Widget _buildThemeSelector() {
    return _buildOptionTile(
      title: 'Chế độ giao diện',
      child: Column(
        children: [
          RadioListTile<ThemeOption>(
            dense: true,
            activeColor: Color(0xFF388E3C),
            title: Row(
              children: [
                Icon(Icons.wb_sunny, color: Color(0xFF388E3C)),
                SizedBox(width: 8),
                Text('Sáng', style: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400)),
              ],
            ),
            value: ThemeOption.light,
            groupValue: _themeOption,
            onChanged: (val) => setState(() => _themeOption = val!),
          ),
          RadioListTile<ThemeOption>(
            dense: true,
            activeColor: Color(0xFF388E3C),
            title: Row(
              children: [
                Icon(Icons.dark_mode, color: Color(0xFF388E3C)),
                SizedBox(width: 8),
                Text('Tối', style: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400)),
              ],
            ),
            value: ThemeOption.dark,
            groupValue: _themeOption,
            onChanged: (val) => setState(() => _themeOption = val!),
          ),
          RadioListTile<ThemeOption>(
            dense: true,
            activeColor: Color(0xFF388E3C),
            title: Row(
              children: [
                Icon(Icons.phone_android, color: Color(0xFF388E3C)),
                SizedBox(width: 8),
                Text('Hệ thống', style: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400)),
              ],
            ),
            value: ThemeOption.system,
            groupValue: _themeOption,
            onChanged: (val) => setState(() => _themeOption = val!),
          ),
        ],
      ),
    );
  }

  Widget _buildLanguageSelector() {
    // Dynamic labels depending on currently selected language
    final String viLabel = _langOption == LanguageOption.vi ? 'Tiếng Việt' : 'Vietnamese';
    final String enLabel = _langOption == LanguageOption.vi ? 'Tiếng Anh' : 'English';

    return _buildOptionTile(
      title: 'Ngôn ngữ',
      child: Row(
        children: [
          Expanded(
            child: RadioListTile<LanguageOption>(
              dense: true,
              contentPadding: EdgeInsets.only(left: 16),
              title: Text(viLabel, style: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400)),
              value: LanguageOption.vi,
              groupValue: _langOption,
              onChanged: (val) => setState(() => _langOption = val!),
              activeColor: Color(0xFF388E3C),
              visualDensity: VisualDensity(horizontal: -4),
            ),
          ),
          Expanded(
            child: RadioListTile<LanguageOption>(
              dense: true,
              contentPadding: EdgeInsets.only(left: 16),
              title: Text(enLabel, style: TextStyle(fontFamily: 'Montserrat', fontSize: 16, fontWeight: FontWeight.w400)),
              value: LanguageOption.en,
              groupValue: _langOption,
              onChanged: (val) => setState(() => _langOption = val!),
              activeColor: Color(0xFF388E3C),
              visualDensity: VisualDensity(horizontal: -4),
            ),
          ),
        ],
      ),
    );
  }

  void _logout() {
    Navigator.pushAndRemoveUntil(
      context,
      MaterialPageRoute(builder: (_) => LoginScreen()),
      (route) => false,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Theme(
      data: Theme.of(context).copyWith(
        unselectedWidgetColor: Color(0xFF388E3C), // green ring for unselected radios
      ),
      child: Scaffold(
        appBar: AppBar(
          title: Text('Settings', style: TextStyle(
            fontFamily: 'Montserrat',
            fontWeight: FontWeight.bold,
            color: Colors.white,
          )),
          backgroundColor: Color(0xFF388E3C),
        ),
        body: Column(
          children: [
            Expanded(
              child: ListView(
                children: [
                  SizedBox(height: 16),
                  _buildThemeSelector(),
                  SizedBox(height: 8),
                  _buildLanguageSelector(),
                ],
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: _logout,
                  child: Text('Đăng xuất', style: TextStyle(
                    fontFamily: 'Montserrat',
                    fontWeight: FontWeight.bold,
                  )),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.red,
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
}

enum ThemeOption { light, dark, system }
enum LanguageOption { vi, en } 