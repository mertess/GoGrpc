package main

import (
	"context"
	"log"

	"github.com/mertess/GoGrpc/gen"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := gen.NewWordsSeenServiceClient(conn)

	response, err := client.CheckWord(context.Background(), &gen.WordsSeenRequest{Word: "test"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
