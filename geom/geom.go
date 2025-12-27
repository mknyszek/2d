package geom

import (
	"fmt"
	"image"
	"math"
)

var Origin = Point{0, 0}

// Point represents a position in R^2.
//
// A Point is a special case of a Vector: a distance vector
// from Origin. The distinction is useful at the type level,
// but functionally the two are essentially equivalent.
type Point struct {
	X, Y float64
}

// Point is shorthand for a new point.
func Pt(x, y float64) Point {
	return Point{x, y}
}

// Add moves a point by a vector.
func (p Point) Add(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y}
}

// Sub subtracts one point's elements from another.
func (p0 Point) Sub(p1 Point) Point {
	return Point{p0.X - p1.X, p0.Y - p1.Y}
}

// Vector converts the point to a vector. This is equivalent to Vec(Origin, p).
func (p Point) Vector() Vector {
	return Vector{p.X, p.Y}
}

// Image returns a rounded point, suitable for use in image processing.
func (p Point) Image() image.Point {
	return image.Point{int(math.Round(p.X)), int(math.Round(p.Y))}
}

// ImagePoint produces a Point from an image.Point.
func ImagePoint(pt image.Point) Point {
	return Point{float64(pt.X), float64(pt.Y)}
}

// Segment represents a line segment between two points.
//
// If both points are identical, then Segment represents a single point.
//
// Segment implements Curve.
type Segment struct {
	Start, End Point
}

// Seg returns a new Segment between the two provided points.
func Seg(a, b Point) Segment {
	return Segment{a, b}
}

// At interpolates between the two points in the segment according to t, a parameter
// from 0 to 1.
func (s Segment) At(t float64) Point {
	return Point{s.Start.X + t*(s.End.X-s.Start.X), s.Start.Y + t*(s.End.Y-s.Start.Y)}
}

// Length returns the length of the segment.
func (s Segment) Length() float64 {
	if s.Start == s.End {
		return 0
	}
	return Vec(s.Start, s.End).Length()
}

// ZeroLength return trues if the segment has zero length.
func (s Segment) ZeroLength() bool {
	return s.Start == s.End
}

// Intersection returns a segment representing the intersection, and whether a segment
// exists at all. If the two segments intersect at only one point, then Segment contains
// that point as both the Start and End. That is, the length is zero.
func (s0 Segment) Intersection(s1 Segment) (Segment, bool) {
	// Credit to Victor Lecomte for this implementation.
	// Taken from https://github.com/vlecomte/cp-geo.
	properIntersection := func(s0, s1 Segment) (Point, bool) {
		a := s0.Start
		b := s0.End
		c := s1.Start
		d := s1.End
		oa := orient(c, d, a)
		ob := orient(c, d, b)
		oc := orient(a, b, c)
		od := orient(a, b, d)

		// Proper intersection exists iff we have opposite signs.
		if oa*ob < 0 && oc*od < 0 {
			return (a.Vector().Scale(ob).Sub(b.Vector().Scale(oa))).Scale(1.0 / (ob - oa)).Point(Origin), true
		}
		return Point{}, false
	}
	if p, ok := properIntersection(s0, s1); ok {
		return Seg(p, p), true
	}

	// Check endpoints and colinearity (geometry sucks).
	if s1.Contains(s0.Start) {
		if s1.Contains(s0.End) {
			return Seg(s0.Start, s0.End), true
		}
		if s0.Contains(s1.Start) {
			return Seg(s0.Start, s1.Start), true
		}
		return Seg(s0.Start, s0.Start), true
	} else if s1.Contains(s0.End) {
		if s0.Contains(s1.Start) {
			return Seg(s1.Start, s0.End), true
		}
		return Seg(s0.End, s0.End), true
	}
	if s0.Contains(s1.Start) {
		if s0.Contains(s1.End) {
			return Seg(s1.Start, s1.End), true
		}
		if s1.Contains(s0.Start) {
			return Seg(s1.Start, s0.Start), true
		}
		return Seg(s1.Start, s1.Start), true
	} else if s0.Contains(s1.End) {
		if s1.Contains(s0.Start) {
			return Seg(s0.Start, s1.End), true
		}
		return Seg(s1.End, s1.End), true
	}
	return Segment{}, false
}

// Contains returns true if point p lies on segment s.
func (s Segment) Contains(p Point) bool {
	return orient(s.Start, s.End, p) == 0 &&
		inDisk(s.Start, s.End, p)
}

func orient(a, b, c Point) float64 {
	return crossMag(Vec(a, b), Vec(a, c))
}

// crossMag returns the magnitude of the cross product of two vectors.
func crossMag(a, b Vector) float64 {
	return a.X*b.Y - a.Y*b.X
}

// inDisk returns whether p is inside a disk with diameter ab.
func inDisk(a, b, p Point) bool {
	return Vec(p, a).Dot(Vec(p, b)) <= 0
}

// Line represents an infinite line of the form y=mx+b where m is the slope
// and b is the y-intercept.
type Line struct {
	M, B float64
}

// LineFromPoints computes the Line that is intersects the two provided points.
func LineFromPoints(p0, p1 Point) Line {
	r := 1 / (p1.X - p0.X)
	return Line{(p1.Y - p0.Y) * r, (p1.X*p0.Y - p0.X*p1.Y) * r}
}

// Intercept returns the intersection point of the two lines.
//
// Returns false if the lines do not intersect, or if they're identical (intersect at every point).
func (l0 Line) Intercept(l1 Line) (Point, bool) {
	if l1.M == l0.M {
		return Point{}, false
	}
	r := 1 / (l1.M - l0.M)
	return Point{(l0.B - l1.B) * r, (l1.M*l0.B - l0.M*l1.B) * r}, true
}

var Zero = Vector{0, 0}

// Vector is a two-dimensional vector.
type Vector struct {
	X, Y float64
}

// Vec create a new Vector in R^2 from a pair of points.
func Vec(origin, p Point) Vector {
	v := p.Sub(origin)
	return Vector{v.X, v.Y}
}

// Add adds vectors v0 and v1 together.
func (v0 Vector) Add(v1 Vector) Vector {
	return Vector{v0.X + v1.X, v0.Y + v1.Y}
}

// Sub subtracts vector v1 from vector v0.
func (v0 Vector) Sub(v1 Vector) Vector {
	return Vector{v0.X - v1.X, v0.Y - v1.Y}
}

// Dot computes the dot product of two vectors.
func (v0 Vector) Dot(v1 Vector) float64 {
	return v0.X*v1.X + v0.Y*v1.Y
}

// Neg negates the vector.
func (v Vector) Neg() Vector {
	return Vector{-v.X, -v.Y}
}

// Scale scales the vector by a.
func (v Vector) Scale(a float64) Vector {
	return Vector{a * v.X, a * v.Y}
}

// Point produces a Point from the Vector, given an origin.
func (v Vector) Point(origin Point) Point {
	return Point{v.X + origin.X, v.Y + origin.Y}
}

// Length2 returns the square of the length of the vector.
func (v Vector) Length2() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Length returns the length of the vector.
func (v Vector) Length() float64 {
	return math.Sqrt(v.Length2())
}

// Normalize returns a unit vector copy of v pointing in the same direction.
func (v Vector) Normalize() Vector {
	l := v.Length()
	return Vector{v.X / l, v.Y / l}
}

// ReflectX reflects the vector over the X axis.
func (v Vector) ReflectX() Vector {
	out := Vector{X: -v.X, Y: v.Y}
	fmt.Println("reflectX", v, out)
	return out
}

// ReflectY reflects the vector over the Y axis.
func (v Vector) ReflectY() Vector {
	out := Vector{X: v.X, Y: -v.Y}
	fmt.Println("reflectY", v, out)
	return out
}

// Rotate rotates the vector about the origin by rad radians.
func (v Vector) Rotate(rad float64) Vector {
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return Vector{
		v.X*cos - v.Y*sin,
		v.X*sin + v.Y*cos,
	}
}

// RightNormal computes the right-normal vector of this vector.
func (v Vector) RightNormal() Vector {
	return Vector{v.Y, -v.X}
}

// ProjectOnto projects vector a onto vector b.
func (a Vector) ProjectOnto(b Vector) Vector {
	return b.Scale(a.Dot(b) / b.Dot(b))
}

// Dimensions is an abstract width and height without a location.
type Dimensions struct {
	X, Y float64
}

// Dim creates a new set of 2D dimensions.
func Dim(x, y float64) Dimensions {
	return Dimensions{x, y}
}

// ImageDim returns the dimensions of an image.Rectangle.
func ImageDim(r image.Rectangle) Dimensions {
	return Dimensions{float64(r.Dx()), float64(r.Dy())}
}

// AABB gives the dimensions a starting location, producing an AABB.
func (d Dimensions) AABB(start Point) AABB {
	return AABB{start, start.Add(d.Vector())}
}

// Vector returns a vector that represents the dimensions.
func (d Dimensions) Vector() Vector {
	return Vector{d.X, d.Y}
}

// AABB describes an axis-aligned bounding box in R^2.
type AABB struct {
	Min, Max Point
}

// ImageAABB returns an AABB for the image.Rectangle.
func ImageAABB(r image.Rectangle) AABB {
	return AABB{ImagePoint(r.Min), ImagePoint(r.Max)}
}

// Image returns an image.Rectangle, rounding the AABB for use in image processing.
func (a AABB) Image() image.Rectangle {
	return image.Rectangle{a.Min.Image(), a.Max.Image()}
}

// Bound creates a new AABB from two points.
func Bound(x0, y0, x1, y1 float64) AABB {
	return AABB{Point{x0, y0}, Point{x1, y1}}
}

func (a AABB) Dx() float64 {
	return a.Max.X - a.Min.X
}

func (a AABB) Dy() float64 {
	return a.Max.Y - a.Min.Y
}

func (a AABB) Dim() Dimensions {
	return Dim(a.Dx(), a.Dy())
}

func (a AABB) Center() Point {
	return a.Min.Add(Vector{a.Dx() / 2, a.Dy() / 2})
}

// MoveTo sets the AABB's minimum point to the given point, updating the maximum accordingly.
func (a AABB) MoveTo(p Point) AABB {
	return AABB{p, p.Add(a.Dim().Vector())}
}

// Translate moves the AABB in the direction of the provided vector.
func (a AABB) Translate(v Vector) AABB {
	return AABB{a.Min.Add(v), a.Max.Add(v)}
}

// Intersects returns true if the two AABBs intersect.
func (a AABB) Intersects(b AABB) bool {
	return !(a.Max.X <= b.Min.X || a.Min.X >= b.Max.X || a.Max.Y <= b.Min.Y || a.Min.Y >= b.Max.Y)
}

// Penetration returns a vector representing the degree and direction of penetration
// of a to b. Returns the zero vector if the two do not intersect.
func (a AABB) Penetration(b AABB) Vector {
	if !a.Intersects(b) {
		return Zero
	}
	md := b.MinkowskiDiff(a)
	p := Vector{X: md.Min.X}
	d := math.Abs(md.Min.X)
	if d0 := math.Abs(md.Max.X); d0 < d {
		d = d0
		p = Vector{X: md.Max.X}
	}
	if d0 := math.Abs(md.Min.Y); d0 < d {
		d = d0
		p = Vector{Y: md.Min.Y}
	}
	if d0 := math.Abs(md.Max.Y); d0 < d {
		d = d0
		p = Vector{Y: md.Max.Y}
	}
	return p
}

// MinkowskiDiff returns the Minkowski difference of the two AABBs,
// which conveniently is another AABB.
func (a AABB) MinkowskiDiff(b AABB) AABB {
	return Dim(a.Dx()+b.Dx(), a.Dy()+b.Dy()).AABB(Pt(a.Min.X-b.Max.X, a.Min.Y-b.Max.Y))
}

// Left returns the left edge of the AABB.
func (a AABB) Left() Segment {
	return Seg(a.Min, Pt(a.Min.X, a.Max.Y))
}

// Top returns the top edge of the AABB.
func (a AABB) Top() Segment {
	return Seg(a.Min, Pt(a.Max.X, a.Min.Y))
}

// Right returns the right edge of the AABB.
func (a AABB) Right() Segment {
	return Seg(Pt(a.Max.X, a.Min.Y), a.Max)
}

// Bottom returns the bottom edge of the AABB.
func (a AABB) Bottom() Segment {
	return Seg(Pt(a.Min.X, a.Max.Y), a.Max)
}

// Rad converts degrees to radians.
func Rad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Deg converts radiasn to degrees.
func Deg(radians float64) float64 {
	return radians * 180 / math.Pi
}
