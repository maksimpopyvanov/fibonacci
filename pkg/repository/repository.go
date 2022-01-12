package repository

import (
	"context"
	"fibonacci"
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

func (r *Repository) GetSequence(input *fibonacci.Input) map[int64]uint64 {
	output := make(map[int64]uint64)
	for index := input.End; index >= input.Start; index-- {
		number := r.GetNumberFibonacci(index)
		output[index] = number
	}
	return output
}

func (r *Repository) GetNumberFibonacci(index int64) uint64 {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}

	if number, err := r.CheckNumber(strconv.FormatInt(index, 10)); number == 0 && err == nil {
		number = r.GetNumberFibonacci(index-2) + r.GetNumberFibonacci(index-1)
		r.SetNumber(strconv.FormatInt(index, 10), strconv.FormatUint(number, 10))
		return number
	} else if number == 0 && err != nil {
		number = r.GetNumberFibonacci(index-2) + r.GetNumberFibonacci(index-1)
		return number
	} else {
		return number
	}
}
