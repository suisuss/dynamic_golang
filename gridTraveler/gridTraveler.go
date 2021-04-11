package gridTraveler

import (
	"fmt"
)

func reverseString(s string) string {
	rns := []rune(s) // convert to array of runes
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

			// swap the letters of the string,
			// like first with last and so on.
			rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func gridTraveler(rows int, columns int) int {
	if rows == 0 || columns == 0 {
		return 0
	}
	if rows == 1 && columns == 1 {
		return 1
	}
	return gridTraveler(rows - 1, columns) + gridTraveler(rows, columns - 1)
}

func gridTravelerMemo(rows int, columns int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", rows, columns)
	key_reverse := reverseString(key)
	if val, ok := memo[key]; ok {
		return val
	}
	if val, ok := memo[key_reverse]; ok {
		return val
	}
	if rows == 0 || columns == 0 {
		return 0
	}
	if rows == 1 && columns == 1 {
		return 1
	}
	memo[key] = gridTraveler(rows - 1, columns) + gridTraveler(rows, columns - 1)
	memo[key_reverse] = gridTraveler(rows - 1, columns) + gridTraveler(rows, columns - 1)
	return memo[key]
}