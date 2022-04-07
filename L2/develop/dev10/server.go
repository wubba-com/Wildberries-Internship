package main

import (
	"fmt"
	"net"
)

const (
	port = ":3000"
)

func main() {
	server, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()

	for {
		// get connect (client)
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go func(conn net.Conn) {
			for {
				input := make([]byte, 1<<(10*1))

				// read client data
				n, err := conn.Read(input)
				if n == 0 || err != nil {
					fmt.Println("[err] read:", err)
					return
				}

				for i, j := 0, n; i < j; i, j = i+1, j-1 {
					input[i], input[j] = input[j], input[i]
				}

				// send data to client
				conn.Write(input)
			}
		}(conn)
	}
}
