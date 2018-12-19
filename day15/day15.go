package day15

import (
	"sort"

	"github.com/qoeg/aoc2018/grid"
)

func totalHp(units []*unit) int {
	var hp int
	for i := range units {
		hp += units[i].hp
	}
	return hp
}

func Answer() int {
	units := []*unit{}
	g := grid.ParseOut(Input, replaceFor(&units))
	grid.Grid(g).Print(renderable(units)...)

	var done bool
	var rounds = -1
	for rounds < 5 {
		rounds++
		sort.Sort(unitByReadOrder(units))

		for _, u := range units {
			var atk *unit
			if atk, done = u.move(&g, units); done {
				break
			}
			if atk != nil {
				killed := u.attack(atk)
				if killed != nil {
					remove(units, killed)
				}
			}
		}

		grid.Grid(g).Print(renderable(units)...)
	}

	return rounds + totalHp(units)
}
