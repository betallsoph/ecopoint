import { AuthenticationError } from 'apollo-server-express';
import jwt from 'jsonwebtoken';
import User, { IUser } from '../models/User';

const JWT_SECRET = process.env.JWT_SECRET || 'fallback-secret';

export interface AuthContext {
  user?: IUser;
}

export const authMiddleware = async (req: any): Promise<AuthContext> => {
  try {
    const token = req.headers.authorization?.replace('Bearer ', '');
    
    if (!token) {
      return {};
    }

    const decoded = jwt.verify(token, JWT_SECRET) as any;
    const user = await User.findById(decoded.userId);
    
    if (!user || !user.isActive) {
      throw new AuthenticationError('Invalid or expired token');
    }

    return { user };
  } catch (error) {
    if (error instanceof jwt.JsonWebTokenError) {
      throw new AuthenticationError('Invalid token');
    }
    throw error;
  }
};

export const requireAuth = (user?: IUser) => {
  if (!user) {
    throw new AuthenticationError('Authentication required');
  }
  return user;
};

export const requireRole = (user: IUser, allowedRoles: string[]) => {
  if (!allowedRoles.includes(user.role)) {
    throw new AuthenticationError('Insufficient permissions');
  }
  return user;
};