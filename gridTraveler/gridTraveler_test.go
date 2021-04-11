package gridTraveler

import (
	"testing"
)

func TestGridTraveler(t *testing.T) {
	input := [][2]int{ 
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
		{3, 3},
		{3, 5},
	}

	tests := []struct{
		expectedResult int
	}{
		{0},
		{0},
		{0},
		{1},
		{6},
		{15},
	}

	for i, tt := range tests {
		result := gridTraveler(input[i][0], input[i][1])
		if result != tt.expectedResult {
			t.Fatalf("tests[%d] - incorrect result. expected=%d, got=%d",
			i, tt.expectedResult, result)
		}
	}
}

func TestGridTravelerMemo(t *testing.T) {
	input := [][2]int{ 
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
		{3, 3},
		{3, 5},
	}

	tests := []struct{
		expectedResult int
	}{
		{0},
		{0},
		{0},
		{1},
		{6},
		{15},
	}

	for i, tt := range tests {
		result := gridTravelerMemo(input[i][0], input[i][1], make(map[string]int))
		if result != tt.expectedResult {
			t.Fatalf("tests[%d] - incorrect result. expected=%d, got=%d",
			i, tt.expectedResult, result)
		}
	}
}