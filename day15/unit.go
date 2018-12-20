package day15

import (
	"container/heap"
	//"fmt"
	//"sort"

	"github.com/qoeg/aoc2018/grid"
)

type unit struct {
	atk int
	hp int
	mark rune
	pos grid.Coordinate
}

func (u *unit) Mark() rune {
	return u.mark
}

func (u *unit) Pos() grid.Coordinate {
	return u.pos
}

func (u *unit) attack(targets []*unit) (*unit, bool) {
	if len(targets) == 0 {
		return nil, false
	}
	
	attacked := targets[0]
	mins := []*unit{attacked}

	for i := 1; i < len(targets); i++ {
		if targets[i].hp < mins[0].hp {
			mins = []*unit{targets[i]}
		} else if targets[i].hp == mins[0].hp {
			mins = append(mins, targets[i])
		}
	}

	if len(mins) > 0 {
		attacked = mins[0]
	}
	
	attacked.hp -= u.atk

	return attacked, attacked.hp <= 0
}

func (u *unit) nextStep(cells [][]grid.Cell, us units, targets units) (grid.Cell, bool) {
	visited := map[grid.Coordinate]bool{}
	visited[u.pos] = true
	
	t := grid.NewTree(cells[u.pos.X][u.pos.Y])
	queue := grid.TreeByReadOrder(*t)
	heap.Init(&queue)

	for queue.Len() > 0 {
		node := heap.Pop(&queue).(*grid.Node)
		if len(targets.where(node.Value)) > 0 {
			path := node.Path()
			return path[1], true
		}
		for _, n := range grid.Grid(cells).With(us.cells()).Neighbors(node, '.', u.enemy()) {
			if !visited[n.Pos()] {
				visited[n.Pos()] = true
				heap.Push(&queue, grid.NewNode(n, node))
			}
		}
	}
	
	return grid.Cell{}, false
}

func (u *unit) enemies(units []*unit) []*unit {
	e := make([]*unit, 0, len(units))
	for _, x := range units {
		if x.mark != u.mark {
			e = append(e, x)
		}
	}
	return e
}

func (u *unit) enemy() rune {
	if u.mark == 'E' { return 'G' }
	return 'E'
}

func (u *unit) move(cells *[][]grid.Cell, us units) (targets units, none bool) {
	targets = units(u.enemies(us))
	if len(targets) == 0 {
		return targets, true
	}

	// if adjacent to an enemy, just attack
	adjEnemies := grid.Grid(*cells).With(targets.cells()).Neighbors(u, u.enemy())
	if len(adjEnemies) > 0 {
		return us.where(adjEnemies...), false
	}

	// find the best path to a target
	if step, yes := u.nextStep(*cells, us, targets); yes {
		u.pos = step.Pos()

		// check if adjacent to enemy after moving
		adjEnemies = grid.Grid(*cells).With(targets.cells()).Neighbors(u, u.enemy())
		if len(adjEnemies) > 0 {
			return us.where(adjEnemies...), false
		}
	}

	return units{}, false
}

type units []*unit

func (us units) cells() []grid.Cell {
	res := []grid.Cell{}
	for _, unit := range us {
		res = append(res, grid.NewCell(unit.mark, unit.pos))
	}
	return res
}

func (us units) where(cs ...grid.Cell) units {
	res := units{}
	for i := range cs {
		for j := range us {
			if cs[i] == grid.NewCell(us[j].mark, us[j].pos) {
				res = append(res, us[j])
			}
		}
	}
	return res
}

func remove(us *[]*unit, match *unit) {
	for i, u := range (*us)[:] {
		if u == match {
			*us = append((*us)[:i], (*us)[i+1:]...)
		}
	}
}

func replaceFor(u *[]*unit) func(grid.Cell) rune {
	return func(c grid.Cell) rune {
		if c.Mark() == 'E' || c.Mark() == 'G' {
			*u = append(*u, &unit{atk: 3, hp: 200, mark: c.Mark(), pos: c.Pos()})
			return '.'
		}
		return c.Mark()
	}
}

type unitByReadOrder []*unit

func (u unitByReadOrder) Len() int {
	return len(u)
}

func (u unitByReadOrder) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u unitByReadOrder) Less(i, j int) bool {
	if u[i].pos.Y == u[j].pos.Y {
		return u[i].pos.X < u[j].pos.X
	}
	return u[i].pos.Y < u[j].pos.Y
}
