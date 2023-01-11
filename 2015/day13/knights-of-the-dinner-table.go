// https://adventofcode.com/2015/day/13
// solution of advent of code 2015, day13

package day13

import (
	"fmt"
	"math"
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

type table struct {
	people    []string
	happiness int
}

type instructions map[string]map[string]int
type permutations []*table

func getOptimumHappiness() int {
	instructions := parseInstructions()
	table := newTable(*instructions)
	permutations := calculatePermutations(table, instructions)

	maxHappiness := math.MinInt
	for _, currentTable := range *permutations {
		if currentTable.happiness > maxHappiness {
			maxHappiness = currentTable.happiness
		}
	}

	return maxHappiness
}

func getOptimumHappinessIncludingMe() int {
	instructions := parseInstructions()
	table := newTable(*instructions)
	addMe(instructions, table)
	permutations := calculatePermutations(table, instructions)

	maxHappiness := math.MinInt
	for _, currentTable := range *permutations {
		if currentTable.happiness > maxHappiness {
			maxHappiness = currentTable.happiness
		}
	}

	return maxHappiness
}

func addMe(inst *instructions, t *table) {
	me := "filojiston"
	t.people = append(t.people, me)
	(*inst)[me] = make(map[string]int)

	for key := range *inst {
		(*inst)[key][me] = 0
		(*inst)[me][key] = 0
	}
}

func parseInstructions() *instructions {
	lines := util.ReadInputFile("input.txt")
	instructions := make(instructions)

	for _, line := range lines {
		line := line[:len(line)-1]
		var person1, gainOrLose, person2 string
		var happiness int
		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s.", &person1, &gainOrLose, &happiness, &person2)
		if strings.EqualFold(gainOrLose, "lose") {
			happiness *= -1
		}
		if _, hasPerson := instructions[person1]; !hasPerson {
			instructions[person1] = make(map[string]int)
			instructions[person1][person2] = happiness
		} else {
			instructions[person1][person2] = happiness
		}
	}

	return &instructions
}

func newTable(instructions instructions) *table {
	var people []string

	for key := range instructions {
		people = append(people, key)
	}

	return &table{
		people:    people,
		happiness: 0,
	}
}

// in this function, we're calculating all permutations of the people in the table
// but we actually don't need all permutations to calculate the optimum happiness because people are sitting on a circular table.
// all permutations = 8! = 40320
// all unique permutations when they're sitting on a circular table = 8! / 8  = 5040
// so we're doing bit of extra work here, but it's ok for this input (8 people)
// it may be even slower if we check for unique permutations
func calculatePermutations(t *table, instructions *instructions) *permutations {
	var helper func([]string, int)
	var result permutations

	helper = func(arr []string, n int) {
		if n == 1 {
			temp := make([]string, len(arr))
			copy(temp, arr)
			result = append(result, &table{people: temp, happiness: calculateHappiness(temp, *instructions)})
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(t.people, len(t.people))
	return &result
}

func calculateHappiness(people []string, instructions instructions) int {
	var happiness int

	for i := 0; i < len(people); i++ {
		var current, next string
		if i == len(people)-1 {
			next = people[0]
		} else {
			next = people[i+1]
		}
		current = people[i]

		happiness += instructions[current][next]
		happiness += instructions[next][current]
	}

	return happiness
}
