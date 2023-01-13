// https://adventofcode.com/2015/day/16
// solution of advent of code 2015, day16

package day16

import (
	"fmt"
	"math"
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

type aunt struct {
	number     int
	props      map[string]int
	likeliness int
}

var mfcsamOutput = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func findCorrectAunt(likelinessFunc func(aunts *[]aunt)) aunt {
	lines := util.ReadInputFile("input.txt")
	aunts := parseAunts(lines)
	likelinessFunc(&aunts)

	maxLikeliness := math.MinInt
	for _, aunt := range aunts {
		if aunt.likeliness > maxLikeliness {
			maxLikeliness = aunt.likeliness
		}
	}

	var result aunt
	for _, aunt := range aunts {
		if aunt.likeliness == maxLikeliness {
			result = aunt
		}
	}
	return result
}

func parseAunts(input []string) (aunts []aunt) {
	for _, line := range input {
		aunts = append(aunts, parseAunt(line))
	}
	return aunts
}

func parseAunt(line string) aunt {
	line = strings.ReplaceAll(line, ":", "")
	var number int
	var prop1, prop2, prop3 string
	var prop1val, prop2val, prop3val int
	fmt.Sscanf(line, "Sue %d %s %d, %s %d, %s %d", &number, &prop1, &prop1val, &prop2, &prop2val, &prop3, &prop3val)

	return aunt{number: number, props: map[string]int{
		prop1: prop1val,
		prop2: prop2val,
		prop3: prop3val,
	}}
}

func calculateLikelinessesPart1(aunts *[]aunt) {
	for idx, aunt := range *aunts {
		(*aunts)[idx].likeliness = calculateLikelinessPart1(aunt)
	}
}

func calculateLikelinessPart1(aunt aunt) (likeliness int) {
	for prop, val := range aunt.props {
		if mfcsamOutput[prop] == val {
			likeliness++
		}
	}
	return likeliness
}

func calculateLikelinessesPart2(aunts *[]aunt) {
	for idx, aunt := range *aunts {
		(*aunts)[idx].likeliness = calculateLikelinessPart2(aunt)
	}
}

func calculateLikelinessPart2(aunt aunt) (likeliness int) {
	for prop, val := range aunt.props {
		switch prop {
		case "cats":
		case "trees":
			if val > mfcsamOutput[prop] {
				likeliness++
			}
		case "pomeranians":
		case "goldfish":
			if mfcsamOutput[prop] > val {
				likeliness++
			}
		default:
			if mfcsamOutput[prop] == val {
				likeliness++
			}
		}
	}
	return likeliness
}
