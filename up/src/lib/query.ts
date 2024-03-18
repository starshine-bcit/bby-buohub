import { DB_USER, DB_HOST, DB_PORT, DB_PASSWORD, DB_NAME } from '$env/static/private';
import mariadb from 'mariadb';

interface Video {
    uuid: string;
    title: string;
    description: string;
    manifest_name: string;
    url: string;
}

async function queryDatabase(): Promise<Video[]> {
    const pool = mariadb.createPool({
        host: DB_HOST,
        user: DB_USER,
        port: +DB_PORT,
        password: DB_PASSWORD,
        database: DB_NAME,
        connectionLimit: 5
    });

    let conn;
    try {
        conn = await pool.getConnection();
        console.log("Connected to MariaDB on localhost:6001")
        const videoQuery = 'SELECT * FROM videos ORDER BY id DESC LIMIT 9';
        const videoResults: Video[] = await conn.query(videoQuery);

        return videoResults.map(video => ({
            url: `http://localhost:9001/stream/${video.uuid}/thumb.png`,
            title: video.title,
            description: video.description,
            videoUrl: `http://localhost:9001/stream/${video.uuid}/${video.manifest_name}`,
            uuid: video.uuid,
            manifest_name: video.manifest_name
        }));
    } catch (error) {
        throw new Error('Error querying database: ');
    } finally {
        if (conn) {
            conn.release(); // Release connection back to the pool
        }
    }
}

export default queryDatabase;
