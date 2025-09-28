import { IUser } from '../models/User';
export interface AuthContext {
    user?: IUser;
}
export declare const authMiddleware: (req: any) => Promise<AuthContext>;
export declare const requireAuth: (user?: IUser) => IUser;
export declare const requireRole: (user: IUser, allowedRoles: string[]) => IUser;
//# sourceMappingURL=auth.d.ts.map