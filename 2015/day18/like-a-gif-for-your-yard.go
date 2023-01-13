// https://adventofcode.com/2015/day/18
// solution of advent of code 2015, day18

// wait, isn't it the conway's game of life?
// well it isn't anymore... (doing part 2)

package day18

import (
	"github.com/filojiston/advent-of-code/2015/util"
)

func calculateLightsOnAfterXSteps(steps int) int {
	lines := util.ReadInputFile("input.txt")
	lights := parseLights(lines)
	for i := 0; i < steps; i++ {
		lights = nextStep(lights)
	}

	return countOnLights(lights)
}

func calculateLightsOnAfterXStepsWithCornersAlwaysOn(steps int) int {
	lines := util.ReadInputFile("input.txt")
	lights := parseLights(lines)
	setCornersAlwaysOn(&lights)

	for i := 0; i < steps; i++ {
		lights = nextStep(lights)
		setCornersAlwaysOn(&lights)
	}

	return countOnLights(lights)
}

func parseLights(input []string) [][]bool {
	lights := make([][]bool, len(input))
	var idx int
	for _, line := range input {
		lights[idx] = make([]bool, len(line))
		for i, c := range line {
			lights[idx][i] = c == '#'
		}
		idx++
	}
	return lights
}

func nextStep(lights [][]bool) [][]bool {
	newLights := clone(lights)
	for i, row := range lights {
		for j, light := range row {
			onNeighborCount := countOnNeighbors(lights, i, j)
			if light {
				if onNeighborCount != 2 && onNeighborCount != 3 {
					newLights[i][j] = false
				}
			} else {
				if onNeighborCount == 3 {
					newLights[i][j] = true
				}
			}
		}
	}
	return newLights
}

func countOnNeighbors(lights [][]bool, i, j int) int {
	var count int
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if inBounds(0, len(lights), x) && inBounds(0, len(lights), y) && !(x == i && y == j) && lights[x][y] {
				count++
			}
		}
	}
	return count
}

func setCornersAlwaysOn(lights *[][]bool) {
	(*lights)[0][0] = true
	(*lights)[0][len((*lights)[0])-1] = true
	(*lights)[len(*lights)-1][0] = true
	(*lights)[len(*lights)-1][len((*lights)[0])-1] = true
}

func inBounds(lowerBound, upperBound, value int) bool {
	return value >= lowerBound && value < upperBound
}

func clone(lights [][]bool) [][]bool {
	newLights := make([][]bool, len(lights))
	for i, row := range lights {
		newLights[i] = make([]bool, len(row))
		copy(newLights[i], row)
	}
	return newLights
}

func countOnLights(lights [][]bool) int {
	var count int
	for _, row := range lights {
		for _, light := range row {
			if light {
				count++
			}
		}
	}
	return count
}
