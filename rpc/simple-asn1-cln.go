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

	buf := []byte{0x30, 0x13, 0x02, 0x01, 0x05, 0x16, 0x0e, 0x41, 0x6e, 0x79, 0x62, 0x6f,
		0x64, 0x79, 0x20, 0x74, 0x68, 0x65, 0x72, 0x65, 0x3f}

	n, err := conn.Write(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes: [% x]\n", n, buf)
}
