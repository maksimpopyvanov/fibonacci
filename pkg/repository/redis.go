package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr     string
	DB       int
	Password string
}

func NewRedisClient(cfg Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		DB:       cfg.DB,
		Password: cfg.Password,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
