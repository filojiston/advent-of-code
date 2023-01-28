// https://adventofcode.com/2015/day/24
// solution of advent of code 2015, day24

package day24

import (
	"math"
	"strconv"

	"github.com/filojiston/advent-of-code/2015/util"
)

func findMinimumQuantumEntanglement(groupSize int) uint64 {
	numbers := readNumbersFromFile()
	target := sum(numbers) / groupSize

	i := 1
	var result []uint64
	for {
		result = myMap(filter(subsetsWithLength(numbers, i), target))
		if len(result) > 0 {
			break
		}
		i++
	}
	return min(result)
}

func readNumbersFromFile() []int {
	input := util.ReadInputFile("input.txt")
	var numbers []int
	for _, line := range input {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	return numbers
}

func subsetsWithLength(numbers []int, length int) [][]int {
	if len(numbers) == 0 || length == 0 {
		return [][]int{{}}
	}

	subs := subsetsWithLength(numbers[1:], length)
	res := make([][]int, len(subs))
	copy(res, subs)

	subs = subsetsWithLength(numbers[1:], length-1)
	for _, s := range subs {
		res = append(res, append([]int{numbers[0]}, s...))
	}

	return res
}

func min(arr []uint64) uint64 {
	var minElem uint64 = math.MaxInt64
	for _, item := range arr {
		if item < minElem {
			minElem = item
		}
	}
	return minElem
}

func filter(arr [][]int, target int) [][]int {
	var result [][]int
	for _, item := range arr {
		if sum(item) == target {
			result = append(result, item)
		}
	}
	return result
}

func myMap(arr [][]int) []uint64 {
	var result []uint64
	for _, item := range arr {
		result = append(result, product(item))
	}
	return result
}

func sum(arr []int) int {
	var result int
	for _, item := range arr {
		result += item
	}
	return result
}

func product(arr []int) uint64 {
	var result uint64 = 1
	for _, item := range arr {
		result *= uint64(item)
	}
	return result
}
