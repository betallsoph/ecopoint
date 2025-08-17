import { getPool } from '../db.ts';
import { UserRepo } from '../repositories/userRepo.ts';
import { RoleRepo } from '../repositories/roleRepo.ts';
import { AddressRepo } from '../repositories/addressRepo.ts';
import { DeviceRepo } from '../repositories/deviceRepo.ts';

async function main() {
  const pool = await getPool();
  const users = new UserRepo(pool);
  const roles = new RoleRepo(pool);
  const addrs = new AddressRepo(pool);
  const devs  = new DeviceRepo(pool);

  // upsert user
  await users.upsertUser({ user_id: 'uid_demo', email: 'demo@example.com', display_name: 'Demo User' });
  // set role
  await roles.setRole('uid_demo', 'customer');
  // insert address
  await addrs.upsert({ user_id: 'uid_demo', full_text: '123 Demo St, HCMC', lat: 10.77, lng: 106.7, is_default: true });
  // device token
  await devs.upsertToken('uid_demo', 'fake-token', 'android');

  const me = await users.getUser('uid_demo');
  const myRoles = await roles.listRoles('uid_demo');
  const myAddrs = await addrs.list('uid_demo');
  const collectorTokens = await devs.listTokensForRole('collector');

  console.log('User:', me);
  console.log('Roles:', myRoles);
  console.log('Addresses:', myAddrs.length);
  console.log('Collector tokens:', collectorTokens.length);

  await pool.end();
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});


