package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"ecopoint-backend/shared/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Service struct {
	Name string
	URL  string
}

var services = map[string]Service{
	"auth":  {Name: "auth-service", URL: "http://localhost:8081"},
	"user":  {Name: "user-service", URL: "http://localhost:8082"},
	"order": {Name: "order-service", URL: "http://localhost:8083"},
	"collector": {Name: "collector-service", URL: "http://localhost:8084"},
}

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Health check for gateway
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "api-gateway",
			"services": map[string]string{
				"auth":      services["auth"].URL + "/health",
				"order":     services["order"].URL + "/health", 
				"user":      services["user"].URL + "/health",
				"collector": services["collector"].URL + "/health",
			},
		})
	})

	// Route all API requests
	api := r.Group("/api/v1")
	{
		// Auth service routes
		auth := api.Group("/auth")
		auth.Any("/*path", func(c *gin.Context) {
			proxyToService(c, "auth")
		})

		// User service routes  
		user := api.Group("/users")
		user.Any("/*path", func(c *gin.Context) {
			proxyToService(c, "user")
		})

		// Order service routes
		order := api.Group("/orders")
		order.Any("/*path", func(c *gin.Context) {
			proxyToService(c, "order")
		})

		// Collector service routes
		collector := api.Group("/collectors")
		collector.Any("/*path", func(c *gin.Context) {
			proxyToService(c, "collector")
		})
	}

	// Catch-all route for unmatched paths
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
			"path":  c.Request.URL.Path,
		})
	})

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("API Gateway starting on port %s", port)
	log.Printf("Routing to services: %+v", services)
	log.Fatal(r.Run(":" + port))
}

func proxyToService(c *gin.Context, serviceName string) {
	service, exists := services[serviceName]
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service not found"})
		return
	}

	// Parse the target URL
	target, err := url.Parse(service.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
		return
	}

	// Create reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Modify the request
	originalPath := c.Request.URL.Path
	
	// Remove the service prefix from path
	// e.g., /api/v1/auth/login -> /api/v1/auth/login
	c.Request.URL.Path = originalPath
	c.Request.URL.Host = target.Host
	c.Request.URL.Scheme = target.Scheme
	c.Request.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))
	c.Request.Host = target.Host

	// Custom error handler
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error for service %s: %v", serviceName, err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"error": "Service unavailable"}`))
	}

	// Custom response modifier
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Add service identification header
		resp.Header.Set("X-Service", service.Name)
		return nil
	}

	// Serve the proxy
	proxy.ServeHTTP(c.Writer, c.Request)
}

// healthCheck checks if a service is healthy
func healthCheck(serviceURL string) bool {
	resp, err := http.Get(serviceURL + "/health")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}