import { writable } from 'svelte/store';
import type { User } from '../models/user';

// stores
const baseUser: User = {
    first_name: "",
    middle_name: "",
    last_name: "",
    nickname: "",
    email: "",
    password: "",
    birthdate: new Date(),
};

export const UserStore = writable<User>(baseUser);