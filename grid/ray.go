package grid

import (
	"iter"
	"math"

	"github.com/mknyszek/2d/geom"
)

// Ray casts a ray over a grid returning an iterator over all
// the grid cells the ray touches.
//
// The ray is assumed to be in the grid's coordinate system.
// Rows correspond to Y position and columns correspond to X
// position.
func Ray(ray geom.Segment) iter.Seq[Index] {
	return func(yield func(Index) bool) {
		dir := geom.Vec(ray.Start, ray.End)
		dirSignX := int(math.Copysign(1, dir.X))
		dirSignY := int(math.Copysign(1, dir.Y))
		offX, offY := 0, 0
		if dir.X > 0 {
			offX = 1
		}
		if dir.Y > 0 {
			offY = 1
		}

		cur := ray.Start
		idx := Idx(int(cur.Y), int(cur.X))
		t := float64(0)
		if dir == geom.Zero {
			yield(idx)
			return
		}
		for t < 1 {
			if !yield(idx) {
				return
			}
			dtX := (float64(idx.Col+offX) - cur.X) / dir.X
			dtY := (float64(idx.Row+offY) - cur.Y) / dir.Y
			if dtX < dtY {
				t += dtX
				idx.Col += dirSignX
			} else {
				t += dtY
				idx.Row += dirSignY
			}
			cur = ray.Start.Add(dir.Scale(t))
		}
	}
}
