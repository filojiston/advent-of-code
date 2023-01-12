// https://adventofcode.com/2015/day/15

package day15

import "testing"

func Test_ShouldGetOptimalCookieScore_PuzzleInput(t *testing.T) {
	result := getOptimalCookieScore()
	if result != 222870 {
		t.Fatalf(`getOptimalCookieScore(): %v, want: %v`, result, 222870)
	}
}

func Test_ShouldGetOptimalCookieScoreWithCalories_PuzzleInput(t *testing.T) {
	result := getOptimalCookieScoreWithCalories(500)
	if result != 117936 {
		t.Fatalf(`getOptimalCookieScoreWithCalories(500): %v, want: %v`, result, 117936)
	}
}
