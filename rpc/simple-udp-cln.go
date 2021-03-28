package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	fmt.Println("connecting...")
	conn, err := net.Dial("udp", "127.0.0.1:7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, 77)
	_, err = conn.Write(buf)
	if err != nil {
		panic(err)
	}
}
