package teaser_test

import (
	"testing"

	teaser "github.com/teasers/find_sub_array"
)

func TestSolution(t *testing.T) {
	var scenarios = []struct {
		array    []int
		subarray []int
		output   int
	}{
		// scenario 1
		{
			array:    []int{1, 3, 5},
			subarray: []int{1},
			output:   0,
		},
		// scenario 2
		{
			array:    []int{4, 9, 3, 7, 8},
			subarray: []int{3, 7},
			output:   2,
		},
		// scenario 3
		{
			array:    []int{7, 8, 9},
			subarray: []int{8, 9, 10},
			output:   -1,
		},
		// scenario 4
		{
			array:    []int{4, 9, 3, 7, 8, 3, 7, 1},
			subarray: []int{3, 7},
			output:   5,
		},
		// scenario 5
		{
			array:    []int{8, 3, 7, 7, 1},
			subarray: []int{3, 7},
			output:   1,
		},
		// scenario 6
		{
			array:    []int{3, 7},
			subarray: []int{3, 7},
			output:   0,
		},
		// scenario 7
		{
			array:    []int{3, 7},
			subarray: []int{3, 7},
			output:   0,
		},
		// scenario 8
		{
			array:    []int{3},
			subarray: []int{3, 7},
			output:   -1,
		},
		// scenario 9
		{
			array:    []int{3, 7},
			subarray: []int{},
			output:   -1,
		},
		// scenario 10
		{
			array:    []int{},
			subarray: []int{3, 4, 5},
			output:   -1,
		},
		// scenario 11
		{
			array:    []int{},
			subarray: []int{},
			output:   -1,
		},
		// scenario 12
		{
			array:    []int{1, 1, 1, 1, 1, 1},
			subarray: []int{1},
			output:   5,
		},
		// scenario 13
		{
			array:    []int{3, 7, 3, 7, 3, 7, 3, 7},
			subarray: []int{3, 7},
			output:   6,
		},
		// scenario 14
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			subarray: []int{5, 6, 7, 8, 9},
			output:   4,
		},
		// scenario 15
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			subarray: []int{5, 6, 7, 8, 9, 10},
			output:   4,
		},
		// scenario 16
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			subarray: []int{11, 12, 13, 14, 15, 16, 17},
			output:   10,
		},
		// scenario 17
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			subarray: []int{13, 14, 15, 16, 17, 18, 19},
			output:   12,
		},
		// scenario 18
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			subarray: []int{1, 2, 3},
			output:   0,
		},
		// scenario 19
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			subarray: []int{2, 3, 4},
			output:   1,
		},
		// scenario 20
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			subarray: []int{2, 3, 4, 5},
			output:   1,
		},
	}

	for index, ts := range scenarios {
		output := teaser.Solve(ts.array, ts.subarray)
		if output != ts.output {
			t.Fatalf("scenario %d, unexpected output returned, expected = %d, got = %d", index+1, ts.output, output)
		}
	}
}
