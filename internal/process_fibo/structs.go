package process_fibo

import "math/big"

type FibonacciRequest struct {
	N int64 `json:"n" msgpack:"n"`
}

type FibonacciRequestV2 struct {
	N *big.Int `json:"n" msgpack:"n"`
}

type FibonacciResponse struct {
	N *big.Int `json:"n" msgpack:"n"`
}
