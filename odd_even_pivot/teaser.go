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
	return numbers
}
