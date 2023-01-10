// https://adventofcode.com/2015/day/12
// solution of advent of code 2015, day12

package day12

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

func sumAllNumbersInInput() int {
	input := util.ReadInputFile("input.txt")[0]
	numbers := getAllNumbers(input)
	return sum(numbers)
}

func getAllNumbers(input string) []int {
	var re = regexp.MustCompile(`-?\d+`)
	numbers := re.FindAllString(input, -1)

	var result []int
	for _, number := range numbers {
		numberToInt, _ := strconv.Atoi(number)
		result = append(result, numberToInt)
	}

	return result
}

func sumAllNumbersInInputExcludingReds() int {
	input := util.ReadInputFile("input.txt")[0]
	var data interface{}
	var numbers []int

	json.Unmarshal([]byte(input), &data)
	traverse(data, &numbers)
	return sum(numbers)
}

func traverse(data interface{}, numbers *[]int) {
	switch v := data.(type) {
	case map[string]interface{}:
		if !hasRed(v) {
			for _, value := range v {
				traverse(value, numbers)
			}
		}
	case []interface{}:
		for _, value := range v {
			traverse(value, numbers)
		}
	case float64:
		*numbers = append(*numbers, int(v))
	}
}

func hasRed(data map[string]interface{}) bool {
	for _, value := range data {
		if reflect.ValueOf(value).Kind() == reflect.String {
			value := fmt.Sprintf("%v", value)
			if strings.EqualFold(value, "red") {
				return true
			}
		}
	}

	return false
}

func sum(arr []int) int {
	var result int
	for _, elem := range arr {
		result += elem
	}

	return result
}
