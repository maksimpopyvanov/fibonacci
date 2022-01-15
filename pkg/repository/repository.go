package repository

import (
	"context"
	"fibonacci"
	"math/big"
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

func (r *Repository) GetNumberFromCache(index string) (string, error) {
	number, err := r.rdb.Get(r.ctx, index).Result()
	if err == redis.Nil {
		return "0", nil
	} else if err != nil {
		return "0", err
	}
	return number, nil
}

func (r *Repository) SetNumberToCache(index string, number string) {
	r.rdb.Set(r.ctx, index, number, 0)
}

func (r *Repository) GetSequence(input *fibonacci.Input) map[int64]string {
	output := make(map[int64]string)
	for index := input.End; index >= input.Start; index-- {
		number := r.GetNumberFibonacci(index)
		if number == "" {
			return nil
		}
		output[index] = number
	}
	return output
}

func (r *Repository) GetNumberFibonacci(index int64) string {
	if index == 0 {
		return "0"
	} else if index == 1 {
		return "1"
	}

	if number, err := r.GetNumberFromCache(strconv.FormatInt(index, 10)); number == "0" && err == nil {
		penultimateNumber, success := new(big.Int).SetString(r.GetNumberFibonacci(index-2), 10)
		if success {
			lastNumber, success := new(big.Int).SetString(r.GetNumberFibonacci(index-1), 10)
			if success {
				number = new(big.Int).Add(penultimateNumber, lastNumber).String()
				r.SetNumberToCache(strconv.FormatInt(index, 10), number)
				return number
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else if number == "0" && err != nil {
		penultimateNumber, success := new(big.Int).SetString(r.GetNumberFibonacci(index-2), 10)
		if success {
			lastNumber, success := new(big.Int).SetString(r.GetNumberFibonacci(index-1), 10)
			if success {
				number = new(big.Int).Add(penultimateNumber, lastNumber).String()
				return number
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return number
	}
}
