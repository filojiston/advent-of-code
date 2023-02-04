// https://adventofcode.com/2015/day/25
// solution of advent of code 2015, day25

package day25

const rowInput int = 2978
const colInput int = 3083

func findCodePosition() int {
	base := rowInput + colInput - 1
	count := (base * (base + 1)) / 2
	return count - rowInput + 1
}

func positionToCode(number int) int {
	code := 20151125
	for i := 1; i < number; i++ {
		code = (code * 252533) % 33554393
	}
	return code
}

func findCode() int {
	pos := findCodePosition()
	code := positionToCode(pos)
	return code
}
