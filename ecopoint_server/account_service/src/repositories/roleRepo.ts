import type { Pool } from 'mysql2/promise';

export class RoleRepo {
  private pool: Pool;
  constructor(pool: Pool) { this.pool = pool; }

  async setRole(user_id: string, role: 'customer'|'collector'|'admin') {
    const sql = `INSERT INTO user_roles (user_id, role)
                 VALUES (?, ?)
                 ON DUPLICATE KEY UPDATE role = VALUES(role)`;
    await this.pool.execute(sql, [user_id, role]);
  }

  async listRoles(user_id: string): Promise<string[]> {
    const [rows] = await this.pool.query(`SELECT role FROM user_roles WHERE user_id = ?`, [user_id]);
    return (rows as any[]).map(r => r.role as string);
  }
}


