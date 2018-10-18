package scene

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type displayer interface {
	update(float64, *pixelgl.Window)
	draw(pixel.Target)
}
