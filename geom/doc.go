/*
Package geom provides 2D geometric primitives.

## Points, vectors, and dimensions

The structure of this package is a little bit strange in that it does
not try to capture all 2D points in the singular concept of a vector,
but treats points, vectors, and dimensions separately:

  - A [Point] is an absolute quantity: a Cartesian coordinate.
  - A [Vector] is a relative quantity: a Euclidean vector.
  - A [Dimensions] is a special case of a vector, representing only the
    width and height of an axis-aligned rectangle.

This approach provides some type-safety at the cost of more verbose
computational expressions.

## Curves

A [Curve] is a parametric 2D space curve defined over the closed interval
[0, 1].
This package provides a few such curves, such as [Segment] and
[QuadraticBezier].
*/
package geom
