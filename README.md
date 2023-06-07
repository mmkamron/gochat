# GoChat - Simple Chat Server

GoChat is a simple chat server implemented in Go. It allows multiple clients to connect and exchange messages in real-time using TCP sockets.

## Features

- Multiple clients can connect to the chat server concurrently.
- Clients can send text messages that are broadcasted to all connected clients.
- Self-host the GoChat server for full control of your chat environment.

## Getting Started

- To get started, download the appropriate binary for your operating system from the [releases](https://github.com/mmkamron/gochat/releases) section. 
- Run the client and provide the server address and port to connect to the GoChat server. If host and port not specified, GoChat connects to my server by default

## Self-hosting
1. Clone the repository:
```bash
git clone https://github.com/your-username/gochat && cd gochat
```
2. (Optional) change the port number in Dockerfile
3. Build and run the docker container
```
docker build -t gochat-server .
docker run -p <port-number>:<port-number> gochat-server
```
The chat server will start and be accessible on specified port number (1337 by default).
