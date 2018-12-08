package util

import (
	"testing"
)

func TestInteract(t *testing.T) {
	type obj struct {
		id int
	}
	testObj := &obj{10}

	type args struct {
		slice interface{}
		item  interface{}
		act   func(int, interface{})
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "strings",
			args: args{
				slice: []interface{}{"a", "b", "c", "d"},
				item: "b",
				act: func(i int, v interface{}) {
					if s, ok := v.(string); ok {
						if s != "b" {
							t.Errorf("Interact() got %v, wanted %s", s, "b")
						}
					} else {
						t.Errorf("Interact() could not convert argument to string")
					}
				},
			},
		},
		{
			name: "structs",
			args: args{
				slice: []interface{}{&obj{1}, &obj{2}, testObj, &obj{3}},
				item: testObj,
				act: func(i int, v interface{}) {
					if o, ok := v.(*obj); ok {
						if o != testObj {
							t.Errorf("Interact() got %v, wanted %v", o, testObj)
						}
					} else {
						t.Errorf("Interact() could not convert argument to string")
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Interact(tt.args.slice, tt.args.item, tt.args.act); (err != nil) != tt.wantErr {
				t.Errorf("Interact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
