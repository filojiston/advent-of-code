// https://adventofcode.com/2015/day/6
// unit tests of advent of code 2015, day6

package day6

import (
	"testing"
)

func Test_ShouldTurnOnLightsPartOne(t *testing.T) {
	var grid [gridSize][gridSize]int

	turnOnPartOne(&grid, 50, 50)
	if grid[50][50] != 1 {
		t.Fatalf(`turnOnPartOne(&grid, 50, 50) = %v, want: %v`, grid[50][50], 1)
	}
}

func Test_ShouldTurnOffLightsPartOne(t *testing.T) {
	var grid [gridSize][gridSize]int

	turnOffPartOne(&grid, 50, 50)
	if grid[50][50] != 0 {
		t.Fatalf(`turnOffPartOne(&grid, 50, 50) = %v, want: %v`, grid[50][50], 0)
	}
}

func Test_ShouldToggleLightsPartOne(t *testing.T) {
	var grid [gridSize][gridSize]int

	turnOnPartOne(&grid, 50, 50)
	if grid[50][50] != 1 {
		t.Fatalf(`turnOnPartOne(&grid, 50, 50) = %v, want: %v`, grid[50][50], 1)
	}

	togglePartOne(&grid, 50, 50)
	if grid[50][50] != 0 {
		t.Fatalf(`togglePartOne(&grid, 50, 50) = %v, want: %v`, grid[50][50], 1)
	}

	turnOffPartOne(&grid, 50, 50)
	if grid[50][50] != 0 {
		t.Fatalf(`turnOffPartOne(&grid, 50, 50) = %v, want: %v`, grid[50][50], 0)
	}

	togglePartOne(&grid, 50, 50)
	if grid[50][50] != 1 {
		t.Fatalf(`togglePartOne(&grid, 50, 50) = %v, want: %v`, grid[50][50], 0)
	}
}

func Test_ShouldTurnOnLightsPartTwo(t *testing.T) {
	var grid [gridSize][gridSize]int

	turnOnPartTwo(&grid, 50, 50)
	if grid[50][50] != 1 {
		t.Fatalf(`turnOnPartTwo(&grid, 50, 50) = %v, want: %v`, grid[50][50], 1)
	}
}

func Test_ShouldTurnOffLightsPartTwo(t *testing.T) {
	var grid [gridSize][gridSize]int

	turnOffPartTwo(&grid, 50, 50)
	if grid[50][50] != 0 {
		t.Fatalf(`turnOffPartTwo(&grid, 50, 50) = %v, want: %v`, grid[50][50], 0)
	}
}

func Test_ShouldToggleLightsPartTwo(t *testing.T) {
	var grid [gridSize][gridSize]int

	togglePartTwo(&grid, 50, 50)
	if grid[50][50] != 2 {
		t.Fatalf(`togglePartOne(&grid, 50, 50) = %v, want: %v`, grid[50][50], 2)
	}
}

func Test_ShouldParseLine(t *testing.T) {
	expected := instruction{command: "turn off", startX: 150, startY: 300, endX: 213, endY: 740}

	result := parseLine("turn off 150,300 through 213,740")
	if result.command != expected.command || result.startX != expected.startX || result.startY != expected.startY || result.endX != expected.endX || result.endY != expected.endY {
		t.Fatalf(`parseLine("turn off 150,300 through 213,740") = %v, want: %v`, result, expected)
	}
}

func Test_ShouldCountLitLights_PuzzleInput(t *testing.T) {
	result := countLitLights()
	if result != 377891 {
		t.Fatalf(`countLitLights() = %v, want: %v`, result, 377891)
	}
}

func Test_ShouldCalculateTotalBrightness_PuzzleInput(t *testing.T) {
	result := calculateTotalBrightness()
	if result != 14110788 {
		t.Fatalf(`calculateTotalBrightness() = %v, want: %v`, result, 14110788)
	}
}
