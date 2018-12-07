package day7

import (
	"fmt"
)

type step struct {
	char string
	pre []*step
	duration int
}

func firstAlpha(steps []*step) (int, *step) {
	var idx int
	var step *step
	for i, s := range steps {
		if step == nil {
			step = s
			idx = i
			continue
		}

		if rune(s.char[0]) < rune(step.char[0]) {
			step = s
			idx = i
		}
	}
	return idx, step
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
				pre: []*step{
					steps[pre],
				},
				duration: int(this[0])-64+baseDuration,
			}
			continue
		}
		
		steps[this].pre = append(steps[this].pre, steps[pre])
	}

	return steps
}

// OrderSteps processes the input steps in the order required by the text.
// The function also allows adding more workers and setting a base duration for each step.
func OrderSteps(input []string, numWorkers, baseDuration int) (string, int) {
	steps := parseSteps(input, baseDuration)
	
	done := []*step{}
	ready := []*step{}
	count := len(steps)
	seconds := 0
	workers := make([]*step, numWorkers)

	for len(done) < count {
		// find ready steps
		for _, step := range steps {
			if len(done) > 0 {
				for i, item := range step.pre {
					if item == done[len(done)-1] {
						step.pre = append(step.pre[:i], step.pre[i+1:]...)
					}
				}
			}

			if len(step.pre) == 0 {
				ready = append(ready, step)
				delete(steps, step.char)
			}
		}

		// assign work
		for idx, worker := range workers {
			if len(ready) > 0 && worker == nil {
				i, step := firstAlpha(ready)

				workers[idx] = step
				ready = append(ready[:i], ready[i+1:]...)
			}
		}

		// do work
		for idx, work := range workers {
			if work == nil {
				continue
			}

			work.duration--

			if work.duration == 0 {
				done = append(done, work)
				workers[idx] = nil
			}
		}

		// tick
		seconds++
	}

	var result []rune
	for _, s := range done {
		result = append(result, rune(s.char[0]))
	}

	return string(result), seconds
}
