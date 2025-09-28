import { GraphQLError } from 'graphql';
import jwt from 'jsonwebtoken';
import User, { IUser } from '../../models/User';

const JWT_SECRET = process.env.JWT_SECRET || 'fallback-secret';

interface Context {
  user?: IUser;
}

export const userResolvers = {
  Query: {
    me: async (_: any, __: any, { user }: Context) => {
      if (!user) {
        throw new GraphQLError('You must be logged in', {
          extensions: { code: 'UNAUTHENTICATED' }
        });
      }
      return user;
    },

    users: async (_: any, { role, isActive }: { role?: string; isActive?: boolean }) => {
      const filter: any = {};
      if (role) filter.role = role;
      if (isActive !== undefined) filter.isActive = isActive;
      
      return User.find(filter).sort({ createdAt: -1 });
    },

    user: async (_: any, { id }: { id: string }) => {
      const user = await User.findById(id);
      if (!user) {
        throw new GraphQLError('User not found', {
          extensions: { code: 'NOT_FOUND' }
        });
      }
      return user;
    }
  },

  Mutation: {
    createUser: async (_: any, { input }: { input: any }) => {
      const { firebaseUid, email, firstName, lastName, phone, role = 'USER', profileImageUrl } = input;

      // Check if user already exists
      const existingUser = await User.findOne({ 
        $or: [{ email }, { firebaseUid }] 
      });
      
      if (existingUser) {
        throw new GraphQLError('User with this email or Firebase UID already exists', {
          extensions: { code: 'BAD_USER_INPUT' }
        });
      }

      // Create new user
      const user = new User({
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

    authenticateUser: async (_: any, { input }: { input: any }) => {
      const { firebaseToken } = input;

      // TODO: Verify Firebase token
      // For now, we'll create a mock user
      const user = await User.findOne({ firebaseUid: 'mock-uid' });
      if (!user) {
        throw new GraphQLError('User not found', {
          extensions: { code: 'NOT_FOUND' }
        });
      }

      return {
        user
      };
    },

    updateProfile: async (_: any, { input }: { input: any }, { user }: Context) => {
      if (!user) {
        throw new GraphQLError('You must be logged in', {
          extensions: { code: 'UNAUTHENTICATED' }
        });
      }

      const updatedUser = await User.findByIdAndUpdate(
        user._id,
        { $set: input },
        { new: true, runValidators: true }
      );

      if (!updatedUser) {
        throw new GraphQLError('User not found', {
          extensions: { code: 'NOT_FOUND' }
        });
      }

      return updatedUser;
    },

    deactivateUser: async (_: any, { id }: { id: string }, { user }: Context) => {
      if (!user || user.role !== 'ADMIN') {
        throw new GraphQLError('Admin access required', {
          extensions: { code: 'FORBIDDEN' }
        });
      }

      const targetUser = await User.findByIdAndUpdate(
        id,
        { isActive: false },
        { new: true }
      );

      if (!targetUser) {
        throw new GraphQLError('User not found', {
          extensions: { code: 'NOT_FOUND' }
        });
      }

      return true;
    },

    activateUser: async (_: any, { id }: { id: string }, { user }: Context) => {
      if (!user || user.role !== 'ADMIN') {
        throw new GraphQLError('Admin access required', {
          extensions: { code: 'FORBIDDEN' }
        });
      }

      const targetUser = await User.findByIdAndUpdate(
        id,
        { isActive: true },
        { new: true }
      );

      if (!targetUser) {
        throw new GraphQLError('User not found', {
          extensions: { code: 'NOT_FOUND' }
        });
      }

      return true;
    }
  }
};