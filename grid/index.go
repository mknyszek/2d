package grid

import "github.com/mknyszek/2d/geom"

// Index is an index into a grid.
type Index struct {
	Row, Col int
}

// Idx is a convenience function for creating an Index.
func Idx(row, col int) Index {
	return Index{row, col}
}

// AABB returns the AABB for the grid cell given the dimension of each cell.
func (i Index) AABB(dim geom.Dimensions) geom.AABB {
	return dim.AABB(i.Min(dim))
}

// Min returns the minimum point in the cell given the cell dimenions.
func (i Index) Min(dim geom.Dimensions) geom.Point {
	return geom.Pt(float64(i.Col)*dim.X, float64(i.Row)*dim.Y)
}

// Max returns the maximum point in the cell given the cell dimenions.
func (i Index) Max(dim geom.Dimensions) geom.Point {
	return geom.Pt(float64(i.Col+1)*dim.X, float64(i.Row+1)*dim.Y)
}
