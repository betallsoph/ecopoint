import express from 'express';
import { ApolloServer } from '@apollo/server';
import { expressMiddleware } from '@apollo/server/express4';
import cors from 'cors';
import helmet from 'helmet';
import morgan from 'morgan';
import dotenv from 'dotenv';

import connectDB from './config/database';
import { userTypeDefs } from './graphql/schemas/userSchema';
import { userResolvers } from './graphql/resolvers/userResolver';
import { authMiddleware } from './middleware/auth';

// Load environment variables
dotenv.config();

const app = express();
const PORT = process.env.PORT || 4000;

async function startServer() {
  // Connect to database
  await connectDB();

  // Create Apollo Server
  const server = new ApolloServer({
    typeDefs: [userTypeDefs],
    resolvers: [userResolvers],
    context: async ({ req }) => {
      return await authMiddleware(req);
    },
    introspection: process.env.NODE_ENV !== 'production',
    plugins: [
      // Add any plugins here
    ]
  });

  // Start the server
  await server.start();

  // Apply middleware
  app.use(helmet());
  app.use(morgan('combined'));
  app.use(cors({
    origin: process.env.CORS_ORIGIN || '*',
    credentials: true
  }));
  app.use(express.json({ limit: '10mb' }));
  app.use(express.urlencoded({ extended: true }));

  // Health check endpoint
  app.get('/health', (req, res) => {
    res.status(200).json({
      status: 'OK',
      service: 'user-service',
      timestamp: new Date().toISOString()
    });
  });

  // Apply GraphQL middleware
  app.use('/graphql', expressMiddleware(server, {
    context: async ({ req }) => {
      return await authMiddleware(req);
    }
  }));

  // Start the Express server
  app.listen(PORT, () => {
    console.log(`🚀 User Service running on port ${PORT}`);
    console.log(`📊 GraphQL endpoint: http://localhost:${PORT}/graphql`);
    console.log(`🏥 Health check: http://localhost:${PORT}/health`);
  });
}

// Handle unhandled promise rejections
process.on('unhandledRejection', (err: Error) => {
  console.error('Unhandled Promise Rejection:', err);
  process.exit(1);
});

// Handle uncaught exceptions
process.on('uncaughtException', (err: Error) => {
  console.error('Uncaught Exception:', err);
  process.exit(1);
});

// Graceful shutdown
process.on('SIGTERM', () => {
  console.log('SIGTERM received. Shutting down gracefully...');
  process.exit(0);
});

process.on('SIGINT', () => {
  console.log('SIGINT received. Shutting down gracefully...');
  process.exit(0);
});

startServer().catch((error) => {
  console.error('Failed to start server:', error);
  process.exit(1);
});