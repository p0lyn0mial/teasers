package teaser

// Teaser:
// Given a variable length array of integers, partition them such that the even
// integers precede the odd integers in the array. Your must operate on the array
// in-place, with a constant amount of extra space. The answer should scale
// linearly in time with respect to the size of the array.
//
// Solution:
//
// | 1 | 2 | 3 | 4 | 5 | 6 | 7
//   ^                       ^
//  even                    odd
//  pointer                 pointer
//
// The basic idea is to maintain two pointers.
// The even pointer placed at the beginning of an array.
// The odd pointer place at the end of an array.
//
// The first step is to advance the even pointer as long as an odd number is detected.
// The next step is to decrease the odd pointer as long as an even number is detected.
// Then swap the numbers and start all over again.
//
// Solve solves the teaser
func Solve(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	ep := 0
	op := len(numbers) - 1
	end := false
	for {
		ep, end = moveEP(ep, op, numbers)
		if end {
			return numbers
		}
		op, end = moveOP(ep, op, numbers)
		if end {
			return numbers
		}
		swap(ep, op, numbers)
	}
	return numbers
}

// moveOP decreases odd pointer
func moveOP(ep int, op int, numbers []int) (int, bool) {
	for {
		if isOdd(numbers[op]) {
			op--
		} else {
			return op, end(ep, op)
		}

		if end(ep, op) {
			return op, true
		}
	}
}

// moveEP advances even pointer
func moveEP(ep int, op int, numbers []int) (int, bool) {
	for {
		if isEven(numbers[ep]) {
			ep++
		} else {
			return ep, end(ep, op)
		}

		if end(ep, op) {
			return ep, true
		}
	}
}

// end determines the end of computation
func end(ep int, op int) bool {
	if ep == op {
		return true
	}
	if ep > op {
		return true
	}
	if op < ep {
		return true
	}
	return false
}

// swap swaps the numbers in place
func swap(ep int, op int, numbers []int) {
	tmp := numbers[ep]
	numbers[ep] = numbers[op]
	numbers[op] = tmp
}

// isEven determines whether a number is an even
func isEven(number int) bool {
	return number%2 == 0
}

// isOdd determines whether a number is an odd
func isOdd(number int) bool {
	return number%2 != 0
}
