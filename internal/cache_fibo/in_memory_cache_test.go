package cache_fibo

import "testing"

func init() {
}

func TestGetResultFromCacheHit(t *testing.T) {
	cache := New()

	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "1"},
		{3, "2"},
		{10, "55"},
	}

	for _, tt := range tests {
		result, err := cache.GetResultFromCache(tt.input)
		if err != nil {
			t.Errorf("GetResultFromCache(%d) returned an error: %v", tt.input, err)
		}
		if result.String() != tt.expected {
			t.Errorf("GetResultFromCache(%d) = %s, want %s", tt.input, result.String(), tt.expected)
		}
	}
}

func TestGetResultFromCacheMiss(t *testing.T) {
	cache := New()

	// Testing a position that does not exist in the cache
	_, err := cache.GetResultFromCache(100)
	if err != nil {
		t.Errorf("GetResultFromCache(100) returned an unexpected error: %v", err)
	}
}

func TestGetResultFromCacheInvalidValue(t *testing.T) {
	// Assuming there's a possibility of invalid values in the cache
	fibonacciCache[1000] = "not_a_number"
	cache := New()

	_, err := cache.GetResultFromCache(1000)
	if err == nil {
		t.Error("GetResultFromCache(1000) expected an error, got nil")
	}
}
