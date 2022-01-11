package service

import (
	"fibonacci/pkg/repository"
	"strconv"
)

type Service struct {
	redis *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{redis: repos}
}

func (s *Service) GetNumberFibonacci(index int) uint64 {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}

	if number, err := s.redis.CheckNumber(strconv.Itoa(index)); number == 0 && err == nil {
		number = s.GetNumberFibonacci(index-2) + s.GetNumberFibonacci(index-1)
		s.redis.SetNumber(strconv.Itoa(index), strconv.FormatUint(number, 10))
		return number
	} else if number == 0 && err != nil {
		number = s.GetNumberFibonacci(index-2) + s.GetNumberFibonacci(index-1)
		return number
	} else {
		return number
	}
}
