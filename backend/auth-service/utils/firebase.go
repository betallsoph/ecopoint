package utils

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type GoogleUserInfo struct {
	UID     string `json:"uid"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

var firebaseAuth *auth.Client

// InitFirebase initializes Firebase Admin SDK
func InitFirebase(configPath string) error {
	opt := option.WithCredentialsFile(configPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	firebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		return err
	}

	log.Println("Firebase Admin SDK initialized")
	return nil
}

// VerifyGoogleIDToken verifies Google ID token and returns user info
func VerifyGoogleIDToken(idToken string) (*GoogleUserInfo, error) {
	// For MVP, we'll use a simple verification
	// In production, you should verify with Firebase Auth
	
	// TODO: Implement actual Firebase verification
	// For now, we'll create a mock user for testing
	return &GoogleUserInfo{
		UID:     "mock_uid_" + idToken[:10],
		Email:   "test@example.com",
		Name:    "Test User",
		Picture: "",
	}, nil

	// Production code would be:
	/*
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	return &GoogleUserInfo{
		UID:     token.UID,
		Email:   token.Claims["email"].(string),
		Name:    token.Claims["name"].(string),
		Picture: token.Claims["picture"].(string),
	}, nil
	*/
}