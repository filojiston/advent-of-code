// https://adventofcode.com/2015/day/2
// solution of advent of code 2015, day2

package day2

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getSmallestTwoInts(x, y, z int) (smallestX, smallestY int) {
	dimensions := []int{x, y, z}
	sort.Ints(sort.IntSlice(dimensions))
	return dimensions[0], dimensions[1]
}

func calculateTotalSquareFeetOfWrappingPaper(length, width, height int) (result int) {
	smallestX, smallestY := getSmallestTwoInts(length, width, height)
	areaOfSmallestSize := smallestX * smallestY
	return (2 * length * width) + (2 * width * height) + (2 * height * length) + areaOfSmallestSize
}

func calculateTotalFeetOfRibbon(length, width, height int) (result int) {
	smallestX, smallestY := getSmallestTwoInts(length, width, height)
	requiredFeetOfRibbonForPresent := (smallestX * 2) + (smallestY * 2)
	requiredFeetOfRibbonForBow := length * width * height

	return requiredFeetOfRibbonForPresent + requiredFeetOfRibbonForBow
}

func readInputFile() (lines []string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func getDimensionsFromLine(line string) (length, width, height int) {
	dimensions := strings.Split(line, "x")
	lengthString, widthString, heightString := dimensions[0], dimensions[1], dimensions[2]
	length, _ = strconv.Atoi(lengthString)
	width, _ = strconv.Atoi(widthString)
	height, _ = strconv.Atoi(heightString)

	return length, width, height
}

func calculateTotalSquareFeetOfWrappingPaperFromFile() (result int) {
	lines := readInputFile()
	for _, line := range lines {
		length, width, height := getDimensionsFromLine(line)
		result += calculateTotalSquareFeetOfWrappingPaper(length, width, height)
	}

	return result
}

func calculateTotalFeetOfRibbonFromFile() (result int) {
	lines := readInputFile()
	for _, line := range lines {
		length, width, height := getDimensionsFromLine(line)
		result += calculateTotalFeetOfRibbon(length, width, height)
	}

	return result
}
