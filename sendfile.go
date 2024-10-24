package main

import (
	"bufio"
	"io"
	"net"
	"os"
)

func SendFile() {
	art()
	TcpClient(FilePath(), IPAddr())
}

func FilePath() string {
	reader := bufio.NewReader(os.Stdin)
	print("File Path: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading input:", err)
		return "" // Return an empty string on error
	}
	return input[:len(input)-1]
}

func IPAddr() string {
	reader := bufio.NewReader(os.Stdin)
	print("Receiver's IP Addr: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading input:", err)
		return "" // Return an empty string on error
	}
	return input[:len(input)-1]
}

func TcpClient(filePath string, serverAddr string) {
	file, err := os.Open(filePath)
	if err != nil {
		println("Error opening file:", err)
		return
	}
	println(file.Name())
	defer file.Close()
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(file.Name() + "\n"))
	if err != nil {
		println("Error sending filename:", err)
		return
	}

	_, err = io.Copy(conn, file)
	if err != nil {
		println("Error sending file:", err)
		return
	}

	println("File sent successfully")
}
