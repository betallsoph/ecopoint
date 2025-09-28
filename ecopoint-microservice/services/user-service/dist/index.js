"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const server_1 = require("@apollo/server");
const express4_1 = require("@apollo/server/express4");
const cors_1 = __importDefault(require("cors"));
const helmet_1 = __importDefault(require("helmet"));
const morgan_1 = __importDefault(require("morgan"));
const dotenv_1 = __importDefault(require("dotenv"));
const database_1 = __importDefault(require("./config/database"));
const userSchema_1 = require("./graphql/schemas/userSchema");
const userResolver_1 = require("./graphql/resolvers/userResolver");
const auth_1 = require("./middleware/auth");
// Load environment variables
dotenv_1.default.config();
const app = (0, express_1.default)();
const PORT = process.env.PORT || 4000;
async function startServer() {
    // Connect to database
    await (0, database_1.default)();
    // Create Apollo Server
    const server = new server_1.ApolloServer({
        typeDefs: [userSchema_1.userTypeDefs],
        resolvers: [userResolver_1.userResolvers],
        introspection: process.env.NODE_ENV !== 'production',
        plugins: [
        // Add any plugins here
        ]
    });
    // Start the server
    await server.start();
    // Apply middleware
    app.use((0, helmet_1.default)());
    app.use((0, morgan_1.default)('combined'));
    app.use((0, cors_1.default)({
        origin: process.env.CORS_ORIGIN || '*',
        credentials: true
    }));
    app.use(express_1.default.json({ limit: '10mb' }));
    app.use(express_1.default.urlencoded({ extended: true }));
    // Health check endpoint
    app.get('/health', (req, res) => {
        res.status(200).json({
            status: 'OK',
            service: 'user-service',
            timestamp: new Date().toISOString()
        });
    });
    // Apply GraphQL middleware
    app.use('/graphql', (0, express4_1.expressMiddleware)(server, {
        context: async ({ req }) => {
            return await (0, auth_1.authMiddleware)(req);
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
process.on('unhandledRejection', (err) => {
    console.error('Unhandled Promise Rejection:', err);
    process.exit(1);
});
// Handle uncaught exceptions
process.on('uncaughtException', (err) => {
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
//# sourceMappingURL=index.js.map