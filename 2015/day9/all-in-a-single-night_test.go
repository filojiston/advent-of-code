// https://adventofcode.com/2015/day/9
// unit tests of advent of code 2015, day9

package day9

import (
	"testing"
)

func Test_ShouldGetMinMaxPathLengths_PuzzleInput(t *testing.T) {
	minDistance, maxDistance := getMinMaxPathLengths()
	if minDistance != 117 {
		t.Fatalf(`getMinMaxPathLengths() = %d, want %d`, minDistance, 117)
	}

	if maxDistance != 909 {
		t.Fatalf(`getMinMaxPathLengths() = %d, want %d`, maxDistance, 909)
	}
}
