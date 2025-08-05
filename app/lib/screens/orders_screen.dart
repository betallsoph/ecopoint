import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import '../services/graphql_service.dart';

class OrdersScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Đơn hàng của tôi'),
        backgroundColor: Color(0xFF388E3C),
        foregroundColor: Colors.white,
      ),
      body: Query(
        options: QueryOptions(
          document: gql(GraphQLQueries.myOrders),
          pollInterval: Duration(seconds: 10), // Auto refresh every 10 seconds
        ),
        builder: (QueryResult result, {VoidCallback? refetch, FetchMore? fetchMore}) {
          if (result.hasException) {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Icon(Icons.error_outline, size: 64, color: Colors.red),
                  SizedBox(height: 16),
                  Text(
                    'Lỗi kết nối',
                    style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                  ),
                  SizedBox(height: 8),
                  Text(
                    result.exception.toString(),
                    textAlign: TextAlign.center,
                    style: TextStyle(color: Colors.grey[600]),
                  ),
                  SizedBox(height: 16),
                  ElevatedButton(
                    onPressed: refetch,
                    child: Text('Thử lại'),
                  ),
                ],
              ),
            );
          }

          if (result.isLoading) {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  CircularProgressIndicator(color: Color(0xFF388E3C)),
                  SizedBox(height: 16),
                  Text('Đang tải đơn hàng...'),
                ],
              ),
            );
          }

          final orders = result.data?['myOrders'] as List?;
          
          if (orders == null || orders.isEmpty) {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Icon(Icons.inbox_outlined, size: 64, color: Colors.grey),
                  SizedBox(height: 16),
                  Text(
                    'Chưa có đơn hàng nào',
                    style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                  ),
                  SizedBox(height: 8),
                  Text(
                    'Hãy tạo đơn hàng đầu tiên của bạn!',
                    style: TextStyle(color: Colors.grey[600]),
                  ),
                ],
              ),
            );
          }

          return RefreshIndicator(
            onRefresh: () async {
              refetch?.call();
            },
            child: ListView.builder(
              padding: EdgeInsets.all(16),
              itemCount: orders.length,
              itemBuilder: (context, index) {
                final order = orders[index];
                return OrderCard(order: order);
              },
            ),
          );
        },
      ),
    );
  }
}

class OrderCard extends StatelessWidget {
  final Map<String, dynamic> order;

  const OrderCard({Key? key, required this.order}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final status = order['status'] as String;
    final wasteTypes = (order['wasteTypes'] as List).cast<String>();
    final estimatedWeight = order['estimatedWeight'] as double;
    final payment = order['payment'] as Map<String, dynamic>;
    final address = order['pickupAddress'] as Map<String, dynamic>;
    final scheduledTime = DateTime.parse(order['scheduledTime']);
    
    Color statusColor;
    String statusText;
    IconData statusIcon;
    
    switch (status) {
      case 'ORDER_STATUS_PENDING':
        statusColor = Colors.orange;
        statusText = 'Đang chờ';
        statusIcon = Icons.schedule;
        break;
      case 'ORDER_STATUS_ACCEPTED':
        statusColor = Colors.blue;
        statusText = 'Đã nhận';
        statusIcon = Icons.check_circle_outline;
        break;
      case 'ORDER_STATUS_IN_PROGRESS':
        statusColor = Colors.purple;
        statusText = 'Đang thu gom';
        statusIcon = Icons.local_shipping;
        break;
      case 'ORDER_STATUS_COMPLETED':
        statusColor = Colors.green;
        statusText = 'Hoàn thành';
        statusIcon = Icons.check_circle;
        break;
      case 'ORDER_STATUS_CANCELLED':
        statusColor = Colors.red;
        statusText = 'Đã hủy';
        statusIcon = Icons.cancel;
        break;
      default:
        statusColor = Colors.grey;
        statusText = 'Không xác định';
        statusIcon = Icons.help_outline;
    }

    return Card(
      margin: EdgeInsets.only(bottom: 16),
      elevation: 2,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      child: Padding(
        padding: EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Header với status
            Row(
              children: [
                Container(
                  padding: EdgeInsets.symmetric(horizontal: 12, vertical: 6),
                  decoration: BoxDecoration(
                    color: statusColor.withOpacity(0.1),
                    borderRadius: BorderRadius.circular(20),
                  ),
                  child: Row(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      Icon(statusIcon, size: 16, color: statusColor),
                      SizedBox(width: 6),
                      Text(
                        statusText,
                        style: TextStyle(
                          color: statusColor,
                          fontWeight: FontWeight.bold,
                          fontSize: 12,
                        ),
                      ),
                    ],
                  ),
                ),
                Spacer(),
                Text(
                  'ID: ${order['id'].toString().substring(0, 8)}...',
                  style: TextStyle(
                    color: Colors.grey[600],
                    fontSize: 12,
                  ),
                ),
              ],
            ),
            SizedBox(height: 12),
            
            // Waste types
            Row(
              children: [
                Icon(Icons.delete_outline, size: 20, color: Color(0xFF388E3C)),
                SizedBox(width: 8),
                Expanded(
                  child: Text(
                    'Loại rác: ${wasteTypes.map((type) => _getWasteTypeName(type)).join(', ')}',
                    style: TextStyle(fontWeight: FontWeight.w500),
                  ),
                ),
              ],
            ),
            SizedBox(height: 8),
            
            // Weight
            Row(
              children: [
                Icon(Icons.scale, size: 20, color: Color(0xFF388E3C)),
                SizedBox(width: 8),
                Text('Khối lượng ước tính: ${estimatedWeight.toStringAsFixed(1)} kg'),
              ],
            ),
            SizedBox(height: 8),
            
            // Address
            Row(
              children: [
                Icon(Icons.location_on_outlined, size: 20, color: Color(0xFF388E3C)),
                SizedBox(width: 8),
                Expanded(
                  child: Text(
                    '${address['street']}, ${address['district']}, ${address['city']}',
                    style: TextStyle(color: Colors.grey[700]),
                  ),
                ),
              ],
            ),
            SizedBox(height: 8),
            
            // Scheduled time
            Row(
              children: [
                Icon(Icons.access_time, size: 20, color: Color(0xFF388E3C)),
                SizedBox(width: 8),
                Text(
                  'Thời gian: ${_formatDateTime(scheduledTime)}',
                  style: TextStyle(color: Colors.grey[700]),
                ),
              ],
            ),
            SizedBox(height: 12),
            
            // Payment info
            Container(
              padding: EdgeInsets.all(12),
              decoration: BoxDecoration(
                color: Colors.grey[50],
                borderRadius: BorderRadius.circular(8),
              ),
              child: Row(
                children: [
                  Icon(Icons.payments_outlined, size: 20, color: Color(0xFF388E3C)),
                  SizedBox(width: 8),
                  Text(
                    'Thanh toán: ${_formatCurrency(payment['amount'])} ${payment['currency']}',
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                      color: Color(0xFF388E3C),
                    ),
                  ),
                  Spacer(),
                  Container(
                    padding: EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                    decoration: BoxDecoration(
                      color: payment['isPaid'] ? Colors.green : Colors.orange,
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: Text(
                      payment['isPaid'] ? 'Đã thanh toán' : 'Chưa thanh toán',
                      style: TextStyle(
                        color: Colors.white,
                        fontSize: 10,
                        fontWeight: FontWeight.bold,
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
  
  String _getWasteTypeName(String type) {
    switch (type) {
      case 'PAPER':
        return 'Giấy';
      case 'PLASTIC':
        return 'Nhựa';
      case 'METAL':
        return 'Kim loại';
      case 'GLASS':
        return 'Thủy tinh';
      case 'ELECTRONIC':
        return 'Điện tử';
      default:
        return type;
    }
  }
  
  String _formatDateTime(DateTime dateTime) {
    return '${dateTime.day}/${dateTime.month}/${dateTime.year} ${dateTime.hour}:${dateTime.minute.toString().padLeft(2, '0')}';
  }
  
  String _formatCurrency(dynamic amount) {
    final value = amount is int ? amount.toDouble() : amount as double;
    return value.toStringAsFixed(0).replaceAllMapped(
      RegExp(r'(\d)(?=(\d{3})+(?!\d))'),
      (Match m) => '${m[1]},',
    );
  }
}