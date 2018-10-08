package player

import (
	"math"

	"github.com/bcvery1/pixelpractice/shooter/bullet"
	"github.com/bcvery1/pixelpractice/shooter/consts"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

const (
	maxAcc   = 6
	accRate  = 12
	slowRate = 10
)

var (
	playerColour = colornames.Yellow
	// SpriteSheet is the loaded ship picture
	SpriteSheet = consts.LoadPicture("assets/ship.png")
	shipSprite  = pixel.NewSprite(SpriteSheet, SpriteSheet.Bounds())
)

// New creates and returns a new Phys
// Initialise at centre of the window
func New(bounds pixel.Rect) *Phys {
	newPlayer := &Phys{
		pos:    pixel.V(consts.WinWidth/2, consts.WinHeight/2),
		score:  0,
		bounds: bounds,
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
	transform    pixel.Matrix
	acceleration pixel.Vec
	// bounds is the rectangle that the player cannot leave
	bounds pixel.Rect
}

// point angles the player towards the vector
func (p *Phys) point(target pixel.Vec) {
	angle := p.pos.To(target).Angle() - (math.Pi / 2)
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
	shipSprite.Draw(t, pixel.IM.Moved(p.Centre()).Chained(p.transform))
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

// reduceTo will reduce i to z by rate
func reduceTo(i, rate, z float64) float64 {
	if i == z || (math.Abs(i)-math.Abs(rate)) < math.Abs(z) {
		return z
	}

	if i < z {
		return i + math.Abs(rate)
	}
	return i - math.Abs(rate)
}

// Update updates the players position according to its' acceleration
func (p *Phys) Update(dt float64, mousePos pixel.Vec) {
	// Apply slow down
	p.acceleration.X = reduceTo(p.acceleration.X, slowRate*dt, 0)
	p.acceleration.Y = reduceTo(p.acceleration.Y, slowRate*dt, 0)

	nextX := p.pos.X + p.acceleration.X
	nextY := p.pos.Y + p.acceleration.Y

	if nextX > p.bounds.Min.X && nextX < p.bounds.Max.X {
		p.pos.X = nextX
	}

	if nextY > p.bounds.Min.Y && nextY < p.bounds.Max.Y {
		p.pos.Y = nextY
	}

	p.point(mousePos)
}

// MoveLeft attempts to move the player left, unless the player would pass the minX
// Returns true if it hits the minX
func (p *Phys) MoveLeft(dt float64) {
	p.acceleration.X = reduceTo(p.acceleration.X, accRate*dt, -1*maxAcc)
}

// MoveRight attempts to move the player right, unless the player would pass the maxX
// Returns true if it hits the minX
func (p *Phys) MoveRight(dt float64) {
	p.acceleration.X = reduceTo(p.acceleration.X, accRate*dt, maxAcc)
}

// MoveUp attempts to move the player up, unless the player would pass the maxY
// Returns true if it hits the minX
func (p *Phys) MoveUp(dt float64) {
	p.acceleration.Y = reduceTo(p.acceleration.Y, accRate*dt, maxAcc)
}

// MoveDown attempts to move the player down, unless the player would pass the minY
// Returns true if it hits the minX
func (p *Phys) MoveDown(dt float64) {
	p.acceleration.Y = reduceTo(p.acceleration.Y, accRate*dt, -1*maxAcc)
}
