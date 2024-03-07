package process_fibo

import (
	"errors"
	"github.com/ali-mohit/simple-fibonacci-server/internal/cache_fibo"
)

type FibonacciProcessHandler interface {
	ProcessFibonacciNumber(req *FibonacciRequest) (*FibonacciResponse, error)
}

type processHandler struct {
	inMemoryCache cache_fibo.InMemoryCache
}

func New(cache cache_fibo.InMemoryCache) FibonacciProcessHandler {
	return &processHandler{
		inMemoryCache: cache,
	}
}

func (s *processHandler) ProcessFibonacciNumber(req *FibonacciRequest) (*FibonacciResponse, error) {
	cacheResult, err := s.inMemoryCache.GetResultFromCacheInt64(req.N)

	if req.N < 0 {
		return nil, errors.New("it is a negative number")
	}

	if err != nil {
		return nil, err
	}

	if cacheResult != nil {
		return &FibonacciResponse{
			N: cacheResult,
		}, nil
	}

	a, _ := s.inMemoryCache.GetResultFromCache(998)
	b, _ := s.inMemoryCache.GetResultFromCache(999)

	for i := int64(1000); i <= req.N; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return &FibonacciResponse{
		N: b,
	}, nil
}
