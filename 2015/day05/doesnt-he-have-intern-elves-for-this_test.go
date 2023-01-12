// https://adventofcode.com/2015/day/5
// unit tests of advent of code 2015, day5

package day5

import "testing"

func Test_ShouldCheckIfStringIsNicePartOne(t *testing.T) {
	result := isStringNicePartOne("ugknbfddgicrmopn")
	if result != true {
		t.Fatalf(`isStringNicePartOne("ugknbfddgicrmopn") = %v, want %v`, result, true)
	}

	result = isStringNicePartOne("aaa")
	if result != true {
		t.Fatalf(`isStringNicePartOne("aaa") = %v, want %v`, result, true)
	}

	result = isStringNicePartOne("jchzalrnumimnmhp")
	if result != false {
		t.Fatalf(`isStringNicePartOne("jchzalrnumimnmhp") = %v, want %v`, result, false)
	}

	result = isStringNicePartOne("haegwjzuvuyypxyu")
	if result != false {
		t.Fatalf(`isStringNicePartOne("haegwjzuvuyypxyu") = %v, want %v`, result, false)
	}

	result = isStringNicePartOne("dvszwmarrgswjxmb")
	if result != false {
		t.Fatalf(`isStringNicePartOne("dvszwmarrgswjxmb") = %v, want %v`, result, false)
	}
}

func Test_ShouldCheckIfStringIsNicePartTwo(t *testing.T) {
	result := isStringNicePartTwo("qjhvhtzxzqqjkmpb")
	if result != true {
		t.Fatalf(`isStringNicePartTwo("qjhvhtzxzqqjkmpb") = %v, want %v`, result, true)
	}

	result = isStringNicePartTwo("xxyxx")
	if result != true {
		t.Fatalf(`isStringNicePartTwo("xxyxx") = %v, want %v`, result, true)
	}

	result = isStringNicePartTwo("aaaa")
	if result != true {
		t.Fatalf(`isStringNicePartTwo("aaaa") = %v, want %v`, result, true)
	}

	result = isStringNicePartTwo("aaa")
	if result != false {
		t.Fatalf(`isStringNicePartTwo("aaa") = %v, want %v`, result, false)
	}

	result = isStringNicePartTwo("uurcxstgmygtbstg")
	if result != false {
		t.Fatalf(`isStringNicePartTwo("uurcxstgmygtbstg") = %v, want %v`, result, false)
	}

	result = isStringNicePartTwo("ieodomkazucvgmuy")
	if result != false {
		t.Fatalf(`isStringNicePartTwo("ieodomkazucvgmuy") = %v, want %v`, result, false)
	}
}

func Test_ShouldCountNiceStrings_PuzzleInput(t *testing.T) {
	result := countNiceStrings(isStringNicePartOne)
	if result != 258 {
		t.Fatalf(`countNiceStrings(isNiceStringPartOne) = %v, want %v`, result, 258)
	}

	result = countNiceStrings(isStringNicePartTwo)
	if result != 53 {
		t.Fatalf(`countNiceStrings(isStringNicePartTwo) = %v, want %v`, result, 53)
	}
}
