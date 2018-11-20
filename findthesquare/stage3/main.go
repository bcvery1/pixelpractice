package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	winWidth  = 1280
	winHeight = 720
)

var (
	backgroundColour = color.RGBA{255, 255, 255, 255}
)

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

	for !win.Closed() {
		win.Clear(backgroundColour)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
