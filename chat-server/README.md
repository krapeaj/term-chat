## API Gateway

**Description**

A light-weight go REST API server and application server that handles client requests and establishes websocket connection.

**REST API**

- Sign up - `PUT /api/signup`. Requires `Authorization` header with Basic-encrypted credentials. 
- Login - `POST /api/login`. Requires `Authorization` header with Basic-encrypted credentials.
- Logout - `POST /api/logout`. Requires header `session-id`.
- Create Chat - `PUT /api/chat`. Requires headers `session-id`, `chat-name`, and `password` (chat password).
- Delete Chat - `DELETE /api/chat`. Requires headers `session-id`, `chat-name`, and `password` (chat password).
- Join Chat - `GET /websocket` (wss). Requires headers `session-id`, `chat-name`, and `password` (chat password). 
    - client is disconnected from chat when websocket connection is closed or lost.

