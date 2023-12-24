package aoc2023

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Coord3F struct {
	x float64
	y float64
	z float64
}

func (c Coord3F) ToDir() Dir3F {
	return Dir3F{c.x, c.y, c.z}
}

type Dir3F struct {
	dx float64
	dy float64
	dz float64
}

func (d Dir3F) Length() float64 {
	return math.Sqrt(d.dx*d.dx + d.dy*d.dy + d.dz*d.dz)
}

func (d Dir3F) Times(f float64) Dir3F {
	return Dir3F{d.dx * f, d.dy * f, d.dz * f}
}

func (d Dir3F) DivideBy(f float64) Dir3F {
	return Dir3F{d.dx / f, d.dy / f, d.dz / f}
}

func (d0 Dir3F) Plus(d1 Dir3F) Dir3F {
	return Dir3F{d0.dx + d1.dx, d0.dy + d1.dy, d0.dz + d1.dz}
}

func (d0 Dir3F) Minus(d1 Dir3F) Dir3F {
	return Dir3F{d0.dx - d1.dx, d0.dy - d1.dy, d0.dz - d1.dz}
}

func (c Coord3F) Plus(d Dir3F) Coord3F {
	return Coord3F{c.x + d.dx, c.y + d.dy, c.z + d.dz}
}

func (c0 Coord3F) MinusToDir(c1 Coord3F) Dir3F {
	return Dir3F{c0.x - c1.x, c0.y - c1.y, c0.z - c1.z}
}

func (c Coord3F) Minus(d Dir3F) Coord3F {
	return Coord3F{c.x - d.dx, c.y - d.dy, c.z - d.dz}
}

type Coord2F struct {
	x float64
	y float64
}

type HailStone struct {
	pos Coord3F
	dir Dir3F
}

func toFloat(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

var hailRegex = regexp.MustCompile(`(-?\d+),\s+(-?\d+),\s+(-?\d+)\s+@\s+(-?\d+),\s+(-?\d+),\s+(-?\d+)`)

func ReadHail(path string) []HailStone {
	lines := ReadFile(path)

	hails := []HailStone{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		hails = append(hails, parseHail(line))
	}

	return hails
}

func parseHail(str string) HailStone {
	matches := hailRegex.FindStringSubmatch(str)

	pos := Coord3F{toFloat(matches[1]), toFloat(matches[2]), toFloat(matches[3])}
	dir := Dir3F{toFloat(matches[4]), toFloat(matches[5]), toFloat(matches[6])}
	return HailStone{pos, dir}
}

func intersectXY(h1 HailStone, h2 HailStone) (float64, float64, *Coord2F) {
	// intersect the line h1 and h2 - for some t1 ant t2 (float):
	// (h1.pos.x + h1.dir.dx * t1, h1.pos.y + h1.dir.dy * t1) = (h2.pos.x + h2.dir.dx * t2, h2.pos.y + h2.dir.dy * t2)
	// i.e.,
	// h1.pos.x + h1.dir.dx * t1 = h2.pos.x + h2.dir.dx * t2 AND
	// h1.pos.y + h1.dir.dy * t1 = h2.pos.y + h2.dir.dy * t2

	// t1 = (h2.pos.x - h1.pos.x + h2.dir.dx * t2) / h1.dir.dx AND
	// t1 = (h2.pos.y - h1.pos.y + h2.dir.dy * t2) / h1.dir.dy
	// hence
	// (h2.pos.x - h1.pos.x + h2.dir.dx * t2) / h1.dir.dx =
	//       (h2.pos.y - h1.pos.y + h2.dir.dy * t2) / h1.dir.dy
	// -> (mutiply by h1.dir.dx * h1.dir.dy on both sides)
	// (h2.pos.x - h1.pos.x + h2.dir.dx * t2) * h1.dir.dy =
	//       (h2.pos.y - h1.pos.y + h2.dir.dy * t2) * h1.dir.dx
	// -> (expand brackets)
	// (h2.pos.x - h1.pos.x) * h1.dir.dy + h2.dir.dx * h1.dir.dy * t2 =
	//       (h2.pos.y - h1.pos.y) * h1.dir.dx + h2.dir.dy * h1.dir.dx * t2
	// -> (t2 to left, other things to right)
	// h2.dir.dx * h1.dir.dy * t2 - h2.dir.dy * h1.dir.dx * t2 =
	//       (h2.pos.y - h1.pos.y) * h1.dir.dx - (h2.pos.x - h1.pos.x) * h1.dir.dy
	// -> (divide both sides by the t2 multiplier)
	// t2 = (h2.pos.y - h1.pos.y) * h1.dir.dx - (h2.pos.x - h1.pos.x) * h1.dir.dy /
	//         (h2.dir.dx * h1.dir.dy - h2.dir.dy * h1.dir.dx)
	denom := h2.dir.dx*h1.dir.dy - h2.dir.dy*h1.dir.dx
	if denom == 0 {
		// parallel
		return 0, 0, nil
	}
	t2 := ((h2.pos.y-h1.pos.y)*h1.dir.dx - (h2.pos.x-h1.pos.x)*h1.dir.dy) / denom
	t1 := (h2.pos.x - h1.pos.x + h2.dir.dx*t2) / h1.dir.dx
	if t1 < 0.0 || t2 < 0.0 {
		// crossed in the past
		return 0, 0, nil
	}
	return t1, t2, &Coord2F{float64(h2.pos.x) + t2*float64(h2.dir.dx), float64(h2.pos.y) + t2*float64(h2.dir.dy)}
}

func CountIntersectionPairs(hails []HailStone, min float64, max float64) int {
	count := 0

	for i := range hails {
		for j := i + 1; j < len(hails); j++ {
			h1 := hails[i]
			h2 := hails[j]

			_, _, c := intersectXY(h1, h2)
			if c != nil && c.x >= min && c.x <= max && c.y >= min && c.y <= max {
				count++
			}
		}
	}

	return count
}

func FindIntersectionPairsXY(hails []HailStone, min float64, max float64) (HailStone, HailStone, HailStone, HailStone) {
	found := []HailStone{}

	for i := range hails {
		for j := i + 1; j < len(hails); j++ {
			h1 := hails[i]
			h2 := hails[j]

			_, _, c := intersectXY(h1, h2)
			if c != nil && c.x >= min && c.x <= max && c.y >= min && c.y <= max {
				found = append(found, h1, h2)
			}

			if len(found) == 4 {
				return found[0], found[1], found[2], found[3]
			}
		}
	}

	panic("need two intersecting pairs in XY")
}

func crossProduct(d0, d1 Dir3F) Dir3F {
	return Dir3F{d0.dy*d1.dz - d0.dz*d1.dy, d0.dz*d1.dx - d0.dx*d1.dz, d0.dx*d1.dy - d0.dy*d1.dx}
}

func dotProduct(d0, d1 Dir3F) float64 {
	return d0.dx*d1.dx + d0.dy*d1.dy + d0.dz*d1.dz
}

func isClose(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}

func findParallelPair(hails []HailStone) (HailStone, HailStone) {
	for i := range hails {
		for j := i + 1; j < len(hails); j++ {
			h0 := hails[i]
			l0 := h0.dir.Length()
			h1 := hails[j]
			l1 := h1.dir.Length()
			if isClose(h0.dir.dx/l0, h1.dir.dx/l1) &&
				isClose(h0.dir.dy/l0, h1.dir.dy/l1) &&
				isClose(h0.dir.dz/l0, h1.dir.dz/l1) {
				return h0, h1
			}
		}
	}
	panic("no parallel pair")
}

func findTwoOtherStones(h0, h1 HailStone, n Dir3F, hails []HailStone) (HailStone, HailStone) {
	found := []HailStone{}

	for _, h := range hails {
		if h != h0 && h != h1 && dotProduct(h.dir, n) != 0 {
			found = append(found, h)
		}
		if len(found) == 2 {
			return found[0], found[1]
		}
	}
	panic("too few hail stones")
}

// intersect the plane defined by the point p0 and its normal vector n
// and the line defined by h
// See https://en.wikipedia.org/wiki/Line%E2%80%93plane_intersection
func intersectPlaneAndLine(p0 Coord3F, n Dir3F, h HailStone) (Coord3F, float64) {
	a := dotProduct(p0.MinusToDir(h.pos), n)
	b := dotProduct(h.dir, n)
	t := a / b
	fmt.Println("h", h, "a", a, "b", b, "t", t)
	p := h.pos.Plus(h.dir.Times(t))
	return p, t
}

func ComputeIntersectingHailStone(hails []HailStone, min, max float64) HailStone {
	// translate everything into a system with hails[0] as the frame of reference
	h0shift := hails[0].pos.ToDir()
	h0dir := hails[0].dir
	hailsR := make([]HailStone, len(hails))
	for i, h := range hails {
		hailsR[i] = HailStone{h.pos.Minus(h0shift), h.dir.Minus(h0dir)}
	}

	// in hailsR, the first hail stone is at (0,0,0) and does not move, so the rock needs
	// to pass through (0,0,0)
	// hailsR[1] follows a trajectory that the rock needs to intersect with. This means that
	// it intersects somewhere in the plane defined by the origin and any 2 points on that line
	// we can compute the unit vector of that plane by taking the vectors from the origin to t=0 and t=1
	// after that, we compute two intersections for hailsR[2] and hailsR[3] and work out dir and p0 from
	// there
	h1 := hailsR[1]
	p10 := h1.pos
	p11 := h1.pos.Plus(h1.dir)
	n := crossProduct(p10.ToDir(), p11.ToDir())
	fmt.Println("h0", hailsR[0], "h1", h1)
	fmt.Println("p10", p10, "p11", p11)
	fmt.Println("n", n)

	h2 := hailsR[2]
	h3 := hailsR[3]
	fmt.Println("h2", h2, "h3", h3)

	p0 := Coord3F{0, 0, 0}
	p2, t2 := intersectPlaneAndLine(p0, n, h2)
	fmt.Println("p2", p2, "t2", t2)
	p3, t3 := intersectPlaneAndLine(p0, n, h3)
	fmt.Println("p3", p3, "t3", t3)

	fmt.Println("t2-t3", t2-t3)
	dir := p2.MinusToDir(p3).DivideBy(t2 - t3)
	fmt.Println("dir before round", dir)
	fmt.Println("t2 diff", t2-math.Round(t2))
	fmt.Println("t3 diff", t3-math.Round(t3))
	// fix some floating point errors by rounding
	dir = Dir3F{math.Round(dir.dx), math.Round(dir.dy), math.Round(dir.dz)}
	//	t2 = math.Round(t2)
	pos := p2.Minus(dir.Times(t2))
	fmt.Println("dir", dir, "pos", pos)

	res := HailStone{pos.Plus(h0shift), dir.Plus(h0dir)}
	fmt.Println("res", res)
	return res
}
