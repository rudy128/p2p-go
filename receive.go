package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Recieve() {
	// Listen on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var filename string
	buf := make([]byte, 256) // Buffer to hold the filename
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading filename:", err)
		return
	}
	filename = string(buf[:n-1]) // Remove the newline character

	// Create the file using the received filename
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// Copy the received data to the file
	_, err = io.Copy(outFile, conn)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("File received successfully")
}
