# term-chat

An chatting application that uses terminal as user interface. Slightly over-engineered for learning purposes.

## Components

Each component is spinned up in a Docker container using `docker-compose`.

* `chat-server`: the backend API server and application that stores chat information and uses websocket to broadcast messages to clients in certain chat rooms. Its API description can be found [here](https://github.com/krapeaj/term-chat/blob/master/chat-server/README.md). **Note that a certificate from CA needs to be obtained for HTTPS. For use without TLS, `nginx` and `chat-server` needs to be reconfigured.**
* `client`: A prototype of a terminal client that can be installed.
* `nginx`: Reverse proxy.
* `redis`: User session storage.
* `mariadb`: User account persistence. (chat history using batch server in consideration)

## For local testing

**A TLS cert needs to be generated and registered to client's computer network security settings as Client to Server communication is in HTTPS.**

1. Clone: `git clone https://github.com/krapeaj/term-chat.git`
2. Server: `cd` into clone directory and `./start_docker-compose` (requires docker)
3. Client: `cd` into `client` directory and `go build` (requires go)
    - for client commands, see [here](https://github.com/krapeaj/term-chat/blob/master/client/README.md)

## Things yet to be done 
* Deployment to the open world on a Cloud platform.
    - Requires a valid certificate from a CA unless TLS is not desired.
    - Kubernetes integration
* Browser client.
* Scalability improvement and service discovery (k8s).

