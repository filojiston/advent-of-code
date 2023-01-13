// https://adventofcode.com/2015/day/17

package day17

import "testing"

func Test_ShouldGetCombinationsCount_PuzzleInput(t *testing.T) {
	result := getCombinationsCount()
	if result != 1304 {
		t.Fatalf(`getCombinationsCount(): %v, want: %v`, result, 1304)
	}
}

func Test_ShouldGetCombinationsCountForMinContainerUsage(t *testing.T) {
	result := getCombinationsCountForMinContainerUsage()
	if result != 18 {
		t.Fatalf(`getCombinationsCountForMinContainerUsage(): %v, want: %v`, result, 18)
	}
}
