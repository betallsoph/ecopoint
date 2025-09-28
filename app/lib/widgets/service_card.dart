import 'package:flutter/material.dart';

class ServiceCard extends StatelessWidget {
  final IconData icon;
  final String title;
  final String subtitle;
  final VoidCallback? onTap;

  const ServiceCard({
    Key? key,
    required this.icon,
    required this.title,
    required this.subtitle,
    this.onTap,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        decoration: BoxDecoration(
          color: Color(0xFF388E3C).withOpacity(0.1),
          borderRadius: BorderRadius.circular(8),
          border: Border.all(
            color: Color(0xFF388E3C).withOpacity(0.3),
            width: 1,
          ),
        ),
        padding: EdgeInsets.all(16),
        child: Column(
          children: [
            Icon(icon, color: Color(0xFF388E3C), size: 32),
            SizedBox(height: 8),
            Text(
              title,
              style: TextStyle(
                fontFamily: 'Montserrat',
                fontWeight: FontWeight.bold,
                fontSize: 14,
                color: Color(0xFF388E3C),
              ),
              textAlign: TextAlign.center,
            ),
            SizedBox(height: 4),
            Text(
              subtitle,
              style: TextStyle(
                fontFamily: 'Montserrat',
                fontSize: 12,
                fontWeight: FontWeight.w400,
                color: Color(0xFF388E3C).withOpacity(0.7),
              ),
              textAlign: TextAlign.center,
            ),
          ],
        ),
      ),
    );
  }
}
