import 'dotenv/config';
import { loadPackageDefinition, Server, ServerCredentials } from '@grpc/grpc-js';
import { loadSync } from '@grpc/proto-loader';
import path from 'node:path';
import { getPool } from '../db.ts';
import { UserRepo } from '../repositories/userRepo.ts';
import { RoleRepo } from '../repositories/roleRepo.ts';
import { AddressRepo } from '../repositories/addressRepo.ts';
import { DeviceRepo } from '../repositories/deviceRepo.ts';

const PROTO_PATH = path.resolve(process.cwd(), '../proto/account.proto');
const pkgDef = loadSync(PROTO_PATH, { keepCase: true, longs: String, enums: String, defaults: true, oneofs: true });
const grpcObj = loadPackageDefinition(pkgDef) as any;
const svc = grpcObj.ecopoint.account.v1;

async function main() {
  const server = new Server();
  const pool = await getPool();
  const users = new UserRepo(pool);
  const roles = new RoleRepo(pool);
  const addrs = new AddressRepo(pool);
  const devs  = new DeviceRepo(pool);

  const impl = {
    async GetUser(call: any, cb: any) {
      const u = await users.getUser(call.request.user_id);
      cb(null, u || {});
    },
    async UpsertUser(call: any, cb: any) {
      await users.upsertUser(call.request);
      cb(null, {});
    },
    async SetRole(call: any, cb: any) {
      await roles.setRole(call.request.user_id, call.request.role);
      cb(null, {});
    },
    async ListAddresses(call: any, cb: any) {
      const arr = await addrs.list(call.request.user_id);
      cb(null, { addresses: arr });
    },
    async UpsertDeviceToken(call: any, cb: any) {
      await devs.upsertToken(call.request.user_id, call.request.token, call.request.platform);
      cb(null, {});
    },
  };

  server.addService(svc.AccountService.service, impl);
  const port = process.env.ACCOUNT_GRPC_PORT || '50051';
  server.bindAsync(`0.0.0.0:${port}`, ServerCredentials.createInsecure(), (err, boundPort) => {
    if (err) throw err;
    console.log(`Account gRPC listening on :${boundPort}`);
    server.start();
  });
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});


