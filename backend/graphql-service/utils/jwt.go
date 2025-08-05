package utils

import (
	"errors"
	"os"

	"ecopoint-backend/shared/models"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   string          `json:"user_id"`
	UserType models.UserType `json:"user_type"`
	jwt.RegisteredClaims
}

// ValidateAccessToken validates and parses an access token
func ValidateAccessToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET not configured")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}