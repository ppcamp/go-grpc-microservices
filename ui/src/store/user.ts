import { writable } from 'svelte/store';
import type { User } from '../models/user';


const load_from_cache = (): User => JSON.parse(localStorage.getItem("user"));
export const UserStore = writable<User>(load_from_cache());

UserStore.subscribe(s => localStorage.setItem("user", JSON.stringify(s)));


export default {
    UserStore,
}