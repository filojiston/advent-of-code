package day2

import "testing"

func Test_ShouldCalculateTotalSquareFeetOfWrappingPaper(t *testing.T) {
	result := CalculateTotalSquareFeetOfWrappingPaper(2, 3, 4)
	if result != 58 {
		t.Fatalf(`CalculateTotalSquareFeetOfWrappingPaper(2, 3, 4) = %v, want %v`, result, 58)
	}

	result = CalculateTotalSquareFeetOfWrappingPaper(1, 1, 10)
	if result != 43 {
		t.Fatalf(`CalculateTotalSquareFeetOfWrappingPaper(1, 1, 10) = %v, want %v`, result, 43)
	}
}

func Test_ShouldCalculateTotalFeetOfRibbon(t *testing.T) {
	result := CalculateTotalFeetOfRibbon(2, 3, 4)
	if result != 34 {
		t.Fatalf(`CalculateTotalFeetOfRibbon(2, 3, 4) = %v, want %v`, result, 34)
	}

	result = CalculateTotalFeetOfRibbon(1, 1, 10)
	if result != 14 {
		t.Fatalf(`CalculateTotalFeetOfRibbon(1, 1, 10) = %v, want %v`, result, 14)
	}
}

func Test_ShouldGetSmallestOfTwoInts(t *testing.T) {
	x, y := GetSmallestTwoInts(1, 1, 10)
	if x != 1 || y != 1 {
		t.Fatalf(`GetSmallestTwoInts(1, 1, 10) = %v, %v, want %v, %v`, x, y, 1, 1)
	}
}

func Test_ShouldGetDimensionsFromLine(t *testing.T) {
	length, width, height := GetDimensionsFromLine("2x3x4")
	if length != 2 || width != 3 || height != 4 {
		t.Fatalf(`GetDimensionsFromLine("2x3x4") = %v, %v, %v want %v, %v, %v`, length, width, height, 2, 3, 4)
	}
}

func Test_ShouldCalculateTotalSquareOfWrappingPaperFromFile_PuzzleInput(t *testing.T) {
	result := CalculateTotalSquareFeetOfWrappingPaperFromFile()
	if result != 1586300 {
		t.Fatalf(`CalculateTotalSquareFeetOfWrappingPaperFromFile() = %v, want %v`, result, 1586300)
	}
}

func Test_ShouldCalculateTotalFeetOfRibbonFromFile_PuzzleInput(t *testing.T) {
	result := CalculateTotalFeetOfRibbonFromFile()
	if result != 3737498 {
		t.Fatalf(`CalculateTotalFeetOfRibbonFromFile() = %v, want %v`, result, 3737498)
	}
}
