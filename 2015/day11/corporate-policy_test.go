// https://adventofcode.com/2015/day/11

package day11

import (
	"strings"
	"testing"
)

func Test_ShouldGetIsPasswordValid(t *testing.T) {
	result := isValidPassword("hijklmmn")
	if result != false {
		t.Fatalf(`isValidPassword("hijklmmn") = %v, want: %v`, result, false)
	}

	result = isValidPassword("abbceffg")
	if result != false {
		t.Fatalf(`isValidPassword("abbceffg") = %v, want: %v`, result, false)
	}

	result = isValidPassword("abbcegjk")
	if result != false {
		t.Fatalf(`isValidPassword("abbcegjk") = %v, want: %v`, result, false)
	}
}

func Test_ShouldGetNextValidPassword(t *testing.T) {
	result := calculateNextValidPassword("abcdefgh")
	if !strings.EqualFold(result, "abcdffaa") {
		t.Fatalf(`getNextValidPassword("abcdefgh"): %v, want: %v`, result, "abcdffaa")
	}

	result = calculateNextValidPassword("ghijklmn")
	if !strings.EqualFold(result, "ghjaabcc") {
		t.Fatalf(`getNextValidPassword("ghjaabcc"): %v, want: %v`, result, "ghjaabcc")
	}
}

func Test_ShouldGetNextValidPassword_PuzzleInput(t *testing.T) {
	result := calculateNextValidPassword("hxbxwxba")
	if !strings.EqualFold(result, "hxbxxyzz") {
		t.Fatalf(`getNextValidPassword("hxbxwxba"): %v, want: %v`, result, "hxbxxyzz")
	}

	result = calculateNextValidPassword(result)
	if !strings.EqualFold(result, "hxcaabcc") {
		t.Fatalf(`getNextValidPassword("hxbxxyzz"): %v, want: %v`, result, "hxcaabcc")
	}
}
