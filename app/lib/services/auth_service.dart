import 'package:firebase_auth/firebase_auth.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import 'graphql_service.dart';

class AuthService {
  final FirebaseAuth _auth = FirebaseAuth.instance;
  final GoogleSignIn _googleSignIn = GoogleSignIn();
  final GraphQLService _graphqlService = GraphQLService.instance;

  // GETTER for USER stream to check auth state
  Stream<User?> get user => _auth.authStateChanges();

  // SIGN IN WITH GOOGLE - Updated to use GraphQL backend
  Future<User?> signInWithGoogle() async {
    try {
      // Trigger the authentication flow.
      final GoogleSignInAccount? googleUser = await _googleSignIn.signIn();

      // If the user cancels the sign-in, return null.
      if (googleUser == null) {
        return null;
      }

      // Obtain the auth details from the request.
      final GoogleSignInAuthentication googleAuth = await googleUser.authentication;

      // Get Firebase ID token
      final idToken = googleAuth.idToken;
      if (idToken == null) {
        print("No ID token received");
        return null;
      }

      // Call GraphQL backend with Firebase ID token
      final result = await _graphqlService.client.mutate(
        MutationOptions(
          document: gql(GraphQLMutations.googleSignIn),
          variables: {
            'input': {
              'idToken': idToken,
              'userType': 'CUSTOMER', // Default for customer app
            },
          },
        ),
      );

      if (result.hasException) {
        print("GraphQL Error: ${result.exception.toString()}");
        return null;
      }

      final data = result.data?['googleSignIn'];
      if (data != null) {
        // Save tokens from backend
        await _graphqlService.saveTokens(
          data['accessToken'],
          data['refreshToken'],
        );

        // Create a Firebase credential for local auth state
        final AuthCredential credential = GoogleAuthProvider.credential(
          accessToken: googleAuth.accessToken,
          idToken: googleAuth.idToken,
        );

        // Sign in to Firebase locally for auth state management
        final UserCredential userCredential = await _auth.signInWithCredential(credential);
        return userCredential.user;
      }

      return null;
    } catch (e) {
      print("Error during Google Sign-In: $e");
      return null;
    }
  }

  // SIGN IN WITH CREDENTIAL (used by both phone and google)
  Future<User?> signInWithCredential(AuthCredential credential) async {
    try {
      final UserCredential userCredential = await _auth.signInWithCredential(credential);
      return userCredential.user;
    } catch (e) {
      print("Error signing in with credential: $e");
      return null;
    }
  }

  // VERIFY PHONE NUMBER
  Future<void> verifyPhoneNumber({
    required String phoneNumber,
    required Function(PhoneAuthCredential) verificationCompleted,
    required Function(FirebaseAuthException) verificationFailed,
    required Function(String, int?) codeSent,
    required Function(String) codeAutoRetrievalTimeout,
  }) async {
    await _auth.verifyPhoneNumber(
      phoneNumber: phoneNumber,
      verificationCompleted: verificationCompleted,
      verificationFailed: verificationFailed,
      codeSent: codeSent,
      codeAutoRetrievalTimeout: codeAutoRetrievalTimeout,
    );
  }

  // SIGN OUT
  Future<void> signOut() async {
    await _googleSignIn.signOut();
    await _auth.signOut();
  }
} 