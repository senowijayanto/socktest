package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"syscall"
)

// Request represents the structure of a client request
type Request struct {
	ID     int                    `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

// Response represents the structure of the server's response
type Response struct {
	ID     int                    `json:"id"`
	Result map[string]interface{} `json:"result"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <socket_path>")
		os.Exit(1)
	}

	socketPath := os.Args[1]

	// If the socket file already exists, remove it
	syscall.Unlink(socketPath)

	// Create a UNIX domain socket
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		// Log client disconnection
		fmt.Println("Client disconnected")
	}()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from connection:", err)
			}
			return
		}

		var req Request
		if err := json.Unmarshal(buffer[:n], &req); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		if req.Method == "echo" {
			// Prepare the response
			response := Response{
				ID: req.ID,
				Result: map[string]interface{}{
					"message": req.Params["message"],
				},
			}

			// Convert response to JSON
			responseJSON, err := json.Marshal(response)
			if err != nil {
				fmt.Println("Error encoding JSON:", err)
				return
			}

			// Write response back to the client
			conn.Write(append(responseJSON, '\n')) // Ensure the response ends with newline
		}
	}
}
