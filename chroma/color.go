package chroma

import "image/color"

// RGB is a packed 24-bit opaque RGB color.
//
// The packed representation allows for easily
// writing out hex codes in plain Go.
// For example, RGB(0xff0000) is a plain red.
type RGB uint32

// RGBA implements color.Color.
func (c RGB) RGBA() (r, g, b, a uint32) {
	r = (uint32(c) >> 16) & 0xff
	g = (uint32(c) >> 8) & 0xff
	b = (uint32(c) >> 0) & 0xff
	a = 0xff
	r |= r << 8
	g |= g << 8
	b |= b << 8
	a |= a << 8
	return
}

// WithAlpha adds an alpha channel to the color and returns a pre-multiplied color.RGBA.
func (c RGB) WithAlpha(a uint8) color.RGBA {
	r := uint32(c) >> 16
	g := uint32(c) >> 8
	b := uint32(c) >> 0
	r = r * uint32(a) / 0xff
	g = g * uint32(a) / 0xff
	b = b * uint32(a) / 0xff
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

// Black is the color black.
const Black = RGB(0)

// White is the color white.
const White = RGB(0xffffff)

// Palette represents a sequential color palette.
type Palette []color.Color

// Spectrum represents a continuous spectrum of color.
type Spectrum interface {
	// At returns a color in the spectrum defined over the closed interval [0, 1].
	At(t float64) color.Color
}
