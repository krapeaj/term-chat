## Client

**Description**

A CLI client that allows user to communicate with api-gateway server.

**Communication with the server**

At the start of the program, user will be asked to log in and the client will send a login request to the server.
- The server will create a sessionId associated with the user and store it. 
- The client will receive this sessionId and use it in its subsequent requests.

The server's ip will be 'hard-coded' in the client and all communication will be in HTTP1.1 (for now).


**Commands**

- `/help` - displays instructions and list of commands
- `/create <name> <password>` - create a chat room
- `/delete <name> <password>` - delete a chat room
- `/join <name> <password>` - enter a chat room
- `/leave` - leave a chat room
- `/quit` or `ctrl-c` - quit from client
