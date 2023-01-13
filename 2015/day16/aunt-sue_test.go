// https://adventofcode.com/2015/day/16

package day16

import "testing"

func Test_ShouldFindCorrectAunt_PuzzleInput(t *testing.T) {
	result := findCorrectAunt(calculateLikelinessesPart1)
	if result.number != 103 {
		t.Fatalf(`findCorrectAunts(): %v, want: %v`, result.number, 103)
	}

	result = findCorrectAunt(calculateLikelinessesPart2)
	if result.number != 405 {
		t.Fatalf(`findCorrectAunts(): %v, want: %v`, result.number, 405)
	}
}
