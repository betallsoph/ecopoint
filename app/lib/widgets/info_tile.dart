import 'package:flutter/material.dart';

class InfoTile extends StatelessWidget {
  final IconData icon;
  final String title;
  final String value;
  final VoidCallback? onTap;

  const InfoTile({
    Key? key,
    required this.icon,
    required this.title,
    required this.value,
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
        child: Row(
          children: [
            Icon(
              icon,
              color: Color(0xFF388E3C),
              size: 24,
            ),
            SizedBox(width: 12),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    title,
                    style: TextStyle(
                      fontFamily: 'Montserrat',
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                      color: Color(0xFF388E3C),
                    ),
                  ),
                  SizedBox(height: 4),
                  Text(
                    value,
                    style: TextStyle(
                      fontFamily: 'Montserrat',
                      fontSize: 14,
                      fontWeight: FontWeight.w400,
                      color: Color(0xFF388E3C).withOpacity(0.8),
                    ),
                  ),
                ],
              ),
            ),
            if (onTap != null)
              Icon(
                Icons.chevron_right,
                color: Color(0xFF388E3C).withOpacity(0.6),
                size: 20,
              ),
          ],
        ),
      ),
    );
  }
}
