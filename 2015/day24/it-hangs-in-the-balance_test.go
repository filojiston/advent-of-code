// https://adventofcode.com/2015/day/24

package day24

import "testing"

func Test_ShouldFindMinimumQuantumEntanglement_PuzzleInput(t *testing.T) {
	result := findMinimumQuantumEntanglement(3)
	if result != 11846773891 {
		t.Fatalf(`findMinimumQuantumEntanglement(3) = %d, want %d`, result, 11846773891)
	}

	result = findMinimumQuantumEntanglement(4)
	if result != 80393059 {
		t.Fatalf(`findMinimumQuantumEntanglement(4) = %d, want %d`, result, 80393059)
	}

}
