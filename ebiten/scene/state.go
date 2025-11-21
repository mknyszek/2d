package scene

import "github.com/hajimehoshi/ebiten/v2"

// State represents a stateful scene in Ebiten.
//
// State is a state in a state machine, returning the next scene
// to display at each call to Update.
type State interface {
	// Update is intended to be executed each tick in the Ebiten game engine.
	//
	// Returns the next State to run.
	Update() (State, error)

	// Draw renders the current frame to the screen.
	Draw(*ebiten.Image)
}
