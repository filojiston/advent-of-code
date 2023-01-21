// https://adventofcode.com/2015/day/21
// solution of advent of code 2015, day21

package day21

type itemType int

const (
	Weapon itemType = iota
	Armor
	Ring
)

type item struct {
	itemType itemType
	name     string
	cost     int
	damage   int
	armor    int
}

type shop struct {
	weapons []item
	armors  []item
	rings   []item
}

type player struct {
	hp     int
	damage int
	armor  int
}

type combination struct {
	weapon item
	armor  item
	ring1  item
	ring2  item
	cost   int
}

func (p *player) equip(combination combination) {
	p.damage = combination.weapon.damage + combination.armor.damage + combination.ring1.damage + combination.ring2.damage
	p.armor = combination.weapon.armor + combination.armor.armor + combination.ring1.armor + combination.ring2.armor
}

func createShop() shop {
	var shop shop
	shop.weapons = []item{
		{Weapon, "Dagger", 8, 4, 0},
		{Weapon, "Shortsword", 10, 5, 0},
		{Weapon, "Warhammer", 25, 6, 0},
		{Weapon, "Longsword", 40, 7, 0},
		{Weapon, "Greataxe", 74, 8, 0},
	}
	shop.armors = []item{
		{Armor, "nothing", 0, 0, 0},
		{Armor, "Leather", 13, 0, 1},
		{Armor, "Chainmail", 31, 0, 2},
		{Armor, "Splintmail", 53, 0, 3},
		{Armor, "Bandedmail", 75, 0, 4},
		{Armor, "Platemail", 102, 0, 5},
	}
	shop.rings = []item{
		{Ring, "nothing", 0, 0, 0},
		{Ring, "Damage +1", 25, 1, 0},
		{Ring, "Damage +2", 50, 2, 0},
		{Ring, "Damage +3", 100, 3, 0},
		{Ring, "Defense +1", 20, 0, 1},
		{Ring, "Defense +2", 40, 0, 2},
		{Ring, "Defense +3", 80, 0, 3},
	}
	return shop
}

func findCheapestWinningCombination() combination {
	shop := createShop()
	boss := player{hp: 103, damage: 9, armor: 2}
	player := player{hp: 100, damage: 0, armor: 0}
	combinations := createCombinations(shop)
	result := minimumCostCombination(filterWinning(combinations, boss, player))
	return result
}

func findMostExpensiveLossingCombination() combination {
	shop := createShop()
	boss := player{hp: 103, damage: 9, armor: 2}
	player := player{hp: 100, damage: 0, armor: 0}
	combinations := createCombinations(shop)
	result := maximumCostCombination(filterLossing(combinations, boss, player))
	return result
}

func createCombinations(shop shop) []combination {
	var combinations []combination
	for _, weapon := range shop.weapons {
		for _, armor := range shop.armors {
			for _, ring1 := range shop.rings {
				for _, ring2 := range shop.rings {
					if ring1 == ring2 {
						continue
					}
					combinations = append(combinations, combination{weapon: weapon, armor: armor, ring1: ring1, ring2: ring2, cost: weapon.cost + armor.cost + ring1.cost + ring2.cost})
				}
			}
		}
	}
	return combinations
}

func filterWinning(combinations []combination, boss player, player player) []combination {
	var winning []combination
	for _, combination := range combinations {
		player.equip(combination)
		if isWinning(combination, boss, player) {
			winning = append(winning, combination)
		}
	}
	return winning
}

func filterLossing(combinations []combination, boss player, player player) []combination {
	var lossing []combination
	for _, combination := range combinations {
		player.equip(combination)
		if !isWinning(combination, boss, player) {
			lossing = append(lossing, combination)
		}
	}
	return lossing
}

func isWinning(combination combination, boss player, player player) bool {
	for player.hp > 0 && boss.hp > 0 {
		boss.hp -= max(1, player.damage-boss.armor)
		if boss.hp <= 0 {
			break
		}
		player.hp -= max(1, boss.damage-player.armor)
	}
	return player.hp > 0
}

func minimumCostCombination(combinations []combination) combination {
	min := combinations[0]
	for _, combination := range combinations {
		if combination.cost < min.cost {
			min = combination
		}
	}
	return min
}

func maximumCostCombination(combinations []combination) combination {
	max := combinations[0]
	for _, combination := range combinations {
		if combination.cost > max.cost {
			max = combination
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
