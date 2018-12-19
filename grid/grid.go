package grid

import (
	"fmt"
)

type Grid [][]Cell

func (g Grid) Neighbors(p TwoDimensional, include ...rune) []Cell {
	n := []Cell{}
	if len(g) == 0 || len(g[0]) == 0 {
		fmt.Println("grid.Neighbors(): the grid is empty")
		return n
	}

	pos := p.Pos()
	sides := make([]Cell, 4)
	if pos.Y != 0 {
		sides[0] = g[pos.X][pos.Y-1]
	}
	if pos.X != 0 {
		sides[1] = g[pos.X-1][pos.Y]
	}
	if pos.X != len(g)-1 {
		sides[2] = g[pos.X+1][pos.Y]
	}
	if pos.Y != len(g[pos.X])-1 {
		sides[3] = g[pos.X][pos.Y+1]
	}
	for _, s := range sides {
		if len(include) == 0 {
			n = append(n, s)
			continue
		}
		for _, in := range include {
			if s.mark == in {
				n = append(n, s)
				break
			}
		}
	}
	return n
}

func (g Grid) Print(objs ...Renderable) {
	for y := 0; y < len(g[0]); y++ {
		for x := 0; x < len(g); x++ {
			// print any objects for the location
			masked := false
			for i, o := range objs {
				if (o.Pos() == Coordinate{x, y}) {
					fmt.Print(string(o.Mark()))
					objs = append(objs[:i], objs[i+1:]...)
					masked = true
					break
				}
			}

			if !masked {
				fmt.Print(string(g[x][y].mark))
			}
		}
		fmt.Println()
	}
}

func (g Grid) With(objs ...Renderable) Grid {
	for _, o := range objs {
		g[o.Pos().X][o.Pos().Y] = Cell{
			mark: o.Mark(),
			pos: o.Pos(),
		}
	}
	return g
}

func Parse(input string) (grid [][]Cell) {
	return ParseOut(input, nil)
}

func ParseOut(input string, replace func(Cell) rune) (grid [][]Cell) {
	var x, y int
	for _, r := range input {
		if r == rune(10) {
			y++
			x = 0
			continue
		}

		c := Cell{mark:r, pos:Coordinate{X:x, Y:y}}
		if replace != nil {
			c.mark = replace(c)
		}

		if y == 0 {
			grid = append(grid, []Cell{})
		}

		grid[x] = append(grid[x], c)
		x++
	}

	return grid
}
