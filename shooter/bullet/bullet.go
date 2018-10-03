package bullet

import (
	"github.com/bcvery1/pixelpractice/shooter/consts"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const (
	bulletSpeed  float64 = 700
	bulletRadius float64 = 1.5
)

var (
	bulletColour = colornames.Yellow
	bullets      = make(map[float64]*bullet)
)

// Add creates then adds a new bullet to the array
func Add(startPos, dest pixel.Vec) {
	newB := new(startPos, dest)
	bullets[newB.ID] = newB
}

// ClearAll allows the game to reset all bullets and start from the beginning
func ClearAll() {
	bullets = make(map[float64]*bullet)
}

// UpdateAll updates all bullets
// Removing bullets which leave the window
func UpdateAll(dt float64) {
	for _, b := range bullets {
		if b.update(dt) {
			delete(bullets, b.ID)
		}
	}
}

// DrawAll will draw all bullets to the target
func DrawAll(t pixel.Target) {
	for _, b := range bullets {
		b.draw(t)
	}
}

// AllPoints provides all the points (as vectors) that make up the bullets
// Used for detecting collision
func AllPoints() (rects []pixel.Vec) {
	for _, b := range bullets {
		rects = append(rects, b.pos)
	}
	return
}

// bullet represents a basic projectile
type bullet struct {
	ID  float64
	pos pixel.Vec
	dir pixel.Vec
}

// new creates a new bullet travelling from startPos to destPos
func new(startPos, destPos pixel.Vec) *bullet {
	id := <-consts.IDs
	pos := startPos.Add(consts.PlayerSize.Scaled(0.5))
	dir := consts.CalcSpeed(bulletSpeed, pos, destPos)

	return &bullet{id, pos, dir}
}

// update will move the bullet position
// Returns true if this bullet should be removed
func (b *bullet) update(dt float64) bool {
	if b.pos.X < 0 || b.pos.Y < 0 || b.pos.X > consts.WinWidth || b.pos.Y > consts.WinHeight {
		return true
	}
	b.pos = b.pos.Add(b.dir.Scaled(dt))
	return false
}

// draw will draw the bullet to the target
func (b *bullet) draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = bulletColour
	imd.Push(b.pos)
	imd.Circle(bulletRadius, 0)

	imd.Draw(t)
}
