package player

import (
	"github.com/bcvery1/pixelpractice/shooter/bullet"
	"github.com/bcvery1/pixelpractice/shooter/consts"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const (
	playerSpeed float64 = 10
)

var (
	playerColour = colornames.Yellow
)

// New creates and returns a new Phys
// Initialise at centre of the window
func New() *Phys {
	newPlayer := &Phys{
		pos:   pixel.V(consts.WinWidth/2, consts.WinHeight/2),
		score: 0,
	}

	// Start listening to the score
	go func(p *Phys) {
		for {
			s := <-consts.Scored
			p.score += s
		}
	}(newPlayer)

	return newPlayer
}

// Phys holds the physics info for the player
// Allows control of the player too
type Phys struct {
	pos   pixel.Vec
	score int
	// transform allows us to rotate or otherwise manipulate the player
	transform pixel.Matrix
}

// Point angles the player towards the vector
func (p *Phys) Point(target pixel.Vec) {
	angle := p.pos.To(target).Angle()
	p.transform = pixel.IM.Rotated(p.Centre(), angle)
}

// Rect returns the pixel rectangle for the player.
// Used for collision
func (p *Phys) Rect() pixel.Rect {
	return pixel.Rect{
		p.pos,
		p.pos.Add(consts.PlayerSize),
	}
}

// Draw will draw the player shape to the target
func (p *Phys) Draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = playerColour
	imd.SetMatrix(p.transform)
	imd.Push(p.pos, p.pos.Add(consts.PlayerSize))
	imd.Rectangle(0)

	imd.Draw(t)
}

// Fire creates a new bullet aimed at `target`
func (p *Phys) Fire(target pixel.Vec) {
	bullet.Add(p.pos, target)
}

// Score returns the players current score
func (p *Phys) Score() int {
	return p.score
}

// Reset returns the player to the centre of the window and resets the score
func (p *Phys) Reset() {
	p.pos = pixel.V(consts.WinWidth/2, consts.WinHeight/2)
	p.score = 0
}

// Centre returns the vector representing the centre of the player
func (p *Phys) Centre() pixel.Vec {
	return p.pos.Add(consts.PlayerSize.Scaled(0.5))
}

// MoveLeft attempts to move the player left, unless the player would pass the minX
// Returns true if it hits the minX
func (p *Phys) MoveLeft(minX float64) bool {
	p.pos.X -= playerSpeed
	if p.pos.X < minX {
		p.pos.X = minX
		return true
	}
	return false
}

// MoveRight attempts to move the player right, unless the player would pass the maxX
// Returns true if it hits the minX
func (p *Phys) MoveRight(maxX float64) bool {
	p.pos.X += playerSpeed
	if p.pos.X > maxX-consts.PlayerSize.X {
		p.pos.X = maxX - consts.PlayerSize.X
		return true
	}
	return false
}

// MoveUp attempts to move the player up, unless the player would pass the maxY
// Returns true if it hits the minX
func (p *Phys) MoveUp(maxY float64) bool {
	p.pos.Y += playerSpeed
	if p.pos.Y > maxY-consts.PlayerSize.Y {
		p.pos.Y = maxY - consts.PlayerSize.Y
		return true
	}
	return false
}

// MoveDown attempts to move the player down, unless the player would pass the minY
// Returns true if it hits the minX
func (p *Phys) MoveDown(minY float64) bool {
	p.pos.Y -= playerSpeed
	if p.pos.Y < minY {
		p.pos.Y = minY
		return true
	}
	return false
}
