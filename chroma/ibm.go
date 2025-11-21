package chroma

// IBM Design Library's colorblind-friendly sequential color palette.
// See https://www.ibm.com/design/language/resources/color-library.
var IBM = struct {
	Blue   RGB
	Violet RGB
	Pink   RGB
	Orange RGB
	Yellow RGB
}{
	Blue:   RGB(0x648FFF),
	Violet: RGB(0x785EF0),
	Pink:   RGB(0xDC267F),
	Orange: RGB(0xFE6100),
	Yellow: RGB(0xFFB000),
}

// IBMPalette is an array version of IBM.
var IBMPalette = Palette{
	RGB(0x648FFF),
	RGB(0x785EF0),
	RGB(0xDC267F),
	RGB(0xFE6100),
	RGB(0xFFB000),
}
