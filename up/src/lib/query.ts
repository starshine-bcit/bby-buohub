import { DB_USER, DB_HOST, DB_PORT, DB_PASSWORD, DB_NAME, CDN_HOST } from '$env/static/private';
import mariadb from 'mariadb';

export interface Video {
	uuid: string;
	title: string;
	description: string;
	manifest_name: string;
	url: string;
}

const pool = mariadb.createPool({
	host: DB_HOST,
	user: DB_USER,
	port: +DB_PORT,
	password: DB_PASSWORD,
	database: DB_NAME,
	connectionLimit: 5
});

export const queryDatabase = async (): Promise<Video[]> => {
	let conn;
	try {
		conn = await pool.getConnection();
		const videoQuery =
			'SELECT * FROM videos WHERE process_complete = true ORDER BY id DESC LIMIT 9';
		const videoResults: Video[] = await conn.query(videoQuery);

		return videoResults.map((video) => ({
			url: `https://${CDN_HOST}/stream/${video.uuid}/thumb.png`,
			title: video.title,
			description: video.description,
			videoUrl: `https://${CDN_HOST}/stream/${video.uuid}/${video.manifest_name}`,
			uuid: video.uuid,
			manifest_name: video.manifest_name
		}));
	} catch (error) {
		throw new Error(`Error querying database: ${error}`);
	} finally {
		if (conn) {
			conn.release();
		}
	}
};
