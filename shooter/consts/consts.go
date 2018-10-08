// Package consts holds common values used throughout this project.
// They are not necessarily all `const`s
package consts

import (
	"bytes"
	"image"
	"math"
	"math/rand"
	"time"

	"github.com/bcvery1/pixelpractice/shooter/bindata"
	"github.com/faiface/pixel"
)

const (
	// WinWidth represents the width in pixels of the window
	WinWidth = 1280
	// WinHeight represents the height in pixels of the window
	WinHeight = 720
)

var (
	// PlayerSize is the size the player rect will be
	PlayerSize = pixel.V(15, 5)
	src        = rand.NewSource(time.Now().UnixNano())
	r          = rand.New(src)
	// Scored allows thread safe score updating, without needed circlular imports
	Scored = make(chan int, 10)
)

// GenRandVec returns a random vector where neither component has an absolute
// length greater than `maxAbs`
func GenRandVec(maxAbs float64) pixel.Vec {
	posMax := math.Abs(maxAbs)
	x := r.Float64() * 2 * posMax
	y := r.Float64() * 2 * posMax

	return pixel.V(x-posMax, y-posMax)
}

// RandRange returns a random value within the range `min`-`max`
func RandRange(min, max float64) float64 {
	scaler := r.Float64()
	return (scaler * (max - min)) + min
}

// RandOffScreenPos will generate a position vector for an off screen location
// This will `dist` pixels from the edge of the screen
func RandOffScreenPos(dist float64) pixel.Vec {
	a := r.Float64()
	var x, y float64

	switch {
	case a <= 0.25:
		// Top
		x = RandRange(0, WinWidth)
		y = WinHeight + dist
	case a > 0.25 && a <= 0.5:
		// Bottom
		x = RandRange(0, WinWidth)
		y = 0 - dist
	case a > 0.5 && a <= 0.75:
		// Left
		x = 0 - dist
		y = RandRange(0, WinHeight)
	default:
		// Right
		x = WinWidth + dist
		y = RandRange(0, WinHeight)
	}

	return pixel.V(x, y)
}

// CalcSpeed provides a scaled vector from `origin` to `dest`
func CalcSpeed(rate float64, origin, dest pixel.Vec) pixel.Vec {
	return origin.To(dest).Unit().Scaled(rate)
}

// Clamp is a mathematical function, used in collision detection
func Clamp(v, min, max float64) float64 {
	return math.Min(math.Max(v, min), max)
}

func LoadPicture(path string) pixel.Picture {
	file, err := bindata.Asset(path)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}
	return pixel.PictureDataFromImage(img)
}
