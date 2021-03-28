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
	for {
		b := make([]byte, 1000)
		n, err := conn.Read(b)
		if err != nil {
			return
		}
		if n > 0 {
			fmt.Printf("connecting... read: %d, addr: %v, data: [% x], decoded: %v\n",
				n, conn.RemoteAddr(), b[:n], string(b[:]))
		}
	}
}
