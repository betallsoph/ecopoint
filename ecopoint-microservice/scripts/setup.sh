#!/bin/bash

# EcoPoint Microservice Setup Script
echo "🚀 Setting up EcoPoint Microservice System..."

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

# Create environment files if they don't exist
echo "📝 Creating environment files..."

if [ ! -f services/user-service/.env ]; then
    cp services/user-service/.env.example services/user-service/.env
    echo "✅ Created services/user-service/.env"
fi

if [ ! -f services/api-gateway/.env ]; then
    cp services/api-gateway/.env.example services/api-gateway/.env
    echo "✅ Created services/api-gateway/.env"
fi

if [ ! -f frontend/web-app/.env.local ]; then
    cp frontend/web-app/.env.local.example frontend/web-app/.env.local
    echo "✅ Created frontend/web-app/.env.local"
fi

# Build and start services
echo "🔨 Building and starting services..."

# Start PostgreSQL and Redis first
echo "Starting databases..."
docker-compose up -d postgres redis mongo

# Wait for databases to be ready
echo "⏳ Waiting for databases to be ready..."
sleep 10

# Start all services
echo "Starting all services..."
docker-compose up -d

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 15

# Check service health
echo "🏥 Checking service health..."

# Check API Gateway
if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ API Gateway is healthy"
else
    echo "❌ API Gateway is not responding"
fi

# Check User Service
if curl -f http://localhost:4000/health > /dev/null 2>&1; then
    echo "✅ User Service is healthy"
else
    echo "❌ User Service is not responding"
fi

# Check Web App
if curl -f http://localhost:3000 > /dev/null 2>&1; then
    echo "✅ Web App is healthy"
else
    echo "❌ Web App is not responding"
fi

echo ""
echo "🎉 Setup complete!"
echo ""
echo "📱 Access your applications:"
echo "   Web App: http://localhost:3000"
echo "   API Gateway: http://localhost:8080"
echo "   GraphQL Playground: http://localhost:4000/graphql"
echo ""
echo "🗄️ Database connections:"
echo "   PostgreSQL: localhost:5432"
echo "   MongoDB: localhost:27017"
echo "   Redis: localhost:6379"
echo ""
echo "📋 Next steps:"
echo "   1. Configure Firebase in your environment files"
echo "   2. Set up Google Maps API key"
echo "   3. Run 'docker-compose logs -f' to view logs"
echo "   4. Run 'docker-compose down' to stop services"
echo ""
echo "Happy coding! 🚀"