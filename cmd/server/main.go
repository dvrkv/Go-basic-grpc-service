package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dvrkv/grpcadder/pkg/adder"
	"github.com/dvrkv/grpcadder/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)
	port := ":8080"
	fmt.Print("Server listen on port", port, "\n")
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
