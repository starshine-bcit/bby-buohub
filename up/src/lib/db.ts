import { DB_USER, DB_HOST, DB_PORT, DB_PASSWORD, DB_NAME } from '$env/static/private';
import { v4 as uuidv4 } from 'uuid';
import mariadb from 'mariadb';

const pool = mariadb.createPool({
	host: DB_HOST,
	user: DB_USER,
	port: +DB_PORT,
	password: DB_PASSWORD,
	database: DB_NAME,
	connectionLimit: 5
});

const defaultUser = 'anon';

export interface insertResponse {
	ok: boolean;
	uuid?: string;
}

export const insertNewVideo = async (
	videoName: string,
	videoDescription: string
): Promise<insertResponse> => {
	let conn;
	const uid = uuidv4();
	const now = new Date().toISOString().slice(0, 19).replace('T', ' ');
	try {
		conn = await pool.getConnection();
		const query =
			'INSERT into videos (uuid, uploaded_by, process_complete, uploaded_at, title, description) value (?, ?, ?, ?, ?, ?)';
		const vals = [uid, defaultUser, false, now, videoName, videoDescription];
		const res = await conn.query(query, vals);

		if (!res || res['affectedRows'] !== 1) {
			throw new Error('Could not insert into database');
		}
	} catch (err) {
		return { ok: false };
	} finally {
		if (conn) {
			conn.end();
		}
	}
	return { ok: true, uuid: uid };
};
