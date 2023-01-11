// https://adventofcode.com/2015/day/13

package day13

import "testing"

func Test_ShouldGetOptimumHappiness_PuzzleInput(t *testing.T) {
	result := getOptimumHappiness()
	if result != 618 {
		t.Fatalf(`getOptimumHappiness(): %v, want: %v`, result, 618)
	}
}

func Test_ShouldGetOptimumHappinessIncludingMe_PuzzleInput(t *testing.T) {
	result := getOptimumHappinessIncludingMe()
	if result != 601 {
		t.Fatalf(`getOptimumHappinessIncludingMe(): %v, want: %v`, result, 601)
	}
}
