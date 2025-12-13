package tile

import (
	"image"
	"iter"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mknyszek/2d/geom"
	"github.com/mknyszek/2d/grid"
)

// Map is a 2D grid of tile indices.
//
// Map will always consist of simple data types that are serializable
// by most encoding packages, like JSON.
type Map grid.Dense[Index]

// NewMap returns a new tile map sized according to rows and cols.
func NewMap(rows, cols int) *Map {
	return (*Map)(grid.New[Index](rows, cols))
}

// At returns the [Index] at the provided row and column in the tile map.
func (m *Map) At(idx grid.Index) Index {
	if idx.Row < 0 || idx.Col < 0 || idx.Row >= m.Rows || idx.Col >= m.Cols {
		return Empty
	}
	return (*grid.Dense[Index])(m).At(idx)
}

// Set mutates the Index at the provided row and column.
func (m *Map) Set(idx grid.Index, i Index) {
	(*grid.Dense[Index])(m).Set(idx, i)
}

// Render returns an iterator that produces the relative coordinate
// to draw each tile at, as well as that tile's image for each non-empty
// tile in the map.
func (m *Map) Render(s Sheet) iter.Seq2[geom.Vector, *ebiten.Image] {
	return func(yield func(geom.Vector, *ebiten.Image) bool) {
		pt := image.Pt(0, 0)
		w := s.TileWidth() * m.Cols
		for _, idx := range m.Data {
			if !idx.Empty() {
				if !yield(geom.ImagePoint(pt).Vector(), s.Select(idx)) {
					return
				}
			}
			pt.X += s.TileWidth()
			if pt.X >= w {
				pt.X = 0
				pt.Y += s.TileHeight()
			}
		}
	}
}

// Size returns the dimensions of the map given a sheet.
func (m *Map) Size(s Sheet) geom.Dimensions {
	return geom.Dim(float64(s.TileWidth()*m.Cols), float64(s.TileHeight()*m.Rows))
}
