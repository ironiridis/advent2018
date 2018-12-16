package main

import (
	"fmt"

	"github.com/ironiridis/advent2018/util/justscan"
)

// ref http://5ko.free.fr/en/jul.php?y=1518
var lastMonthDay = map[uint]uint{
	1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31,
}

func must(s string, e error) {
	if e != nil {
		fmt.Printf("failed to %s: %s\n", s, e)
		panic(e)
	}
}

type ts struct {
	month  uint
	day    uint
	hour   uint
	minute uint
}

type guardminute struct {
	guard  uint
	minute uint
}

func (e *ts) inc() bool {
	if e.minute < 59 {
		e.minute++
	} else if e.hour < 23 {
		e.minute = 0
		e.hour++
	} else if e.day < lastMonthDay[e.month] {
		e.minute = 0
		e.hour = 0
		e.day++
	} else {
		e.minute = 0
		e.hour = 0
		e.day = 1
		e.month++
	}
	return (e.hour != 1)
}

func main() {
	var guard uint
	var evts ts

	logGuardStart := map[ts]uint{}
	logGuardSleep := map[ts]bool{}
	logGuardWake := map[ts]bool{}

	// ingest input data and parse into timestamp-based structures
	for s := range justscan.Chan("../input.txt") {
		var kind string

		_, err := fmt.Sscanf(s, "[1518-%d-%d %d:%d] %s", &evts.month, &evts.day, &evts.hour, &evts.minute, &kind)
		must("scan event timestamp", err)
		switch kind {
		case "Guard":
			_, err := fmt.Sscanf(s[19:], "Guard #%d begins shift", &guard)
			must("scan guard id", err)
			logGuardStart[evts] = guard
		case "wakes":
			logGuardWake[evts] = true
		case "falls":
			logGuardSleep[evts] = true
		}
	}

	// scan logs to tabulate sleeping time per guard
	var longestSleepTime int
	var longestSleepGuard uint
	guardSleepTime := map[uint]int{}
	guardMinuteCount := map[guardminute]int{}
	for evts, guard = range logGuardStart {
		sleeping := false
		for evts.inc() {
			// new guard starts
			if logGuardStart[evts] != 0 {
				break
			}
			if logGuardSleep[evts] {
				sleeping = true
			}
			if logGuardWake[evts] {
				sleeping = false
			}
			if sleeping {
				guardSleepTime[guard]++
				// per-minute tabulation only applies during midnight hour
				// shouldn't matter, since sleep/wake events only appear during midnight
				if evts.hour == 0 {
					guardMinuteCount[guardminute{guard: guard, minute: evts.minute}]++
				}
				if guardSleepTime[guard] > longestSleepTime {
					longestSleepTime = guardSleepTime[guard]
					longestSleepGuard = guard
				}
			}
		}
	}

	// find minute during which longestSleepGuard was asleep most
	var guardMinuteMostMinute uint
	var guardMinuteMostDays int
	for gm := range guardMinuteCount {
		// only consider the guard with the most overall sleep minutes
		if gm.guard != longestSleepGuard {
			continue
		}
		if guardMinuteCount[gm] > guardMinuteMostDays {
			guardMinuteMostMinute = gm.minute
			guardMinuteMostDays = guardMinuteCount[gm]
		}
	}

	fmt.Printf("longest sleeping guard is %d with %d minutes\n", longestSleepGuard, longestSleepTime)
	fmt.Printf("slept most often at 00:%02d (%d days)\n", guardMinuteMostMinute, guardMinuteMostDays)
	fmt.Printf("challenge response should be %d\n", longestSleepGuard*guardMinuteMostMinute)
}
