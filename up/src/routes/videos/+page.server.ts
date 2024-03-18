import queryDatabase from '$lib/query'; 

/** @type {import('.types').PageServerLoad} */
export async function load() {
    const urlArray = await queryDatabase(); 

	return {
		urlArray: urlArray.map((video) => ({
			 url: `http://localhost:9001/stream/${video.uuid}/thumb.png`,
            title: video.title,
            description: video.description,
            videoUrl: `http://localhost:9001/stream/${video.uuid}/${video.manifest_name}`,
            uuid: video.uuid,
            manifest_name: video.manifest_name
		}))
	};
}

//import * as db from '$lib/db';


//export async function load() {
//    return {
//        post: await db.getPost(params.slug)
//    }/
//}


