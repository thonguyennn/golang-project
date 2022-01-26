package cache

import (
	"context"
	"log"

	// "encoding/json"

	"time"

	"github.com/go-redis/redis/v8"
	// "github.com/thonguyen/rest-api-golang/entity"
)

type PostCache interface {
	Set(key string, value string)
	Get(key string)
}

func GetRedisDB() redis.UniversalClient {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisCtx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	if err := redisDB.Ping(redisCtx).Err(); err != nil {
		log.Fatal("Cannot connect to redis")
	}
	return redisDB
}
