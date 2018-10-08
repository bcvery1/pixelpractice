package enemies

import (
	"bytes"
	"image"
	"math"
	"math/rand"
	"time"

	// blank import needed for loading the png rock image
	_ "image/png"

	"github.com/bcvery1/pixelpractice/shooter/bindata"
	"github.com/bcvery1/pixelpractice/shooter/bullet"
	"github.com/bcvery1/pixelpractice/shooter/consts"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

const (
	rockMinSpeed  = 120
	rockMinRadius = 10
	rockMaxRadius = 45
)

var (
	objects    = make(map[float64]target)
	rockColour = colornames.Grey
	src        = rand.NewSource(time.Now().UnixNano())
	r          = rand.New(src)
	// SpriteSheet is the loaded rock picture
	SpriteSheet = loadPicture("assets/rock.png")
	rockSprite  = pixel.NewSprite(SpriteSheet, SpriteSheet.Bounds())
)

type target interface {
	Pos() pixel.Vec
	RotationRate() float64
	Movement() pixel.Vec
	Update(float64)
	Draw(pixel.Target)
	Collides(pixel.Rect) bool
}

func loadPicture(path string) pixel.Picture {
	file, err := bindata.Asset(path)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}
	return pixel.PictureDataFromImage(img)
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
func NewRock(sinceStart float64) {
	id := <-consts.IDs
	radius := consts.RandRange(rockMinRadius, rockMaxRadius)
	pos := consts.RandOffScreenPos(radius / 2)
	rotationRate := math.Pi * r.Float64()
	angle := math.Pi * r.Float64()

	// Calculate speed based on how many seconds have passed
	speed := rockMinSpeed + (sinceStart * 2)

	// Pick a random point across the middle of the screen, aim the rock towards
	// that.  This will ensure it is travelling in a random direction, but will
	// always pass over the visible area
	x := consts.RandRange(0, consts.WinWidth)
	y := consts.RandRange(0, consts.WinHeight)
	movement := consts.CalcSpeed(speed*r.Float64(), pos, pixel.V(x, y))

	rock := &rock{
		id,
		pos,
		rotationRate,
		angle,
		movement,
		radius,
		speed,
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
	speed        float64
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
	r.angle = math.Mod(r.angle+(r.RotationRate()*dt), 2*math.Pi)
}

func (r *rock) Draw(t pixel.Target) {
	rockSprite.Draw(t, pixel.IM.Moved(r.Pos()).Rotated(r.Pos(), r.angle).Scaled(r.Pos(), r.radius/26))
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
