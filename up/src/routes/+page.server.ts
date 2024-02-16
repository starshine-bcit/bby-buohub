import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';
import { postAuth } from '$lib/auth';
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
						maxAge: 60 * 15,
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
