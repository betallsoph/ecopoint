package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ecopoint/api-gateway/internal/middleware"
)

// GraphQLProxy handles GraphQL requests by proxying to user service
func GraphQLProxy(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// TODO: Implement GraphQL proxy
	// 1. Forward request to user service
	// 2. Add user context to GraphQL request
	// 3. Return response from user service
	
	c.JSON(http.StatusOK, gin.H{
		"message": "GraphQLProxy endpoint - TODO: Implement",
		"user_id": user.UID,
	})
}