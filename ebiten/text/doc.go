/*
Package text is a wrapper around the text rendering API part of
github.com/hajimehoshi/ebiten/v2/text/v2.

It provides several additional conveniences that match how GUI text
editing interfaces tend to work, including:

  - Basic types for styling text.
  - A styled [String] type that may interleave multiple styles along
    with helpers to efficiently construct them.
  - A text [Box] that renders text strictly within a well-defined
    area with overflow hidden.
  - An automatically-sized text box, [AutoBox] which allows simply
    drawing text and anchoring its position to a relative point in
    the box.
*/
package text
