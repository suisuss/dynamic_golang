package canSum

import(
	//"fmt"
)


func canSum(targetSum int, numbers []int) bool {
	if targetSum < 0 {
		return false
	}
	if targetSum == 0 {
		return true
	}

	for _, number := range numbers {
		if canSum(targetSum - number, numbers) == true {
			return true
		}
	}

	return false
}


