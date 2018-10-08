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
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
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
	// This will increase with score
	maxEnemies = 20
	minEnemies = maxEnemies
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
	// Writing
	writingCanvas := pixelgl.NewCanvas(win.Bounds())
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(win.Bounds().Center(), atlas)
	txt.Color = colornames.Black
	fmt.Fprintf(txt, "You Lose!\nPress space to restart")
	// batch to draw rock sprites to
	rockBatch := pixel.NewBatch(&pixel.TrianglesData{}, enemies.SpriteSheet)

	p = player.New(win.Bounds())

	last := time.Now()
	// startTime is used to increase the enemy speed
	startTime := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(backgroundColour)
		buff.Clear()
		rockBatch.Clear()
		overlayBuff.Clear()
		writingCanvas.Clear(color.Transparent)

		switch gamestate {
		case GAMEPLAYING:
			// Increase enemy count based on score
			maxEnemies = minEnemies + p.Score()/50

			// Update enemies
			for enemies.Count() < maxEnemies {
				enemies.NewRock(time.Since(startTime).Seconds())
			}

			enemies.UpdateAll(dt)

			bullet.UpdateAll(dt)

			// Point the player towards the mouse
			p.Update(dt, win.MousePosition())

			// Try accelerate the player right
			if win.Pressed(pixelgl.KeyD) {
				p.MoveRight(dt)
			}
			// Try accelerate player left
			if win.Pressed(pixelgl.KeyA) {
				p.MoveLeft(dt)
			}
			// Try accelerate player down
			if win.Pressed(pixelgl.KeyS) {
				p.MoveDown(dt)
			}
			// Try accelerate player up
			if win.Pressed(pixelgl.KeyW) {
				p.MoveUp(dt)
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

			// Write to screen
			txt.Draw(writingCanvas, pixel.IM.Scaled(txt.Orig, 3).Moved(pixel.V(100, 150)))

		case GAMERESET:
			gamestate = GAMEPLAYING
			enemies.ClearAll()
			bullet.ClearAll()
			p.Reset()
			alphaState = 0
		}

		enemies.DrawAll(rockBatch)
		bullet.DrawAll(buff)
		p.Draw(win)

		buff.Draw(win)
		rockBatch.Draw(win)
		overlayBuff.Draw(win)
		writingCanvas.Draw(win, pixel.IM)

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
