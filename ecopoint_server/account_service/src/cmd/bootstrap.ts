import { getPool } from '../db.ts';
import fs from 'node:fs';
import path from 'node:path';

async function main() {
  const pool = await getPool();
  const schemaPath = path.resolve(process.cwd(), 'src/schema.sql');
  const sql = fs.readFileSync(schemaPath, 'utf8');
  // Split on ; for simple execution
  const statements = sql.split(/;\s*\n/).map(s => s.trim()).filter(Boolean);
  for (const stmt of statements) {
    await pool.query(stmt);
  }
  console.log('MySQL schema ensured.');
  await pool.end();
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});


