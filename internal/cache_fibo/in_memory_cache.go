package cache_fibo

import (
	"errors"
	"github.com/sirupsen/logrus"
	"math/big"
)

type InMemoryCache interface {
	GetResultFromCache(i int) (*big.Int, error)
	GetResultFromCacheInt64(i int64) (*big.Int, error)
}

type impl struct {
}

func New() InMemoryCache {
	return &impl{}
}

func (s *impl) GetResultFromCache(i int) (*big.Int, error) {
	r := int64(i)

	return s.GetResultFromCacheInt64(r)
}

func (s *impl) GetResultFromCacheInt64(i int64) (*big.Int, error) {
	if _, ok := fibonacciCache[i]; !ok {
		return nil, nil
	}
	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(fibonacciCache[i], 10)

	if !ok {
		logrus.Fatal("Failed to convert string to big.Int")
		return nil, errors.New("failed to convert string to big.Int")
	}

	return bigInt, nil
}
