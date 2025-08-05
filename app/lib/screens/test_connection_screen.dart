import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import '../services/graphql_service.dart';

class TestConnectionScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Test Backend Connection'),
        backgroundColor: Color(0xFF388E3C),
        foregroundColor: Colors.white,
      ),
      body: Padding(
        padding: EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Test GraphQL Connection
            Card(
              child: Padding(
                padding: EdgeInsets.all(16),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'GraphQL Connection Test',
                      style: TextStyle(
                        fontSize: 18,
                        fontWeight: FontWeight.bold,
                        color: Color(0xFF388E3C),
                      ),
                    ),
                    SizedBox(height: 12),
                    Query(
                      options: QueryOptions(
                        document: gql('''
                          query {
                            __schema {
                              queryType {
                                name
                              }
                            }
                          }
                        '''),
                      ),
                      builder: (QueryResult result, {VoidCallback? refetch, FetchMore? fetchMore}) {
                        if (result.hasException) {
                          return Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Row(
                                children: [
                                  Icon(Icons.error, color: Colors.red),
                                  SizedBox(width: 8),
                                  Text('Connection Failed', style: TextStyle(color: Colors.red, fontWeight: FontWeight.bold)),
                                ],
                              ),
                              SizedBox(height: 8),
                              Text(
                                result.exception.toString(),
                                style: TextStyle(fontSize: 12, color: Colors.grey[600]),
                              ),
                              SizedBox(height: 8),
                              ElevatedButton(
                                onPressed: refetch,
                                child: Text('Retry'),
                              ),
                            ],
                          );
                        }

                        if (result.isLoading) {
                          return Row(
                            children: [
                              SizedBox(
                                width: 20,
                                height: 20,
                                child: CircularProgressIndicator(strokeWidth: 2),
                              ),
                              SizedBox(width: 12),
                              Text('Connecting to backend...'),
                            ],
                          );
                        }

                        if (result.data != null) {
                          return Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Row(
                                children: [
                                  Icon(Icons.check_circle, color: Colors.green),
                                  SizedBox(width: 8),
                                  Text('Connected Successfully!', style: TextStyle(color: Colors.green, fontWeight: FontWeight.bold)),
                                ],
                              ),
                              SizedBox(height: 8),
                              Text('GraphQL Schema Available', style: TextStyle(color: Colors.grey[600])),
                            ],
                          );
                        }

                        return Text('No data received');
                      },
                    ),
                  ],
                ),
              ),
            ),
            
            SizedBox(height: 16),
            
            // Test Google Sign In
            Card(
              child: Padding(
                padding: EdgeInsets.all(16),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Authentication Test',
                      style: TextStyle(
                        fontSize: 18,
                        fontWeight: FontWeight.bold,
                        color: Color(0xFF388E3C),
                      ),
                    ),
                    SizedBox(height: 12),
                    Text('Test Google Sign-In with mock data:'),
                    SizedBox(height: 12),
                    Mutation(
                      options: MutationOptions(
                        document: gql(GraphQLMutations.googleSignIn),
                        onCompleted: (data) {
                          if (data != null && data['googleSignIn'] != null) {
                            ScaffoldMessenger.of(context).showSnackBar(
                              SnackBar(
                                content: Text('Authentication successful!'),
                                backgroundColor: Colors.green,
                              ),
                            );
                          }
                        },
                        onError: (error) {
                          ScaffoldMessenger.of(context).showSnackBar(
                            SnackBar(
                              content: Text('Auth error: ${error.toString()}'),
                              backgroundColor: Colors.red,
                            ),
                          );
                        },
                      ),
                      builder: (runMutation, result) {
                        return Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            ElevatedButton(
                              onPressed: result?.isLoading == true ? null : () {
                                runMutation({
                                  'input': {
                                    'idToken': 'test_token_flutter_${DateTime.now().millisecondsSinceEpoch}',
                                    'userType': 'CUSTOMER'
                                  }
                                });
                              },
                              style: ElevatedButton.styleFrom(
                                backgroundColor: Color(0xFF388E3C),
                                foregroundColor: Colors.white,
                              ),
                              child: result?.isLoading == true
                                  ? Row(
                                      mainAxisSize: MainAxisSize.min,
                                      children: [
                                        SizedBox(
                                          width: 16,
                                          height: 16,
                                          child: CircularProgressIndicator(
                                            color: Colors.white,
                                            strokeWidth: 2,
                                          ),
                                        ),
                                        SizedBox(width: 8),
                                        Text('Testing...'),
                                      ],
                                    )
                                  : Text('Test Mock Sign-In'),
                            ),
                            if (result?.data != null) ...[
                              SizedBox(height: 12),
                              Container(
                                padding: EdgeInsets.all(12),
                                decoration: BoxDecoration(
                                  color: Colors.green[50],
                                  borderRadius: BorderRadius.circular(8),
                                ),
                                child: Column(
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: [
                                    Text('✅ Authentication Success!', style: TextStyle(fontWeight: FontWeight.bold, color: Colors.green[700])),
                                    SizedBox(height: 4),
                                    Text('User: ${result!.data!['googleSignIn']['user']['displayName']}'),
                                    Text('Email: ${result!.data!['googleSignIn']['user']['email']}'),
                                    Text('Type: ${result!.data!['googleSignIn']['user']['userType']}'),
                                  ],
                                ),
                              ),
                            ],
                            if (result?.hasException == true) ...[
                              SizedBox(height: 12),
                              Container(
                                padding: EdgeInsets.all(12),
                                decoration: BoxDecoration(
                                  color: Colors.red[50],
                                  borderRadius: BorderRadius.circular(8),
                                ),
                                child: Column(
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: [
                                    Text('❌ Authentication Failed', style: TextStyle(fontWeight: FontWeight.bold, color: Colors.red[700])),
                                    SizedBox(height: 4),
                                    Text(result!.exception.toString(), style: TextStyle(fontSize: 12)),
                                  ],
                                ),
                              ),
                            ],
                          ],
                        );
                      },
                    ),
                  ],
                ),
              ),
            ),
            
            SizedBox(height: 24),
            
            // Instructions
            Card(
              color: Colors.blue[50],
              child: Padding(
                padding: EdgeInsets.all(16),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      children: [
                        Icon(Icons.info, color: Colors.blue),
                        SizedBox(width: 8),
                        Text('Test Instructions', style: TextStyle(fontWeight: FontWeight.bold, color: Colors.blue[700])),
                      ],
                    ),
                    SizedBox(height: 8),
                    Text('1. First test should show "Connected Successfully" if backend is running'),
                    Text('2. Second test will create a mock user and return authentication data'),
                    Text('3. If both pass, your Flutter app can communicate with the backend!'),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}