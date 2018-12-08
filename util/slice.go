package util

import (
	"errors"
)

// Interact allows a caller to take an action on an item in a slice if the item exists
func Interact(slice interface{}, item interface{}, act func(int, interface{})) error {
	if s, ok := slice.([]interface{}); ok {
		for i, v := range s {
			if v == item {
				act(i, v)
			}
		}
	} else {
		return errors.New("util: unrecognized slice")
	}
	return nil
}
