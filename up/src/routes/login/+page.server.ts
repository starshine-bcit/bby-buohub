import { redirect } from '@sveltejs/kit';
import { postLogin } from '$lib/auth.js';
import { RefreshCookieName, AuthCookieName } from '$lib/cookies.js';

export const actions = {
	default: async ({ request, cookies }) => {
		const formData = await request.formData();
		const e = { err: 'Invalid Login' };
		const username = formData.get('username');
		const password = formData.get('password');
		if (!username || !password) {
			return e;
		}
		if (String(username).length < 5 || String(password).length < 5) {
			return e;
		}
		const loginRes = await postLogin(String(username), String(password));
		if (loginRes.valid === true) {
			cookies.set(AuthCookieName, `Bearer ${loginRes.authToken}`, {
				httpOnly: true,
				sameSite: 'lax',
				secure: false,
				maxAge: 60 * 60 * 24,
				path: '/'
			});
			cookies.set(RefreshCookieName, `Refresh ${loginRes.refreshToken}`, {
				httpOnly: true,
				sameSite: 'lax',
				secure: false,
				maxAge: 60 * 60 * 24,
				path: '/'
			});
			throw redirect(303, '/');
		} else {
			return e;
		}
	}
};
