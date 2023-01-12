// https://adventofcode.com/2015/day/4
// unit tests of advent of code 2015, day4

package day4

import "testing"

func Test_ShouldFindLowestNumberToGetMd5HashStartsWithXZeroes(t *testing.T) {
	result := findLowestNumberToGetMd5HashStartsWithXZeroes("abcdef", 5)
	if result != 609043 {
		t.Fatalf(`findLowestNumberToGetMd5HashStartsWithXZeroes("abcdef") = %v, want %v`, result, 609043)
	}

	result = findLowestNumberToGetMd5HashStartsWithXZeroes("pqrstuv", 5)
	if result != 1048970 {
		t.Fatalf(`findLowestNumberToGetMd5HashStartsWithXZeroes("pqrstuv") = %v, want %v`, result, 1048970)
	}
}

func Test_ShouldFindLowestNumberToGetMd5HashStartsWithXZeroes_PuzzleInput(t *testing.T) {
	result := findLowestNumberToGetMd5HashStartsWithXZeroes("yzbqklnj", 5)
	if result != 282749 {
		t.Fatalf(`findLowestNumberToGetMd5HashStartsWithXZeroes("yzbqklnj") = %v, want %v`, result, 282749)
	}

	result = findLowestNumberToGetMd5HashStartsWithXZeroes("yzbqklnj", 6)
	if result != 9962624 {
		t.Fatalf(`findLowestNumberToGetMd5HashStartsWithXZeroes("yzbqklnj") = %v, want %v`, result, 9962624)
	}
}
