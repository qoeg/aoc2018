package day15

import (
	"reflect"
	"testing"

	"github.com/qoeg/aoc2018/grid"
)

func Test_unit_bestPath(t *testing.T) {
	var input = 
`#######
#E..G.#
#...#.#
#.G.#G#
#######`
	var units []*unit
	g := grid.ParseOut(input, replaceFor(&units))
	
	type args struct {
		cells   [][]grid.Cell
		units   []*unit
		targets []*unit
	}
	tests := []struct {
		name string
		u    *unit
		args args
		want grid.Coordinate
	}{
		{
			name: "1",
			u: units[0],
			args: args{
				cells: g,
				units: units,
				targets: units[0].enemies(units),
			},
			want: grid.Coordinate{X:2, Y:1},
		},
		{
			name: "2",
			u: units[1],
			args: args{
				cells: g,
				units: units,
				targets: units[1].enemies(units),
			},
			want: grid.Coordinate{X:3, Y:1},
		},
		{
			name: "3",
			u: units[2],
			args: args{
				cells: g,
				units: units,
				targets: units[2].enemies(units),
			},
			want: grid.Coordinate{X:2, Y:2},
		},
		{
			name: "4",
			u: units[3],
			args: args{
				cells: g,
				units: units,
				targets: units[3].enemies(units),
			},
			want: grid.Coordinate{X:0, Y:0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.u.nextStep(tt.args.cells, tt.args.units, tt.args.targets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unit.nextStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
