import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = () => {
	// do use authentication
	const user = false;
	if (!user) {
		const redirNum: 300 | 301 | 302 | 303 | 304 | 305 | 306 | 307 | 308 = 302;
		throw redirect(redirNum, '/login');
	}
};
