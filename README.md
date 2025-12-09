# Simple Go TCP HTTP Client

This project demonstrates how to implement a minimal HTTP client in Go using **raw TCP sockets**, without relying on the built-in `net/http` package.

It manually:

* Opens a TCP connection to a server
* Constructs and sends an HTTP GET request
* Reads and parses the HTTP status line
* Prints custom status messages
* Prints the full HTTP response

---

## 🚀 Features

* Connects to any host and port via TCP
* Sends a custom HTTP GET request
* Supports custom request headers
* Parses and interprets HTTP status codes
* Prints the complete server response

---

## 📄 Code Overview

The client:

1. Connects to `example.com` on port `80`
2. Sends a GET request to the root path `/`
3. Reads the HTTP response line-by-line
4. Parses the status code to show friendly output:

   * `200 OK`
   * `404 Not Found`
   * `500 Internal Server Error`
   * Any other status code is labeled as "unknown"

---

## 🧪 Example Output

```
Status Line: HTTP/1.1 200 OK
✅ HTTP Status: 200 OK

=== Full Response ===
HTTP/1.1 200 OK
Content-Type: text/html; charset=UTF-8
...
<html>...</html>
```

---

## 🛠 How to Run

Make sure you have Go installed.

```bash
go run main.go
```

The program will:

* Connect to the server
* Send the HTTP request
* Print the parsed status and the full response

---

## 🧩 Customization

You can modify these variables:

```go
host := "example.com"
port := "80"
path := "/"
```

Add or change custom headers:

```go
headers := map[string]string{
    "User-Agent": "MyClient/1.0",
    "Accept": "*/*",
}
```

