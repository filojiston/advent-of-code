// https://adventofcode.com/2015/day/5
// solution of advent of code 2015, day5

package day5

import (
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

func countNiceStrings(isStringNiceFunc func(string) bool) (count uint32) {
	lines := util.ReadInputFile("input.txt")

	for _, line := range lines {
		if isStringNiceFunc(line) {
			count++
		}
	}

	return count
}

func isStringNicePartOne(input string) bool {
	return containsAtLeastThreeVowels(input) && containsSubsequentLetters(input) && !containsForbiddenSubstring(input)
}

func isStringNicePartTwo(input string) bool {
	return containsPairsTwiceWithoutOverlapping(input) && containsALetterRepeatsWithExactlyOneLetterBetween(input)
}

func containsAtLeastThreeVowels(input string) bool {
	vowels := getVowels()
	vowelsCount := 0

	for _, letter := range input {
		if strings.Contains(vowels, string(letter)) {
			vowelsCount++
		}
	}

	return vowelsCount >= 3
}

func containsSubsequentLetters(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			return true
		}
	}

	return false
}

func containsForbiddenSubstring(input string) bool {
	forbiddenSubstrings := getForbiddenSubstrings()
	for _, forbiddenSubstring := range forbiddenSubstrings {
		if strings.Contains(input, forbiddenSubstring) {
			return true
		}
	}
	return false
}

func containsPairsTwiceWithoutOverlapping(input string) bool {
	pairs := make(map[string]uint16)
	var lastPair string
	var lastIndex int

	for i := 0; i+1 < len(input); i++ {
		key := input[i : i+2]
		if pairs[key] == 0 {
			pairs[key]++
			lastPair = key
			lastIndex = i
		} else {
			if strings.EqualFold(lastPair, key) && lastIndex == i-1 {
				continue
			} else {
				pairs[key]++
			}
		}
	}

	for _, value := range pairs {
		if value >= 2 {
			return true
		}
	}

	return false
}

func containsALetterRepeatsWithExactlyOneLetterBetween(input string) bool {
	for i, j := 0, 1; i+2 < len(input) || j+2 < len(input); i, j = i+2, j+2 {
		if (input[i] == input[i+2]) || (input[j] == input[j+2]) {
			return true
		}
	}

	return false
}

func getForbiddenSubstrings() []string {
	return []string{"ab", "cd", "pq", "xy"}
}

func getVowels() string {
	return "aeiou"
}
