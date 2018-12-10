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
	var freq int
	for s := range justscan.Chan("../input.txt") {
		var delta int
		_, err := fmt.Sscanf(s, "%d", &delta)
		must("scan for frequency delta", err)
		freq += delta
	}
	fmt.Printf("final frequency: %d\n", freq)
}
