package auth

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var firebaseAuth *auth.Client

type FirebaseUserInfo struct {
	UID         string `json:"uid"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Verified    bool   `json:"email_verified"`
}

// InitFirebase initializes Firebase Admin SDK
func InitFirebase(configPath string) error {
	opt := option.WithCredentialsFile(configPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("failed to initialize firebase app: %v", err)
	}

	firebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		return fmt.Errorf("failed to initialize firebase auth: %v", err)
	}

	log.Println("Firebase Admin SDK initialized successfully")
	return nil
}

// VerifyIDToken verifies Firebase ID token and returns user info
func VerifyIDToken(idToken string) (*FirebaseUserInfo, error) {
	if firebaseAuth == nil {
		return nil, fmt.Errorf("firebase auth not initialized")
	}

	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %v", err)
	}

	// Extract user info from token claims
	userInfo := &FirebaseUserInfo{
		UID:      token.UID,
		Verified: true, // Firebase tokens are already verified
	}

	if email, ok := token.Claims["email"].(string); ok {
		userInfo.Email = email
	}

	if name, ok := token.Claims["name"].(string); ok {
		userInfo.Name = name
	}

	if picture, ok := token.Claims["picture"].(string); ok {
		userInfo.Picture = picture
	}

	if emailVerified, ok := token.Claims["email_verified"].(bool); ok {
		userInfo.Verified = emailVerified
	}

	return userInfo, nil
}

// GetUserByUID gets user info by Firebase UID
func GetUserByUID(uid string) (*auth.UserRecord, error) {
	if firebaseAuth == nil {
		return nil, fmt.Errorf("firebase auth not initialized")
	}

	user, err := firebaseAuth.GetUser(context.Background(), uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}