import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';
import { postAuth, postUpload } from '$lib/auth';
import { insertNewVideo } from '$lib/db';
import { RefreshCookieName, AuthCookieName } from '$lib/cookies';

export const load: PageServerLoad = async ({ cookies }) => {
	let okay = false;
	const authCookie = cookies.get(AuthCookieName);
	const refreshCookie = cookies.get(RefreshCookieName);
	if (authCookie && refreshCookie) {
		const authCookieSplit = authCookie.split(' ');
		const refreshCookieSplit = refreshCookie.split(' ');
		if (authCookieSplit.length === 2 && refreshCookieSplit.length === 2) {
			const res = await postAuth(authCookieSplit[1], refreshCookieSplit[1]);
			if (res.valid === true) {
				okay = true;
				if (res.newToken) {
					cookies.set(AuthCookieName, `Bearer ${res.newToken}`, {
						httpOnly: true,
						sameSite: 'lax',
						secure: false,
						maxAge: 60 * 60 * 24,
						path: '/'
					});
				}
			}
		}
	}
	if (!okay) {
		throw redirect(302, '/login');
	}
};

export const actions = {
	default: async ({ request, cookies }) => {
		let okay = false;
		const authCookie = cookies.get(AuthCookieName);
		const refreshCookie = cookies.get(RefreshCookieName);
		if (authCookie && refreshCookie) {
			const authCookieSplit = authCookie.split(' ');
			const refreshCookieSplit = refreshCookie.split(' ');
			if (authCookieSplit.length === 2 && refreshCookieSplit.length === 2) {
				const res = await postAuth(authCookieSplit[1], refreshCookieSplit[1]);
				if (res.valid === true) {
					okay = true;
					if (res.newToken) {
						cookies.set(AuthCookieName, `Bearer ${res.newToken}`, {
							httpOnly: true,
							sameSite: 'lax',
							secure: false,
							maxAge: 60 * 60 * 24,
							path: '/'
						});
					}
				}
			}
		}
		if (!okay) {
			throw redirect(302, '/login');
		}
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

