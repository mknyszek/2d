package chroma

// Wong is Bang Wong's colorblind-friendly sequential color palette.
// See https://www.nature.com/articles/nmeth.1618.
var Wong = struct {
	Black     RGB
	Orange    RGB
	LightBlue RGB
	Green     RGB
	Yellow    RGB
	Blue      RGB
	Amber     RGB
	Pink      RGB
}{
	Black:     Black,
	Orange:    RGB(0xE69F00),
	LightBlue: RGB(0x56B4E9),
	Green:     RGB(0x009E73),
	Yellow:    RGB(0xF0E442),
	Blue:      RGB(0x0072B2),
	Amber:     RGB(0xD55E00),
	Pink:      RGB(0xCC79A7),
}

// WongPalette is an array version of Wong.
var WongPalette = Palette{
	Black,
	RGB(0xE69F00),
	RGB(0x56B4E9),
	RGB(0x009E73),
	RGB(0xF0E442),
	RGB(0x0072B2),
	RGB(0xD55E00),
	RGB(0xCC79A7),
}
