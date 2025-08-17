import type { Pool } from 'mysql2/promise';

export class DeviceRepo {
  private pool: Pool;
  constructor(pool: Pool) { this.pool = pool; }

  async upsertToken(user_id: string, token: string, platform: 'ios'|'android'|'web') {
    const sql = `INSERT INTO devices (user_id, fcm_token, platform, last_seen_at)
                 VALUES (?, ?, ?, NOW())`;
    await this.pool.execute(sql, [user_id, token, platform]);
  }

  async listTokensForRole(role: 'collector'|'customer') {
    const [rows] = await this.pool.query(`
      SELECT d.fcm_token
      FROM devices d
      JOIN user_roles ur ON ur.user_id = d.user_id
      WHERE ur.role = ?
    `, [role]);
    return (rows as any[]).map(r => r.fcm_token as string);
  }
}


