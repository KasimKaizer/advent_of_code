package day22

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const noOfSpells = 5

type Player struct {
	hp, mana, armor int
	active          [noOfSpells]int
}

type spell struct {
	name                          string
	cost                          int
	dmg, heal, armor, mana, timer int
}

var spells = [noOfSpells]spell{
	0: {"Magic Missile", 53, 4, 0, 0, 0, 0},
	1: {"Drain", 73, 2, 2, 0, 0, 0},
	2: {"Shield", 113, 0, 0, 7, 0, 6},
	3: {"Poison", 173, 3, 0, 0, 0, 6},
	4: {"Recharge", 229, 0, 0, 0, 101, 5},
}

func (p *Player) useSpell(s spell) int {
	p.hp += s.heal
	p.mana += s.mana
	p.armor += s.armor
	return s.dmg
}

func (p *Player) attack(idx int) (int, bool) {
	if !validateSpell(*p, idx) {
		return 0, false
	}
	p.mana -= spells[idx].cost
	if spells[idx].timer == 0 {
		return p.useSpell(spells[idx]), true
	}
	p.active[idx] = spells[idx].timer
	return 0, true
}

func (p *Player) spellRun() int {
	var totalDmg int
	p.armor = 0 // Reset armor at the start of effects
	for idx, count := range p.active {
		if count == 0 {
			continue
		}
		totalDmg += p.useSpell(spells[idx])
		p.active[idx]--
	}
	return totalDmg
}

type enemy struct {
	hp, dmg int
}

func SolveOne(player, enemy string) (int, error) {
	return solve(
		player,
		enemy,
		func(p Player) Player {
			return p
		})
}

func SolveTwo(player, enemy string) (int, error) {
	return solve(
		player,
		enemy,
		func(p Player) Player {
			p.hp--
			return p
		})
}

func solve(ply, enm string, f func(Player) Player) (int, error) {
	p, err := createPlayer(ply)
	if err != nil {
		return 0, err
	}
	e, err := createEnemy(enm)
	if err != nil {
		return 0, err
	}
	mem := make(map[string]int)
	return fight(p, e, mem, f), nil
}

var re = regexp.MustCompile(`(\d+)`)

func createPlayer(p string) (Player, error) {
	var ply Player
	var err error
	numChars := re.FindAllString(p, -1)
	ply.hp, err = strconv.Atoi(numChars[0])
	if err != nil {
		return ply, err
	}
	ply.mana, err = strconv.Atoi(numChars[1])
	return ply, err
}

func createEnemy(e string) (enemy, error) {
	var emy enemy
	var err error
	numChars := re.FindAllString(e, -1)
	emy.hp, err = strconv.Atoi(numChars[0])
	if err != nil {
		return emy, err
	}
	emy.dmg, err = strconv.Atoi(numChars[1])
	return emy, err
}

func fight(p Player, e enemy, mem map[string]int, f func(Player) Player) int {
	// memoization check
	hash := createHash(p, e)
	if num, ok := mem[hash]; ok {
		return num
	}
	// pre-Condition for player
	p = f(p)

	// Player's turn
	// apply spell effects
	e.hp -= p.spellRun()
	if e.hp <= 0 {
		return 0 // Boss defeated
	}

	// try casting each spell
	minCost := math.MaxInt32
	for i := range noOfSpells {
		if !validateSpell(p, i) {
			continue
		}

		// Copy player and enemy states for recursion
		newP := p
		newE := e

		// Cast spell and apply its immediate effect if applicable
		curAtk, ok := newP.attack(i)
		if !ok {
			continue // Invalid spell cast
		}
		newE.hp -= curAtk

		// Check if boss is defeated after casting spell
		if newE.hp <= 0 {
			minCost = min(minCost, spells[i].cost)
			continue
		}

		// Boss's turn
		// apply ongoing effects again
		newE.hp -= newP.spellRun()
		if newE.hp <= 0 {
			minCost = min(minCost, spells[i].cost)
			continue
		}

		// Boss attacks player
		newP.hp -= max(newE.dmg-newP.armor, 1)
		if newP.hp <= 0 {
			continue // Player defeated
		}
		// Recurse into next round with updated states
		minCost = min(minCost, spells[i].cost+fight(newP, newE, mem, f))
	}

	// Memoize result
	mem[hash] = minCost
	return minCost
}

func createHash(p Player, e enemy) string {
	return fmt.Sprintf("%d:%d:%d:%d:%d:%d:%d:%d",
		p.hp, p.mana,
		p.active[0], p.active[1], p.active[2], p.active[3], p.active[4],
		e.hp)
}

func validateSpell(p Player, spellIdx int) bool {
	return p.mana >= spells[spellIdx].cost && p.active[spellIdx] == 0
}

// alternative approach for part 1
//
// func SolveOne(player, enemy string) (int, error) {
// 	p, err := createPlayer(player)
// 	if err != nil {
// 		return 0, err
// 	}
// 	e, err := createEnemy(enemy)
// 	if err != nil {
// 		return 0, err
// 	}
// 	mem := make(map[string]int)
// 	minCost := math.MaxInt32
// 	for i := range noOfSpells {
// 		minCost = min(minCost, fight(p, e, i, mem))
// 	}
// 	return minCost, nil
// }
//
// func createHash(p player, e enemy, nxtSpell int) string {
// 	return fmt.Sprintf("%d:%d:%d:%d:%d:%d:%d:%d:%d",
// 		p.hp, p.mana,
// 		p.active[0], p.active[1], p.active[2], p.active[3], p.active[4],
// 		e.hp, nxtSpell)
// }
//
// func fight(p player, e enemy, spellID int, mem map[string]int) int {
// 	// momoize check
// 	hash := createHash(p, e, spellID)
// 	if num, ok := mem[hash]; ok {
// 		return num
// 	}
//
// 	// player turn
// 	// players turn effect
// 	e.hp -= p.spellRun()
// 	if e.hp <= 0 {
// 		return 0
// 	}
//
// 	// cast new spell
// 	curAtk, ok := p.attack(spellID)
// 	if !ok {
// 		return math.MaxInt32
// 	}
// 	e.hp -= curAtk
// 	if e.hp <= 0 {
// 		return spells[spellID].cost
// 	}
//
// 	// boss turn
// 	// boss turn effect
// 	e.hp -= p.spellRun()
// 	if e.hp <= 0 {
// 		return spells[spellID].cost
// 	}
//
// 	// boss attacks player
// 	p.hp -= max((e.dmg - p.armor), 1)
// 	if p.hp <= 0 {
// 		return math.MaxInt32
// 	}
//
// 	// next possible player turns
// 	minCost := math.MaxInt32
// 	for i := range noOfSpells {
// 		nextCost := fight(p, e, i, mem)
// 		minCost = min(minCost, spells[spellID].cost+nextCost)
// 	}
// 	// memoize
// 	mem[hash] = minCost
// 	return minCost
// }
