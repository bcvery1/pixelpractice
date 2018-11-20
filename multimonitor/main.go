package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	// quit is used to exit out of the run function when one of the window loops
	// closes
	quit = make(chan struct{}, 1)
)

// createCfgs creates a full-screen config for every monitor
func createCfgs() []pixelgl.WindowConfig {
	ret := []pixelgl.WindowConfig{}
	for _, m := range pixelgl.Monitors() {
		mX, mY := m.Size()
		ret = append(ret, pixelgl.WindowConfig{
			Bounds:  pixel.R(0, 0, mX, mY),
			VSync:   true,
			Monitor: m,
		})
	}

	return ret
}

// winLoop is a render loop run on one monitor
func winLoop(cfg pixelgl.WindowConfig) {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		// For the test, just set the background to black
		win.Clear(colornames.Black)

		// check if the esc key has been pressed
		if win.JustPressed(pixelgl.KeyEscape) {
			// exit out of the win.Closed() loop
			break
		}
		win.Update()
	}

	// When the win.Closed() loop exists, close the `quit` channel
	close(quit)
}

func run() {
	// Create all the configs
	cfgs := createCfgs()

	// Run a winLoop on each monitor
	for _, cfg := range cfgs {
		go winLoop(cfg)
	}

	// Wait until the quit channel closes
	<-quit
}

func main() {
	pixelgl.Run(run)
}
