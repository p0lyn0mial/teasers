package teaser

import "errors"

// Solve will mix the arrays in accordance with startPosition, firstGap and secondGap
// For example:
//
//     firstArray:   []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
//     secondArray:  []string{"A", "B"},
//     startPostion: 3,
//     firstGap: 0,
//     secondGap: 2,
//     output:       []string{"2", "A", "3", "A", "4", "5", "B", "6", "7", "8", "9"},
//
func Solve(firstArray []string, secondArray []string, startPostion, firstGap, secondGap int) ([]string, error) {
	if startPostion == 0 {
		return nil, errors.New("invalid start position, must be > 0")
	}
	if len(firstArray) == 0 || len(secondArray) == 0 {
		return []string{}, errors.New("firsArray or secondArray is empty")
	}
	//TODO: validate startPosition, len, firstGap, secondGap

	// a helper function that returns elements from the second array
	// be doing so we don't have to maintain additional index and complexity
	getNextItemFuncIndex := 0
	getNextItemFunc := func() string {
		if getNextItemFuncIndex >= len(secondArray) {
			return ""
		}
		nextItem := secondArray[getNextItemFuncIndex]
		getNextItemFuncIndex++
		return nextItem

	}

	// first half from the startPosition to the end
	firstHalfArray, _ := solveFirstHalf(firstArray, getNextItemFunc, startPostion, firstGap, secondGap)

	// second half from the beginning to the startPosition
	getNextItemFuncIndex = 0
	secondHalfArray, _ := solveSecondHalf(firstArray, getNextItemFunc, startPostion, firstGap, secondGap)

	// assemble the final result
	// first append secondHalfArray then firstHalfArray
	// note that returned arrays can be bigger so we need to shirnk them
	// and take as much as needed.
	finalArray := []string{}
	finalArray = append(finalArray, secondHalfArray...)
	finalArray = finalArray[len(finalArray)-startPostion+1:]
	finalArray = append(finalArray, firstArray[startPostion-1])
	finalArray = append(finalArray, firstHalfArray...)
	finalArray = finalArray[:len(firstArray)]

	return finalArray, nil
}

// solveFirstHalf will mix arrays from startPoint to the length of the firstArray
// to make things more easier, items from the second array are delivered by a function getNextItemFromSecond
// this method breaks up the first array(firstArray) into two smaller arrays:
//   the first new array holds the itmes from starting point(startPoint) to the first gap(firstGap) from the first array(firstArray)
//   the second new array holds the remaining items
func solveFirstHalf(firstArray []string, getNextItemFromSecond func() string, startPostion, firstGap, secondGap int) ([]string, error) {
	// the length of the array is shorter than the starting position(startPosition) and the first gap(firstGap)
	// i.e. array = [1,2,3], startPostion = 1, and the gap is = 3
	// it means that we want to place an item on 4th position which does not exists
	if len(firstArray) < startPostion+firstGap {
		return firstArray[startPostion:], nil
	}

	// from startPosition to the first gap
	firstHalfFirstGapArray, firstArrayIndex := mixArrays(firstArray[startPostion:startPostion+firstGap], getNextItemFromSecond, firstGap)

	// from startPosition + the number of items that were taken from the firstArray to from firstHalfFirstGapArray this number is effectively a global index
	// in the first array
	firstArrayNewPosition := startPostion + firstArrayIndex
	firstHalfSecondGapArray, firstArrayIndex := mixArrays(firstArray[firstArrayNewPosition:], getNextItemFromSecond, secondGap)

	finalArray := []string{}
	finalArray = append(finalArray, firstHalfFirstGapArray...)
	finalArray = append(finalArray, firstHalfSecondGapArray...)
	return finalArray, nil
}

// solveSecondHalf will mix arrays from the beginning to the starting point(startPoint) the first array(firstArray)
// to make things more easier, items from the second array are delivered by a function getNextItemFromSecond
// the key idea here is to reverse the second half so that it looks exactly as the first half then:
// this method breaks up the first array(firstArray) into two smaller arrays:
//   the first new array holds the itmes from starting point(startPoint) to the first gap(firstGap) from the first array(firstArray)
//   the second new array holds the remaining items
func solveSecondHalf(firstArray []string, getNextItemFromSecond func() string, startPostion, firstGap, secondGap int) ([]string, error) {
	newStartPosition := 0
	// reverse the first array(firstArray)
	secondHalfReveredArray := reverseArray(firstArray[0 : startPostion-1])

	// the length of the array is shorter than the new starting position(newStartPosition) and the first gap(firstGap)
	// i.e. array = [1,2,3], startPostion = 1, and the gap is = 3
	// it means that we want to place an item on 4th position which does not exists
	if len(secondHalfReveredArray) < newStartPosition+firstGap {
		return firstArray[:startPostion-1], nil
	}

	// from the beginning to the new starting position (newStartPosition)
	secondHalfFirstGapReveredArray, secondHalfReveresedArrayIndex := mixArrays(secondHalfReveredArray[newStartPosition:newStartPosition+firstGap], getNextItemFromSecond, firstGap)

	// from the new starting position (newStartPosition) + the number of items that were taken from the fist array(firstArray) to form secondHalfFirstGapRevertedArray
	// this number is effectively a global index
	secondHalfReversedNewPosition := newStartPosition + secondHalfReveresedArrayIndex
	secondHalfSecondGapReveredArray, _ := mixArrays(secondHalfReveredArray[secondHalfReversedNewPosition:], getNextItemFromSecond, secondGap)

	// reverse the arrays
	secondHalfFirstGapArray := reverseArray(secondHalfFirstGapReveredArray)
	secondHalfSecondGapArray := reverseArray(secondHalfSecondGapReveredArray)

	// assemble the final resutl
	finalArray := []string{}
	finalArray = append(finalArray, secondHalfSecondGapArray...)
	finalArray = append(finalArray, secondHalfFirstGapArray...)
	return finalArray, nil
}

// mixArrays adds items obtained from getNextItemFromSecond function to the array in accordance with the gap
// note that this method creates a new array and appends the elements from the input array
// the length of the new array might be different than the lenght of the input array
func mixArrays(array []string, getNextItemFromSecond func() string, gap int) ([]string, int) {
	newArray := []string{}
	firstArrayIndex := 0
	numberOfItemsFromFirstArray := 0
	hitGap := 0

	for firstArrayIndex, hitGap = 0, 0; firstArrayIndex < len(array); firstArrayIndex, hitGap = firstArrayIndex+1, hitGap+1 {
		nextItem := ""
		if hitGap >= gap {
			nextItem = getNextItemFromSecond()
			if len(nextItem) == 0 {
				nextItem = array[numberOfItemsFromFirstArray]
				numberOfItemsFromFirstArray = numberOfItemsFromFirstArray + 1
			}
			hitGap = 0
		} else {
			nextItem = array[numberOfItemsFromFirstArray]
			numberOfItemsFromFirstArray = numberOfItemsFromFirstArray + 1
		}

		newArray = append(newArray, nextItem)
	}

	if hitGap == gap {
		nextItem := getNextItemFromSecond()
		if len(nextItem) > 0 {
			newArray = append(newArray, nextItem)
		}
	}

	return newArray, numberOfItemsFromFirstArray
}

func reverseArray(array []string) []string {
	reversedArray := make([]string, len(array))
	for i, j := len(array)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversedArray[j] = array[i]
	}
	return reversedArray
}
