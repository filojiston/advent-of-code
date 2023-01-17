// https://adventofcode.com/2015/day/20

package day20

import (
	"testing"
)

func Test_ShouldGetLowestHouseNumber_PuzzleInput(t *testing.T) {
	result := getLowestHouseNumber()
	if result != 831600 {
		t.Fatalf(`getLowestHouseNumber(): %v, want: %v`, result, 831600)
	}

}

func Test_ShouldGetLowestHouseNumberWithLimit_PuzzleInput(t *testing.T) {
	result := getLowestHouseNumberWithLimit()
	if result != 884520 {
		t.Fatalf(`getLowestHouseNumberWithLimit(): %v, want: %v`, result, 884520)
	}
}
