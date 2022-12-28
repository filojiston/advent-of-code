package day2

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetSmallestTwoInts(x, y, z int) (smallestX, smallestY int) {
	dimensions := []int{x, y, z}
	sort.Ints(sort.IntSlice(dimensions))
	return dimensions[0], dimensions[1]
}

func CalculateTotalSquareFeetOfWrappingPaper(length, width, height int) (result int) {
	smallestX, smallestY := GetSmallestTwoInts(length, width, height)
	areaOfSmallestSize := smallestX * smallestY
	return (2 * length * width) + (2 * width * height) + (2 * height * length) + areaOfSmallestSize
}

func CalculateTotalFeetOfRibbon(length, width, height int) (result int) {
	smallestX, smallestY := GetSmallestTwoInts(length, width, height)
	requiredFeetOfRibbonForPresent := (smallestX * 2) + (smallestY * 2)
	requiredFeetOfRibbonForBow := length * width * height

	return requiredFeetOfRibbonForPresent + requiredFeetOfRibbonForBow
}

func ReadInputFile() (lines []string) {
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

func GetDimensionsFromLine(line string) (length, width, height int) {
	dimensions := strings.Split(line, "x")
	lengthString, widthString, heightString := dimensions[0], dimensions[1], dimensions[2]
	length, _ = strconv.Atoi(lengthString)
	width, _ = strconv.Atoi(widthString)
	height, _ = strconv.Atoi(heightString)

	return length, width, height
}

func CalculateTotalSquareFeetOfWrappingPaperFromFile() (result int) {
	lines := ReadInputFile()
	for _, line := range lines {
		length, width, height := GetDimensionsFromLine(line)
		result += CalculateTotalSquareFeetOfWrappingPaper(length, width, height)
	}

	return result
}

func CalculateTotalFeetOfRibbonFromFile() (result int) {
	lines := ReadInputFile()
	for _, line := range lines {
		length, width, height := GetDimensionsFromLine(line)
		result += CalculateTotalFeetOfRibbon(length, width, height)
	}

	return result
}
