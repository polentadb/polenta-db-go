package main

import (
	"fmt"
	polenta "github.com/polentadb/polenta-core-go/polenta"
	"net"
	"os"
	"strings"
)

func port(args []string) string {
	if len(args) > 1 && strings.HasPrefix(args[1], "port=") {
		return strings.Split(args[1], "=")[1]
	}
	return "9000"
}

func main() {
	port := port(os.Args)
	address := "localhost:" + port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("PolentaDB is waiting for statements on port " + port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go func(c net.Conn) {
			handleClient(c)
		}(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		statement := string(buffer[:n])
		fmt.Println("received statement:", statement)
		response := polenta.Run(statement)
		conn.Write(encodeResponse(response))
	}
}

func encodeResponse(response polenta.Response) []byte {
	responseStr := response.Message
	return []byte(responseStr)
}
