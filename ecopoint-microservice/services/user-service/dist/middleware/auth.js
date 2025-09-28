"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.requireRole = exports.requireAuth = exports.authMiddleware = void 0;
const graphql_1 = require("graphql");
const jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
const User_1 = __importDefault(require("../models/User"));
const JWT_SECRET = process.env.JWT_SECRET || 'fallback-secret';
const authMiddleware = async (req) => {
    try {
        const token = req.headers.authorization?.replace('Bearer ', '');
        if (!token) {
            return {};
        }
        const decoded = jsonwebtoken_1.default.verify(token, JWT_SECRET);
        const user = await User_1.default.findById(decoded.userId);
        if (!user || !user.isActive) {
            throw new graphql_1.GraphQLError('Invalid or expired token', {
                extensions: { code: 'UNAUTHENTICATED' }
            });
        }
        return { user };
    }
    catch (error) {
        if (error instanceof jsonwebtoken_1.default.JsonWebTokenError) {
            throw new graphql_1.GraphQLError('Invalid token', {
                extensions: { code: 'UNAUTHENTICATED' }
            });
        }
        throw error;
    }
};
exports.authMiddleware = authMiddleware;
const requireAuth = (user) => {
    if (!user) {
        throw new graphql_1.GraphQLError('Authentication required', {
            extensions: { code: 'UNAUTHENTICATED' }
        });
    }
    return user;
};
exports.requireAuth = requireAuth;
const requireRole = (user, allowedRoles) => {
    if (!allowedRoles.includes(user.role)) {
        throw new graphql_1.GraphQLError('Insufficient permissions', {
            extensions: { code: 'FORBIDDEN' }
        });
    }
    return user;
};
exports.requireRole = requireRole;
//# sourceMappingURL=auth.js.map