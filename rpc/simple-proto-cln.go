package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/tbocek/VSS-RPC/go-gen2"
	"net"
)

func main() {
	fmt.Println("Connecting...")

	m := &pb.AMessage{Id: 5, Message: "Anybody there?"}
	out, err := proto.Marshal(m)

	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	n, err := conn.Write(out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes: [% x]\n", n, out)
}
