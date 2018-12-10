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

func main() {
	type pt struct{ x, y int }
	fab := map[pt]int{}

	for s := range justscan.Chan("../input.txt") {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		must("scan claim string to mark", err)
		for ptx := 0; ptx < w; ptx++ {
			for pty := 0; pty < h; pty++ {
				fab[pt{x: x + ptx, y: y + pty}]++
			}
		}
	}

ClaimCheckRange:
	for s := range justscan.Chan("../input.txt") {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		must("scan claim string to check", err)
		for ptx := 0; ptx < w; ptx++ {
			for pty := 0; pty < h; pty++ {
				if fab[pt{x: x + ptx, y: y + pty}] != 1 {
					continue ClaimCheckRange
				}
			}
		}
		fmt.Printf("claim id with no conflicts: %d\n", id)
	}

}
