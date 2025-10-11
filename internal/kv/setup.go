package kv

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var Client *redis.Client

func ConnectRedis() {
	redis_url := os.Getenv("REDIS_URL")
	options, err := redis.ParseURL(redis_url)
	if err != nil {
		log.Fatal().Err(err).Send()
	} else {
		log.Debug().Msg("Redis Connected")
	}

	Client = redis.NewClient(options)

	ping := Client.Ping(context.Background())
	if ping.Err() != nil {
		log.Fatal().Err(ping.Err()).Msg("Redis not working")
	}
}
