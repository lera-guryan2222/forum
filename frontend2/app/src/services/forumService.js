import { ForumServiceClient } from '../proto/forum_grpc_web_pb';
import { MessageRequest, GetMessagesRequest } from '../proto/forum_pb';

const client = new ForumServiceClient('http://localhost:50052', null, null);

export const sendMessage = (userId, content) => {
    const request = new MessageRequest();
    request.setUserId(userId);
    request.setContent(content);
    return new Promise((resolve, reject) => {
        client.sendMessage(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response);
            }
        });
    });
};

export const getMessages = (limit, offset) => {
    const request = new GetMessagesRequest();
    request.setLimit(limit);
    request.setOffset(offset);
    return new Promise((resolve, reject) => {
        client.getMessages(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response.getMessagesList());
            }
        });
    });
};
