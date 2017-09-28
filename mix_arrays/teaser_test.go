package teaser_test

import (
	"reflect"
	"testing"

	teaser "github.com/teasers/mix_arrays"
)

//
func TestSolution(t *testing.T) {
	var scenarios = []struct {
		firstArray   []string
		secondArray  []string
		output       []string
		startPostion int
		firstGap     int
		secondGap    int
	}{
		// scenario 0
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			secondArray:  []string{"A", "B"},
			output:       []string{"2", "3", "A", "4", "5", "6", "A", "7", "8", "B"},
			startPostion: 5, firstGap: 1, secondGap: 2,
		},
		// scenario 1
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"2", "3", "A", "4", "5", "6", "A", "7", "8", "B", "9"},
			startPostion: 5, firstGap: 1, secondGap: 2,
		},
		// scenario 2
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"A", "2", "3", "4", "A", "5", "6", "B", "7", "8", "9"},
			startPostion: 3, firstGap: 1, secondGap: 2,
		},
		// scenario 3
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"2", "A", "3", "A", "4", "5", "B", "6", "7", "8", "9"},
			startPostion: 3, firstGap: 0, secondGap: 2,
		},
		// scenario 4
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"1", "2", "3", "4", "5", "6", "A", "7", "8", "B", "9"},
			startPostion: 3, firstGap: 3, secondGap: 2,
		},
		//scenario 5
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"1", "2", "3", "A", "4", "5", "B", "6", "7", "8", "9"},
			startPostion: 2, firstGap: 1, secondGap: 2,
		},
		// scenario 6
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"1", "2", "A", "3", "4", "B", "5", "6", "7", "8", "9"},
			startPostion: 1, firstGap: 1, secondGap: 2,
		},
		// scenario 7
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"3", "B", "4", "5", "A", "6", "7", "8", "A", "9", "10"},
			startPostion: 7, firstGap: 1, secondGap: 2,
		},
		// scenario 8
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"3", "B", "4", "5", "A", "6", "7", "8", "9", "10", "11"},
			startPostion: 9, firstGap: 3, secondGap: 2,
		},
		// scenario 9
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B"},
			output:       []string{"1", "A", "B", "2", "3", "4", "5", "6", "7", "8", "9"},
			startPostion: 1, firstGap: 0, secondGap: 0,
		},
		// scenario 10
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B", "B", "B", "B", "B", "B"},
			output:       []string{"1", "A", "B", "B", "B", "B", "B", "B", "2", "3", "4"},
			startPostion: 1, firstGap: 0, secondGap: 0,
		},
		// scenario 11
		{
			firstArray:   []string{"1", "2"},
			secondArray:  []string{"A", "B", "B", "B", "B", "B", "B"},
			output:       []string{"1", "A"},
			startPostion: 1, firstGap: 0, secondGap: 0,
		},
		// scenario 12
		{
			firstArray:   []string{"1"},
			secondArray:  []string{"A", "B", "B", "B", "B", "B", "B"},
			output:       []string{"1"},
			startPostion: 1, firstGap: 0, secondGap: 0,
		},
		// scenario 13
		{
			firstArray:   []string{"1", "2", "3"},
			secondArray:  []string{"A", "B", "B", "B", "B", "B", "B"},
			output:       []string{"1", "2", "3"},
			startPostion: 3, firstGap: 3, secondGap: 0,
		},
		// scenario 14
		{
			firstArray:   []string{"1", "2", "3"},
			secondArray:  []string{"A", "B", "B", "B", "B", "B", "B"},
			output:       []string{"1", "2", "3"},
			startPostion: 3, firstGap: 3, secondGap: 0,
		},
		// scenario 15
		{
			firstArray:   []string{"1", "2", "3"},
			secondArray:  []string{"A", "B", "B", "B", "B", "B", "B"},
			output:       []string{"A", "2", "3"},
			startPostion: 3, firstGap: 1, secondGap: 3,
		},
		// scenario 16
		{
			firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
			secondArray:  []string{"A", "B", "B", "B", "B", "B", "B"},
			output:       []string{"B", "B", "B", "A", "5", "A", "B", "B", "B", "B", "B"},
			startPostion: 5, firstGap: 0, secondGap: 0,
		},
	}
	for index, ts := range scenarios {
		output, err := teaser.Solve(ts.firstArray, ts.secondArray, ts.startPostion, ts.firstGap, ts.secondGap)
		if err != nil {
			t.Fatalf("scenario %d, unexpeced error ocurred = %v", err)
		}
		if !reflect.DeepEqual(output, ts.output) {
			t.Fatalf("scenario %d, incorrect output returned = %v, expected = %v", index, output, ts.output)
		}
	}

}
