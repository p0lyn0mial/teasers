package teaser

// Teaser:
// Implement a method that given two arrays as parameters will find the starting index where the second parameter occurs as a sub-array in the array given as the first parameter.
// If given sub-array (second parameter) occurs more than once, then the method should return the starting index of the last occurrence
// Your implementations should return -1 if the sub-array cannot be found.
// Your implementation must implement the FindArray interface.
// For extra points: implement it in an efficient way for large input arrays.
//
// Sample Input:
// [4,9,3,7,8] and [3,7] should return 2.
// [1,3,5] and [1] should return 0.
// [7,8,9] and [8,9,10] should return -1.
// [4,9,3,7,8,3,7,1] and [3,7] should return 5.
func Solve(array []int, subArray []int) int {
	if len(subArray) > len(array) {
		return -1
	}
	if len(array) == 0 || len(subArray) == 0 {
		return -1
	}

	subArrayLen := len(subArray)
	arrayLen := len(array)
	skip := subArrayLen

	for i := arrayLen - skip; i >= 0; i = i - skip {
		// if there is a match on the first index,
		// check the remaining positions
		if array[i] == subArray[0] {
			foundSolution := true
			for j := 1; j < subArrayLen; j++ {
				if array[i+j] != subArray[j] {
					foundSolution = false
					break
				}
			}
			if foundSolution {
				return i
			}
			skip = subArrayLen
			continue
		}

		// the idea is to check whether numbers from the input array
		// overlap with the numbers from the sub-array on any positions
		// in that case we need to recalculate our skip.
		newSkip := 0
		potentialSolution := false
		nextArrayPosition := 0
		for j := 0; j < subArrayLen; j++ {
			if array[i+nextArrayPosition] != subArray[j] {
				if potentialSolution {
					potentialSolution = false
					break
				}
				newSkip = newSkip + 1
			} else {
				potentialSolution = true
				nextArrayPosition = nextArrayPosition + 1
			}
		}

		if potentialSolution {
			skip = newSkip
		} else {
			skip = subArrayLen
		}
	}

	return -1
}
