package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

// VerifyGoogleIDToken verifies Google ID token with Google's OAuth2 API
func VerifyGoogleIDToken(idToken string) (*GoogleUserInfo, error) {
	// Use Google's tokeninfo endpoint to verify the ID token
	url := "https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid token: status code %d", resp.StatusCode)
	}

	var tokenInfo struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
		Aud           string `json:"aud"`
		Iss           string `json:"iss"`
		Exp           string `json:"exp"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to decode token info: %v", err)
	}

	// Verify the token is from Google
	if tokenInfo.Iss != "https://accounts.google.com" && tokenInfo.Iss != "accounts.google.com" {
		return nil, fmt.Errorf("invalid issuer: %s", tokenInfo.Iss)
	}

	// Verify email is verified
	if tokenInfo.EmailVerified != "true" {
		return nil, fmt.Errorf("email not verified")
	}

	return &GoogleUserInfo{
		UID:     tokenInfo.Sub,
		Email:   tokenInfo.Email,
		Name:    tokenInfo.Name,
		Picture: tokenInfo.Picture,
	}, nil
}