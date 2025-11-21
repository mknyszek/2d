package geom

// Curve is an interface representing a 2D parametric space curve.
type Curve interface {
	// At returns the [Point] along the curve over the closed interval [0, 1].
	At(t float64) Point
}
