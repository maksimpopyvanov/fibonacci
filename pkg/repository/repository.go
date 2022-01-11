package repository

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRepository(redis *redis.Client) *Repository {
	return &Repository{
		rdb: redis,
		ctx: context.Background(),
	}
}

func (r *Repository) CheckNumber(index string) (uint64, error) {
	number, err := r.rdb.Get(r.ctx, index).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	} else if result, err := strconv.ParseUint(number, 10, 64); err == nil {
		return result, nil
	} else {
		return 0, err
	}
}

func (r *Repository) SetNumber(index string, number string) {
	r.rdb.Set(r.ctx, index, number, 0)
}
