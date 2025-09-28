import { IUser } from '../../models/User';
interface Context {
    user?: IUser;
}
export declare const userResolvers: {
    Query: {
        me: (_: any, __: any, { user }: Context) => Promise<IUser>;
        users: (_: any, { role, isActive }: {
            role?: string;
            isActive?: boolean;
        }) => Promise<(import("mongoose").Document<unknown, {}, IUser, {}, {}> & IUser & Required<{
            _id: string;
        }> & {
            __v: number;
        })[]>;
        user: (_: any, { id }: {
            id: string;
        }) => Promise<import("mongoose").Document<unknown, {}, IUser, {}, {}> & IUser & Required<{
            _id: string;
        }> & {
            __v: number;
        }>;
    };
    Mutation: {
        createUser: (_: any, { input }: {
            input: any;
        }) => Promise<{
            user: import("mongoose").Document<unknown, {}, IUser, {}, {}> & IUser & Required<{
                _id: string;
            }> & {
                __v: number;
            };
        }>;
        authenticateUser: (_: any, { input }: {
            input: any;
        }) => Promise<{
            user: import("mongoose").Document<unknown, {}, IUser, {}, {}> & IUser & Required<{
                _id: string;
            }> & {
                __v: number;
            };
        }>;
        updateProfile: (_: any, { input }: {
            input: any;
        }, { user }: Context) => Promise<import("mongoose").Document<unknown, {}, IUser, {}, {}> & IUser & Required<{
            _id: string;
        }> & {
            __v: number;
        }>;
        deactivateUser: (_: any, { id }: {
            id: string;
        }, { user }: Context) => Promise<boolean>;
        activateUser: (_: any, { id }: {
            id: string;
        }, { user }: Context) => Promise<boolean>;
    };
};
export {};
//# sourceMappingURL=userResolver.d.ts.map