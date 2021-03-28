package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("connecting...")
	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 15)
	buf[0] = 5
	copy(buf[1:], []byte("Anybody there?"))
	n, err := conn.Write(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes: [% x]\n", n, buf)
}
