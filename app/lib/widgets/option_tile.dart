import 'package:flutter/material.dart';

class OptionTile extends StatelessWidget {
  final String title;
  final Widget child;

  const OptionTile({
    Key? key,
    required this.title,
    required this.child,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
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
}
