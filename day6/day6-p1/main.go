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
	infinite bool // whether this region has infinite area
	start    pt   // the position where this region starts
	count    int  // the number of points (invalid if infinite is true)
}

func (p *region) String() string {
	if p.infinite {
		return fmt.Sprintf("[region at %s]", p.start)
	}
	return fmt.Sprintf("[region at %s of count %d]", p.start, p.count)

}

type pos struct {
	point   pt
	dists   map[*region]int // distance to each region
	nearest *region         // nearest region
}

func (p *pos) String() string {
	if p.nearest == nil {
		return fmt.Sprintf("~%s", p.point)
	}
	return fmt.Sprintf("%s near %s", p.point, p.nearest)

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

func markInfinite(x, y int) {
	p := regionMap[pt{x: x, y: y}]
	if p.nearest == nil {
		return
	}
	p.nearest.infinite = true
}

func makeRegion(x, y int) *region {
	var r region
	r.start.x = x
	r.start.y = y
	return &r
}

func makePos(x, y int) *pos {
	var p pos
	p.nearest = nil
	p.point.x = x
	p.point.y = y
	return &p
}

func (p *pos) calc() {
	mindist := int(^uint(0) >> 1)
	p.dists = make(map[*region]int)
	for _, r := range regions {
		p.dists[r] = p.point.lineardistance(r.start)
		if p.dists[r] < mindist {
			mindist = p.dists[r]
			p.nearest = r
		}
	}
	// check for multiple regions at same (minimum) distance
	for _, r := range regions {
		if p.dists[r] == mindist && r != p.nearest {
			p.nearest = nil
			break
		}
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
			if p.nearest != nil {
				p.nearest.count++
			}
		}
	}
	for x = minX - 1; x <= maxX+1; x++ {
		markInfinite(x, minY)
		markInfinite(x, maxY)
	}
	for y = minY - 1; y <= maxY+1; y++ {
		markInfinite(minX, y)
		markInfinite(maxX, y)
	}

	var maxsize int
	for _, r := range regions {
		if !r.infinite && r.count > maxsize {
			maxsize = r.count
		}
	}

	fmt.Printf("largest region is %d\n", maxsize)
}
