package day7

import "testing"

func Test_ShouldGetSignalOfWirePartOne_PuzzleInput(t *testing.T) {
	emulateCircuit("input1.txt")
	result := getWire("a")
	if result.signal.value != 46065 {
		t.Fatalf(`signal of the wire a should %v, get %v`, 46095, result)
	}
	resetInstructions()
	resetWires()
}

func Test_ShouldGetSignalOfWirePartTwo_PuzzleInput(t *testing.T) {
	emulateCircuit("input2.txt")
	result := getWire("a")
	if result.signal.value != 14134 {
		t.Fatalf(`signal of the wire a should %v, get %v`, 14134, result)
	}
	resetInstructions()
	resetWires()
}
