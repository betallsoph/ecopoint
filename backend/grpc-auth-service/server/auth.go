package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"ecopoint-backend/auth-service/utils"
	"ecopoint-backend/shared/config"
	"ecopoint-backend/shared/database"
	"ecopoint-backend/shared/models"
	pb_auth "ecopoint-backend/proto"
	pb_user "ecopoint-backend/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthServer struct {
	pb_auth.UnimplementedAuthServiceServer
	config *config.Config
	userDB *mongo.Collection
}

func NewAuthServer(cfg *config.Config) *AuthServer {
	db := database.GetDatabase("ecopoint")
	return &AuthServer{
		config: cfg,
		userDB: db.Collection("users"),
	}
}

func (s *AuthServer) GoogleSignIn(ctx context.Context, req *pb_auth.GoogleSignInRequest) (*pb_auth.GoogleSignInResponse, error) {
	// TODO: Verify Firebase ID token
	// For now, create a mock user
	
	// Check if user exists by some identifier (in real implementation, use Firebase UID)
	mockUID := "firebase_" + req.IdToken[:10]
	
	var existingUser models.User
	err := s.userDB.FindOne(ctx, bson.M{"firebase_uid": mockUID}).Decode(&existingUser)
	
	if err == mongo.ErrNoDocuments {
		// Create new user
		now := time.Now()
		newUser := models.User{
			ID:          primitive.NewObjectID(),
			FirebaseUID: mockUID,
			Email:       "user@example.com", // In real implementation, get from Firebase
			DisplayName: "Test User",
			UserType:    models.UserType(req.UserType.String()),
			IsActive:    true,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		_, err = s.userDB.InsertOne(ctx, newUser)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
		}

		existingUser = newUser
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	// Generate JWT tokens (simplified)
	// Generate real JWT tokens
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET not configured")
	}
	
	accessToken, err := utils.GenerateAccessToken(existingUser.ID.Hex(), existingUser.UserType, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}
	
	refreshToken, err := utils.GenerateRefreshToken(existingUser.ID.Hex(), secret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	// Convert to protobuf User
	pbUser := &pb_user.User{
		Id:          existingUser.ID.Hex(),
		FirebaseUid: existingUser.FirebaseUID,
		Email:       existingUser.Email,
		DisplayName: existingUser.DisplayName,
		UserType:    pb_user.UserType(pb_user.UserType_value[string(existingUser.UserType)]),
		IsActive:    existingUser.IsActive,
		CreatedAt:   timestamppb.New(existingUser.CreatedAt),
		UpdatedAt:   timestamppb.New(existingUser.UpdatedAt),
	}

	return &pb_auth.GoogleSignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         pbUser,
	}, nil
}

func (s *AuthServer) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenRequest) (*pb_auth.RefreshTokenResponse, error) {
	// TODO: Implement token refresh logic
	return &pb_auth.RefreshTokenResponse{
		AccessToken: "new_access_token",
	}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *pb_auth.ValidateTokenRequest) (*pb_auth.ValidateTokenResponse, error) {
	// TODO: Implement token validation logic
	// For now, return mock validation
	return &pb_auth.ValidateTokenResponse{
		Valid:    true,
		UserId:   "mock_user_id",
		UserType: pb_user.UserType_USER_TYPE_CUSTOMER,
	}, nil
}

func (s *AuthServer) GetProfile(ctx context.Context, req *pb_auth.GetProfileRequest) (*pb_auth.GetProfileResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID: %v", err)
	}

	var user models.User
	err = s.userDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, "User not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	// Convert to protobuf User
	pbUser := &pb_user.User{
		Id:          user.ID.Hex(),
		FirebaseUid: user.FirebaseUID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		UserType:    pb_user.UserType(pb_user.UserType_value[string(user.UserType)]),
		IsActive:    user.IsActive,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
	}

	return &pb_auth.GetProfileResponse{
		User: pbUser,
	}, nil
}