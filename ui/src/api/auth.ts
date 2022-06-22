import environment from '../environment';
import auth from '../stores/auth';
import { UserPasswordServiceClient } from "../gRPC/web/User_passwordServiceClientPb";
import { CreateInput, CreateOutput } from "../gRPC/web/user_password_pb";

let token: string;
auth.Token.subscribe(t => token = t);

interface LoginPayload {
    nick: string;
    password: string;
}

const getToken = (): string => token.slice();

/**
 * Login into system
 *
 * @throws Exceptions if failed to fetch data
 */
async function login(data: LoginPayload): Promise<void>{
    let new_token: string;
    let service = new UserPasswordServiceClient(environment.AuthHost);

    const request = new CreateInput();
    const header = {};
    let response: CreateOutput;

    try {
        response = await service.create(request, header);
    } catch (err) {
        console.error(err)
    }

    // auth.Token.set(new_token);
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
    auth.Token.set(new_token);
}


/**
 * Logout and invalidate the current session
 */
 async function logout(): Promise<void>{
    // 1. do request to API to invalidate the token

    // 2. remove token
    auth.clearToken();
}



export default {
    logout,
    refresh,
    login,
    getToken,
}