// unit tests of advent of code 2015, day2

package day2

import "testing"

func Test_ShouldCalculateTotalSquareFeetOfWrappingPaper(t *testing.T) {
	result := calculateTotalSquareFeetOfWrappingPaper(2, 3, 4)
	if result != 58 {
		t.Fatalf(`calculateTotalSquareFeetOfWrappingPaper(2, 3, 4) = %v, want %v`, result, 58)
	}

	result = calculateTotalSquareFeetOfWrappingPaper(1, 1, 10)
	if result != 43 {
		t.Fatalf(`calculateTotalSquareFeetOfWrappingPaper(1, 1, 10) = %v, want %v`, result, 43)
	}
}

func Test_ShouldCalculateTotalFeetOfRibbon(t *testing.T) {
	result := calculateTotalFeetOfRibbon(2, 3, 4)
	if result != 34 {
		t.Fatalf(`calculateTotalFeetOfRibbon(2, 3, 4) = %v, want %v`, result, 34)
	}

	result = calculateTotalFeetOfRibbon(1, 1, 10)
	if result != 14 {
		t.Fatalf(`calculateTotalFeetOfRibbon(1, 1, 10) = %v, want %v`, result, 14)
	}
}

func Test_ShouldGetSmallestOfTwoInts(t *testing.T) {
	x, y := getSmallestTwoInts(1, 1, 10)
	if x != 1 || y != 1 {
		t.Fatalf(`getSmallestTwoInts(1, 1, 10) = %v, %v, want %v, %v`, x, y, 1, 1)
	}
}

func Test_ShouldGetDimensionsFromLine(t *testing.T) {
	length, width, height := getDimensionsFromLine("2x3x4")
	if length != 2 || width != 3 || height != 4 {
		t.Fatalf(`getDimensionsFromLine("2x3x4") = %v, %v, %v want %v, %v, %v`, length, width, height, 2, 3, 4)
	}
}

func Test_ShouldCalculateTotalSquareOfWrappingPaperFromFile_PuzzleInput(t *testing.T) {
	result := calculateTotalSquareFeetOfWrappingPaperFromFile()
	if result != 1586300 {
		t.Fatalf(`calculateTotalSquareFeetOfWrappingPaperFromFile() = %v, want %v`, result, 1586300)
	}
}

func Test_ShouldCalculateTotalFeetOfRibbonFromFile_PuzzleInput(t *testing.T) {
	result := calculateTotalFeetOfRibbonFromFile()
	if result != 3737498 {
		t.Fatalf(`calculateTotalFeetOfRibbonFromFile() = %v, want %v`, result, 3737498)
	}
}
