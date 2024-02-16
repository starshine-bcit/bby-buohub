import { AUTH_HOST, AUTH_PORT } from '$env/static/private';

interface loginResponse {
	valid: boolean;
	authToken?: string;
	refreshToken?: string;
}

interface authResponse {
	valid: boolean;
	newToken?: string;
}

const head = {
	'Content-Type': 'application/json'
};

export const postAuth = async (authToken: string, refreshToken: string): Promise<authResponse> => {
	const req = new Request(`http://${AUTH_HOST}:${AUTH_PORT}/auth`, {
		method: 'POST',
		body: `{"accessToken": "${authToken}", "refreshToken": "${refreshToken}"}`,
		headers: head
	});
	const res = await fetch(req);
	if (res.ok === true) {
		const js = await res.json();
		if (js.ok === true) {
			if (js.refreshed === true) {
				return {
					valid: true,
					newToken: js.newToken
				};
			} else {
				return { valid: true };
			}
		}
	}
	return { valid: false };
};

export const postRefresh = async (
	authToken: string,
	refreshToken: string
): Promise<authResponse> => {
	const req = new Request(`http://${AUTH_HOST}:${AUTH_PORT}/refresh`, {
		method: 'POST',
		body: `{"accessToken": "${authToken}", "refreshToken": "${refreshToken}"}`,
		headers: head
	});
	const res = await fetch(req);
	if (res.ok === true) {
		const js = await res.json();
		if (js.ok === true) {
			if (js.refreshed === true) {
				return {
					valid: true,
					newToken: js.newToken
				};
			} else {
				return { valid: true };
			}
		}
	}
	return { valid: false };
};

export const postLogin = async (username: string, password: string): Promise<loginResponse> => {
	const req = new Request(`http://${AUTH_HOST}:${AUTH_PORT}/login`, {
		method: 'POST',
		body: `{"username": "${username}", "refreshToken": "${password}"}`,
		headers: head
	});
	const res = await fetch(req);
	if (res.ok === true) {
		const js = await res.json();
		if (js.valid === true) {
			return {
				valid: true,
				authToken: js.accessToken,
				refreshToken: js.refreshToken
			};
		}
	}
	return { valid: false };
};

export const postRegister = async (username: string, password: string): Promise<loginResponse> => {
	const req = new Request(`http://${AUTH_HOST}:${AUTH_PORT}/create`, {
		method: 'POST',
		body: `{"username": "${username}", "refreshToken": "${password}"}`,
		headers: head
	});
	const res = await fetch(req);
	if (res.ok === true) {
		const js = await res.json();
		if (js.created === true) {
			return {
				valid: true,
				authToken: js.accessToken,
				refreshToken: js.refreshToken
			};
		}
	}
	return { valid: false };
};
