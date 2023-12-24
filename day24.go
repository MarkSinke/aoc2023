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

func intersectXY(h1 HailStone, h2 HailStone) *Coord2F {
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
		return nil
	}
	t2 := ((h2.pos.y-h1.pos.y)*h1.dir.dx - (h2.pos.x-h1.pos.x)*h1.dir.dy) / denom
	t1 := (h2.pos.x - h1.pos.x + h2.dir.dx*t2) / h1.dir.dx
	if t1 < 0.0 || t2 < 0.0 {
		// crossed in the past
		return nil
	}
	return &Coord2F{float64(h2.pos.x) + t2*float64(h2.dir.dx), float64(h2.pos.y) + t2*float64(h2.dir.dy)}
}

func CountIntersectionPairs(hails []HailStone, min float64, max float64) int {
	count := 0

	for i := range hails {
		for j := i + 1; j < len(hails); j++ {
			h1 := hails[i]
			h2 := hails[j]

			c := intersectXY(h1, h2)
			if c != nil && c.x >= min && c.x <= max && c.y >= min && c.y <= max {
				count++
			}
		}
	}

	return count
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

func ComputeIntersectingHailStone(hails []HailStone) HailStone {
	// if the new stone h intersects with h0, h1, h2, ... it intersects with each of the paths:
	// FORALL i in N : h.pos + h.dir * t_i = h_i.pos + t_i * h_i.dir
	// This means it needs to hit the first two hailstones in the plane defined by their paths (i.e.,
	// the plane that goes through both paths). To make sure we have a plane P that holds both lines,
	// we look for a pair where two paths are parallel, say h0 and h1.
	// A 3rd hail stone h2 then hits this plane P somewhere and that
	// is where h intersects with h2. We can compute that intersection point and timestamp.
	// If we do the same for another hail stone h3, we have two positions and timestamps,
	// and hence we can compute h.pos and h.dir.
	h0, h1 := findParallelPair(hails)
	fmt.Println("h0", h0, "h1", h1)
	n := crossProduct(h0.dir, h1.pos.MinusToDir(h0.pos)) // normal of the intersection plane
	//n = n.DivideBy(n.Length())
	fmt.Println("n", n)
	p0 := h0.pos // a point in the intersection plane
	fmt.Println("p0", p0)
	h2, h3 := findTwoOtherStones(h0, h1, n, hails)
	fmt.Println("h2", h2, "h3", h3)

	fmt.Println("dp2", dotProduct(h2.dir, n))
	fmt.Println("dp3", dotProduct(h3.dir, n))
	p2, t2 := intersectPlaneAndLine(p0, n, h2)
	fmt.Println("p2", p2, "t2", t2)
	p3, t3 := intersectPlaneAndLine(p0, n, h3)
	fmt.Println("p3", p3, "t3", t3)

	// direction vector travels from p2 to p3 in t2-t3 time
	dir := (p2.MinusToDir(p3)).DivideBy(t2 - t3)
	fmt.Println("dir", dir)
	p := p2.Minus(dir.Times(t2))
	fmt.Println("p", p)

	fmt.Println("intersect h2", h2.pos.Plus(h2.dir.Times(t2)), p.Plus(dir.Times(t2)))
	fmt.Println("intersect h3", h3.pos.Plus(h3.dir.Times(t3)), p.Plus(dir.Times(t3)))

	return HailStone{p, dir}
	//
	// In this equations, h.pos (3 unknowns), h.dir (3 unknowns), and FORALL i: t_i are unknown
	// each of the above equations effectively turns into 3 equations
	// so we have 3N equations with 6 + N unknowns. This is solvable for N >= 3.
	// solving for h.pos and h.dir hence involves solving a system of equations of size at least 3N=21.

	// h.pos + h.dir * t0 = h0.pos + h0.dir * t0
	// h.pos + h.dir * t1 = h1.pos + h1.dir * t1
	// h.pos + h.dir * t2 = h2.pos + h2.dir * t2
	//
	// writing out the equations yields:
	// h.pos.x + h.dir.dx * t0 = h0.pos.x + h0.dir.dx * t0
	// h.pos.y + h.dir.dy * t0 = h0.pos.y + h0.dir.dy * t0
	// h.pos.z + h.dir.dz * t0 = h0.pos.z + h0.dir.dz * t0
	//
	// h.pos.x + h.dir.dx * t1 = h1.pos.x + h1.dir.dx * t1
	// h.pos.y + h.dir.dy * t1 = h1.pos.y + h1.dir.dy * t1
	// h.pos.z + h.dir.dz * t1 = h1.pos.z + h1.dir.dz * t1
	//
	// h.pos.x + h.dir.dx * t2 = h2.pos.x + h2.dir.dx * t2
	// h.pos.y + h.dir.dy * t2 = h2.pos.y + h2.dir.dy * t2
	// h.pos.z + h.dir.dz * t2 = h2.pos.z + h2.dir.dz * t2
	//
	// Or,
	// (h.dir.dx - h0.dir.dx) * t0 = h0.pos.x - h.pos.x
	// (h.dir.dy - h0.dir.dy) * t0 = h0.pos.y - h.pos.y
	// (h.dir.dz - h0.dir.dz) * t0 = h0.pos.z - h.pos.z
	// [...]
	// ->
	// t0 = (h0.pos.x - h.pos.x) / (h.dir.dx - h0.dir.dx)
	// t0 = (h0.pos.y - h.pos.y) / (h.dir.dy - h0.dir.dy)
	// t0 = (h0.pos.z - h.pos.z) / (h.dir.dz - h0.dir.dz)
	// [...]
	// -> (eliminate t0, t1, t2)
	// (h0.pos.x - h.pos.x) / (h.dir.dx - h0.dir.dx) = (h0.pos.y - h.pos.y) / (h.dir.dy - h0.dir.dy)
	// (h0.pos.y - h.pos.y) / (h.dir.dy - h0.dir.dy) = (h0.pos.z - h.pos.z) / (h.dir.dz - h0.dir.dz)
	// (h0.pos.z - h.pos.z) / (h.dir.dz - h0.dir.dz) = (h0.pos.x - h.pos.x) / (h.dir.dx - h0.dir.dx)
	// [...]
	// -> (multiply both sides to get rid of denominators)
	// (h0.pos.x - h.pos.x) * (h.dir.dy - h0.dir.dy) = (h0.pos.y - h.pos.y) * (h.dir.dx - h0.dir.dx)
	// (h0.pos.y - h.pos.y) * (h.dir.dz - h0.dir.dz) = (h0.pos.z - h.pos.z) * (h.dir.dy - h0.dir.dy)
	// (h0.pos.z - h.pos.z) * (h.dir.dx - h0.dir.dx) = (h0.pos.x - h.pos.x) * (h.dir.dz - h0.dir.dz)
	// [...]
	// h0.pos.x * h.dir.dy - h0.pos.x * h0.dir.dy - h.pos.x * h.dir.dy + h.pos.x * h0.dir.dy
	//
	// Consider these equations in a Newtonian system S moving with constant velocity h.dir.
	// Then we are looking for the first 3 hail stone paths to travel through a single point h.pos.
	// I.e. the intersection point of the first 3 lines in S. From there we can compute back.
	// h.pos = (h0.dir - h.dir) * t0 + h0.pos
	// h.pos = (h1.dir - h.dir) * t1 + h1.pos
	// h.pos = (h2.dir - h.dir) * t2 + h2.pos

	// We can compute the intersection point h.pos from the first 2 lines.
	// Consider the plane rooted at h0.pos, extending in all directions orthogonal to h0.dir. Similarly,
	// for h1, and h2. Then the intersection point P is the point where
	// d0' . h0.dir = 0
	// h0.pos + d0' = P
	// d1' . h1.dir = 0
	// h1.pos + d1' = P
	// d2' . h2.dir = 0
	// h2.pos + d2' = P =>
	//
	// d0'.dx * h0.dir.dx + d0'.dy * h0.dir.dy + d2' * h0.dir.dz = 0
	// h0.pos.x + d0'.dx = P.x
	// h0.pos.y + d0'.dy = P.y
	// h0.pos.z + d0'.dz = P.z
	//
	// d1'.dx * h1.dir.dx + d1'.dy * h1.dir.dy + d2' * h1.dir.dz = 0
	// h1.pos.x + d1'.dx = P.x
	// h1.pos.y + d1'.dy = P.y
	// h1.pos.z + d1'.dz = P.z
	//
	// d2'.dx * h2.dir.dx + d2'.dy * h2.dir.dy + d2' * h2.dir.dz = 0
	// h2.pos.x + d2'.dx = P.x
	// h2.pos.y + d2'.dy = P.y
	// h2.pos.z + d2'.dz = P.z

}
