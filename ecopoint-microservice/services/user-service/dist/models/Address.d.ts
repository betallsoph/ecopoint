import mongoose, { Document } from 'mongoose';
export interface IAddress extends Document {
    _id: string;
    userId: string;
    street: string;
    city: string;
    state: string;
    zipCode: string;
    latitude: number;
    longitude: number;
    isDefault: boolean;
    createdAt: Date;
    updatedAt: Date;
}
declare const _default: mongoose.Model<IAddress, {}, {}, {}, mongoose.Document<unknown, {}, IAddress, {}, {}> & IAddress & Required<{
    _id: string;
}> & {
    __v: number;
}, any>;
export default _default;
//# sourceMappingURL=Address.d.ts.map