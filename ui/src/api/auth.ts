const _tokenKey = "token";


export function getToken(): string {
    return localStorage.getItem(_tokenKey) || "";
}

export function setToken(token: string): void {
    localStorage.setItem(_tokenKey, token);
}

export function clearToken(): void {
    localStorage.removeItem(_tokenKey);
}


interface LoginPayload {
    nick: string;
    password: string;
}

/**
 * Login into system
 *
 * @throws Exceptions if failed to fetch data
 */
async function login(data: LoginPayload): Promise<void>{
    let new_token: string;
    setToken(new_token);
}

/**
 * Refresh the current JWT token basing on some that already stored
 *
 * Note that if fails, you must to redirect to login page.
 *
 * @throws Exceptions if failed to fetch data
 */
async function refresh(): Promise<void>{
    // 1. do request to API
    let new_token: string;

    // 2. update token
    setToken(new_token);
}


/**
 * Logout and invalidate the current session
 */
 async function logout(): Promise<void>{
    // 1. do request to API to invalidate the token

    // 2. remove token
    clearToken();
}



export default {
    getToken,
    clearToken,
    setToken,
    logout,
    refresh,
    login
}