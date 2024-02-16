interface CookieOptions {
	name: string;
	value: string;
	expiresIn?: number;
	path?: string;
	domain?: string;
	secure?: boolean;
	httpOnly?: boolean;
}

export const AuthCookieName = 'AuthorizationToken';
export const RefreshCookieName = 'RefreshToken';

export const setAuthToken = (token: string) => {
	const opts: CookieOptions = {
		name: 'AuthorizationToken',
		value: `Bearer ${token}`,
		expiresIn: 15,
		httpOnly: true,
		domain: window.location.hostname
	};
	setCookie(opts);
};

export const setRefreshToken = (token: string) => {
	const opts: CookieOptions = {
		name: 'RefreshToken',
		value: `Refresh ${token}`,
		expiresIn: 15,
		httpOnly: true,
		domain: window.location.hostname
	};
	setCookie(opts);
};

const getCookie = (name: string): undefined | string => {
	const nameLenPlus = name.length + 1;
	return document.cookie
		.split(';')
		.map((cookie) => cookie.trim())
		.find((cookie) => cookie.substring(0, nameLenPlus) === `${name}=`)
		?.split('=')[1];
};

const setCookie = (options: CookieOptions): void => {
	let cookieString = `
  ${encodeURIComponent(options.name)}=
  ${encodeURIComponent(options.value)};
  path=${options.path || '/'};
  `;
	if (options.expiresIn) {
		const d = new Date();
		d.setTime(d.getTime() + options.expiresIn * 60 * 1000);
		cookieString += `expires=
    ${d.toUTCString()};`;
	}
	if (options.secure) cookieString += 'secure;';
	if (options.httpOnly) cookieString += 'HttpOnly;';
	document.cookie = cookieString.trim();
};

export const getRefreshToken = (): undefined | string => {
	return getCookie('RefreshToken');
};

export const getAuthToken = (): undefined | string => {
	return getCookie('AuthorizationToken');
};
