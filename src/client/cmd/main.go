package main

import (
	"context"
	"log"
	"time"

	"github.com/mertess/GoGrpc/gen"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := gen.NewWordsSeenServiceClient(conn)

	for {
		response, err := client.CheckWord(context.Background(), &gen.WordsSeenRequest{Word: "test"})
		if err != nil {
			log.Fatal(err)
		}

		log.Println(response)

		time.Sleep(2 * time.Second)
	}
}
