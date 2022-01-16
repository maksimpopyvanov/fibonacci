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

func (r *Repository) GetSequence(input *fibonacci.Input) map[int64]string {
	sequence := make(map[int64]string)
	for index := input.End; index >= input.Start; index-- {
		number := r.GetNumberFibonacci(index)
		if number == "" {
			return nil
		}
		sequence[index] = number
	}
	return sequence
}

func (r *Repository) GetNumberFibonacci(index int64) string {
	if index == 0 {
		return "0"
	} else if index == 1 {
		return "1"
	}

	if number, err := r.GetNumberFromCache(strconv.FormatInt(index, 10)); number == "0" && err == nil {
		//Достучались до кеша, числа с данным индексом в нем нет => нужно вычислить значение и добавить в кеш
		number = r.AddPreviousNumbers(index)
		if number != "" {
			//Нет ошибки в сложении 2х предыдущих чисел
			r.rdb.Set(r.ctx, strconv.FormatInt(index, 10), number, 0)
			return number
		}
		//Ошибка в сложении двух предыдущих чисел
		return number
	} else if number == "0" && err != nil {
		//Ошибка в получении значения из кеша
		number = r.AddPreviousNumbers(index)
		return number
	} else {
		//Возвращаем значение из кеша
		return number
	}
}

func (r *Repository) AddPreviousNumbers(index int64) string {
	lastNumber, success := new(big.Int).SetString(r.GetNumberFibonacci(index-1), 10)
	if success {
		penultimateNumber, success := new(big.Int).SetString(r.GetNumberFibonacci(index-2), 10)
		if success {
			number := new(big.Int).Add(penultimateNumber, lastNumber).String()
			return number
		}
		return ""
	}
	return ""
}
