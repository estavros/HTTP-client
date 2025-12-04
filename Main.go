package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	host := "example.com"
	port := "80"
	path := "/"

	// Connect via TCP
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Build HTTP GET request
	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", path, host)

	// Send request
	_, err = conn.Write([]byte(request))
	if err != nil {
		panic(err)
	}

	// Read response
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}
