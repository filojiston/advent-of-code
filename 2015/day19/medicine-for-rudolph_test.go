// https://adventofcode.com/2015/day/19

package day19

import (
	"testing"
)

func Test_ShouldCalculateAllDistinctMolecules_PuzzleInput(t *testing.T) {
	result := calculateAllDistinctMolecules()
	if len(result) != 518 {
		t.Fatalf(`calculateAllDistinctMolecules() = %v, want %v`, len(result), 518)
	}
}

func Test_ShouldCalculateFewestStepsForCreatingMedicine_PuzzleInput(t *testing.T) {
	result := calculateFewestStepsForCreatingMedicine()
	if result != 200 {
		t.Fatalf(`calculateFewestStepsForCreatingMedicine() = %v, want %v`, result, 200)
	}
}
