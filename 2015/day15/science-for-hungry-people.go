// https://adventofcode.com/2015/day/15
// solution of advent of code 2015, day15

package day15

import (
	"fmt"
	"math"

	"github.com/filojiston/advent-of-code/2015/util"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
	quantity   int
}

type cookie struct {
	ingredients []ingredient
	score       int
	calories    int
}

const teaspoons int = 100

func getOptimalCookieScore() int {
	lines := util.ReadInputFile("input.txt")
	ingredients := parseIngredients(lines)
	var possibleCookies []cookie
	curr := make([]int, len(ingredients))
	createCookies(ingredients, teaspoons, curr, 0, &possibleCookies)

	optimalScore := math.MinInt
	for _, cookie := range possibleCookies {
		if cookie.score > optimalScore {
			optimalScore = cookie.score
		}
	}

	return optimalScore
}

func getOptimalCookieScoreWithCalories(calories int) int {
	lines := util.ReadInputFile("input.txt")
	ingredients := parseIngredients(lines)
	var possibleCookies []cookie
	curr := make([]int, len(ingredients))
	createCookies(ingredients, teaspoons, curr, 0, &possibleCookies)

	optimalScore := math.MinInt
	for _, cookie := range possibleCookies {
		if cookie.score > optimalScore && cookie.calories == calories {
			optimalScore = cookie.score
		}
	}

	return optimalScore
}

func parseIngredients(input []string) (ingredients []ingredient) {
	for _, line := range input {
		ingredients = append(ingredients, parseIngredient(line))
	}

	return ingredients
}

func parseIngredient(line string) ingredient {
	var name string
	var capacity, durability, flavor, texture, calories int
	fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)

	return ingredient{
		name:       name,
		capacity:   capacity,
		durability: durability,
		flavor:     flavor,
		texture:    texture,
		calories:   calories,
	}
}

func createCookies(ingredients []ingredient, teaspoonsLeft int, curr []int, idx int, results *[]cookie) {
	k := len(ingredients)
	if idx == k-1 {
		curr[idx] = teaspoonsLeft
		newIngredients := make([]ingredient, k)
		copy(newIngredients, ingredients)
		for i, qty := range curr {
			newIngredients[i].quantity = qty
		}
		*results = append(*results, cookie{ingredients: newIngredients, score: calculateScore(newIngredients), calories: calculateCalories(newIngredients)})
		return
	}
	for i := 0; i <= teaspoonsLeft; i++ {
		curr[idx] = i
		createCookies(ingredients, teaspoonsLeft-i, curr, idx+1, results)
	}
}

func calculateScore(ingredients []ingredient) int {
	var totalCapacity, totalDurability, totalFlavor, totalTexture int
	for _, ingredient := range ingredients {
		totalCapacity += ingredient.capacity * ingredient.quantity
		totalDurability += ingredient.durability * ingredient.quantity
		totalFlavor += ingredient.flavor * ingredient.quantity
		totalTexture += ingredient.texture * ingredient.quantity
	}

	totalCapacity = int(math.Max(float64(totalCapacity), 0))
	totalDurability = int(math.Max(float64(totalDurability), 0))
	totalFlavor = int(math.Max(float64(totalFlavor), 0))
	totalTexture = int(math.Max(float64(totalTexture), 0))

	return totalCapacity * totalDurability * totalFlavor * totalTexture
}

func calculateCalories(ingredients []ingredient) int {
	var totalCalories int
	for _, ingredient := range ingredients {
		totalCalories += ingredient.calories * ingredient.quantity
	}
	return totalCalories
}
