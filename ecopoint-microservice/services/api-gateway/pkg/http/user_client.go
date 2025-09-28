package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type UserClient struct {
	baseURL    string
	httpClient *http.Client
}

type User struct {
	ID            string `json:"id"`
	FirebaseUID   string `json:"firebase_uid"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	Role          string `json:"role"`
	IsActive      bool   `json:"is_active"`
	ProfileImageURL string `json:"profile_image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Address struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Street    string  `json:"street"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsDefault bool    `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	FirebaseUID    string `json:"firebase_uid"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Role           string `json:"role"`
	ProfileImageURL string `json:"profile_image_url"`
}

type UpdateUserRequest struct {
	FirstName      *string `json:"first_name,omitempty"`
	LastName       *string `json:"last_name,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	ProfileImageURL *string `json:"profile_image_url,omitempty"`
}

type CreateAddressRequest struct {
	Street    string  `json:"street"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsDefault bool    `json:"is_default"`
}

type UpdateAddressRequest struct {
	Street    *string  `json:"street,omitempty"`
	City      *string  `json:"city,omitempty"`
	State     *string  `json:"state,omitempty"`
	ZipCode   *string  `json:"zip_code,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	IsDefault *bool    `json:"is_default,omitempty"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	User    *User       `json:"user,omitempty"`
	Address *Address    `json:"address,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *UserClient) makeRequest(ctx context.Context, method, endpoint string, body interface{}, headers map[string]string) (*APIResponse, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s (status: %d)", apiResp.Message, resp.StatusCode)
	}

	return &apiResp, nil
}

// User methods
func (c *UserClient) CreateUser(ctx context.Context, req *CreateUserRequest, token string) (*User, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.makeRequest(ctx, "POST", "/users", req, headers)
	if err != nil {
		return nil, err
	}

	if resp.User == nil {
		return nil, fmt.Errorf("user not found in response")
	}

	return resp.User, nil
}

func (c *UserClient) GetCurrentUser(ctx context.Context, token string) (*User, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.makeRequest(ctx, "GET", "/users/me", nil, headers)
	if err != nil {
		return nil, err
	}

	if resp.User == nil {
		return nil, fmt.Errorf("user not found in response")
	}

	return resp.User, nil
}

func (c *UserClient) UpdateUser(ctx context.Context, req *UpdateUserRequest, token string) (*User, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.makeRequest(ctx, "PUT", "/users/me", req, headers)
	if err != nil {
		return nil, err
	}

	if resp.User == nil {
		return nil, fmt.Errorf("user not found in response")
	}

	return resp.User, nil
}

func (c *UserClient) GetUser(ctx context.Context, userID, token string) (*User, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.makeRequest(ctx, "GET", "/users/"+userID, nil, headers)
	if err != nil {
		return nil, err
	}

	if resp.User == nil {
		return nil, fmt.Errorf("user not found in response")
	}

	return resp.User, nil
}

// Address methods
func (c *UserClient) GetAddresses(ctx context.Context, token string) ([]Address, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.makeRequest(ctx, "GET", "/addresses", nil, headers)
	if err != nil {
		return nil, err
	}

	return resp.Addresses, nil
}

func (c *UserClient) CreateAddress(ctx context.Context, req *CreateAddressRequest, token string) (*Address, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.makeRequest(ctx, "POST", "/addresses", req, headers)
	if err != nil {
		return nil, err
	}

	if resp.Address == nil {
		return nil, fmt.Errorf("address not found in response")
	}

	return resp.Address, nil
}

func (c *UserClient) UpdateAddress(ctx context.Context, addressID string, req *UpdateAddressRequest, token string) (*Address, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.makeRequest(ctx, "PUT", "/addresses/"+addressID, req, headers)
	if err != nil {
		return nil, err
	}

	if resp.Address == nil {
		return nil, fmt.Errorf("address not found in response")
	}

	return resp.Address, nil
}

func (c *UserClient) DeleteAddress(ctx context.Context, addressID, token string) error {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	_, err := c.makeRequest(ctx, "DELETE", "/addresses/"+addressID, nil, headers)
	return err
}