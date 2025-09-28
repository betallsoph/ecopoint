import { AuthenticationError, UserInputError } from 'apollo-server-express';
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
        throw new AuthenticationError('You must be logged in');
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
        throw new UserInputError('User not found');
      }
      return user;
    }
  },

  Mutation: {
    register: async (_: any, { input }: { input: any }) => {
      const { email, password, firstName, lastName, phone, role = 'USER', address, preferences } = input;

      // Check if user already exists
      const existingUser = await User.findOne({ 
        $or: [{ email }, { phone }] 
      });
      
      if (existingUser) {
        throw new UserInputError('User with this email or phone already exists');
      }

      // Create new user
      const user = new User({
        email,
        password,
        firstName,
        lastName,
        phone,
        role,
        address,
        preferences: {
          language: 'vi',
          notifications: true,
          theme: 'light',
          ...preferences
        }
      });

      await user.save();

      // Generate JWT token
      const token = jwt.sign(
        { userId: user._id, email: user.email, role: user.role },
        JWT_SECRET,
        { expiresIn: '7d' }
      );

      return {
        token,
        user
      };
    },

    login: async (_: any, { input }: { input: any }) => {
      const { email, password } = input;

      // Find user by email
      const user = await User.findOne({ email }).select('+password');
      if (!user) {
        throw new AuthenticationError('Invalid credentials');
      }

      // Check if user is active
      if (!user.isActive) {
        throw new AuthenticationError('Account is deactivated');
      }

      // Verify password
      const isValidPassword = await user.comparePassword(password);
      if (!isValidPassword) {
        throw new AuthenticationError('Invalid credentials');
      }

      // Generate JWT token
      const token = jwt.sign(
        { userId: user._id, email: user.email, role: user.role },
        JWT_SECRET,
        { expiresIn: '7d' }
      );

      return {
        token,
        user
      };
    },

    updateProfile: async (_: any, { input }: { input: any }, { user }: Context) => {
      if (!user) {
        throw new AuthenticationError('You must be logged in');
      }

      const updatedUser = await User.findByIdAndUpdate(
        user._id,
        { $set: input },
        { new: true, runValidators: true }
      );

      if (!updatedUser) {
        throw new UserInputError('User not found');
      }

      return updatedUser;
    },

    changePassword: async (_: any, { currentPassword, newPassword }: { currentPassword: string; newPassword: string }, { user }: Context) => {
      if (!user) {
        throw new AuthenticationError('You must be logged in');
      }

      const userWithPassword = await User.findById(user._id).select('+password');
      if (!userWithPassword) {
        throw new UserInputError('User not found');
      }

      // Verify current password
      const isValidPassword = await userWithPassword.comparePassword(currentPassword);
      if (!isValidPassword) {
        throw new UserInputError('Current password is incorrect');
      }

      // Update password
      userWithPassword.password = newPassword;
      await userWithPassword.save();

      return true;
    },

    deactivateUser: async (_: any, { id }: { id: string }, { user }: Context) => {
      if (!user || user.role !== 'ADMIN') {
        throw new AuthenticationError('Admin access required');
      }

      const targetUser = await User.findByIdAndUpdate(
        id,
        { isActive: false },
        { new: true }
      );

      if (!targetUser) {
        throw new UserInputError('User not found');
      }

      return true;
    },

    activateUser: async (_: any, { id }: { id: string }, { user }: Context) => {
      if (!user || user.role !== 'ADMIN') {
        throw new AuthenticationError('Admin access required');
      }

      const targetUser = await User.findByIdAndUpdate(
        id,
        { isActive: true },
        { new: true }
      );

      if (!targetUser) {
        throw new UserInputError('User not found');
      }

      return true;
    }
  }
};