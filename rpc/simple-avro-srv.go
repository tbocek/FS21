package main

import (
	"fmt"
	"github.com/linkedin/goavro"
	"io/ioutil"
	"net"
)

func main() {
	fmt.Println("Listening...")

	schema, err := ioutil.ReadFile("schema.avsc")
	if err != nil {
		fmt.Println(err)
	}
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		fmt.Println(err)
	}

	tcpConn, err := net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}
	conn, err := tcpConn.Accept() //do this in a go routine

	if err != nil {
		panic(err)
	}

	binary := make([]byte, 100)
	n, _ := conn.Read(binary)

	native, _, err := codec.NativeFromBinary(binary[:n])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("read: %v\n", native)
}
