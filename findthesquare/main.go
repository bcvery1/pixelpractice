package main

import (
	"fmt"
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

	boxSize = pixel.V(boxWidth, boxHeight)
)

type box struct {
	colour color.RGBA
	bounds pixel.Rect
}

func randUInt8MinMax(min, max int) uint8 {
	randInt := r.Intn(max-min) + min
	return uint8(randInt)
}

func randFloat64(min, max float64) float64 {
	return (r.Float64() * (max - min)) + min
}

func (b *box) Move() {
	x := randFloat64(0, winWidth-boxWidth)
	y := randFloat64(0, winHeight-boxHeight)
	b.bounds = pixel.R(x, y, x+boxWidth, y+boxHeight)

	red := randUInt8MinMax(0, 155)
	green := randUInt8MinMax(0, 155)
	blue := randUInt8MinMax(0, 155)
	b.colour = color.RGBA{red, green, blue, 255}
}

func (b *box) Draw(target pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = b.colour
	imd.Push(b.bounds.Min, b.bounds.Max)
	imd.Rectangle(0)
	imd.Draw(target)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Find the square - Then click it",
		Bounds: pixel.R(0, 0, 1280, 720),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	b := &box{}
	b.Move()

	score := 0

	for !win.Closed() {
		win.Clear(backgroundColour)

		if b.bounds.Contains(win.MousePosition()) {
			b.Draw(win)

			if win.JustPressed(pixelgl.MouseButtonLeft) {
				b.Move()
				score++
			}
		}

		win.SetTitle(fmt.Sprintf("%s Score: %d", cfg.Title, score))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
