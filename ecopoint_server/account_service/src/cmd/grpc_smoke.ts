import { credentials, loadPackageDefinition } from '@grpc/grpc-js';
import { loadSync } from '@grpc/proto-loader';
import path from 'node:path';

const PROTO_PATH = path.resolve(process.cwd(), '../proto/account.proto');
const pkgDef = loadSync(PROTO_PATH, { keepCase: true, longs: String, enums: String, defaults: true, oneofs: true });
const grpcObj = loadPackageDefinition(pkgDef) as any;
const svc = grpcObj.ecopoint.account.v1;

async function main() {
  const client = new svc.AccountService('localhost:' + (process.env.ACCOUNT_GRPC_PORT || '50051'), credentials.createInsecure());

  await new Promise<void>((resolve, reject) => {
    client.UpsertUser({ user_id: 'u_grpc', display_name: 'GRPC User' }, (err: any) => err ? reject(err) : resolve());
  });

  const user = await new Promise<any>((resolve, reject) => {
    client.GetUser({ user_id: 'u_grpc' }, (err: any, res: any) => err ? reject(err) : resolve(res));
  });
  console.log('GetUser:', user);
}

main().catch((e) => { console.error(e); process.exit(1); });


