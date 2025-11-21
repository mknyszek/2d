/*
Package anim contains primitives for animation.

The core of the package is the [Playable] interface, which just
represents a series of keyframes.
What is at each keyframe is completely abstract and generic.
For example, it could be a row/column pair indexing into a
tilesheet, or it could be point on a parametric curve.

The package provides several wrappers for a [Playable] to
generically modify its behavior, such as:

  - [Reversed], which reverses the order of the keyframes.
  - [Delayed], which inserts additional frames between each keyframe.
  - [PingPonged], which plays the keyframes in reverse order after
    playing them forward once.

These wrappers compose with one another.
This means that an animation's keyframes can generally be separated
from how they're actually played back.

[Player] is a type which can actually play back a [Playable], expressed
loosely as a pull iterator (which is more flexible in, for example,
game engine contexts).
*/
package anim
