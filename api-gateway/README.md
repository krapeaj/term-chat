## API Gateway

**Description**

A light-weight go REST API server that handles client requests and allows communication between different services.

**REST API**

All requests requires header `userId`.
- Login - `POST /api/login` Requires header encrypted `password`.
- Logout - `POST /api/logout`
- Create Chat - `PUT /api/chat`
- Delete Chat - `DELETE /api/chat/{chatId}`
- Enter Chat - `GET /api/chat/{chatId}` 
- Leave Chat - `DELETE /api/chat/{chatId}`
- Send Message - `POST /api/chat/{chatId}`

**Etc**

- Timeout: 30 seconds

