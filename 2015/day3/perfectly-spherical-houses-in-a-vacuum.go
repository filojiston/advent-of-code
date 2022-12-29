// https://adventofcode.com/2015/day/3
// solution of advent of code 2015, day3

package day3

type position struct {
	x int
	y int
}

func calculateHomeCountReceivedAtLeastOnePresent(input string) (homeCount uint32) {
	visitedHomes := make([]position, 0)
	var santaLocation position

	visitedHomes = append(visitedHomes, santaLocation)

	for _, direction := range input {
		getNewLocationFromDirection(direction, &santaLocation)
		if isVisited(&visitedHomes, santaLocation) {
			continue
		}
		visitedHomes = append(visitedHomes, santaLocation)
	}

	return uint32(len(visitedHomes))
}

func calculateHomeCountReceivedAtLeastOnePresentWithRoboSanta(input string) (homeCount uint32) {
	visitedHomes := make([]position, 0)
	var santaLocation position
	var roboSantaLocation position

	visitedHomes = append(visitedHomes, santaLocation)

	for idx, direction := range input {
		if idx%2 == 0 {
			getNewLocationFromDirection(direction, &santaLocation)
			if isVisited(&visitedHomes, santaLocation) {
				continue
			}
			visitedHomes = append(visitedHomes, santaLocation)
		} else {
			getNewLocationFromDirection(direction, &roboSantaLocation)
			if isVisited(&visitedHomes, roboSantaLocation) {
				continue
			}
			visitedHomes = append(visitedHomes, roboSantaLocation)
		}
	}

	return uint32(len(visitedHomes))
}

func getNewLocationFromDirection(direction rune, currentLocation *position) {
	if direction == '^' {
		currentLocation.y++
	} else if direction == '>' {
		currentLocation.x++
	} else if direction == '<' {
		currentLocation.x--
	} else if direction == 'v' {
		currentLocation.y--
	}
}

func isVisited(visitedHomes *[]position, location position) (result bool) {
	for _, home := range *visitedHomes {
		if home.x == location.x && home.y == location.y {
			return true
		}
	}

	return false
}
