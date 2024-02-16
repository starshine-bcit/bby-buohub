import { redirect } from '@sveltejs/kit';

export const actions = {
	default: async ({ request }) => {
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
		// do post request
		throw redirect(303, '/');
	}
};
