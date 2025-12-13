package grid

import "iter"

// Bools is a 2D grid of values, one per cell.
// The values are assumed to be dense (all present).
// Bools is guaranteed to be plain old data if T is
// also plain old data.
type Bools struct {
	data []uint64
	Rows int
	Cols int
}

// NewBools returns a new 2D grid of booleans with the provided number of rows and columns.
func NewBools(rows, cols int) *Bools {
	return &Bools{Rows: rows, Cols: cols, data: make([]uint64, (rows*cols+63)/64)}
}

// At returns the value at the provided row and column in the grid.
func (b *Bools) At(idx Index) bool {
	if idx.Row < 0 || idx.Col < 0 || idx.Row >= b.Rows || idx.Col >= b.Cols {
		panic("row and/or column is out of bounds")
	}
	d := idx.Row*b.Cols + idx.Col
	return b.data[d/64]&(uint64(1)<<(d%64)) != 0
}

// Set sets the value at the provided row and column.
func (b *Bools) Set(idx Index, value bool) {
	d := idx.Row*b.Cols + idx.Col
	if value {
		b.data[d/64] |= uint64(1) << (d % 64)
	} else {
		b.data[d/64] &^= uint64(1) << (d % 64)
	}
}

// SetAll sets the value of every tile in the grid.
func (d *Bools) SetAll(value bool) {
	if value {
		for i := range d.data {
			d.data[i] = ^uint64(0)
		}
	} else {
		for i := range d.data {
			d.data[i] = 0
		}
	}
}

// All returns an iterator over all entries in the grid.
func (b *Bools) All() iter.Seq2[Index, bool] {
	return func(yield func(Index, bool) bool) {
		row := 0
		col := 0
		for range b.data {
			i := Idx(row, col)
			if !yield(i, b.At(i)) {
				return
			}
			col++
			if col >= b.Cols {
				col = 0
				row++
			}
		}
	}
}
