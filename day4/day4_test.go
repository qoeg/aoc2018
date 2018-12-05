package day4

import (
	"testing"
)

var input = []string{
	"[1518-11-01 00:00] Guard #10 begins shift",
	"[1518-11-01 00:05] falls asleep",
	"[1518-11-01 00:25] wakes up",
	"[1518-11-01 00:30] falls asleep",
	"[1518-11-01 00:55] wakes up",
	"[1518-11-01 23:58] Guard #99 begins shift",
	"[1518-11-02 00:40] falls asleep",
	"[1518-11-02 00:50] wakes up",
	"[1518-11-03 00:05] Guard #10 begins shift",
	"[1518-11-03 00:24] falls asleep",
	"[1518-11-03 00:29] wakes up",
	"[1518-11-04 00:02] Guard #99 begins shift",
	"[1518-11-04 00:36] falls asleep",
	"[1518-11-04 00:46] wakes up",
	"[1518-11-05 00:03] Guard #99 begins shift",
	"[1518-11-05 00:45] falls asleep",
	"[1518-11-05 00:55] wakes up",
}

func TestParseEvent(t *testing.T) {
	e := parseEvent("[1518-09-20 00:43] falls asleep")

	expE := "1518-09-20 00:43:00 +0000 UTC"
	if e.Timestamp.String() != expE {
		t.Logf("Expected %s, but got %s", expE, e.Timestamp)
		t.Fail()
	}

	expD := "falls asleep"
	if e.Description != expD {
		t.Logf("Expected %s, but got %s", expD, e.Description)
		t.Fail()
	}
}

func TestGetGuardSchedules(t *testing.T) {
	guards := GetGuardSchedules(input)
	if guards == nil {
		t.FailNow()
	}

	// for _, g := range guards {
	// 	fmt.Printf("Guard #%d\n", g.ID)

	// 	for n, m := range g.MinutesAsleep {
	// 		fmt.Printf("%d-%d: ", n.month, n.day)

	// 		sort.Ints(m)
	// 		fmt.Printf("%v\n", m)
	// 	}
	// }
}

func TestMostSleepy(t *testing.T) {
	guards := GetGuardSchedules(input)
	guard := MostSleepy(guards)

	if guard.ID != 10 {
		t.Logf("Expected %d, but got %d", 10, guard.ID)
		t.Fail()
	}
}

func TestMostCommonMinute(t *testing.T) {
	guards := GetGuardSchedules(input)
	guard := MostSleepy(guards)
	minute, _ := MostCommonMinute(guard)

	if minute != 24 {
		t.Logf("Expected %d, but got %d", 24, minute)
		t.Fail()
	}
}
