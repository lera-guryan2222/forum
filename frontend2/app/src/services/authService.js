import { AuthServiceClient } from '../proto/auth_grpc_web_pb';
import { LoginRequest, RegisterRequest } from '../proto/auth_pb';

const client = new AuthServiceClient('http://localhost:50051', null, null);

export const login = (username, password) => {
    const request = new LoginRequest();
    request.setUsername(username);
    request.setPassword(password);
    return new Promise((resolve, reject) => {
        client.login(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response.getToken());
            }
        });
    });
};

export const register = (username, password) => {
    const request = new RegisterRequest();
    request.setUsername(username);
    request.setPassword(password);
    return new Promise((resolve, reject) => {
        client.register(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response.getMessage());
            }
        });
    });
};
