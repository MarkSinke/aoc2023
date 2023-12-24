package aoc2023

import (
	"math/big"
	"regexp"
)

type Coord3F struct {
	x *big.Int
	y *big.Int
	z *big.Int
}

func (c Coord3F) String() string {
	return "{" + c.x.String() + "," + c.y.String() + "," + c.z.String() + "}"
}

func (c Coord3F) ToDir() Dir3F {
	return Dir3F{c.x, c.y, c.z}
}

type Dir3F struct {
	dx *big.Int
	dy *big.Int
	dz *big.Int
}

func (d Dir3F) String() string {
	return "{" + d.dx.String() + "," + d.dy.String() + "," + d.dz.String() + "}"
}

func (d Dir3F) Times(f *big.Int) Dir3F {
	var dx big.Int
	dx.Mul(d.dx, f)
	var dy big.Int
	dy.Mul(d.dy, f)
	var dz big.Int
	dz.Mul(d.dz, f)
	return Dir3F{&dx, &dy, &dz}
}

func (d Dir3F) DivideBy(f *big.Int) Dir3F {
	var dx big.Int
	dx.Div(d.dx, f)
	var dy big.Int
	dy.Div(d.dy, f)
	var dz big.Int
	dz.Div(d.dz, f)
	return Dir3F{&dx, &dy, &dz}
}

func (d0 Dir3F) Plus(d1 Dir3F) Dir3F {
	var dx big.Int
	dx.Add(d0.dx, d1.dx)
	var dy big.Int
	dy.Add(d0.dy, d1.dy)
	var dz big.Int
	dz.Add(d0.dz, d1.dz)
	return Dir3F{&dx, &dy, &dz}
}

func (d0 Dir3F) Minus(d1 Dir3F) Dir3F {
	var dx big.Int
	dx.Sub(d0.dx, d1.dx)
	var dy big.Int
	dy.Sub(d0.dy, d1.dy)
	var dz big.Int
	dz.Sub(d0.dz, d1.dz)
	return Dir3F{&dx, &dy, &dz}
}

func (c Coord3F) Plus(d Dir3F) Coord3F {
	var x big.Int
	x.Add(c.x, d.dx)
	var y big.Int
	y.Add(c.y, d.dy)
	var z big.Int
	z.Add(c.z, d.dz)
	return Coord3F{&x, &y, &z}
}

func (c Coord3F) Minus(d Dir3F) Coord3F {
	var x big.Int
	x.Sub(c.x, d.dx)
	var y big.Int
	y.Sub(c.y, d.dy)
	var z big.Int
	z.Sub(c.z, d.dz)
	return Coord3F{&x, &y, &z}
}

func (c0 Coord3F) MinusToDir(c1 Coord3F) Dir3F {
	var x big.Int
	x.Sub(c0.x, c1.x)
	var y big.Int
	y.Sub(c0.y, c1.y)
	var z big.Int
	z.Sub(c0.z, c1.z)
	return Dir3F{&x, &y, &z}
}

type Coord2F struct {
	x *big.Int
	y *big.Int
}

type HailStone struct {
	pos Coord3F
	dir Dir3F
}

func (h HailStone) String() string {
	return "{" + h.pos.String() + " " + h.dir.String() + "}"
}

func toBigInt(str string) *big.Int {
	var x big.Int
	x.SetString(str, 10)
	return &x
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

	pos := Coord3F{toBigInt(matches[1]), toBigInt(matches[2]), toBigInt(matches[3])}
	dir := Dir3F{toBigInt(matches[4]), toBigInt(matches[5]), toBigInt(matches[6])}
	return HailStone{pos, dir}
}

var zero *big.Int = big.NewInt(0)

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
	//
	// denom := h2.dir.dx * h1.dir.dy - h2.dir.dy * h1.dir.dx
	var denom1 big.Int
	denom1.Mul(h2.dir.dx, h1.dir.dy)
	var denom2 big.Int
	denom2.Mul(h2.dir.dy, h1.dir.dx)
	var denom big.Int
	denom.Sub(&denom1, &denom2)
	if denom.Cmp(zero) == 0 {
		// parallel
		return nil
	}
	// t2 := ((h2.pos.y-h1.pos.y)*h1.dir.dx - (h2.pos.x-h1.pos.x)*h1.dir.dy) / denom
	var t2a big.Int
	t2a.Sub(h2.pos.y, h1.pos.y)
	t2a.Mul(&t2a, h1.dir.dx)
	var t2b big.Int
	t2b.Sub(h2.pos.x, h1.pos.x)
	t2b.Mul(&t2b, h1.dir.dy)
	var t2 big.Int
	t2.Sub(&t2a, &t2b)
	t2.Div(&t2, &denom)

	// t1 := (h2.pos.x - h1.pos.x + h2.dir.dx*t2) / h1.dir.dx
	var t1 big.Int
	t1.Mul(h2.dir.dx, &t2)
	t1.Add(&t1, h2.pos.x)
	t1.Sub(&t1, h1.pos.x)
	t1.Sub(&t1, h1.dir.dx)
	if t1.Cmp(zero) < 0 || t2.Cmp(zero) < 0 {
		// crossed in the past
		return nil
	}

	var x big.Int // h2.pos.x + t2*h2.dir.dx
	x.Mul(&t2, h2.dir.dx)
	x.Add(&x, h2.pos.x)
	var y big.Int // h2.pos.y + t2*h2.dir.dy
	y.Mul(&t2, h2.dir.dy)
	y.Add(&y, h2.pos.y)
	return &Coord2F{&x, &y}
}

func CountIntersectionPairs(hails []HailStone, min int64, max int64) int {
	count := 0

	bigMin := big.NewInt(min)
	bigMax := big.NewInt(max)
	for i := range hails {
		for j := i + 1; j < len(hails); j++ {
			h1 := hails[i]
			h2 := hails[j]

			c := intersectXY(h1, h2)
			if c != nil && c.x.Cmp(bigMin) >= 0 && c.x.Cmp(bigMax) <= 0 &&
				c.y.Cmp(bigMin) >= 0 && c.y.Cmp(bigMax) <= 0 {
				count++
			}
		}
	}

	return count
}

func crossProduct(d0, d1 Dir3F) Dir3F {
	// dx := d0.dy*d1.dz - d0.dz*d1.dy
	var dx1 big.Int
	dx1.Mul(d0.dy, d1.dz)
	var dx2 big.Int
	dx2.Mul(d0.dz, d1.dy)
	var dx big.Int
	dx.Sub(&dx1, &dx2)

	// dy := d0.dz*d1.dx - d0.dx*d1.dz
	var dy1 big.Int
	dy1.Mul(d0.dz, d1.dx)
	var dy2 big.Int
	dy2.Mul(d0.dx, d1.dz)
	var dy big.Int
	dy.Sub(&dy1, &dy2)

	// dz := d0.dx*d1.dy - d0.dy*d1.dx
	var dz1 big.Int
	dz1.Mul(d0.dx, d1.dy)
	var dz2 big.Int
	dz2.Mul(d0.dy, d1.dx)
	var dz big.Int
	dz.Sub(&dz1, &dz2)

	return Dir3F{&dx, &dy, &dz}
}

func dotProduct(d0, d1 Dir3F) *big.Int {
	// r := d0.dx*d1.dx + d0.dy*d1.dy + d0.dz*d1.dz
	var rx big.Int
	rx.Mul(d0.dx, d1.dx)
	var ry big.Int
	ry.Mul(d0.dy, d1.dy)
	var rz big.Int
	rz.Mul(d0.dz, d1.dz)

	var r big.Int
	r.Add(&rx, &ry)
	r.Add(&r, &rz)

	return &r
}

// intersect the plane defined by the point p0 and its normal vector n
// and the line defined by h
// See https://en.wikipedia.org/wiki/Line%E2%80%93plane_intersection
func intersectPlaneAndLine(p0 Coord3F, n Dir3F, h HailStone) (Coord3F, *big.Int) {
	a := dotProduct(p0.MinusToDir(h.pos), n)
	b := dotProduct(h.dir, n)
	var t big.Int
	t.Div(a, b)
	p := h.pos.Plus(h.dir.Times(&t))
	return p, &t
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

	h2 := hailsR[2]
	h3 := hailsR[3]

	p0 := Coord3F{zero, zero, zero}
	p2, t2 := intersectPlaneAndLine(p0, n, h2)
	p3, t3 := intersectPlaneAndLine(p0, n, h3)

	var diffT big.Int
	diffT.Sub(t2, t3)
	dir := p2.MinusToDir(p3).DivideBy(&diffT)
	pos := p2.Minus(dir.Times(t2))

	return HailStone{pos.Plus(h0shift), dir.Plus(h0dir)}
}
