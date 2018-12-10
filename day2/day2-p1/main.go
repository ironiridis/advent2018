package main

import (
	"fmt"

	"github.com/ironiridis/advent2018/util/justscan"
)

func must(s string, e error) {
	if e != nil {
		fmt.Printf("failed to %s: %s\n", s, e)
		panic(e)
	}
}

func analyze(s string) (rune2, rune3 bool) {
	// a map of all runes that occur in the provided string
	runes := map[rune]int{}

	// scan the string and count every rune
	for _, r := range s {
		runes[r]++
	}

	// scan the map
	for _, v := range runes {
		// if any symbol has exactly two or three instances,
		// set the return values to true
		if v == 2 {
			rune2 = true
		}
		if v == 3 {
			rune3 = true
		}
	}
	return
}

func main() {
	var twos, threes int

	for s := range justscan.Chan("../input.txt") {
		rtwo, rthree := analyze(s)
		if rtwo {
			twos++
		}
		if rthree {
			threes++
		}
	}
	fmt.Printf("checksum of %d*%d: %d\n", twos, threes, twos*threes)
}
