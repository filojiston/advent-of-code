// https://adventofcode.com/2015/day/8
// unit tests of advent of code 2015, day8

package day8

import (
	"testing"
)

func Test_ShouldGetTotalCharsOfStringCodeMinusTotalCharsInMemory_PuzzleInput(t *testing.T) {
	result := totalCharsOfStringCodeMinusTotalCharsInMemory()
	if result != 1371 {
		t.Fatalf(`("totalCharsOfStringCodeMinusTotalCharsInMemory(): %v, want: %v`, result, 1371)
	}
}

func Test_ShouldGetTotalCharsOfEncodedMinusTotalCharsOfStringCode_PuzzleInput(t *testing.T) {
	result := totalCharsOfEncodedMinusTotalCharsOfStringCode()
	if result != 2117 {
		t.Fatalf(`("totalCharsOfEncodedMinusTotalCharsOfStringCode(): %v, want: %v`, result, 2117)
	}
}
