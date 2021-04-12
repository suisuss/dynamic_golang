package canSum

import (
	"testing"
	// "fmt"
)

func TestCanSum(t *testing.T) {
	input := []struct{
		targetSum int
		numbers []int
	}{
		{55, []int{ 5, 7 ,8, 2, 1 }},
		{0, []int{ 5, 7 ,8, 2, 1 }},
		{1, []int{ 5, 7 ,8, 2}},
	}

	tests := []struct{
		expectedResult bool
	}{
		{ true },
		{ true },
		{ false },
	}

	for i, tt := range tests {
		result := canSum(input[i].targetSum, input[i].numbers)
		if result != tt.expectedResult {
			t.Fatalf("tests[%d] - incorrect result. expected=%t, got=%t",
			i, tt.expectedResult, result)
		}
	}
}