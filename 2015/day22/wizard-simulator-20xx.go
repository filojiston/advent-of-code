// https://adventofcode.com/2015/day/22
// solution of advent of code 2015, day22

package day22

import "math"

type spell struct {
	name string
	cost int
}

type player struct {
	hp, mana, armor, damage int
}

type game struct {
	player, boss                         player
	shieldLeft, poisonLeft, rechargeLeft int
}

func (g *game) processEffects() {
	if g.shieldLeft > 0 {
		g.player.armor = 7
		g.shieldLeft--
	} else {
		g.player.armor = 0
	}
	if g.poisonLeft > 0 {
		g.boss.hp -= 3
		g.poisonLeft--
	}
	if g.rechargeLeft > 0 {
		g.player.mana += 101
		g.rechargeLeft--
	}
}

func calculateLeastManaForWin(difficulty string) int {
	me := player{hp: 50, mana: 500}
	boss := player{hp: 51, damage: 9}
	g := game{player: me, boss: boss}

	return calculateLeastManaForWinRecursive(g, true, 0, difficulty, math.MaxInt)
}

func calculateLeastManaForWinRecursive(g game, turn bool, manaSpent int, difficulty string, leastMana int) int {
	if manaSpent >= leastMana {
		return leastMana
	}

	g.processEffects()

	if g.boss.hp <= 0 {
		return min(leastMana, manaSpent)
	}

	if turn {
		if difficulty == "hard" {
			g.player.hp--
			if g.player.hp <= 0 {
				return leastMana
			}
		}
		if g.player.mana < 53 {
			return leastMana
		}

		for _, s := range getSpells() {
			if canCastSpell(g, s) {
				leastMana = min(leastMana, calculateLeastManaForWinRecursive(castSpell(g, s), !turn, manaSpent+s.cost, difficulty, leastMana))
			}
		}
	} else {
		g.player.hp -= max(1, g.boss.damage-g.player.armor)
		if g.player.hp <= 0 {
			return leastMana
		}
		leastMana = min(leastMana, calculateLeastManaForWinRecursive(g, !turn, manaSpent, difficulty, leastMana))
	}
	return leastMana
}

func castSpell(g game, s spell) game {
	switch s.name {
	case "Magic Missile":
		g.boss.hp -= 4
	case "Drain":
		g.boss.hp -= 2
		g.player.hp += 2
	case "Shield":
		g.shieldLeft = 6
	case "Poison":
		g.poisonLeft = 6
	case "Recharge":
		g.rechargeLeft = 5
	}
	g.player.mana -= s.cost
	return g
}

func canCastSpell(g game, s spell) bool {
	if g.player.mana < s.cost {
		return false
	}
	switch s.name {
	case "Shield":
		return g.shieldLeft == 0
	case "Poison":
		return g.poisonLeft == 0
	case "Recharge":
		return g.rechargeLeft == 0
	}
	return true
}

func getSpells() []spell {
	return []spell{
		{"Magic Missile", 53},
		{"Drain", 73},
		{"Shield", 113},
		{"Poison", 173},
		{"Recharge", 229},
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
