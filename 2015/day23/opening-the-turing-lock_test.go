// https://adventofcode.com/2015/day/23

package day23

import "testing"

func Test_ShouldOpenTuringLock_PuzzleInput(t *testing.T) {
	_, registerBValue := openTuringLock(0, 0)
	if registerBValue != 255 {
		t.Fatalf(`openTuringLock(0, 0): %v, want: %v`, registerBValue, 255)
	}

	_, registerBValue = openTuringLock(1, 0)
	if registerBValue != 334 {
		t.Fatalf(`openTuringLock(1, 0): %v, want: %v`, registerBValue, 334)
	}
}
