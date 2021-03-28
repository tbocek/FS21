package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("listening...")
	tcpConn, err := net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}
	conn, err := tcpConn.Accept() //do this in a go routine

	if err != nil {
		panic(err)
	}
	b := make([]byte, 15)
	n, _ := conn.Read(b)
	fmt.Printf("connecting... read: %d, addr: %v, data: [% x], decoded: %v\n",
		n, conn.RemoteAddr(), b[:n], string(b[1:]))
}
