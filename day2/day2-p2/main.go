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

func isMatch(ids []string, id string) *string {
MapRange:
	for k := range ids {
		// ids must be the same length to be comparable
		if len(ids[k]) != len(id) {
			continue
		}
		m := false
		// compare each byte in the new id to the ids map
		for p := range id {
			// if this byte differs...
			if id[p] != ids[k][p] {
				// and we've already seen a differing byte...
				if m {
					// then give up; this differs too much.
					continue MapRange
				}
				// otherwise
				m = true
			}

		}
		if m {
			r := common(ids[k], id)
			return &r
		}
	}
	return nil
}

func common(a, b string) (r string) {
	if len(a) != len(b) {
		panic("unequal id lengths")
	}
	for p := range a {
		if a[p] == b[p] {
			r += string(a[p])
		}
	}
	return
}

func main() {
	ids := []string{}

	for s := range justscan.Chan("../input.txt") {
		if matchedID := isMatch(ids, s); matchedID != nil {
			fmt.Printf("found common id string: %s\n", *matchedID)
			return
		}
		ids = append(ids, s)
	}

	panic("didn't find a common id string")
}
