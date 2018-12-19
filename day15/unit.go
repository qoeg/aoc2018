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

func (u *unit) attack(target *unit) (killed *unit) {
	// TODO: damage target
	return killed
}

func (u *unit) nextStep(cells [][]grid.Cell, units []*unit, targets []*unit) (grid.Coordinate, bool) {
	var visited [][]bool
	for x := range cells {
		visited = append(visited, []bool{})
		for y := 0; y < len(cells[x]); y++ {
			visited[x] = append(visited[x], false)
		}
	}
	
	t := grid.NewTree(cells[u.pos.X][u.pos.Y])
	queue := grid.TreeByReadOrder(*t)
	heap.Init(&queue)

	visited[u.pos.X][u.pos.Y] = true

	for queue.Len() > 0 {
		//fmt.Println("Queue", queue)
		node := heap.Pop(&queue).(*grid.Node)
		//fmt.Println("Using", node)
		if findUnit(targets, node.Pos()) != nil {
			path := node.Path()
			//fmt.Println("Found path", path)
			return path[1], true
		}
		for _, n := range grid.Grid(cells).With(renderable(units)...).Neighbors(node, '.', u.enemy()) {
			if !visited[n.Pos().X][n.Pos().Y] {
				visited[n.Pos().X][n.Pos().Y] = true
				//fmt.Println("Visited ", n)
				heap.Push(&queue, grid.NewNode(grid.Coordinate{X:n.Pos().X, Y:n.Pos().Y}, node))
			}
		}
	}
	
	return grid.Coordinate{}, false
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

func (u *unit) move(cells *[][]grid.Cell, units []*unit) (target *unit, none bool) {
	targets := u.enemies(units)
	if len(targets) == 0 {
		return nil, true
	}

	// if adjacent to an enemy, just attack
	adjEnemies := grid.Grid(*cells).With(renderable(targets)...).Neighbors(u, u.enemy())
	if len(adjEnemies) > 0 {
		return findUnit(units, adjEnemies[0].Pos()), false
	}

	// find the best path to a target
	if step, yes := u.nextStep(*cells, units, targets); yes {
		u.pos = step.Pos()

		// check if adjacent to enemy after moving
		adjEnemies = grid.Grid(*cells).With(renderable(targets)...).Neighbors(u, u.enemy())
		if len(adjEnemies) > 0 {
			return findUnit(units, adjEnemies[0].Pos()), false
		}
	}

	return nil, false
}

func findUnit(units []*unit, c grid.Coordinate) *unit {
	for _, u := range units {
		if u.pos == c {
			return u
		}
	}
	return nil
}

func remove(units []*unit, match *unit) {
	for i, u := range units[:] {
		if u == match {
			units = append(units[:i], units[i+1:]...)
		}
	}
}

func renderable(units []*unit) []grid.Renderable {
	r := make([]grid.Renderable, len(units))
	for i, u := range units {
		r[i] = u
	}
	return r
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
