package main

import (
	"net"
)

func main() {
	srv, _ := net.ResolveUDPAddr("udp","127.0.0.1:8081")
	local, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	conn, _ := net.DialUDP("udp", local, srv)
	defer conn.Close()
	conn.Write([]byte("5Anybody there?"))
}