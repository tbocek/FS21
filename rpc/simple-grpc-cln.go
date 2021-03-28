package main

import (
	"context"
	"fmt"
	schema_grpc "github.com/tbocek/VSS-RPC/go-gen3"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Connecting...")

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := schema_grpc.NewMessageServiceClient(conn)

	response, err := c.SendMessage(context.Background(), &schema_grpc.AMessage{Code: 5, Message: "Anybody there?"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %v", response)
}
