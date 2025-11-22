/*
Package scene provides a simple scene multiplexer for Ebitengine games.

Each scene implements a stateful [State], whose implementation includes
a state machine for selecting the next [State].
The network of [State] values in turn represents a state machine.
*/
package scene
