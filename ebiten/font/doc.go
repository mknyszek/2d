/*
Package font is a lightweight wrapper around the Ebiten fonts API of
the github.com/hajimehoshi/ebiten/v2/text/v2 package.

It provides several conveniences that are absent in the relatively
low-level text/v2 package:

- A registry which caches loaded fonts.
- Support for finding and using system fonts (provided by [sysfont]).

It still provides access to the low-level text/v2 representation
for rendering.

[sysfont]: https://pkg.go.dev/github.com/adrg/sysfont
*/
package font
