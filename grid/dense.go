package grid

import (
	"iter"
	"strconv"
)

// Dense is a 2D grid of values, one per cell.
// The values are assumed to be dense (all present).
// Dense is guaranteed to be plain old data if T is
// also plain old data.
type Dense[T any] struct {
	Data []T
	Rows int
	Cols int
}

// New returns a new 2D grid of values with the provided number of rows and columns.
func New[T any](rows, cols int) *Dense[T] {
	return &Dense[T]{Rows: rows, Cols: cols, Data: make([]T, rows*cols)}
}

// At returns the value at the provided row and column in the grid.
func (d *Dense[T]) At(idx Index) T {
	if idx.Row < 0 || idx.Col < 0 || idx.Row >= d.Rows || idx.Col >= d.Cols {
		panic("row and/or column is out of bounds: " + strconv.Itoa(idx.Row) + ", " + strconv.Itoa(idx.Col))
	}
	return d.Data[idx.Row*d.Cols+idx.Col]
}

// Set sets the value at the provided row and column.
func (d *Dense[T]) Set(idx Index, t T) {
	d.Data[idx.Row*d.Cols+idx.Col] = t
}

// SetAll sets the value of every cell in the grid.
func (d *Dense[T]) SetAll(value T) {
	for i := range d.Data {
		d.Data[i] = value
	}
}

// All returns an iterator over all entries in the grid.
func (d *Dense[T]) All() iter.Seq2[Index, T] {
	return func(yield func(Index, T) bool) {
		row := 0
		col := 0
		for _, value := range d.Data {
			if !yield(Idx(row, col), value) {
				return
			}
			col++
			if col >= d.Cols {
				col = 0
				row++
			}
		}
	}
}
