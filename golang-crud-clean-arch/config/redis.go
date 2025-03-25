package config

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func ConnectRedis() *redis.Client {
	LoadConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_ADDR", "localhost:6379"),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	fmt.Println("Connected to Redis!")
	return client
}
