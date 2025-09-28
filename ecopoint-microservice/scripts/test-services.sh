#!/bin/bash

# EcoPoint Services Test Script
echo "🧪 Testing EcoPoint Services..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✅ $2${NC}"
    else
        echo -e "${RED}❌ $2${NC}"
    fi
}

# Test User Service
echo "🔍 Testing User Service..."
cd services/user-service

# Check if node_modules exists
if [ -d "node_modules" ]; then
    print_status 0 "Node modules installed"
else
    print_status 1 "Node modules not found"
fi

# Test TypeScript compilation
if npm run build > /dev/null 2>&1; then
    print_status 0 "User Service TypeScript compilation"
else
    print_status 1 "User Service TypeScript compilation failed"
fi

# Test Booking Service
echo "🔍 Testing Booking Service..."
cd ../booking-service

# Check if go.mod exists
if [ -f "go.mod" ]; then
    print_status 0 "Go module found"
else
    print_status 1 "Go module not found"
fi

# Test Go compilation
if go build -o booking-service ./cmd > /dev/null 2>&1; then
    print_status 0 "Booking Service Go compilation"
    rm -f booking-service
else
    print_status 1 "Booking Service Go compilation failed"
fi

# Test API Gateway
echo "🔍 Testing API Gateway..."
cd ../api-gateway

# Test Go compilation
if go build -o api-gateway ./cmd > /dev/null 2>&1; then
    print_status 0 "API Gateway Go compilation"
    rm -f api-gateway
else
    print_status 1 "API Gateway Go compilation failed"
fi

# Test Web App
echo "🔍 Testing Web App..."
cd ../../frontend/web-app

# Check if node_modules exists
if [ -d "node_modules" ]; then
    print_status 0 "Node modules installed"
else
    print_status 1 "Node modules not found"
fi

# Test Next.js build
if npm run build > /dev/null 2>&1; then
    print_status 0 "Web App Next.js build"
else
    print_status 1 "Web App Next.js build failed"
fi

# Test Flutter App
echo "🔍 Testing Flutter App..."
cd ../../mobile/flutter-app

# Check if pubspec.yaml exists
if [ -f "pubspec.yaml" ]; then
    print_status 0 "Flutter project found"
else
    print_status 1 "Flutter project not found"
fi

# Check if Flutter is available
if command -v flutter &> /dev/null; then
    print_status 0 "Flutter CLI available"
else
    print_status 1 "Flutter CLI not available"
fi

# Test Database Schema
echo "🔍 Testing Database Schema..."
cd ../../docs

if [ -f "database-schema.sql" ]; then
    print_status 0 "Database schema file found"
    
    # Check if schema has required tables
    if grep -q "CREATE TABLE users" database-schema.sql; then
        print_status 0 "Users table definition found"
    else
        print_status 1 "Users table definition not found"
    fi
    
    if grep -q "CREATE TABLE bookings" database-schema.sql; then
        print_status 0 "Bookings table definition found"
    else
        print_status 1 "Bookings table definition not found"
    fi
else
    print_status 1 "Database schema file not found"
fi

# Test Docker Files
echo "🔍 Testing Docker Files..."
cd ..

# Check if docker-compose.yml exists
if [ -f "docker-compose.yml" ]; then
    print_status 0 "Docker Compose file found"
else
    print_status 1 "Docker Compose file not found"
fi

# Check if all Dockerfiles exist
services=("user-service" "booking-service" "api-gateway")
for service in "${services[@]}"; do
    if [ -f "services/$service/Dockerfile" ]; then
        print_status 0 "Dockerfile for $service found"
    else
        print_status 1 "Dockerfile for $service not found"
    fi
done

if [ -f "frontend/web-app/Dockerfile" ]; then
    print_status 0 "Dockerfile for web-app found"
else
    print_status 1 "Dockerfile for web-app not found"
fi

echo ""
echo "🎯 Test Summary:"
echo "   - All services should compile successfully"
echo "   - Database schema should be complete"
echo "   - Docker files should be present"
echo ""
echo "📋 Next Steps:"
echo "   1. Configure environment variables"
echo "   2. Set up Firebase project"
echo "   3. Run 'docker-compose up -d' to start services"
echo "   4. Test API endpoints"
echo ""
echo "Happy coding! 🚀"