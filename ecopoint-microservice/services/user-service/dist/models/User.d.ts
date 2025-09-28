import mongoose, { Document } from 'mongoose';
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
declare const _default: mongoose.Model<IUser, {}, {}, {}, mongoose.Document<unknown, {}, IUser, {}, {}> & IUser & Required<{
    _id: string;
}> & {
    __v: number;
}, any>;
export default _default;
//# sourceMappingURL=User.d.ts.map