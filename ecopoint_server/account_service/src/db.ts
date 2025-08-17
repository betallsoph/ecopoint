import mysql from 'mysql2/promise';
import { config } from './config.ts';

export async function getPool() {
  return mysql.createPool({
    host: config.mysql.host,
    port: config.mysql.port,
    user: config.mysql.user,
    password: config.mysql.password,
    database: config.mysql.database,
    connectionLimit: 5,
    namedPlaceholders: true,
  });
}


