package tile

// Index represents a specific tile in a Sheet.
type Index struct {
	Row int
	Col int
}

// Idx creates a new Index.
func Idx(row, col int) Index {
	return Index{Row: row, Col: col}
}

// Row returns a subset of a row of tiles from colStart to colEnd.
func Row(row, colStart, colEnd int) []Index {
	pts := make([]Index, 0, colEnd-colStart+1)
	for i := colStart; i <= colEnd; i++ {
		pts = append(pts, Index{Row: row, Col: i})
	}
	return pts
}

// Col returns a subset of a column of tiles from rowStart to rowEnd.
func Col(col, rowStart, rowEnd int) []Index {
	pts := make([]Index, 0, rowEnd-rowStart+1)
	for i := rowStart; i <= rowEnd; i++ {
		pts = append(pts, Index{Row: i, Col: col})
	}
	return pts
}
