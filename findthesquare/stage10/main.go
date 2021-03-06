package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	winWidth  = 1280
	winHeight = 720

	boxWidth  = 100
	boxHeight = 100
)

var (
	backgroundColour = color.RGBA{255, 255, 255, 255}

	src = rand.NewSource(time.Now().UnixNano())
	r   = rand.New(src)
)

func randomFloat64(min, max float64) float64 {
	return (r.Float64() * (max - min)) + min
}

type square struct {
	bounds pixel.Rect
}

func (s *square) Draw(target pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = color.RGBA{0, 0, 0, 255}
	imd.Push(s.bounds.Min, s.bounds.Max)
	imd.Rectangle(0)
	imd.Draw(target)
}

func (s *square) move() {
	x := randomFloat64(0, winWidth-boxWidth)
	y := randomFloat64(0, winHeight-boxHeight)
	s.bounds = pixel.R(x, y, x+boxWidth, y+boxHeight)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Find the square",
		Bounds: pixel.R(0, 0, 1280, 720),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	s := square{
		// bounds: pixel.R(200, 200, 200+boxWidth, 200+boxHeight),
	}
	s.move()

	for !win.Closed() {
		win.Clear(backgroundColour)

		if s.bounds.Contains(win.MousePosition()) {
			s.Draw(win)
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
