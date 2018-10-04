package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/bcvery1/pixelpractice/shooter/bullet"
	"github.com/bcvery1/pixelpractice/shooter/consts"
	"github.com/bcvery1/pixelpractice/shooter/enemies"
	"github.com/bcvery1/pixelpractice/shooter/player"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	maxEnemies = 15
)

const (
	// GAMEPLAYING indicates the game is currently playing.  Update all entities
	GAMEPLAYING = iota
	// GAMELOST indicates the game is stopped, do not update
	GAMELOST
	// GAMERESET indicates the game should be reset
	GAMERESET
)

var (
	p                *player.Phys
	backgroundColour = colornames.Black
	// frames is used to measure the fps
	frames = 0
	// second is a ticker meauring seconds for fps
	second    = time.Tick(time.Second)
	gamestate = GAMEPLAYING
	// alphaState allows for fading out of the screen on losing
	alphaState uint8
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Shoot",
		Bounds: pixel.R(0, 0, consts.WinWidth, consts.WinHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	// buff is a canvas to draw all objects to for efficiency
	buff := imdraw.New(nil)
	// overlayBuff is a canvas drawn over buff
	overlayBuff := imdraw.New(nil)

	p = player.New()

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(backgroundColour)
		buff.Clear()
		overlayBuff.Clear()

		switch gamestate {
		case GAMEPLAYING:
			// Update enemies
			for enemies.Count() < maxEnemies {
				enemies.NewRock()
			}

			enemies.UpdateAll(dt)

			bullet.UpdateAll(dt)

			// Try move the player right
			if win.Pressed(pixelgl.KeyD) {
				p.MoveRight(win.Bounds().Max.X)
			}
			// Try move player left
			if win.Pressed(pixelgl.KeyA) {
				p.MoveLeft(0)
			}
			// Try move player down
			if win.Pressed(pixelgl.KeyS) {
				p.MoveDown(0)
			}
			// Try move player up
			if win.Pressed(pixelgl.KeyW) {
				p.MoveUp(win.Bounds().Max.Y)
			}

			if win.JustPressed(pixelgl.MouseButtonLeft) {
				p.Fire(win.MousePosition())
			}

			// Check if the player has collided
			if enemies.Collides(p.Rect()) {
				gamestate = GAMELOST
			}

		case GAMELOST:
			if win.JustPressed(pixelgl.KeySpace) {
				gamestate = GAMERESET
			}

			// Draw a faded background
			imd := imdraw.New(nil)
			imd.Color = color.RGBA{200, 200, 200, 0x00}
			imd.Push(pixel.ZV, win.Bounds().Max)
			imd.Rectangle(0)
			imd.Draw(overlayBuff)

			// Increase alphaState to fade out
			if alphaState < 200 {
				alphaState++
			}

		case GAMERESET:
			gamestate = GAMEPLAYING
			enemies.ClearAll()
			bullet.ClearAll()
			p.Reset()
			alphaState = 0
		}

		enemies.DrawAll(buff)
		bullet.DrawAll(buff)
		p.Draw(buff)

		buff.Draw(win)
		overlayBuff.Draw(win)

		win.Update()

		// Calculate and display FPS
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d | Score: %d", cfg.Title, frames, p.Score()))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
