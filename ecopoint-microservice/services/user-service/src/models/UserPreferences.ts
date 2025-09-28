import mongoose, { Document, Schema } from 'mongoose';

export interface IUserPreferences extends Document {
  _id: string;
  userId: string;
  language: string;
  notificationsEnabled: boolean;
  theme: 'light' | 'dark';
  createdAt: Date;
  updatedAt: Date;
}

const UserPreferencesSchema = new Schema<IUserPreferences>({
  userId: {
    type: String,
    required: true,
    unique: true,
    ref: 'User'
  },
  language: {
    type: String,
    default: 'vi',
    trim: true
  },
  notificationsEnabled: {
    type: Boolean,
    default: true
  },
  theme: {
    type: String,
    enum: ['light', 'dark'],
    default: 'light'
  }
}, {
  timestamps: true
});

export default mongoose.model<IUserPreferences>('UserPreferences', UserPreferencesSchema);