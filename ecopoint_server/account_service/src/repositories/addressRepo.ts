import type { Pool } from 'mysql2/promise';

export interface Address {
  id?: number;
  user_id: string;
  label?: string | null;
  full_text: string;
  lat?: number | null;
  lng?: number | null;
  is_default?: boolean;
}

export class AddressRepo {
  private pool: Pool;
  constructor(pool: Pool) { this.pool = pool; }

  async list(user_id: string) {
    const [rows] = await this.pool.query(`SELECT * FROM addresses WHERE user_id = ? ORDER BY id DESC`, [user_id]);
    return rows as any[];
  }

  async upsert(addr: Address) {
    const sql = `INSERT INTO addresses (user_id, label, full_text, lat, lng, is_default)
                 VALUES (:user_id, :label, :full_text, :lat, :lng, :is_default)
                 ON DUPLICATE KEY UPDATE
                   label = VALUES(label),
                   full_text = VALUES(full_text),
                   lat = VALUES(lat),
                   lng = VALUES(lng),
                   is_default = VALUES(is_default)`;
    const params = {
      user_id: addr.user_id,
      label: addr.label ?? null,
      full_text: addr.full_text,
      lat: addr.lat ?? null,
      lng: addr.lng ?? null,
      is_default: addr.is_default ? 1 : 0,
    } as const;
    await this.pool.execute(sql, params as any);
  }
}


