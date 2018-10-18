package main

import (
	"fmt"
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
	winBounds      = pixel.R(0, 0, 1280, 720)
	defaultClicked = pixel.V(-1, -1)
	grid           map[pixel.Rect]*square
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

	clickedPos := defaultClicked

	batch := imdraw.New(nil)

	overlayCanvas := pixelgl.NewCanvas(winBounds)

	gamestate := gsPlaying

	generateGrid()

	for !win.Closed() {
		win.Clear(borderColour)
		overlayCanvas.Clear(color.Transparent)
		batch.Clear()

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			clickedPos = win.MousePosition()
		}

		for pos, s := range grid {
			if clickedPos != defaultClicked && pos.Contains(clickedPos) {
				if s.click() {
					gamestate = gsLost
					fmt.Println(gamestate)
				}
			}

			s.Draw(batch, overlayCanvas)
		}

		// set clickedPos to a place off the grid
		clickedPos = defaultClicked

		batch.Draw(win)
		overlayCanvas.Draw(win, pixel.IM)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
