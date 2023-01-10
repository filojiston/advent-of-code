// https://adventofcode.com/2015/day/11
// solution of advent of code 2015, day11

package day11

import (
	"strings"
)

func calculateNextValidPassword(currentPassword string) string {
	result := calculateNextPassword(currentPassword)
	for !isValidPassword(result) {
		result = calculateNextPassword(result)
	}
	return result
}

func calculateNextPassword(password string) string {
	var result []byte
	for i := len(password) - 1; i >= 0; i-- {
		if password[i] == 'z' {
			result = append([]byte{'a'}, result...)
			continue
		}
		result = append([]byte{password[i] + 1}, result...)
		for j := i - 1; j >= 0; j-- {
			result = append([]byte{password[j]}, result...)
		}
		return string(result)
	}
	return string(result)
}

func isValidPassword(s string) bool {
	return containsAtLeastTwoDifferentNonOverlappingPairs(s) && containsAtLeastThreeIncreasingLetters(s) && !hasForbiddenLetter(s)
}

func containsAtLeastTwoDifferentNonOverlappingPairs(s string) bool {
	pairs := make(map[string]uint16)
	var lastPair string
	var lastIndex int

	for i := 0; i+1 < len(s); i++ {
		pair := s[i : i+2]
		if pair[0] != pair[1] {
			continue
		}
		if pairs[pair] == 0 {
			pairs[pair]++
			lastPair = pair
			lastIndex = i
		} else {
			if strings.EqualFold(lastPair, pair) && lastIndex == i-1 {
				continue
			} else {
				pairs[pair]++
			}
		}
	}

	return len(pairs) >= 2
}

func containsAtLeastThreeIncreasingLetters(s string) bool {
	lastChar := s[0]
	s = s[1:]
	increaseCounter := 1

	for i := 0; i < len(s); i++ {
		if increaseCounter == 3 {
			return true
		}

		currentChar := s[i]
		if lastChar+1 == currentChar {
			increaseCounter++
		} else {
			increaseCounter = 1
		}

		lastChar = currentChar
	}

	return false
}

func hasForbiddenLetter(s string) bool {
	for _, letter := range getForbiddenLetters() {
		if strings.Contains(s, letter) {
			return true
		}
	}

	return false
}

func getForbiddenLetters() []string {
	return []string{"i", "o", "l"}
}
