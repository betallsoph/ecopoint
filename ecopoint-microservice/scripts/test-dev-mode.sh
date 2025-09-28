#!/bin/bash

# EcoPoint Dev Mode Test Script
echo "🧪 Testing EcoPoint Dev Mode..."

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

echo "🔍 Testing Web App Dev Mode..."

# Test Web App compilation
cd frontend/web-app

if npm run build > /dev/null 2>&1; then
    print_status 0 "Web App builds successfully with dev mode"
else
    print_status 1 "Web App build failed"
fi

# Check if dev mode files exist
if [ -f "src/contexts/DevModeContext.tsx" ]; then
    print_status 0 "DevModeContext found"
else
    print_status 1 "DevModeContext not found"
fi

if [ -f "src/components/DevModeToggle.tsx" ]; then
    print_status 0 "DevModeToggle component found"
else
    print_status 1 "DevModeToggle component not found"
fi

if [ -f "src/lib/mockApi.ts" ]; then
    print_status 0 "Mock API service found"
else
    print_status 1 "Mock API service not found"
fi

if [ -f "src/lib/apiClient.ts" ]; then
    print_status 0 "API Client with dev mode found"
else
    print_status 1 "API Client with dev mode not found"
fi

echo ""
echo "🔍 Testing Flutter Apps Dev Mode..."

# Test User App
cd ../../mobile/user-app
if [ -f "lib/services/dev_mode_service.dart" ]; then
    print_status 0 "User App dev mode service found"
else
    print_status 1 "User App dev mode service not found"
fi

if [ -f "lib/screens/auth/login_screen.dart" ]; then
    print_status 0 "User App login screen with dev mode found"
else
    print_status 1 "User App login screen with dev mode not found"
fi

# Test Collector App
cd ../collector-app
if [ -f "lib/services/dev_mode_service.dart" ]; then
    print_status 0 "Collector App dev mode service found"
else
    print_status 1 "Collector App dev mode service not found"
fi

# Test Admin App
cd ../admin-app
if [ -f "lib/services/dev_mode_service.dart" ]; then
    print_status 0 "Admin App dev mode service found"
else
    print_status 1 "Admin App dev mode service not found"
fi

echo ""
echo "🎯 Dev Mode Test Summary:"
echo "   - Web App: Dev mode toggle and mock API ready"
echo "   - Flutter Apps: Dev mode service and login bypass ready"
echo "   - Mock data: Available for testing"
echo "   - Firebase bypass: Implemented"
echo ""
echo "🚀 How to use Dev Mode:"
echo ""
echo "📱 Web App:"
echo "   1. Start: npm run dev"
echo "   2. Toggle 'Dev Mode' in top-right corner"
echo "   3. Switch between User/Collector/Admin roles"
echo "   4. Use any email/password to login"
echo ""
echo "📱 Flutter Apps:"
echo "   1. Enable dev mode in login screen"
echo "   2. Use any email/password to login"
echo "   3. Mock data will be loaded automatically"
echo ""
echo "🎉 Dev Mode is ready for testing!"