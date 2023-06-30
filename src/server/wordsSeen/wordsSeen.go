package wordsSeen

import (
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis"
	gen "github.com/mertess/GoGrpc/gen"
	rc "github.com/mertess/GoGrpc/src/server/redisClient"
)

type WordsSeenServiceServer struct {
	gen.UnsafeWordsSeenServiceServer
	words   map[string]int32
	redisDb *rc.RedisClient
}

func NewWordsSeenServiceServer() *WordsSeenServiceServer {
	return &WordsSeenServiceServer{
		words:   make(map[string]int32),
		redisDb: rc.NewRedisClient("localhost:6379"),
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (s *WordsSeenServiceServer) CheckWord(ctx context.Context, req *gen.WordsSeenRequest) (*gen.WordsSeenResponse, error) {
	s.words[req.Word]++

	log.Println("Current state: " + strconv.Itoa(int(s.words[req.Word])))

	cachedValue, err := s.redisDb.Get(req.Word)
	if err == redis.Nil {
		err := s.redisDb.Set(req.Word, s.words[req.Word])
		handleError(err)
		cachedValue, err = s.redisDb.Get(req.Word)
		handleError(err)
	}

	log.Println("Cached value: " + cachedValue)

	cachedValueInt, err := strconv.Atoi(cachedValue)
	handleError(err)

	return &gen.WordsSeenResponse{Count: int32(cachedValueInt)}, nil
}
