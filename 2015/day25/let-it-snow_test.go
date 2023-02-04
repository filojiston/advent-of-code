// https://adventofcode.com/2015/day/25

package day25

import "testing"

func Test_ShouldFindCode_PuzzleInput(t *testing.T) {
	result := findCode()
	if result != 2650453 {
		t.Fatalf(`findCode() = %d, want %d`, result, 2650453)
	}
}
