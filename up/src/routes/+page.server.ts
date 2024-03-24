import { queryDatabase } from '$lib/query';
import { CDN_HOST } from '$env/static/private';

export const load = async () => {
	const urlArray = await queryDatabase();
	return {
		urlArray: urlArray.map((video) => ({
			url: `https://${CDN_HOST}/stream/${video.uuid}/thumb.png`,
			title: video.title,
			description: video.description,
			videoUrl: `https://${CDN_HOST}/stream/${video.uuid}/${video.manifest_name}`,
			uuid: video.uuid,
			manifest_name: video.manifest_name
		}))
	};
};
