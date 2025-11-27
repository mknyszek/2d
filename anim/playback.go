package anim

import "strconv"

// Playable is an interface that describes an animation in
// terms of keyframes.
//
// Really, it just describes any finite sequence.
type Playable[T any] interface {
	// At returns the value at the i'th keyframe.
	At(i int) T

	// Len returns the number of keyframes.
	Len() int
}

// Player plays back an animation.
//
// Really, it is a pull iterator that produces values from a
// finite sequence.
type Player[T any] struct {
	anim  Playable[T]
	frame int
	loop  bool
	done  bool
}

// Start begins playing back the provided animation. The animation
// will play through once, then stop.
//
// This immediately overrides any current animation playing.
func (p *Player[T]) Start(anim Playable[T]) {
	p.anim = anim
	p.frame = 0
	p.loop = false
}

// Loop begins playing back the provided animation. The animation
// will loop forever.
//
// This immediately overrides any current animation playing.
func (p *Player[T]) Loop(anim Playable[T]) {
	p.Start(anim)
	p.loop = true
}

// Current returns the value of the current frame in the animation.
func (p *Player[T]) Current() T {
	return p.anim.At(p.frame)
}

// Next advances the animation.
//
// Returns true if the animation has completed.
func (p *Player[T]) Next() bool {
	if p.Done() {
		return true
	}
	p.frame++
	if p.frame >= p.anim.Len() {
		if p.loop {
			p.frame = 0
		} else {
			p.frame = p.anim.Len() - 1
			p.done = true
			return true
		}
	}
	return false
}

// Done returns true if the animation has completed.
func (p *Player[T]) Done() bool {
	return p.done
}

// Reverse reverses an animation by wrapping it with a Reversed.
func Reverse[A Playable[T], T any](anim A) Reversed[A, T] {
	return Reversed[A, T]{anim: anim}
}

// Reversed is a wrapper type that represents a reversed animation.
type Reversed[A Playable[T], T any] struct {
	anim A
}

// At implements Playable.
func (r Reversed[A, T]) At(i int) T {
	return r.anim.At(r.anim.Len() - 1 - i)
}

// Len implements Playable.
func (r Reversed[A, T]) Len() int {
	return r.anim.Len()
}

// PingPong causes the animation to ping pong (proceed forward, then backward)
// before completing.
func PingPong[A Playable[T], T any](anim A) PingPonged[A, T] {
	return PingPonged[A, T]{anim: anim}
}

// PingPonged is a wrapper type that represents a reversed animation.
type PingPonged[A Playable[T], T any] struct {
	anim A
}

// At implements Playable.
func (p PingPonged[A, T]) At(i int) T {
	switch {
	case i < p.anim.Len():
		return p.anim.At(i)
	case i < p.anim.Len()*2:
		return p.anim.At(2*p.anim.Len() - 1 - i)
	default:
		panic("out of bounds: " + strconv.Itoa(i) + " exceeds " + strconv.Itoa(p.anim.Len()))
	}
}

// Len implements Playable.
func (p PingPonged[A, T]) Len() int {
	return p.anim.Len() * 2
}

// Delay inserts delays in between each frame.
//
// That is, Next will advance through multiples of the same frame
// before advancing to the next one.
//
// A delay of 0 represents no change. A delay of 1 is a doubling of
// every frame (half the frame rate).
func Delay[A Playable[T], T any](anim A, delay int) Delayed[A, T] {
	return Delayed[A, T]{anim: anim, delay: delay}
}

// Delayed is a wrapper type that represents an animation with additional
// delays inserted between frames.
type Delayed[A Playable[T], T any] struct {
	anim  A
	delay int
}

// At implements Playable.
func (d Delayed[A, T]) At(i int) T {
	return d.anim.At(i / (d.delay + 1))
}

// Len implements Playable.
func (d Delayed[A, T]) Len() int {
	return d.anim.Len() * (d.delay + 1)
}
