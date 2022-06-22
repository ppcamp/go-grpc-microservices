import { writable } from 'svelte/store';

const _tokenKey = "token";


const getToken = (): string =>  localStorage.getItem(_tokenKey) || "";
const setToken = (token: string): void => localStorage.setItem(_tokenKey, token);
const clearToken = (): void => localStorage.removeItem(_tokenKey);


const Token = writable<string>(getToken());

console.info(getToken());
Token.subscribe(s => setToken(s));

export default {
    clearToken,
    Token
}