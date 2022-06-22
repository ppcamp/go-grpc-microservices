import { writable } from 'svelte/store';

export interface User {
    first_name: string;
    middle_name: string;
    last_name: string;
    nickname: string;
    email: string;
    password: string;
    birthdate: Date;
}


const load_from_cache = (): User => JSON.parse(localStorage.getItem("user"));
export const UserStore = writable<User>(load_from_cache());

UserStore.subscribe(s => localStorage.setItem("user", JSON.stringify(s)));


export default {
    UserStore,
}