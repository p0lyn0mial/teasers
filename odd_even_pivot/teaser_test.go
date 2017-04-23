package teaser_test

import (
	"testing"

	teaser "github.com/teasers/odd_even_pivot"
)

//TODO: add description
func TestStart(t *testing.T) {
	var scenarios = []struct {
		input  []int
		output []int
	}{
		{
			input: []int{1, 2, 3},
		},
	}

	for index, ts := range scenarios {
		ret := teaser.Solve(ts.input)
		if !assertSolution(ret) {
			t.Fatalf("scenario %d, unexpected output returned = %v", index, ret)
		}
	}
}

// assertSolution checks whether even numbers precede odd in the array.
// TODO: implement
func assertSolution(numbers []int) bool {
	if len(numbers) <= 1 {
		return true
	}
	return false
}
