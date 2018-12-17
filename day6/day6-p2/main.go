package main

import (
	"fmt"

	"github.com/ironiridis/advent2018/util/justscan"
)

type pt struct {
	x int
	y int
}

func (p pt) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type region struct {
	start pt // the position where this region starts
}

func (p *region) String() string {
	return fmt.Sprintf("[region at %s]", p.start)
}

type pos struct {
	point   pt
	sumdist int
}

func (p *pos) String() string {
	return fmt.Sprintf("~%s", p.point)
}

var regions []*region
var regionMap map[pt]*pos

func must(s string, e error) {
	if e != nil {
		fmt.Printf("failed to %s: %s\n", s, e)
		panic(e)
	}
}

func (p pt) lineardistance(o pt) int {
	var w int
	if p.x < o.x {
		w = o.x - p.x
	} else {
		w = p.x - o.x
	}

	if p.y < o.y {
		return w + (o.y - p.y)
	}
	return w + (p.y - o.y)
}

func makeRegion(x, y int) *region {
	var r region
	r.start.x = x
	r.start.y = y
	return &r
}

func makePos(x, y int) *pos {
	var p pos
	p.point.x = x
	p.point.y = y
	return &p
}

func (p *pos) calc() {
	for _, r := range regions {
		p.sumdist += p.point.lineardistance(r.start)
	}
}

func main() {
	regions = make([]*region, 0, 50)
	regionMap = make(map[pt]*pos)
	var minX, minY, maxX, maxY int
	var x, y int

	for s := range justscan.Chan("../input.txt") {
		_, err := fmt.Sscanf(s, "%d, %d", &x, &y)
		must("scan coordinates", err)
		if minX == 0 && minY == 0 && maxX == 0 && maxY == 0 {
			minX = x
			minY = y
			maxX = x
			maxY = y
		} else {
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
		regions = append(regions, makeRegion(x, y))
	}

	for x = minX - 1; x <= maxX+1; x++ {
		for y = minY - 1; y <= maxY+1; y++ {
			p := makePos(x, y)
			regionMap[pt{x: x, y: y}] = p
			p.calc()
		}
	}

	var poscount int
	for _, p := range regionMap {
		if p.sumdist < 10000 {
			poscount++
		}
	}

	fmt.Printf("positions with a sum distance less than 10000: %d\n", poscount)
}
