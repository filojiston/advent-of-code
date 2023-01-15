// https://adventofcode.com/2015/day/19
// solution of advent of code 2015, day19

package day19

import (
	"regexp"
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

type replacement struct {
	from string
	to   string
}

const input string = "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"

func calculateAllDistinctMolecules() map[string]int {
	replacementsInput := util.ReadInputFile("replacements.txt")
	replacements := parseReplacements(replacementsInput)
	molecules := parseMolecules(input)

	distinctMolecules := make(map[string]int)
	for _, replacement := range replacements {
		addAll(&distinctMolecules, createMolecules(replacement, molecules))
	}
	return distinctMolecules
}

// i know this method won't work for all cases, but it works for this puzzle
// and it's fast
func calculateFewestStepsForCreatingMedicine() int {
	replacementsInput := util.ReadInputFile("replacements.txt")
	replacements := parseReplacements(replacementsInput)
	molecule := input

	var count int
	for molecule != "e" {
		var longestReplacement replacement
		for _, replacement := range replacements {
			if strings.Contains(molecule, replacement.to) {
				if len(replacement.to) > len(longestReplacement.to) {
					longestReplacement = replacement
				}
			}
		}
		molecule = strings.Replace(molecule, longestReplacement.to, longestReplacement.from, 1)
		count++
	}
	return count
}

func parseReplacements(input []string) []replacement {
	var replacements []replacement
	for _, line := range input {
		data := strings.Split(line, " => ")
		replacements = append(replacements, replacement{data[0], data[1]})
	}
	return replacements
}

func parseMolecules(input string) []string {
	var re = regexp.MustCompile(`([A-Z][a-z]?)`)
	return re.FindAllString(input, -1)
}

func addAll(distinctMolecules *map[string]int, molecules []string) {
	for _, molecule := range molecules {
		(*distinctMolecules)[molecule]++
	}
}

func createMolecules(replacement replacement, molecules []string) []string {
	var result []string
	for i, molecule := range molecules {
		if molecule == replacement.from {
			newMolecule := make([]string, len(molecules))
			copy(newMolecule, molecules)
			newMolecule[i] = replacement.to
			result = append(result, strings.Join(newMolecule, ""))
		}
	}
	return result
}

// * this was my initial thought to solve the problem, but it was too slow
// func calculateFewestStepsForCreatingMedicine() int {
// 	replacementsInput := util.ReadInputFile("replacements.txt")
// 	replacements := parseReplacements(replacementsInput)
// 	startReplacements := findAll(replacements, "e")

// 	minSteps := math.MaxInt
// 	for _, replacement := range startReplacements {
// 		steps := getMedicineSteps(replacements, parseMolecules(replacement.to), 1)
// 		if steps < minSteps {
// 			minSteps = steps
// 		}
// 	}
// 	return minSteps
// }

// func getMedicineSteps(replacements []replacement, molecules []string, steps int) int {
// 	medicine := strings.Join(molecules, "")
// 	if medicine == input {
// 		return steps
// 	}
// 	if len(medicine) >= len(input) {
// 		return math.MaxInt
// 	}

// 	var stepsForMolecules []int
// 	for _, replacement := range replacements {
// 		currentMolecules := createMolecules(replacement, parseMolecules(medicine))
// 		for _, molecule := range currentMolecules {
// 			stepsForMolecules = append(stepsForMolecules, getMedicineSteps(replacements, parseMolecules(molecule), steps+1))
// 		}
// 	}
// 	return min(stepsForMolecules)
// }

// func min(nums []int) int {
// 	min := math.MaxInt
// 	for _, num := range nums {
// 		if num < min {
// 			min = num
// 		}
// 	}
// 	return min
// }

// func findAll(replacements []replacement, molecule string) []replacement {
// 	var result []replacement
// 	for _, replacement := range replacements {
// 		if replacement.from == molecule {
// 			result = append(result, replacement)
// 		}
// 	}
// 	return result
// }
