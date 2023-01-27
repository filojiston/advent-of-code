// https://adventofcode.com/2015/day/22

package day22

import "testing"

func Test_ShouldCalculateLeastManaForWin_EasyMode_PuzzleInput(t *testing.T) {
	result := calculateLeastManaForWin("easy")
	if result != 900 {
		t.Fatalf(`calculateLeastManaForWin("easy") = %v; want %v`, result, 900)
	}
}

func Test_ShouldCalculateLeastManaForWin_HardMode_PuzzleInput(t *testing.T) {
	result := calculateLeastManaForWin("hard")
	if result != 1216 {
		t.Fatalf(`calculateLeastManaForWin("hard") = %v; want %v`, result, 1216)
	}
}
