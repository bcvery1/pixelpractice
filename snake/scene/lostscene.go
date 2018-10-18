package scene

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	lost *lostScene
)

func init() {
	lost = &lostScene{}
}

type lostScene struct{}

func (t *lostScene) update(dt float64, win *pixelgl.Window) {}

func (t *lostScene) draw(target pixel.Target) {}
