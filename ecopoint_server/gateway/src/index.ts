import 'dotenv/config';
import { ApolloServer } from '@apollo/server';
import { startStandaloneServer } from '@apollo/server/standalone';
import { loadPackageDefinition, credentials } from '@grpc/grpc-js';
import { loadSync } from '@grpc/proto-loader';
import path from 'node:path';

const ACCOUNT_ADDR = process.env.ACCOUNT_ADDR || 'localhost:50051';
const COLLECT_ADDR = process.env.COLLECT_ADDR || 'localhost:50052';

const accountDef = loadSync(path.resolve(process.cwd(), '../proto/account.proto'), { keepCase: true, longs: String, enums: String, defaults: true, oneofs: true });
const accountPkg = loadPackageDefinition(accountDef) as any;
const accountSvc = new accountPkg.ecopoint.account.v1.AccountService(ACCOUNT_ADDR, credentials.createInsecure());

const collectDef = loadSync(path.resolve(process.cwd(), '../proto/collecting.proto'), { keepCase: true, longs: String, enums: String, defaults: true, oneofs: true });
const collectPkg = loadPackageDefinition(collectDef) as any;
const collectSvc = new collectPkg.ecopoint.collecting.v1.CollectingService(COLLECT_ADDR, credentials.createInsecure());

const typeDefs = `#graphql
  type User { user_id: ID!, email: String, phone: String, display_name: String, avatar_url: String }
  type Address { id: ID!, full_text: String!, lat: Float, lng: Float, is_default: Boolean }
  type Order { id: ID!, status: String!, customer_id: ID!, accepted_by: String, note: String }

  type Query {
    me(user_id: ID!): User
    myAddresses(user_id: ID!): [Address!]!
    availableOrders(limit: Int): [Order!]!
  }

  input CreateOrderInput {
    customer_id: ID!
    full_text: String!
    lat: Float
    lng: Float
    display_name: String
    phone: String
    note: String
  }

  type Mutation {
    createOrder(input: CreateOrderInput!): Order!
    acceptOrder(order_id: ID!, collector_id: ID!): Order!
    updateOrderStatus(order_id: ID!, status: String!, collector_id: ID!): Order!
  }
`;

const resolvers = {
  Query: {
    me: async (_: any, { user_id }: any) => {
      return await new Promise((resolve, reject) => {
        accountSvc.GetUser({ user_id }, (err: any, res: any) => err ? reject(err) : resolve(res));
      });
    },
    myAddresses: async (_: any, { user_id }: any) => {
      const res = await new Promise<any>((resolve, reject) => {
        accountSvc.ListAddresses({ user_id }, (err: any, r: any) => err ? reject(err) : resolve(r));
      });
      return res.addresses || [];
    },
    availableOrders: async (_: any, { limit }: any) => {
      const res = await new Promise<any>((resolve, reject) => {
        collectSvc.ListAvailableOrders({ limit: limit ?? 10 }, (err: any, r: any) => err ? reject(err) : resolve(r));
      });
      return res.orders || [];
    },
  },
  Mutation: {
    createOrder: async (_: any, { input }: any) => {
      const payload = {
        customer_id: input.customer_id,
        pick_address: { full_text: input.full_text, lat: input.lat ?? 0, lng: input.lng ?? 0 },
        customer_snapshot: { display_name: input.display_name ?? '', phone: input.phone ?? '' },
        items: [], total_weight: 0, estimated_price: 0, note: input.note ?? ''
      };
      const res = await new Promise<any>((resolve, reject) => {
        collectSvc.CreateOrder(payload, (err: any, r: any) => err ? reject(err) : resolve(r));
      });
      return res;
    },
    acceptOrder: async (_: any, { order_id, collector_id }: any) => {
      const res = await new Promise<any>((resolve, reject) => {
        collectSvc.AcceptOrder({ order_id, collector_id }, (err: any, r: any) => err ? reject(err) : resolve(r));
      });
      return res;
    },
    updateOrderStatus: async (_: any, { order_id, status, collector_id }: any) => {
      const res = await new Promise<any>((resolve, reject) => {
        collectSvc.UpdateOrderStatus({ order_id, status, collector_id }, (err: any, r: any) => err ? reject(err) : resolve(r));
      });
      return res;
    },
  }
};

const server = new ApolloServer({ typeDefs, resolvers });
const port = Number(process.env.PORT || 4000);
startStandaloneServer(server, { listen: { port } }).then(({ url }) => {
  console.log('GraphQL Gateway running at', url);
});


