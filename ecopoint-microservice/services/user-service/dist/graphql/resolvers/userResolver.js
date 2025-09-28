"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.userResolvers = void 0;
const graphql_1 = require("graphql");
const User_1 = __importDefault(require("../../models/User"));
const JWT_SECRET = process.env.JWT_SECRET || 'fallback-secret';
exports.userResolvers = {
    Query: {
        me: async (_, __, { user }) => {
            if (!user) {
                throw new graphql_1.GraphQLError('You must be logged in', {
                    extensions: { code: 'UNAUTHENTICATED' }
                });
            }
            return user;
        },
        users: async (_, { role, isActive }) => {
            const filter = {};
            if (role)
                filter.role = role;
            if (isActive !== undefined)
                filter.isActive = isActive;
            return User_1.default.find(filter).sort({ createdAt: -1 });
        },
        user: async (_, { id }) => {
            const user = await User_1.default.findById(id);
            if (!user) {
                throw new graphql_1.GraphQLError('User not found', {
                    extensions: { code: 'NOT_FOUND' }
                });
            }
            return user;
        }
    },
    Mutation: {
        createUser: async (_, { input }) => {
            const { firebaseUid, email, firstName, lastName, phone, role = 'USER', profileImageUrl } = input;
            // Check if user already exists
            const existingUser = await User_1.default.findOne({
                $or: [{ email }, { firebaseUid }]
            });
            if (existingUser) {
                throw new graphql_1.GraphQLError('User with this email or Firebase UID already exists', {
                    extensions: { code: 'BAD_USER_INPUT' }
                });
            }
            // Create new user
            const user = new User_1.default({
                firebaseUid,
                email,
                firstName,
                lastName,
                phone,
                role,
                profileImageUrl
            });
            await user.save();
            return {
                user
            };
        },
        authenticateUser: async (_, { input }) => {
            const { firebaseToken } = input;
            // TODO: Verify Firebase token
            // For now, we'll create a mock user
            const user = await User_1.default.findOne({ firebaseUid: 'mock-uid' });
            if (!user) {
                throw new graphql_1.GraphQLError('User not found', {
                    extensions: { code: 'NOT_FOUND' }
                });
            }
            return {
                user
            };
        },
        updateProfile: async (_, { input }, { user }) => {
            if (!user) {
                throw new graphql_1.GraphQLError('You must be logged in', {
                    extensions: { code: 'UNAUTHENTICATED' }
                });
            }
            const updatedUser = await User_1.default.findByIdAndUpdate(user._id, { $set: input }, { new: true, runValidators: true });
            if (!updatedUser) {
                throw new graphql_1.GraphQLError('User not found', {
                    extensions: { code: 'NOT_FOUND' }
                });
            }
            return updatedUser;
        },
        deactivateUser: async (_, { id }, { user }) => {
            if (!user || user.role !== 'ADMIN') {
                throw new graphql_1.GraphQLError('Admin access required', {
                    extensions: { code: 'FORBIDDEN' }
                });
            }
            const targetUser = await User_1.default.findByIdAndUpdate(id, { isActive: false }, { new: true });
            if (!targetUser) {
                throw new graphql_1.GraphQLError('User not found', {
                    extensions: { code: 'NOT_FOUND' }
                });
            }
            return true;
        },
        activateUser: async (_, { id }, { user }) => {
            if (!user || user.role !== 'ADMIN') {
                throw new graphql_1.GraphQLError('Admin access required', {
                    extensions: { code: 'FORBIDDEN' }
                });
            }
            const targetUser = await User_1.default.findByIdAndUpdate(id, { isActive: true }, { new: true });
            if (!targetUser) {
                throw new graphql_1.GraphQLError('User not found', {
                    extensions: { code: 'NOT_FOUND' }
                });
            }
            return true;
        }
    }
};
//# sourceMappingURL=userResolver.js.map