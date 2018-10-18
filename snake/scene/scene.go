package scene

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	// Scene is the singleton scene instance
	Scene scene
)

func init() {
	Scene = scene{}

	Start()
}

// scene represents the current game state
type scene struct {
	currentDisplayer displayer
}

// Update runs all updates on the current scene
func (s *scene) Update(dt float64, win *pixelgl.Window) {
	s.currentDisplayer.update(dt, win)
}

// Draw performs all buffer drawings to the target
func (s *scene) Draw(target pixel.Target) {
	s.currentDisplayer.draw(target)
}

// Start will start the game from the beginning
func Start() {
	resetGame()
	Scene.currentDisplayer = game
}

// Lose will put the game into a loosing state
func Lose() {
	log.Println("Lost")
	Scene.currentDisplayer = lost
}
