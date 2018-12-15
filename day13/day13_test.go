package day13

import (
	"testing"

	"github.com/qoeg/aoc2018/util"
)

var input = 
`/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

var input2 = 
`/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`

func Test_Input(t *testing.T) {
	g := NewGrid(input)
	g.Render(true)

	for g.Move() {
		g.Render(true)
	}
	g.Render(true)

	if len(g.Crashes) != 1 {
		t.FailNow()
	}

	expected := util.Coordinate{X:7, Y:3}
	if (g.Crashes[0] != expected) {
		t.Errorf("Expected %v, got %v", expected, g.Crashes[0])
	}
}

func Test_Input2(t *testing.T) {
	g := NewGrid(input2)
	g.Render(false)

	for len(g.Carts) > 1 {
		g.Move()
		g.Render(false)
	}

	if len(g.Carts) != 1 {
		t.FailNow()
	}

	expected := util.Coordinate{X:6, Y:4}
	if (g.Carts[0].Pos != expected) {
		t.Errorf("Expected %v, got %v", expected, g.Carts[0].Pos)
	}
}
