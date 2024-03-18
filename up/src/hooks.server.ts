import { redirect } from '@sveltejs/kit';
import { postAuth } from '$lib/auth';
import { RefreshCookieName, AuthCookieName, ProtectedRoutes } from '$lib/cookies';
import type { HandleFetch, Handle } from '@sveltejs/kit';

export const handle = (async ({ event, resolve }) => {
	if (ProtectedRoutes.includes(event.url.pathname)) {
		let okay = false;
		const authCookie = event.cookies.get(AuthCookieName);
		const refreshCookie = event.cookies.get(RefreshCookieName);
		if (authCookie && refreshCookie) {
			const authCookieSplit = authCookie.split(' ');
			const refreshCookieSplit = refreshCookie.split(' ');
			if (authCookieSplit.length === 2 && refreshCookieSplit.length === 2) {
				const res = await postAuth(authCookieSplit[1], refreshCookieSplit[1]);
				if (res.valid === true) {
					okay = true;
					if (res.newToken) {
						event.cookies.set(AuthCookieName, `Bearer ${res.newToken}`, {
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
	}
	const response = await resolve(event);
	return response;
}) satisfies Handle;

export const handleFetch = (async ({ request, fetch, event }) => {
	const url = new URL(request.url);
	if (ProtectedRoutes.includes(url.pathname)) {
		let okay = false;
		const authCookie = event.cookies.get(AuthCookieName);
		const refreshCookie = event.cookies.get(RefreshCookieName);
		if (authCookie && refreshCookie) {
			const authCookieSplit = authCookie.split(' ');
			const refreshCookieSplit = refreshCookie.split(' ');
			if (authCookieSplit.length === 2 && refreshCookieSplit.length === 2) {
				const res = await postAuth(authCookieSplit[1], refreshCookieSplit[1]);
				if (res.valid === true) {
					okay = true;
					if (res.newToken) {
						event.cookies.set(AuthCookieName, `Bearer ${res.newToken}`, {
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
	}
	return await fetch(request);
}) satisfies HandleFetch;
