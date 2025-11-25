package tile

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sheet represents a tile sheet where each tile has uniform rectangular dimensions.
type Sheet struct {
	image      *ebiten.Image
	tileWidth  int
	tileHeight int
}

// SheetFrom creates a new Sheet from an image the dimensions of each tile in pixels
// in that image.
func SheetFrom(image *ebiten.Image, tileWidth, tileHeight int) Sheet {
	return Sheet{image: image, tileWidth: tileWidth, tileHeight: tileHeight}
}

// TileWidth returns the width of each tile in pixels.
func (s Sheet) TileWidth() int {
	return s.tileWidth
}

// TileHeight returns the height of each tile in pixels.
func (s Sheet) TileHeight() int {
	return s.tileHeight
}

// Select returns a subset of the Sheet's underlying image corresponding to the
// provided tile index.
//
// Returns nil if tile is [Empty].
func (s Sheet) Select(tile Index) *ebiten.Image {
	if tile.Empty() {
		return nil
	}
	p := image.Point{tile.Col() * s.tileWidth, tile.Row() * s.tileHeight}
	return s.image.SubImage(image.Rect(p.X, p.Y, p.X+s.tileWidth, p.Y+s.tileHeight)).(*ebiten.Image)
}
