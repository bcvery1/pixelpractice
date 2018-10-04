package enemies

import (
	"math"
	"math/rand"
	"time"

	"github.com/bcvery1/pixelpractice/shooter/bullet"
	"github.com/bcvery1/pixelpractice/shooter/consts"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const (
	rockMaxSpeed  = 120
	rockMinRadius = 10
	rockMaxRadius = 45
)

var (
	objects    = make(map[float64]target)
	rockColour = colornames.Grey
	src        = rand.NewSource(time.Now().UnixNano())
	r          = rand.New(src)
)

type target interface {
	Pos() pixel.Vec
	RotationRate() float64
	Movement() pixel.Vec
	Update(float64)
	Draw(pixel.Target)
	Collides(pixel.Rect) bool
}

// ClearAll allows the game to reset all enemies and start from the beginning
func ClearAll() {
	objects = make(map[float64]target)
}

// UpdateAll will update the position of all enemies
func UpdateAll(dt float64) {
	for _, o := range objects {
		o.Update(dt)
	}
}

// DrawAll will draw all enemies to the target
func DrawAll(t pixel.Target) {
	for _, o := range objects {
		o.Draw(t)
	}
}

// Collides checks if the supplied rectangle collides with any of the objects
func Collides(subj pixel.Rect) bool {
	for _, o := range objects {
		if o.Collides(subj) {
			return true
		}
	}
	return false
}

// Count returns how many enemies the game is currently tracking
func Count() int {
	return len(objects)
}

// NewRock adds a rock to the `objects` map
// Will be started at a random location off the screen
func NewRock() {
	id := <-consts.IDs
	radius := consts.RandRange(rockMinRadius, rockMaxRadius)
	pos := consts.RandOffScreenPos(radius / 2)
	rotationRate := r.Float64()
	angle := 360 * r.Float64()

	// Pick a random point across the middle of the screen, aim the rock towards
	// that.  This will ensure it is travelling in a random direction, but will
	// always pass over the visible area
	x := consts.RandRange(0, consts.WinWidth)
	y := consts.RandRange(0, consts.WinHeight)
	movement := consts.CalcSpeed(rockMaxSpeed*r.Float64(), pos, pixel.V(x, y))

	rock := &rock{
		id,
		pos,
		rotationRate,
		angle,
		movement,
		radius,
	}

	objects[id] = rock
}

type rock struct {
	id           float64
	pos          pixel.Vec
	rotationRate float64
	angle        float64
	movement     pixel.Vec
	radius       float64
}

func (r *rock) Pos() pixel.Vec {
	return r.pos
}

func (r *rock) RotationRate() float64 {
	return r.rotationRate
}

func (r *rock) Movement() pixel.Vec {
	return r.movement
}

func (r *rock) Update(dt float64) {
	// Check if the rock is completely off the screen
	if r.Pos().X < 0-r.radius || r.Pos().X > consts.WinWidth+r.radius || r.Pos().Y < 0-r.radius || r.Pos().Y > consts.WinHeight+r.radius {
		r.remove()
		return
	}
	// Check if a bullet collides
	for _, b := range bullet.AllPoints() {
		if r.pointCollides(b) {
			// Add the inverse of the radius as a score
			// The smaller the rock, the harder to hit
			consts.Scored <- int(rockMaxRadius - r.radius)
			r.remove()
			return
		}
	}
	r.pos = r.Pos().Add(r.Movement().Scaled(dt))
	r.angle = math.Mod((r.angle+r.RotationRate())*dt, 360)
}

func (r *rock) Draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = rockColour
	imd.Push(r.Pos())
	imd.Circle(r.radius, 0)

	imd.Draw(t)
}

func (r *rock) Collides(rect pixel.Rect) bool {
	closest := pixel.V(
		consts.Clamp(r.Pos().X, rect.Min.X, rect.Max.X),
		consts.Clamp(r.Pos().Y, rect.Min.Y, rect.Max.Y))
	dist := r.Pos().Sub(closest)

	return dist.Dot(dist) < r.radius*r.radius
}

// pointCollides returns whether a given point is within the circle
func (r *rock) pointCollides(point pixel.Vec) bool {
	d := math.Pow(point.X-r.pos.X, 2) + math.Pow(point.Y-r.pos.Y, 2)
	return math.Sqrt(d) < r.radius
}

func (r *rock) remove() {
	delete(objects, r.id)
}
