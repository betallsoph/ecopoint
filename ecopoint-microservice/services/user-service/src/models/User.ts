import mongoose, { Document, Schema } from 'mongoose';
import bcrypt from 'bcryptjs';

export interface IUser extends Document {
  _id: string;
  firebaseUid: string;
  email: string;
  firstName: string;
  lastName: string;
  phone?: string;
  role: 'USER' | 'COLLECTOR' | 'ADMIN';
  isActive: boolean;
  profileImageUrl?: string;
  createdAt: Date;
  updatedAt: Date;
}

const UserSchema = new Schema<IUser>({
  firebaseUid: {
    type: String,
    required: true,
    unique: true,
    index: true
  },
  email: {
    type: String,
    required: true,
    unique: true,
    lowercase: true,
    trim: true
  },
  firstName: {
    type: String,
    required: true,
    trim: true
  },
  lastName: {
    type: String,
    required: true,
    trim: true
  },
  phone: {
    type: String,
    unique: true,
    sparse: true
  },
  role: {
    type: String,
    enum: ['USER', 'COLLECTOR', 'ADMIN'],
    default: 'USER'
  },
  isActive: {
    type: Boolean,
    default: true
  },
  profileImageUrl: {
    type: String
  }
}, {
  timestamps: true
});

// Remove sensitive fields from JSON output
UserSchema.methods.toJSON = function() {
  const userObject = this.toObject();
  return userObject;
};

export default mongoose.model<IUser>('User', UserSchema);