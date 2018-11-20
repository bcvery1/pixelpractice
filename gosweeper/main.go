package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	gridSize = 20
)

const (
	gsPlaying = iota
	gsLost
)

var (
	winBounds = pixel.R(0, 0, 1280, 720)
	grid      map[pixel.Rect]*square
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "GoSweeper",
		Bounds: winBounds.Scaled(2),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	batch := imdraw.New(nil)

	overlayCanvas := pixelgl.NewCanvas(winBounds)

	gamestate := gsPlaying

	generateGrid()

	for !win.Closed() {
		win.Clear(borderColour)
		overlayCanvas.Clear(color.Transparent)
		batch.Clear()

		for pos, s := range grid {
			if win.JustPressed(pixelgl.MouseButtonLeft) && pos.Contains(win.MousePosition()) {
				if s.click() {
					gamestate = gsLost
				}
			}

			if win.JustPressed(pixelgl.MouseButtonRight) && pos.Contains(win.MousePosition()) {
				s.flag()
			}

			s.Draw(batch, overlayCanvas)
		}

		batch.Draw(win)
		overlayCanvas.Draw(win, pixel.IM)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
