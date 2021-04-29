package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	fmt.Println("listening...")
	inet := &net.UDPAddr{IP: net.IPv4zero, Port: 7000}
	udpConn, err := net.ListenUDP("udp", inet)
	if err != nil {
		panic(err)
	}
	b := make([]byte, 4)
	n, b2, err := udpConn.ReadFromUDP(b)

	if err != nil {
		panic(err)
	}
	fmt.Printf("connecting... read: %d, addr: %v, data: %v, decoded: %v\n",
		n, b2, b[:n], binary.LittleEndian.Uint32(b))
	//fmt.Printf("connecting... read: %d, addr: %v, data: %v, decoded: %v\n",
	//	n, b2, b[:n], binary.LittleEndian.Uint32(b))
}
