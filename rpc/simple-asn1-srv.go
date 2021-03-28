package main

import (
	"encoding/asn1"
	"fmt"
	"math/big"
	"net"
)

type TestASN struct {
	Code    *big.Int
	Message string
}

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
	b := make([]byte, 21)
	_, _ = conn.Read(b)
	var t TestASN
	_, err = asn1.Unmarshal(b, &t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read: %v\n", t)
}
