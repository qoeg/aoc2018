package day15

import (
	"fmt"
	"sort"
	"time"

	"github.com/qoeg/aoc2018/grid"
)

var sample =
`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`

func totalHp(units []*unit) int {
	var hp int
	for i := range units {
		hp += units[i].hp
	}
	return hp
}

func Answer() string {
	us := []*unit{}
	g := grid.ParseOut(sample, replaceFor(&us))

	var round = 0
	Print(g, us, round)

	var done bool
	for !done {
		immUnits := make([]*unit, len(us))
		copy(immUnits, us)
		
		for _, u := range immUnits {
			if u.hp <= 0 {
				continue
			}

			var targets []*unit
			if targets, done = u.move(&g, us); done {
				break
			}
			if atk, dead := u.attack(targets); dead {
				remove(&us, atk)
			}
		}

		if done {
			break
		}

		sort.Sort(unitByReadOrder(us))

		round++
		Print(g, us, round)

		time.Sleep(time.Millisecond * 50)
	}

	return fmt.Sprintf("%d (%d * %d)", round * totalHp(us), round, totalHp(us))
}

func Print(g grid.Grid, us units, round int) {
	if round < 1 {
		fmt.Println("Initially:")
	} else {
		fmt.Printf("After %d rounds:\n", round)
	}

	for y := 0; y < len(g[0]); y++ {
		for x := 0; x < len(g); x++ {
			masked := false
			for _, u := range us {
				if (u.Pos() == grid.Coordinate{X:x, Y:y}) {
					fmt.Print(string(u.mark))
					masked = true
					break
				}
			}

			if !masked {
				fmt.Print(g[x][y])
			}
		}
		for _, u := range us {
			if u.pos.Y == y {
				fmt.Printf(" %s(%d)", string(u.mark), u.hp)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
