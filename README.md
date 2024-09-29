# UNIX Domain Socket Echo Server

This is a simple UNIX domain socket echo server implemented in Go. The server listens on a UNIX domain socket and handles incoming JSON requests, echoing back the message it receives.

## Features

- Implements a UNIX domain socket server.
- Accepts JSON requests from clients with an `echo` method.
- Returns JSON responses with the echoed message.
- Logs client connections and disconnections.

## Requirements

- Go 1.16 or later
- socat (for testing, optional)
- A Unix-like environment (Linux, macOS)

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/<your-username>/<your-repo>.git
   cd <your-repo>
   ```

2. Build the Go server:

   ```bash
   go build -o unix-echo-server main.go
   ```

3. Alternatively, run the server directly using `go run`:

   ```bash
   go run main.go /tmp/my_socket.sock
   ```

## Usage

To start the server, provide the path to the UNIX domain socket as an argument:

```bash
./unix-echo-server /tmp/my_socket.sock
```

### Client Example (Using `socat`)

You can use `socat` to send a test message to the server:

```bash
echo '{"id": 42, "method": "echo", "params": {"message": "Hello"}}' | socat - UNIX-CONNECT:/tmp/my_socket.sock
```

The server will respond with:

```json
{"id": 42, "result": {"message": "Hello"}}
```

### Client Example (Using Python)

You can also create a simple Python client to send a request:

```python
import socket
import json

socket_path = "/tmp/my_socket.sock"

client_socket = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
client_socket.connect(socket_path)

request = {
    "id": 42,
    "method": "echo",
    "params": {"message": "Hello from Python"}
}

client_socket.sendall((json.dumps(request) + "\n").encode('utf-8'))

response = client_socket.recv(1024)
print("Response:", response.decode('utf-8'))

client_socket.close()
```

## Logging

- The server will log `Client connected` when a new client connects.
- It will log `Client disconnected` when the client disconnects.

## Cleanup

If the server terminates or you stop it, the socket file (`.sock`) may remain on your file system. You can remove it manually if necessary:

```bash
rm /tmp/my_socket.sock
```

## Contributing

Feel free to open issues or submit pull requests if you have suggestions or improvements.

## License

This project is licensed under the MIT License.