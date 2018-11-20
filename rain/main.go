package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	maxDrops   = 2000
	maxBounces = 3

	minDropRadius = 0.5
	maxDropRadius = 4

	minDropXVel = -5
	maxDropXVel = 5

	minInitialDropYVel = 0
	maxInitialDropYVel = 10

	minDropBounce = -0.6
	maxDropBounce = -0.4

	gravity         = 25
	airResistFactor = 0.2

	boxMinSize            = 150
	boxMaxSize            = 400
	boxGrowRate   float64 = 700
	boxShrinkRate float64 = -250
)

const (
	BOXACTION_GROWING = iota
	BOXACTION_SHRINKING
	BOXACTION_STATIC
)

var (
	backgroundColour = color.RGBA{0, 0, 20, 255}
	dropColour       = color.RGBA{85, 177, 224, 200}

	drops = make(map[int]*drop)
	ids   = make(chan int, 5)

	src = rand.NewSource(time.Now().UnixNano())
	r   = rand.New(src)

	monitor            *pixelgl.Monitor
	monitorX, monitorY float64

	box      pixel.Rect
	boxSize  float64 = boxMinSize
	boxState         = BOXACTION_STATIC
)

func randFloat64(min, max float64) float64 {
	return (r.Float64() * (max - min)) + min
}

func genColour() color.RGBA {
	red := r.Int()
	green := r.Int()
	blue := r.Int()
	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}

func genIDs() {
	id := 0
	for {
		ids <- id
		id++
	}
}

type drop struct {
	id       int
	radius   float64
	pos      pixel.Vec
	velocity pixel.Vec
	bounces  int
	colour   color.RGBA
}

func (d *drop) draw(imd *imdraw.IMDraw) {
	imd.Color = d.colour
	imd.Push(d.pos)
	imd.Circle(d.radius, 0)
}

func (d *drop) update(dt float64) {
	d.velocity.Y -= gravity * dt
	d.velocity.Y += airResistFactor * dt * d.radius

	d.pos = d.pos.Add(d.velocity)

	if box.Max.Y >= d.pos.Y-d.radius && box.Min.Y <= d.pos.Y+d.radius && box.Min.X <= d.pos.X+d.radius && box.Max.X >= d.pos.X-d.radius {
		// Hits box
		d.bounce()
		d.velocity.X *= -1
	}

	if rectShrink(0.9, box).Contains(d.pos) {
		d.remove()
	} else if d.pos.Y < 0 && d.bounces >= maxBounces {
		d.remove()
	} else if d.pos.Y < d.radius && d.bounces < maxBounces {
		d.bounce()
	}
}

func (d *drop) remove() {
	delete(drops, d.id)
}

func (d *drop) bounce() {
	if d.radius > minDropRadius {
		d.radius /= 2
	}
	d.velocity.Y *= randFloat64(minDropBounce, maxDropBounce) * (1 / d.radius)
	d.bounces++
}

func newDrop() *drop {
	return &drop{
		id:       <-ids,
		radius:   randFloat64(minDropRadius, maxDropRadius),
		pos:      pixel.V(randFloat64(0, monitorX), monitorY+(maxDropRadius*3)),
		velocity: pixel.V(randFloat64(minDropXVel, maxDropXVel), randFloat64(minInitialDropYVel, maxInitialDropYVel)),
		bounces:  0,
		colour:   genColour(),
	}
}

func boxUpdate(dt float64) {
	if boxState == BOXACTION_STATIC {
		return
	}

	var rate float64
	if boxState == BOXACTION_GROWING {
		rate = boxGrowRate
		if boxSize >= boxMaxSize {
			boxSize = boxMaxSize
			boxState = BOXACTION_SHRINKING
			return
		}
	}

	if boxState == BOXACTION_SHRINKING {
		rate = boxShrinkRate
		if boxSize <= boxMinSize {
			boxSize = boxMinSize
			boxState = BOXACTION_STATIC
			return
		}
	}

	boxSize += rate * dt
}

func rectShrink(scale float64, r pixel.Rect) pixel.Rect {
	newSize := r.Size().Scaled(scale)
	centre := r.Center()
	halfX := newSize.X / 2
	halfY := newSize.Y / 2
	return pixel.R(centre.X-halfX, centre.Y-halfY, centre.X+halfX, centre.Y+halfY)
}

func run() {
	go genIDs()

	monitor = pixelgl.PrimaryMonitor()
	monitorX, monitorY = monitor.Size()
	// monitorX, monitorY = 1440, 900

	cfg := pixelgl.WindowConfig{
		Title:   "Rain",
		Bounds:  pixel.R(0, 0, monitorX, monitorY),
		VSync:   true,
		Monitor: monitor,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)
	win.SetCursorVisible(false)

	imd := imdraw.New(nil)
	imd.Color = dropColour

	last := time.Now()

	box = pixel.R(0, 0, boxMinSize, boxMinSize)

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(backgroundColour)
		imd.Clear()

		mousePos := win.MousePosition()
		box = pixel.R(mousePos.X-(boxSize/2), mousePos.Y-(boxSize/2), mousePos.X+(boxSize/2), mousePos.Y+(boxSize/2))

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			boxState = BOXACTION_GROWING
		}

		boxUpdate(dt)

		for len(drops) < maxDrops {
			d := newDrop()
			drops[d.id] = d
		}

		for _, d := range drops {
			d.update(dt)
			d.draw(imd)
		}

		if win.JustPressed(pixelgl.KeyEscape) {
			break
		}

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
