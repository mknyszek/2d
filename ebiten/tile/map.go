package tile

import (
	"image"
	"iter"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mknyszek/2d/geom"
)

// Map is a 2D grid of tile indices.
//
// Map will always consist of simple data types that are serializable
// by most encoding packages, like JSON.
type Map struct {
	Tiles []Index
	Rows  int
	Cols  int
}

// NewMap returns a new tile map sized according to rows and cols.
func NewMap(rows, cols int) *Map {
	return &Map{Rows: rows, Cols: cols, Tiles: make([]Index, rows*cols)}
}

// At returns the [Index] at the provided row and column in the tile map.
func (m *Map) At(row, col int) Index {
	if row < 0 || col < 0 || row >= m.Rows || col >= m.Cols {
		return Empty
	}
	return m.Tiles[row*m.Cols+col]
}

// Set mutates the Index at the provided row and column.
func (m *Map) Set(row, col int, i Index) {
	m.Tiles[row*m.Cols+col] = i
}

// Render returns an iterator that produces the relative coordinate
// to draw each tile at, as well as that tile's image for each non-empty
// tile in the map.
func (m *Map) Render(s Sheet) iter.Seq2[geom.Vector, *ebiten.Image] {
	return func(yield func(geom.Vector, *ebiten.Image) bool) {
		pt := image.Pt(0, 0)
		w := s.TileWidth() * m.Cols
		for _, idx := range m.Tiles {
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
