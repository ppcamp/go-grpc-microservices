import auth from './auth';
import type { User } from '../models/user';

export interface CreateUser extends User {}
export interface UpdateUser extends User {}

async function create_user(data: CreateUser): Promise<void>{}

async function update_user(data: UpdateUser): Promise<void>{
    const token: string = auth.getToken();
}

async function delete_user(): Promise<void>{
    const token: string = auth.getToken();
}
