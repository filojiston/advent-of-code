// https://adventofcode.com/2015/day/6
// solution of advent of code 2015, day6

package day6

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

type instruction struct {
	command        string
	startX, startY int
	endX, endY     int
}

const gridSize int = 1000

func countLitLights() (litCount uint64) {
	var grid [gridSize][gridSize]int

	lines := util.ReadInputFile("input.txt")
	processInstructionsPartOne(&grid, lines)

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 1 {
				litCount++
			}
		}
	}

	return litCount
}

func calculateTotalBrightness() (totalBrightness uint64) {
	var grid [gridSize][gridSize]int

	lines := util.ReadInputFile("input.txt")
	processInstructionsPartTwo(&grid, lines)

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			totalBrightness += uint64(grid[i][j])
		}
	}

	return totalBrightness
}

func processInstructionsPartOne(grid *[gridSize][gridSize]int, input []string) {
	for _, line := range input {
		instruction := parseLine(line)
		for i := instruction.startX; i <= instruction.endX; i++ {
			for j := instruction.startY; j <= instruction.endY; j++ {
				if strings.EqualFold(instruction.command, "turn on") {
					turnOnPartOne(grid, i, j)
				} else if strings.EqualFold(instruction.command, "turn off") {
					turnOffPartOne(grid, i, j)
				} else {
					togglePartOne(grid, i, j)
				}
			}
		}
	}
}

func processInstructionsPartTwo(grid *[gridSize][gridSize]int, input []string) {
	for _, line := range input {
		instruction := parseLine(line)
		for i := instruction.startX; i <= instruction.endX; i++ {
			for j := instruction.startY; j <= instruction.endY; j++ {
				if strings.EqualFold(instruction.command, "turn on") {
					turnOnPartTwo(grid, i, j)
				} else if strings.EqualFold(instruction.command, "turn off") {
					turnOffPartTwo(grid, i, j)
				} else {
					togglePartTwo(grid, i, j)
				}
			}
		}
	}
}

func parseLine(line string) instruction {
	re := regexp.MustCompile(`(?P<command>\w+( \w+)?) (?P<startPos>\d+,\d+) through (?P<endPos>\d+,\d+)`)
	match := re.FindStringSubmatch(line)
	command := match[1]
	startPos := strings.Split(match[3], ",")
	endPos := strings.Split(match[4], ",")
	startX, startY, endX, endY := parsePositions(startPos, endPos)

	return instruction{command: command, startX: startX, startY: startY, endX: endX, endY: endY}
}

func parsePositions(startPos, endPos []string) (int, int, int, int) {
	startX, _ := strconv.Atoi(startPos[0])
	startY, _ := strconv.Atoi(startPos[1])
	endX, _ := strconv.Atoi(endPos[0])
	endY, _ := strconv.Atoi(endPos[1])

	return startX, startY, endX, endY
}

func turnOnPartOne(grid *[gridSize][gridSize]int, x, y int) {
	if (*grid)[x][y] == 0 {
		(*grid)[x][y] = 1
	}
}

func turnOffPartOne(grid *[gridSize][gridSize]int, x, y int) {
	if (*grid)[x][y] == 1 {
		(*grid)[x][y] = 0
	}
}

func togglePartOne(grid *[gridSize][gridSize]int, x, y int) {
	if (*grid)[x][y] == 0 {
		(*grid)[x][y] = 1
	} else {
		(*grid)[x][y] = 0
	}
}

func turnOnPartTwo(grid *[gridSize][gridSize]int, x, y int) {
	(*grid)[x][y]++
}

func turnOffPartTwo(grid *[gridSize][gridSize]int, x, y int) {
	if (*grid)[x][y] > 0 {
		(*grid)[x][y]--
	}
}

func togglePartTwo(grid *[gridSize][gridSize]int, x, y int) {
	(*grid)[x][y] += 2
}
