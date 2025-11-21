package tween

// Function represents a parametric function.
type Function[T any] interface {
	// At returns the value of the function along the closed interval [0, 1].
	At(t float64) T
}

// Make creates a new tween that plays out over some number of steps.
//
// The provided Function will be evaluated along 0 <= t <= 1 across
// these steps. The first and last steps are always guaranteed to be
// 0 and 1. steps must be at least 2.
func Make[F Function[T], T any](f F, steps int) Sequence[F, T] {
	if steps < 2 {
		panic("tween must have at least 2 steps")
	}
	return Sequence[F, T]{fn: f, n: steps - 1}
}

// Sequence is a tween.
//
// Concretely, it is the discretization of a Function over some number of steps.
type Sequence[F Function[T], T any] struct {
	fn F
	n  int
}

// At returns the value of the function at step i.
func (s Sequence[F, T]) At(i int) T {
	return s.fn.At(float64(i) / float64(s.n))
}

// Len returns the number of steps in the sequence.
func (s Sequence[F, T]) Len() int {
	return s.n
}

// Ease applies Easing to a Function, returning a new Function.
func Ease[F Function[T], T any](f F, ease Easing) Eased[F, T] {
	return Eased[F, T]{fn: f, ease: ease}
}

// Eased is a wrapper type that applies Easing to a Function.
type Eased[F Function[T], T any] struct {
	fn   F
	ease Easing
}

// At implements Function.
func (e Eased[F, T]) At(t float64) T {
	return e.fn.At(e.ease(t))
}

// Easing is a function that maps 0 <= t <= 1 to a range from 0 to 1.
type Easing func(t float64) (tp float64)

func EaseInQuad(t float64) float64 {
	return t * t
}

func EaseOutQuad(t float64) float64 {
	return 1 - (1-t)*(1-t)
}

func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return 1 - ((-2*t+2)*(-2*t+2))/2
}

func EaseInCubic(t float64) float64 {
	return t * t * t
}

func EaseOutCubic(t float64) float64 {
	return 1 - ((1 - t) * (1 - t) * (1 - t))
}

func EaseInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	}
	return 1 - ((-2*t+2)*(-2*t+2)*(-2*t+2))/2
}

func EaseInQuart(t float64) float64 {
	return t * t * t * t
}

func EaseOutQuart(t float64) float64 {
	return 1 - ((1 - t) * (1 - t) * (1 - t) * (1 - t))
}

func EaseInOutQuart(t float64) float64 {
	if t < 0.5 {
		return 8 * t * t * t * t
	}
	return 1 - ((-2*t+2)*(-2*t+2)*(-2*t+2)*(-2*t+2))/2
}

func EaseInQuint(t float64) float64 {
	return t * t * t * t * t
}

func EaseOutQuint(t float64) float64 {
	return 1 - ((1 - t) * (1 - t) * (1 - t) * (1 - t) * (1 - t))
}

func EaseInOutQuint(t float64) float64 {
	if t < 0.5 {
		return 16 * t * t * t * t * t
	}
	return 1 - ((-2*t+2)*(-2*t+2)*(-2*t+2)*(-2*t+2)*(-2*t+2))/2
}
