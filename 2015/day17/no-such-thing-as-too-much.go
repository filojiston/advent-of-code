// https://adventofcode.com/2015/day/17
// solution of advent of code 2015, day17

// TIL that this is known as knapsack problem, but in this version we don't have values, only weights.

package day17

import (
	"math"
	"strconv"

	"github.com/filojiston/advent-of-code/2015/util"
)

const capacity int = 150

func getCombinationsCount() int {
	lines := util.ReadInputFile("input.txt")
	var items []int
	for _, line := range lines {
		items = append(items, atoi(line))
	}

	var combinations [][]int
	calculateCombinations(items, []int{}, 0, &combinations)

	return len(combinations)
}

func getCombinationsCountForMinContainerUsage() int {
	lines := util.ReadInputFile("input.txt")
	var items []int
	for _, line := range lines {
		items = append(items, atoi(line))
	}

	var combinations [][]int
	calculateCombinations(items, []int{}, 0, &combinations)
	minNumberRequired := getMinNumberOfContainers(combinations)

	var count int
	for _, combination := range combinations {
		if len(combination) == minNumberRequired {
			count++
		}
	}

	return count
}

func getMinNumberOfContainers(combinations [][]int) int {
	result := math.MaxInt
	for _, combination := range combinations {
		if len(combination) < result {
			result = len(combination)
		}
	}
	return result
}

func calculateCombinations(items []int, current []int, index int, combinations *[][]int) {
	currentWeight := sum(current)
	if currentWeight == capacity {
		*combinations = append(*combinations, current)
		return
	}
	if currentWeight > capacity || index == len(items) {
		return
	}

	calculateCombinations(items, current, index+1, combinations)
	calculateCombinations(items, append(current, items[index]), index+1, combinations)
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func sum(arr []int) int {
	var result int
	for _, elem := range arr {
		result += elem
	}
	return result
}
