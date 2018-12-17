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

type dep struct {
	pred byte
	step byte
}

type deptrack map[dep]bool
type depset map[byte]bool

func findPred(s byte, dt *deptrack) bool {
	for p := byte(0x41); p <= 0x5A; p++ {
		if (*dt)[dep{pred: p, step: s}] {
			return true
		}
	}
	return false
}

func deletePred(p byte, dt *deptrack) {
	for s := byte(0x41); s <= 0x5A; s++ {
		delete(*dt, dep{pred: p, step: s})
	}

}

func main() {
	deps := deptrack{}
	want := depset{}

	for s := range justscan.Chan("../input.txt") {
		var a, b byte
		_, err := fmt.Sscanf(s, "Step %c must be finished before step %c can begin.", &a, &b)
		must("scan rule dependency", err)
		deps[dep{pred: a, step: b}] = true
		want[a] = true
		want[b] = true
	}

	for len(want) > 0 {
		for step := byte(0x41); step <= 0x5A; step++ {
			if want[step] && !findPred(step, &deps) {
				fmt.Printf("%c", step)
				delete(want, step)
				deletePred(step, &deps)
				break
			}
		}
	}

	fmt.Println()
}
