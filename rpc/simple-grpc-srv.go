package main

import (
	"context"
	"fmt"
	schemagrpc "github.com/tbocek/VSS-RPC/go-gen3"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func (s *Server) SendMessage(_ context.Context, in *schemagrpc.AMessage) (*schemagrpc.Empty, error) {
	log.Printf("Receive message %s %d", in.Message, in.Code)
	return &schemagrpc.Empty{}, nil
}

func main() {
	fmt.Println("Listening...")

	tcpConn, err := net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	s := Server{}
	schemagrpc.RegisterMessageServiceServer(grpcServer, &s)
	grpcServer.Serve(tcpConn)

	fmt.Printf("done\n")
}
