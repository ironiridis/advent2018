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
	return fmt.Sprintf("[%d:%d]", p.x, p.y)
}

type region struct {
	infinite bool // whether this region has infinite area
	start    pt   // the position where this region starts
	count    int  // the number of points (invalid if infinite is true)
}

func (p *region) String() string {
	if p.infinite {
		return fmt.Sprintf("infinite region at %s", p.start)
	}
	return fmt.Sprintf("bounded region at %s of size %d", p.start, p.count)

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

func (p *pos) closerToAnyRegionThan(p2 *pos) bool {
	p.calc()
	p2.calc()
	for _, r := range regions {
		if p.dists[r] < p2.dists[r] {
			fmt.Printf("%s is closer to %s than %s\n", p.point, r, p2.point)
			return true
		}
	}
	return false
}

func (p *pos) exapandEdges() {
	p.calc()
	if p.nearest != nil {
		if p.nearest.infinite {
			return
		}
	}

	var p2 [4]*pos
	p2[0] = getPos(p.point.x, p.point.y-1)
	p2[1] = getPos(p.point.x, p.point.y+1)
	p2[2] = getPos(p.point.x-1, p.point.y)
	p2[3] = getPos(p.point.x+1, p.point.y)

	if p2[0].dists != nil && p2[1].dists != nil && p2[2].dists != nil && p2[3].dists != nil {
		// already explored all of these points
		return
	}

	if p2[0].dists == nil {
		if p2[0].closerToAnyRegionThan(p) {
			return
		}
	}
	if p2[1].dists == nil {
		if p2[1].closerToAnyRegionThan(p) {
			return
		}
	}
	if p2[2].dists == nil {
		if p2[2].closerToAnyRegionThan(p) {
			return
		}
	}
	if p2[3].dists == nil {
		if p2[3].closerToAnyRegionThan(p) {
			return
		}
	}

	// p is in a region that extends infinitely
	p.nearest.infinite = true
}

func makeRegion(x, y int) *region {
	var r region
	r.start.x = x
	r.start.y = y
	makePos(x, y)
	return &r
}

func makePos(x, y int) *pos {
	var p pos
	p.dists = make(map[*region]int)
	p.nearest = nil
	p.point.x = x
	p.point.y = y
	regionMap[p.point] = &p
	return &p
}

func getPos(x, y int) *pos {
	p, ok := regionMap[pt{x: x, y: y}]
	if !ok {
		return makePos(x, y)
	}
	return p
}

func (p *pos) calc() {
	if p.nearest != nil {
		return
	}

	for _, r := range regions {
		p.dists[r] = p.point.lineardistance(r.start)
		if p.nearest == nil {
			p.nearest = r
		} else if p.dists[r] < p.dists[p.nearest] {
			p.nearest = r
		}
	}
	p.nearest.count++
}

func main() {
	regions = make([]*region, 0, 50)
	regionMap = make(map[pt]*pos)

	for s := range justscan.Chan("../input.txt") {
		var x, y int
		_, err := fmt.Sscanf(s, "%d, %d", &x, &y)
		must("scan coordinates", err)
		regions = append(regions, makeRegion(x, y))
	}

	for i := 0; i < 30; i++ {
		var prevSize int
		prevSize = len(regionMap)
		for _, pos := range regionMap {
			pos.exapandEdges()
		}
		if len(regionMap) == prevSize {
			break
		}
		//fmt.Printf("regionMap size: %d\n", len(regionMap))
	}
	fmt.Printf("regions: %v\n", regions)
	fmt.Printf("regionMap: %v\n", regionMap)

}
