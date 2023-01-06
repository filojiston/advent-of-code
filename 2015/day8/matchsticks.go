// https://adventofcode.com/2015/day/8
// solution of advent of code 2015, day8

package day8

import (
	"strconv"

	"github.com/filojiston/advent-of-code/2015/util"
)

func totalCharsOfStringCodeMinusTotalCharsInMemory() int {
	lines := util.ReadInputFile("input.txt")
	totalCharsInMemory := 0
	totalCharsOfStringCode := 0

	for _, line := range lines {
		totalCharsInMemory += totalNumberOfCharactersInMemory(line)
		totalCharsOfStringCode += totalNumberOfCharactersOfStringCode(line)
	}

	return totalCharsOfStringCode - totalCharsInMemory
}

func totalCharsOfEncodedMinusTotalCharsOfStringCode() int {
	lines := util.ReadInputFile("input.txt")
	totalCharsOfEncoded := 0
	totalCharsOfStringCode := 0

	for _, line := range lines {
		totalCharsOfEncoded += totalNumberOfCharactersOfEncodedString(line)
		totalCharsOfStringCode += totalNumberOfCharactersOfStringCode(line)
	}

	return totalCharsOfEncoded - totalCharsOfStringCode
}

func totalNumberOfCharactersInMemory(str string) int {
	str, _ = strconv.Unquote(str)
	return len(str)
}

func totalNumberOfCharactersOfStringCode(str string) int {
	return len(str)
}

func totalNumberOfCharactersOfEncodedString(str string) int {
	str = strconv.Quote(str)
	return len(str)
}
