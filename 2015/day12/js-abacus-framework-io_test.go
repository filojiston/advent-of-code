// https://adventofcode.com/2015/day/12

package day12

import "testing"

func Test_ShouldSumAllNumbersInInput_PuzzleInput(t *testing.T) {
	result := sumAllNumbersInInput()
	if result != 111754 {
		t.Fatalf(`sumAllNumbersInInput(): %v, want: %v`, result, 111754)
	}
}

func Test_ShouldSumAllNumbersInInputExcludingReds_PuzzleInput(t *testing.T) {
	result := sumAllNumbersInInputExcludingReds()
	if result != 65402 {
		t.Fatalf(`sumAllNumbersInInputExcludingReds(): %v, want: %v`, result, 65402)
	}
}
