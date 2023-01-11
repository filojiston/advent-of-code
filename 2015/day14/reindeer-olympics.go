// https://adventofcode.com/2015/day/14
// solution of advent of code 2015, day14

package day14

import (
	"fmt"
	"math"

	"github.com/filojiston/advent-of-code/2015/util"
)

func simulateRace(secs int) int {
	lines := util.ReadInputFile("input.txt")
	reindeers := parseReindeers(lines)

	maxDistance := math.MinInt
	for i := 1; i <= secs; i++ {
		for _, reindeer := range reindeers {
			reindeer.race(i)
			if reindeer.distanceTraveled > maxDistance {
				maxDistance = reindeer.distanceTraveled
			}
		}
	}

	return maxDistance
}

func simulateRaceWithLeadPoints(secs int) int {
	lines := util.ReadInputFile("input.txt")
	reindeers := parseReindeers(lines)

	maxDistance := math.MinInt
	for i := 1; i <= secs; i++ {
		for _, reindeer := range reindeers {
			reindeer.race(i)
			if reindeer.distanceTraveled > maxDistance {
				maxDistance = reindeer.distanceTraveled
			}
		}
		for _, reindeer := range reindeers {
			if reindeer.distanceTraveled == maxDistance {
				reindeer.points++
			}
		}
	}

	return getMaxPoints(reindeers)
}

func parseLine(line string) *reindeer {
	var name string
	var speed, flyDuration, restDuration int
	fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &flyDuration, &restDuration)

	return &reindeer{name: name, speed: speed, flyDuration: flyDuration, restDuration: restDuration, flyDurationLeft: flyDuration, restDurationLeft: restDuration, distanceTraveled: 0}
}

func parseReindeers(input []string) []*reindeer {
	var reindeers []*reindeer
	for _, line := range input {
		reindeer := parseLine(line)
		reindeers = append(reindeers, reindeer)
	}

	return reindeers
}

func getMaxPoints(reindeers []*reindeer) int {
	maxPoints := math.MinInt
	for _, reindeer := range reindeers {
		if reindeer.points > maxPoints {
			maxPoints = reindeer.points
		}
	}

	return maxPoints
}
