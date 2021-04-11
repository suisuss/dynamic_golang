package fib

import (
	"testing"
	//"fmt"
)

func TestFib(t *testing.T) {
	input := 10
	tests := []struct {
		expectedResult int
	}{
		{55},
	}

	for i, tt := range tests {
		result := fib(input)
		if result != tt.expectedResult {
			t.Fatalf("tests[%d] - incorrect result. expected=%d, got=%d",
			i, tt.expectedResult, result)
		}
	}
}

func TestFibMemo(t *testing.T) {
	input := 10
	tests := []struct {
		expectedResult int
	}{
		{55},
	}

	for i, tt := range tests {
		result := fibMemo(input, make(map[int]int))
		if result != tt.expectedResult {
			t.Fatalf("tests[%d] - incorrect result. expected=%d, got=%d",
			i, tt.expectedResult, result)
		}
	}
}