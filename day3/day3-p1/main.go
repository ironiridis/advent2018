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
		must("scan claim string", err)
		for ptx := 0; ptx < w; ptx++ {
			for pty := 0; pty < h; pty++ {
				fab[pt{x: x + ptx, y: y + pty}]++
			}
		}
	}

	var overcommit int
	for _, tileClaimed := range fab {
		if tileClaimed > 1 {
			overcommit++
		}
	}

	fmt.Printf("fabric overcommit is: %d\n", overcommit)
}
