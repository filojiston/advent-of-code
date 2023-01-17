// https://adventofcode.com/2015/day/20
// solution of advent of code 2015, day20

package day20

const input int = 36000000
const target int = input / 10

func getLowestHouseNumber() int {
	var houses = make([]int, target)
	for i := 1; i < target; i++ {
		for j := i; j < target; j += i {
			houses[j] += i * 10
		}
	}
	result := findIndex(houses)
	return result
}

func getLowestHouseNumberWithLimit() int {
	var houses = make([]int, target)
	for i := 1; i < target; i++ {
		for j := i; j < target && j <= i*50; j += i {
			houses[j] += i * 11
		}
	}
	result := findIndex(houses)
	return result
}

func findIndex(array []int) int {
	for idx, value := range array {
		if value >= input {
			return idx
		}
	}
	return -1
}
