// https://adventofcode.com/2015/day/10
// solution of advent of code 2015, day10

package day10

import (
	"fmt"
)

func lookAndSayXTimes(s string, times int) string {
	for i := 0; i < times; i++ {
		s = lookAndSayPerformant(s)
	}
	return s
}

// my inital thought was use regex to calculate next iteration of look and say
// but go regular expressions does not support backreferences
// this function calculates for 40 iterations around 17 secs on my computer, 50... let's not talk about it
func lookAndSay(s string) string {
	var result string
	var lastChar rune
	var count int = 1
	for index, char := range s {
		if index == 0 {
			lastChar = char
			continue
		}
		if char != lastChar {
			result += fmt.Sprintf("%d%c", count, lastChar)
			lastChar = char
			count = 1
		} else {
			count++
		}
	}

	result += fmt.Sprintf("%d%c", count, lastChar)
	return result
}

// same logic, just by using byte arrays it reduced 17 secs of runtime to 0.07 seconds
// for 50 seconds, it takes about 1.1 secs
func lookAndSayPerformant(s string) string {
	var result []byte
	var lastChar byte
	var count int = 1
	for i := 0; i < len(s); i++ {
		if i == 0 {
			lastChar = s[i]
			continue
		}
		if s[i] != lastChar {
			result = append(result, fmt.Sprintf("%d%c", count, lastChar)...)
			lastChar = s[i]
			count = 1
		} else {
			count++
		}
	}

	result = append(result, fmt.Sprintf("%d%c", count, lastChar)...)
	return string(result)
}
