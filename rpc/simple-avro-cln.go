package main

import (
	"fmt"
	"github.com/linkedin/goavro"
	"io/ioutil"
	"net"
)

func main() {
	fmt.Println("Connecting...")

	schema, err := ioutil.ReadFile("schema.avsc")
	if err != nil {
		fmt.Println(err)
	}
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		fmt.Println(err)
	}

	map1 := map[string]interface{}{
		"request": "Anybody there?",
		"code":    5,
	}
	binary, err := codec.BinaryFromNative(nil, map1)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	n, err := conn.Write(binary)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes: [% x]\n", n, binary)
}
