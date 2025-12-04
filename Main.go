package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
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

	reader := bufio.NewReader(conn)

	// Read status line
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print("Status Line: ", statusLine)

	// Parse status code
	parts := strings.SplitN(statusLine, " ", 3)
	if len(parts) >= 2 {
		statusCode, err := strconv.Atoi(parts[1])
		if err == nil {
			if statusCode == 200 {
				fmt.Println("✅ HTTP Status: 200 OK")
			} else {
				fmt.Printf("⚠ HTTP Status: %d\n", statusCode)
			}
		}
	}

	// Print the rest of the response
	fmt.Println("\n=== Full Response ===")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}
