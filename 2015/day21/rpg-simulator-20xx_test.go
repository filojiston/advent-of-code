// https://adventofcode.com/2015/day/21

package day21

import "testing"

func Test_ShouldFindCheapestWinningCombination_PuzzleInput(t *testing.T) {
	result := findCheapestWinningCombination()
	if result.cost != 121 {
		t.Fatalf(`findCheapestWinningCombination() = %v, want %v`, result.cost, 121)
	}
}

func Test_ShouldFindMostExpensiveLossingCombination_PuzzleInput(t *testing.T) {
	result := findMostExpensiveLossingCombination()
	if result.cost != 201 {
		t.Fatalf(`findMostExpensiveLossingCombination() = %v, want %v`, result.cost, 201)
	}
}
