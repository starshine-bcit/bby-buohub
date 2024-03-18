import { postUpload } from '$lib/auth';
import { insertNewVideo } from '$lib/db';

export const actions = {
	default: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('file') as File;
		const title = formData.get('title');
		const description = formData.get('description');
		if (!file.name) {
			return { err: 'Bad choice!' };
		}
		if (!title || !description) {
			return { err: 'Bad input!' };
		}
		const insRes = await insertNewVideo(String(title), String(description));
		if (!insRes.ok) {
			return { err: 'Failed to insert into database!' };
		}
		const uploadOk = await postUpload(file, String(insRes.uuid));
		if (uploadOk) {
			return;
		}
		return { err: 'Invalid upload!' };
	}
};
