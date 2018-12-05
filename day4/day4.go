package day4

import (
	"fmt"
	"sort"
	"time"
)

type event struct {
	Timestamp time.Time
	Description string
}

type night struct {
	month int
	day int
}

func parseEvent(e string) event {
	var year, month, day, hour, minute int
	fmt.Sscanf(e, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)

	return event{
		Timestamp: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
		Description: e[19:],
	}
}

// Guard represents an elf who sleeps on the job
type Guard struct {
	ID int
	MinutesAsleep map[night][]int
}

// MostSleepy finds the Guard that has slept the most amount of minutes total
func MostSleepy(guards map[int]*Guard) *Guard {
	var max int
	var guard *Guard
	
	for _, g := range guards {
		count := 0
		for _, mins := range g.MinutesAsleep {
			count += len(mins)
		}

		if count > max {
			guard = g
			max = count
		}
	}

	return guard
}

// MostCommonMinute finds the minute a specific Guard sleeps during the most
func MostCommonMinute(guard *Guard) (int, int) {
	minMap := map[int]int{}
	for _, mins := range guard.MinutesAsleep {
		for _, m := range mins {
			minMap[m]++
		}
	}

	max := 0
	minute := 0
	for min, count := range minMap {
		if count > max {
			minute = min
			max = count
		}
	}

	return minute, max
}

// GetGuardSchedules gets a structured map of Guards from a list of event strings
func GetGuardSchedules(list []string) map[int]*Guard {
	l := make([]string, len(list))
	copy(l, list)

	sort.Sort(sort.StringSlice(l))

	guards := map[int]*Guard{}

	var guard *Guard
	var asleepAt time.Time
	for _, item := range l {
		e := parseEvent(item)

		var first, second string
		fmt.Sscan(e.Description, &first, &second)

		if first == "Guard" {
			var id int
			fmt.Sscanf(second, "#%d", &id)

			guard = guards[id]
			if guard == nil {
				guard = &Guard{
					ID: id,
					MinutesAsleep: map[night][]int{},
				}
				guards[id] = guard
			}
			continue
		}

		if first == "falls" {
			asleepAt = e.Timestamp
			continue
		}

		if first == "wakes" {
			n := night{
				month: int(e.Timestamp.Month()),
				day: int(e.Timestamp.Day()),
			}

			if guard.MinutesAsleep[n] == nil {
				guard.MinutesAsleep[n] = []int{}
			}

			for m := asleepAt.Minute(); m < e.Timestamp.Minute(); m++ {
				guard.MinutesAsleep[n] = append(guard.MinutesAsleep[n], m)
			}
		}
	}
	
	return guards
}
