package tile

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mknyszek/2d/anim"
)

// Animation is a series of tile indices representing
// an animation in some Sheet.
//
// Implements [github.com/mknyszek/2d/anim.Playable].
type Animation struct {
	tiles []Index
}

// At implements [github.com/mknyszek/2d/anim.Playable].
func (a Animation) At(i int) Index {
	return a.tiles[i]
}

// Len implements [github.com/mknyszek/2d/anim.Playable].
func (a Animation) Len() int {
	return len(a.tiles)
}

// Anim creates a new animation from the provided tile indices.
func Anim(tiles ...Index) Animation {
	return Animation{tiles: tiles}
}

// Sprite represents a simple animated sprite that iteratively draws
// different tiles from a tile sheet.
//
// It consists of an embedded [github.com/mknyszek/2d/anim.Player]
// that plays back [Animation]s, and a [Sheet] to select images from
// at each frame.
// [Sprite.Image] returns the image for the current frame.
type Sprite struct {
	anim.Player[Index]
	Sheet
}

// Image returns the image for the current frame.
func (s *Sprite) Image() *ebiten.Image {
	return s.Sheet.Select(s.Player.Current())
}
