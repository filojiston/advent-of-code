// https://adventofcode.com/2015/day/14

package day14

import (
	"testing"
)

func Test_ShouldSimulateRace(t *testing.T) {
	result := simulateRace(2503)
	if result != 2655 {
		t.Fatalf(`simulateRace(2503): %v, want: %v`, result, 2655)
	}
}

func Test_ShouldSimulateRaceWithLeadPoints(t *testing.T) {
	result := simulateRaceWithLeadPoints(2503)
	if result != 1059 {
		t.Fatalf(`simulateRaceWithLeadPoints(2503): %v, want: %v`, result, 1059)
	}
}
