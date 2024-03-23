import { CDN_HOST, CDN_PORT } from '$env/static/private';

export const load = ({ params }) => {
	const uuid = params['uuid'] as string;
	const manifest_name = params['manifest_name'] as string;
	return {
		uuid,
		manifest_name,
		CDN_HOST,
		CDN_PORT
	};
};
