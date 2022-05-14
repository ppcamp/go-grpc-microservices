import auth from './auth';
import { UserPasswordServiceClient } from '../gRPC/web/User_passwordServiceClientPb';
import { CreateInput, CreateOutput } from '../gRPC/web/user_password_pb';
import environment from '../environment';
import metadata from '../gRPC/metadata';

interface User {
    first_name: string;
    middle_name: string;
    last_name: string;
    nickname: string;
    email: string;
    password: string;
    birthdate: Date;
}

export interface CreateUser extends User {}
export interface UpdateUser extends User {}

export async function create_user(data: CreateUser): Promise<void>{
    let service = new UserPasswordServiceClient(environment.AuthHost);

    const request = new CreateInput();
    request.setPassword(data.password);
    request.setUser("f53c11ed-2b33-4aab-86da-bd9b8c26a997");

    const m = metadata.with_deadline(10);

    let response: CreateOutput;
    try {
        response = await service.create(request, m);
        console.log(response);
    } catch (err) {
        console.log(err);
        console.log("error");
        return;
    }
    console.log(response);
    console.log("success");
}

export async function update_user(data: UpdateUser): Promise<void>{
    const token: string = auth.getToken();
}

export async function delete_user(): Promise<void>{
    const token: string = auth.getToken();
}
