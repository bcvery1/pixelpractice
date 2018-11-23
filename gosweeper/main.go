package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	gridSize float64 = 30
)

const (
	gsPlaying = iota
	gsLost
	gsWon
)

var (
	winBounds = pixel.R(0, 0, 1200, 600)
	grid      map[pixel.Vec]*square
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "GoSweeper",
		Bounds: winBounds,
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
		win.SetTitle(fmt.Sprintf("%s | Bombs: %d | Press 'r' to restart", cfg.Title, numBombs))

		win.Clear(borderColour)
		overlayCanvas.Clear(color.Transparent)
		batch.Clear()

		if win.JustPressed(pixelgl.KeyR) {
			generateGrid()
			gamestate = gsPlaying
		}

		if win.JustPressed(pixelgl.MouseButtonLeft) && gamestate == gsPlaying {
			if grid[vecToRowCol(win.MousePosition())].click() {
				gamestate = gsLost
			}
		}

		if win.JustPressed(pixelgl.MouseButtonRight) && gamestate == gsPlaying {
			if grid[vecToRowCol(win.MousePosition())].flag() {
				gamestate = gsWon
			}
		}

		for _, s := range grid {
			s.Draw(batch, overlayCanvas)
		}

		batch.Draw(win)
		overlayCanvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
