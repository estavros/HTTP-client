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

	// Optional custom headers
	headers := map[string]string{
		"User-Agent": "MyClient/1.0",
		"Accept":     "*/*",
	}

	// Connect via TCP
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Build HTTP GET request
	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\n", path, host)
	for k, v := range headers {
		request += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	request += "Connection: close\r\n\r\n"

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
			switch statusCode {
			case 200:
				fmt.Println("✅ HTTP Status: 200 OK")
			case 404:
				fmt.Println("❌ HTTP Status: 404 Not Found")
			case 500:
				fmt.Println("💥 HTTP Status: 500 Internal Server Error")
			default:
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
