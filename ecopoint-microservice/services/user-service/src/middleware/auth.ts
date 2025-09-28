import { GraphQLError } from 'graphql';
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
      throw new GraphQLError('Invalid or expired token', {
        extensions: { code: 'UNAUTHENTICATED' }
      });
    }

    return { user };
  } catch (error) {
    if (error instanceof jwt.JsonWebTokenError) {
      throw new GraphQLError('Invalid token', {
        extensions: { code: 'UNAUTHENTICATED' }
      });
    }
    throw error;
  }
};

export const requireAuth = (user?: IUser) => {
  if (!user) {
    throw new GraphQLError('Authentication required', {
      extensions: { code: 'UNAUTHENTICATED' }
    });
  }
  return user;
};

export const requireRole = (user: IUser, allowedRoles: string[]) => {
  if (!allowedRoles.includes(user.role)) {
    throw new GraphQLError('Insufficient permissions', {
      extensions: { code: 'FORBIDDEN' }
    });
  }
  return user;
};