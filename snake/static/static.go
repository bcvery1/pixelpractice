// Package static contains values that a used in various places throughout the
// application
package static

import "github.com/faiface/pixel"

const (
	// WinWidth is the number of pixels to make the window wide
	WinWidth = 1280
	// WinHeight is the number of pixels to make the window high
	WinHeight = 720
	// SegmentWidth is the pixels length of one side of a segment of the player
	SegmentWidth = 20
)

var (
	// WinSegmentCount is the number of segments on the window
	WinSegmentCount = pixel.V(WinWidth/SegmentWidth, WinHeight/SegmentWidth)
	// WinBounds is the bounds of the window
	WinBounds = pixel.R(0, 0, WinWidth, WinHeight)
)
