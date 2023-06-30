package wordsSeen

import (
	"context"

	gen "github.com/mertess/GoGrpc/gen"
)

type WordsSeenServiceServer struct {
	gen.UnsafeWordsSeenServiceServer
	words map[string]int32
}

func NewWordsSeenServiceServer() *WordsSeenServiceServer {
	return &WordsSeenServiceServer{words: make(map[string]int32)}
}

func (s *WordsSeenServiceServer) CheckWord(ctx context.Context, req *gen.WordsSeenRequest) (*gen.WordsSeenResponse, error) {
	s.words[req.Word]++
	return &gen.WordsSeenResponse{Count: s.words[req.Word]}, nil
}
