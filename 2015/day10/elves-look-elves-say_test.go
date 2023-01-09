// unit tests of advent of code 2015, day10

package day10

import "testing"

func Test_ShouldLookAndSayXTimes_PuzzleInput(t *testing.T) {
	result := lookAndSayXTimes("1321131112", 40)
	if len(result) != 492982 {
		t.Fatalf(`lookAndSayXTimes("1321131112", 40) = %v, want %v`, result, 492982)
	}

	// pick up where we left off
	result = lookAndSayXTimes(result, 10)
	if len(result) != 6989950 {
		t.Fatalf(`lookAndSayXTimes("1321131112", 50) = %v, want %v`, result, 6989950)
	}
}

func Test_ShouldLookAndSay(t *testing.T) {
	result := lookAndSay("1")
	if result != "11" {
		t.Fatalf(`lookAndSay("1") = %v, want %v`, result, "11")
	}

	result = lookAndSay("11")
	if result != "21" {
		t.Fatalf(`lookAndSay("11") = %v, want %v`, result, "21")
	}

	result = lookAndSay("111221")
	if result != "312211" {
		t.Fatalf(`lookAndSay("111221") = %v, want %v`, result, "312211")
	}
}

func Test_ShouldLookAndSayPerformant(t *testing.T) {
	result := lookAndSayPerformant("1")
	if result != "11" {
		t.Fatalf(`lookAndSayPerformant("1") = %v, want %v`, result, "11")
	}

	result = lookAndSayPerformant("11")
	if result != "21" {
		t.Fatalf(`lookAndSayPerformant("11") = %v, want %v`, result, "21")
	}

	result = lookAndSayPerformant("111221")
	if result != "312211" {
		t.Fatalf(`lookAndSayPerformant("111221") = %v, want %v`, result, "312211")
	}
}
