import type { Pool } from 'mysql2/promise';

export interface User {
  user_id: string;
  email?: string | null;
  phone?: string | null;
  display_name?: string | null;
  avatar_url?: string | null;
}

export class UserRepo {
  private pool: Pool;
  constructor(pool: Pool) { this.pool = pool; }

  async upsertUser(u: User) {
    const sql = `INSERT INTO users (user_id, email, phone, display_name, avatar_url)
                 VALUES (:user_id, :email, :phone, :display_name, :avatar_url)
                 ON DUPLICATE KEY UPDATE
                   email = VALUES(email),
                   phone = VALUES(phone),
                   display_name = VALUES(display_name),
                   avatar_url = VALUES(avatar_url)`;
    const params = {
      user_id: u.user_id,
      email: u.email ?? null,
      phone: u.phone ?? null,
      display_name: u.display_name ?? null,
      avatar_url: u.avatar_url ?? null,
    } as const;
    await this.pool.execute(sql, params as any);
  }

  async getUser(user_id: string): Promise<User | null> {
    const [rows] = await this.pool.query(`SELECT * FROM users WHERE user_id = ?`, [user_id]);
    const arr = rows as any[];
    return arr[0] || null;
  }
}


