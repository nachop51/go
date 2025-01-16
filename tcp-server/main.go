package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
		}

		// conn.Write([]byte("Hello from TCP server\n"))

		// io.WriteString(conn, "Hello from TCP server\n")
		// fmt.Fprintln(conn, "How are you doing?")
		// fmt.Fprintf(conn, "%v", "Well, I hope!")

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Panicln("CONN TIMEOUT")
	}

	read(conn)

	response(conn)

	defer conn.Close()
}

func read(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	var i = 0
	var host = ""
	var method = ""
	var uri = ""

	for scanner.Scan() {
		ln := scanner.Text()
		// fmt.Println(ln)
		// fmt.Fprintf(conn, "I heard you say: %s\n", ln)

		if i == 0 {
			tokens := strings.Fields(ln)

			method = tokens[0]
			uri = tokens[1]
		} else if i == 1 {
			host = strings.Fields(ln)[1]
		}

		if ln == "" {
			break
		}

		i++
	}

	fmt.Printf("%v %v %v\n", method, uri, host)
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello from TCP server</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
