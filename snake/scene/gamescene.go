package scene

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/bcvery1/pixelpractice/snake/static"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	// playerStartLength is the number of segments the player starts with
	// including the head segment
	playerStartLength = 6
)

var (
	game        *gamescene
	defaultGame *gamescene
	// playerBorder is the border around the coloured area of a player segment to
	// leave blank
	playerBorder = pixel.V(.5, .5)
	// The size of a player segment, is the full segment size, minus the border at
	// the top/right, minus border at the bottom/left
	playerSegmentSize = pixel.V(static.SegmentWidth, static.SegmentWidth).Sub(playerBorder).Sub(playerBorder)

	// food are the items the player must attempt to get
	foodRadius = static.SegmentWidth / 2
	foodColour = colornames.Blue

	currectBackground    color.RGBA
	backgroundColour     = colornames.Whitesmoke
	lostBackgroundColour = colornames.Red

	// The following vectors set the next movement of the player
	up    = pixel.V(0, 1)
	down  = pixel.V(0, -1)
	left  = pixel.V(-1, 0)
	right = pixel.V(1, 0)

	// Create random source for food placement
	s = rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
)

func init() {
	// Get the segment in the middle of the screen. Then go down the screen the
	// length of the player
	centreSegment := static.WinSegmentCount.Scaled(0.5).Sub(pixel.V(0, playerStartLength))

	// Create tail segment
	s := &segment{
		pos:         centreSegment,
		nextSegment: nil,
	}
	// Create the segments from the tail to the head
	for i := 0; i < playerStartLength; i++ {
		// Move the vector up the screen by 1
		centreSegment = centreSegment.Add(pixel.V(0, 1))
		// Create new segment with pointer to previous
		s = &segment{
			pos:         centreSegment,
			nextSegment: s,
		}
	}

	defaultGame = &gamescene{
		direction:     up,
		headSegment:   s,
		headColour:    colornames.Navy,
		bodyColour:    colornames.Black,
		movementRate:  0.3,
		lastMoved:     0,
		prevDirection: up,
	}

	defaultGame.addFood()

	// Set the background colour
	currectBackground = backgroundColour
}

// resetGame sets the singleton game to the default settings
func resetGame() {
	game = defaultGame
}

type gamescene struct {
	// direction is the direction the player is currently moving
	direction   pixel.Vec
	headSegment *segment
	// headColour is the colour that the head should be drawn
	headColour color.RGBA
	// bodyColour is the colour that the head should be drawn
	bodyColour color.RGBA
	// movementRate is the amount of seconds before the player moves
	// this should decrease as the game goes on
	movementRate float64
	// lastMoved keeps track of how long ago the player moved
	lastMoved float64
	// prevDirection is the last direction previously moved in this can be used to
	// prevent the snake folding back on itself
	prevDirection pixel.Vec
	food          food
}

func (t *gamescene) update(dt float64, win *pixelgl.Window) {
	t.lastMoved += dt

	// test if we should change direction
	t.changeDirection(win)

	// Ensure there is a piece of food on the map
	if t.food.eaten {
		t.addFood()
	}

	if t.lastMoved > t.movementRate {
		// enough time has passed, move the player and update the lastMoved
		t.lastMoved = 0

		lastPos := t.headSegment.pos
		t.headSegment.pos = t.headSegment.pos.Add(t.direction)
		// Check if the game has been lost
		t.checkHealth()
		s := t.headSegment.next()

		// Update the last direction the player moved in
		t.prevDirection = t.direction

		// Check if the head has eaten
		if t.headSegment.pos == t.food.pos {
			t.food.eaten = true

			// Insert new piece behind head
			s = t.grow(lastPos)
			lastPos = s.next().pos
			t.headSegment.nextSegment = s
			// No need to move other segments
			return
		}

		// Move all other segments
		for s != nil {
			tmpPos := s.pos
			s.pos = lastPos
			lastPos = tmpPos

			// Process the next segment
			s = s.next()
		}
	}
}

// checkHealth will check the snake has not bumped into anything
func (t *gamescene) checkHealth() {
	// check the map borders and collision with head piece
	if t.collides(t.headSegment.pos) ||
		t.headSegment.pos.X < 0 ||
		t.headSegment.pos.Y < 0 ||
		t.headSegment.pos.X >= static.WinSegmentCount.X ||
		t.headSegment.pos.Y >= static.WinSegmentCount.Y {
		// Set the background colour
		currectBackground = lostBackgroundColour
		Lose()
	}
}

// collides checks whether the provided vector position collides with any part
// of the players body.  Note, this does not check for collision with the head
func (t *gamescene) collides(v pixel.Vec) bool {
	s := t.headSegment.next()
	for s != nil {
		if s.pos == v {
			return true
		}
		s = s.next()
	}

	return false
}

// changeDirection checks which buttons are being pressed and updates the
// direction to travel in next if appropriate. It uses the property
// `prevDirection` to prevent the snake folding back on itself
func (t *gamescene) changeDirection(win *pixelgl.Window) {
	// Test if up pressed
	if win.JustPressed(pixelgl.KeyW) {
		if t.prevDirection != down {
			t.direction = up
		}
	}

	// Test if down pressed
	if win.JustPressed(pixelgl.KeyS) {
		if t.prevDirection != up {
			t.direction = down
		}
	}

	// Test if left pressed
	if win.JustPressed(pixelgl.KeyA) {
		if t.prevDirection != right {
			t.direction = left
		}
	}

	// Test if right pressed
	if win.JustPressed(pixelgl.KeyD) {
		if t.prevDirection != left {
			t.direction = right
		}
	}
}

func (t *gamescene) draw(target pixel.Target) {
	// canvas is used to buffer the drawing
	canvas := pixelgl.NewCanvas(static.WinBounds)
	canvas.Clear(backgroundColour)

	t.headSegment.draw(t.headColour, canvas)
	s := t.headSegment.next()

	for s != nil {
		s.draw(t.bodyColour, canvas)
		s = s.next()
	}

	t.food.draw(canvas)

	canvas.Draw(target, pixel.IM.Moved(static.WinBounds.Center()))
}

// grow will add a new segment to the end of the player
// this function should be called immediately after the head has moved as it
// inserts a segment after the head
func (t *gamescene) grow(insertAt pixel.Vec) *segment {
	firstBodySeg := t.headSegment.nextSegment
	s := &segment{
		pos:         insertAt,
		nextSegment: firstBodySeg,
	}
	t.headSegment.nextSegment = s

	return s
}

// addFood will create a piece of food and add it to the map.  This function
// will ensure that it does not appear under any body piece, but is otherwise
// in a random location
func (t *gamescene) addFood() {
	for {
		x := r.Intn(int(static.WinSegmentCount.X))
		y := r.Intn(int(static.WinSegmentCount.Y))
		v := pixel.V(float64(x), float64(y))

		if !t.collides(v) {
			t.food = food{v, false}
			return
		}
	}
}

// segment represents either the head or body piece of a player
// It is part of a single-linked list, pointing to the next segment
type segment struct {
	// pos is the segment ordinal - not the pixel placement
	// This is scaled in `draw` to get the correct pixel position
	pos         pixel.Vec
	nextSegment *segment
}

// Draws a segment of the player to the target
func (s *segment) draw(colour color.RGBA, t pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = colour
	imd.Push(
		s.pos.Scaled(static.SegmentWidth).Add(playerBorder),
		s.pos.Scaled(static.SegmentWidth).Add(playerBorder).Add(playerSegmentSize),
	)
	imd.Rectangle(0)

	imd.Draw(t)
}

func (s *segment) next() *segment {
	return s.nextSegment
}

type food struct {
	pos   pixel.Vec
	eaten bool
}

func (f *food) draw(t pixel.Target) {
	// Don't draw if eaten
	if f.eaten {
		return
	}

	imd := imdraw.New(nil)
	imd.Color = foodColour
	centreShift := pixel.V(static.SegmentWidth/2, static.SegmentWidth/2)
	imd.Push(f.pos.Scaled(static.SegmentWidth).Add(centreShift))
	imd.Circle(static.SegmentWidth/4, 0)

	imd.Draw(t)
}
