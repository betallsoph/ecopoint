import mongoose, { Document } from 'mongoose';
export interface IUserPreferences extends Document {
    _id: string;
    userId: string;
    language: string;
    notificationsEnabled: boolean;
    theme: 'light' | 'dark';
    createdAt: Date;
    updatedAt: Date;
}
declare const _default: mongoose.Model<IUserPreferences, {}, {}, {}, mongoose.Document<unknown, {}, IUserPreferences, {}, {}> & IUserPreferences & Required<{
    _id: string;
}> & {
    __v: number;
}, any>;
export default _default;
//# sourceMappingURL=UserPreferences.d.ts.map