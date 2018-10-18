package scene

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	title *titleScene
)

func init() {
	title = &titleScene{}
}

type titleScene struct{}

func (t *titleScene) update(dt float64, win *pixelgl.Window) {}

func (t *titleScene) draw(target pixel.Target) {}
