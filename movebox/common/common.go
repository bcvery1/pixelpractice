package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	// PlayerWidth is the pixel width of the player
	PlayerWidth = 50
	// PlayerHeight is the pixel height of the player
	PlayerHeight = 50
)

var (
	playerSizeV = pixel.V(PlayerWidth, PlayerHeight)
)

// Player is a message sent from a client, providing an update to postition on
// the screen
type Player struct {
	X, Y   float64
	Colour color.RGBA
}

// DecodeFrom decodes a byte slice into an Player struct
func (p *Player) DecodeFrom(inpBytes []byte) error {
	if err := gob.NewDecoder(bytes.NewBuffer(inpBytes)).Decode(p); err != nil {
		return err
	}
	return nil
}

// Byte encodes an Player struct into a byte slice
func (p *Player) Byte() ([]byte, error) {
	// As standard we send 1024 always
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(p); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

// V creates a pixel vector from the coords
func (p *Player) V() pixel.Vec {
	return pixel.V(p.X, p.Y)
}

// Draw draws the player to the target
func (p *Player) Draw(t pixel.Target) {
	icon := imdraw.New(nil)
	icon.Color = (p.Colour)
	icon.Push(p.V(), p.V().Add(playerSizeV))
	icon.Rectangle(0)

	icon.Draw(t)
}

func (p *Player) String() string {
	return fmt.Sprintf("Player at (%.2f, %.2f)", p.X, p.Y)
}
