package tile

// Index represents a specific tile in a [Sheet].
//
// 0 represents an empty tile.
type Index uint32

// Empty returns an index for which Empty is true.
var Empty = Index(0)

// Idx creates a new Index.
//
// Negative indices will cause the index to be considered empty.
func Idx(row, col uint16) Index {
	return Index((uint32(row) << 16) | uint32(col+1))
}

// Empty returns true if the Index doesn't represent any selection.
func (i Index) Empty() bool {
	return i == 0
}

// Row returns the row index.
//
// This is only meaningful if [Empty] returns false.
func (i Index) Row() int {
	return int(uint32(i) >> 16)
}

// Col returns the column index.
//
// This is only meaningful if [Empty] returns false.
func (i Index) Col() int {
	return int(uint32(i)&uint32(^uint16(0))) - 1
}

// Row returns a subset of a row of tiles from colStart to colEnd.
func Row(row, colStart, colEnd uint16) []Index {
	pts := make([]Index, 0, colEnd-colStart+1)
	for i := colStart; i <= colEnd; i++ {
		pts = append(pts, Idx(row, i))
	}
	return pts
}

// Col returns a subset of a column of tiles from rowStart to rowEnd.
func Col(col, rowStart, rowEnd uint16) []Index {
	pts := make([]Index, 0, rowEnd-rowStart+1)
	for i := rowStart; i <= rowEnd; i++ {
		pts = append(pts, Idx(i, col))
	}
	return pts
}
