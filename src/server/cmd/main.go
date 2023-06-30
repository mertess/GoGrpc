package main

import (
	"log"
	"net"

	"github.com/mertess/GoGrpc/gen"
	"github.com/mertess/GoGrpc/src/server/wordsSeen"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	gen.RegisterWordsSeenServiceServer(server, wordsSeen.NewWordsSeenServiceServer())

	server.Serve(listener)
}
