package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, errPong := redisClient.Ping(ctx).Result()
	if errPong != nil {
		panic(errPong)
	}

	fmt.Printf("\nRedis connection successfully: {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userID string) {
	errSave := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if errSave != nil {
		panic(errSave)
	}
}

func RetrieveInitUrl(shortUrl string) string {
	result, errGet := storeService.redisClient.Get(ctx, shortUrl).Result()
	if errGet != nil {
		panic(errGet)
	}
	return result
}
