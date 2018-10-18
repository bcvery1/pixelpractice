package main

import (
	"log"
	"time"

	"github.com/bcvery1/pixelpractice/snake/scene"
	"github.com/bcvery1/pixelpractice/snake/static"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: static.WinBounds,
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Panic(err)
	}
	win.SetSmooth(true)

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		scene.Scene.Update(dt, win)
		scene.Scene.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
