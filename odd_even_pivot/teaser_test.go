package teaser_test

import (
	"fmt"
	"testing"

	teaser "github.com/teasers/odd_even_pivot"
)

const (
	oddMode  = "om"
	evenMode = "em"
)

// TestSoulution checks if the solution for a given input is correct
func TestSolution(t *testing.T) {
	var scenarios = []struct {
		input []int
	}{
		{
			input: []int{1, 2, 3},
		},
	}

	for index, ts := range scenarios {
		ret := teaser.Solve(ts.input)
		if err := assertSolution(ret, t); err != nil {
			t.Fatalf("scenario %d, unexpected error returned = %v, input was = %v", index, err, ts.input)
		}
	}
}

// TestAssertSolution checks if assertSolution function works correctly
func TestAssertSolution(t *testing.T) {
	var scenarios = []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 2, 3},
			output: false,
		},
		{
			input:  []int{1, 2},
			output: false,
		},
		{
			input:  []int{1, 3, 5, 7, 2},
			output: false,
		},
		{
			input:  []int{2, 1},
			output: true,
		},
		{
			input:  []int{2, 4, 6, 1, 3, 5, 7},
			output: true,
		},
		{
			input:  []int{-2, -4, 0, 6, 1, 3, 5, 7},
			output: true,
		},
	}

	for index, ts := range scenarios {
		err := assertSolution(ts.input, t)
		if ts.output && err != nil {
			t.Fatalf("scenario %d, unexpected error returned = %v, input was = %v", index, err, ts.input)
		}
	}
}

// assertSolution checks whether even numbers precede odd in an array.
func assertSolution(numbers []int, t *testing.T) error {
	if len(numbers) <= 1 {
		return nil
	}

	prevMode := determineMode(numbers[0])
	for index, number := range numbers {
		nextMode := determineMode(number)
		if !checkOrder(prevMode, nextMode) {
			return fmt.Errorf("wrong number = %d at the position = %d, previous was = %d", number, index+1, numbers[index-1])
		}

		prevMode = nextMode
	}

	return nil
}

// determineMode determines a mode based on a number
func determineMode(number int) string {
	if number%2 == 0 {
		return evenMode
	}
	return oddMode
}

// checkOrder checks the order of the numbers
// returns true if the order is correct
func checkOrder(prev, next string) bool {
	switch prev {
	case oddMode:
		return next == oddMode
	case evenMode:
		return next == oddMode || next == evenMode
	default:
		return false
	}
}
