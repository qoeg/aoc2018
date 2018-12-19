package day15

import (
	"testing"

	"github.com/qoeg/aoc2018/grid"
)

var test = 
`#######
#E..G.#
#...#.#
#.G.#G#
#######`

func Test_Enemies(t *testing.T) {
	units := []*unit{}
	grid.ParseOut(test, replaceFor(&units))
}
