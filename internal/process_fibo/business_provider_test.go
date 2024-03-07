package process_fibo

import (
	"github.com/ali-mohit/simple-fibonacci-server/internal/cache_fibo"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessFibonacciNumber(t *testing.T) {
	// Create a new mock cache instance
	memoryCache := cache_fibo.New()
	handler := New(memoryCache)

	// Define test cases
	tests := []struct {
		name      string
		n         int64
		mockSetup func()
		want      *big.Int
		wantErr   bool
	}{
		{
			name:      "negative number",
			n:         -1,
			mockSetup: func() {},
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "cached number",
			n:         10,
			mockSetup: func() {},
			want:      big.NewInt(55),
			wantErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup() // Setup mock behavior for this test case

			req := &FibonacciRequest{N: tc.n}
			resp, err := handler.ProcessFibonacciNumber(req)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, resp.N)
			}
		})
	}
}
