// https://adventofcode.com/2015/day/1
// solution of advent of code 2015, day1

package day1

func getFloorForSanta(input string) (floor int) {
	for _, ch := range input {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
	}

	return floor
}

func findPositionOfTheFirstCharacterToBasement(input string) (floor int) {
	for idx, ch := range input {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}

		if floor == -1 {
			return idx + 1
		}
	}

	return -1
}
