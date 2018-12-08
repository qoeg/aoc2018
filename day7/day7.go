package day7

import (
	"fmt"
)

type step struct {
	char string
	pre []*step
	duration int
}

func remove(slice *[]*step, item *step) *step {
	s := *slice
	for i, v := range s {
		if v == item {
			*slice = append(s[:i], s[i+1:]...)
		}
	}
	return item
}

func firstAlpha(steps []*step) *step {
	var step *step
	for _, s := range steps {
		if step == nil {
			step = s
			continue
		}

		if rune(s.char[0]) < rune(step.char[0]) {
			step = s
		}
	}
	return step
}

func parseSteps(input []string, baseDuration int) map[string]*step {
	steps := map[string]*step{}
	
	for _, line := range input {
		var pre, this, z string
		fmt.Sscan(line, &z, &pre, &z, &z, &z, &z, &z, &this, &z, &z)
		
		if (steps[pre] == nil) {
			steps[pre] = &step{
				char: pre,
				pre: []*step{},
				duration: int(pre[0])-64+baseDuration,
			}
		}

		if (steps[this] == nil) {
			steps[this] = &step{
				char: this,
				pre: []*step{},
				duration: int(this[0])-64+baseDuration,
			}
		}
		
		steps[this].pre = append(steps[this].pre, steps[pre])
	}

	return steps
}

func updateReadyState(todos *map[string]*step, ready *[]*step, done []*step) {
	for _, step := range *todos {
		if len(done) > 0 {
			remove(&step.pre, done[len(done)-1])
		}

		if len(step.pre) == 0 {
			*ready = append(*ready, step)
			delete(*todos, step.char)
		}
	}
}

func updateWorkerState(ready, workers, done *[]*step) {
	w := *workers
	for i := range w {
		if w[i] == nil {
			if len(*ready) == 0 {
				continue
			}
			w[i] = remove(ready, firstAlpha(*ready))
		} 

		w[i].duration--
		if w[i].duration == 0 {
			*done = append(*done, w[i])
			w[i] = nil
		}
	}
}

// OrderSteps processes the input steps in the order required by the text.
// The function also allows adding more workers and setting a base duration for each step.
func OrderSteps(input []string, numWorkers, baseDuration int) (string, int) {
	todos := parseSteps(input, baseDuration)
	total := len(todos)
	
	done := []*step{}
	ready := []*step{}
	workers := make([]*step, numWorkers)
	seconds := 0

	for len(done) < total {
		updateReadyState(&todos, &ready, done)
		updateWorkerState(&ready, &workers, &done)
		seconds++
	}

	var result []rune
	for _, s := range done {
		result = append(result, rune(s.char[0]))
	}

	return string(result), seconds
}
