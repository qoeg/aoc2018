package grid

import (
	"reflect"
	"testing"
)

var input = 
`#######
#E..G.#
#...#.#
#.G.#G#
#######`

func TestGrid_Neighbors(t *testing.T) {
	grid := Parse(input)

	type args struct {
		p       TwoDimensional
		include []rune
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want []Cell
	}{
		{
			name: "Corner",
			g: grid,
			args: args{Cell{pos:Coordinate{X:0, Y:0}}, []rune{'.'}},
			want: []Cell{},
		},
		{
			name: "Partial",
			g: grid,
			args: args{Cell{pos:Coordinate{X:4, Y:1}}, []rune{'.'}},
			want: []Cell{
				Cell{mark:'.', pos:Coordinate{X:3, Y:1}},
				Cell{mark:'.', pos:Coordinate{X:5, Y:1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Neighbors(tt.args.p, tt.args.include...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.Neighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}
