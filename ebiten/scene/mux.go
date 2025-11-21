package scene

import "github.com/hajimehoshi/ebiten/v2"

// Mux is an [github.com/hajimehoshi/ebiten/v2.Game] that multiplexes States.
//
// It simply takes the next State to run at each Update and runs it on the next
// Update, dropping the old State.
type Mux struct {
	current State
	layout  LayoutFunc
}

// NewMux creates a new State multiplexer.
//
// LayoutFunc provides the implementation of the Layout method of
// [github.com/hajimehoshi/ebiten/v2.Game].
func NewMux(init State, layout LayoutFunc) *Mux {
	return &Mux{current: init, layout: layout}
}

// Update implements [github.com/hajimehoshi/ebiten/v2.Game].
func (m *Mux) Update() (err error) {
	// Write your game's logical update.
	m.current, err = m.current.Update()
	return
}

// Draw implements [github.com/hajimehoshi/ebiten/v2.Game].
func (m *Mux) Draw(screen *ebiten.Image) {
	m.current.Draw(screen)
}

// Layout implements [github.com/hajimehoshi/ebiten/v2.Game].
func (m *Mux) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return m.layout(outsideWidth, outsideHeight)
}

// Enforce interface implementation.
var _ ebiten.Game = &Mux{}

// LayoutFunc takes the outside size (e.g., the window size) and returns the (logical) screen size.
type LayoutFunc func(outWidth, outHeight int) (screenWidth, screenHeight int)
