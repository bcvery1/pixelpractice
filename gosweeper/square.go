package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

const (
	borderWidth = 1
	// TODO(Set this based on grid size)
	numBombs = 50
)

var (
	borderColour    = color.RGBA{0xd2, 0xd5, 0xb6, 0xff}
	coveredColour   = color.RGBA{0x94, 0x94, 0x94, 0xff}
	uncoveredColour = color.RGBA{0xc6, 0xc9, 0xac, 0xff}
	lostColour      = color.RGBA{0xb3, 0x00, 0x00, 0xff}

	src = rand.NewSource(time.Now().Unix())
	r   = rand.New(src)

	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Direction vectors
	// lateral
	up    = pixel.V(0, gridSize)
	down  = pixel.V(0, -gridSize)
	left  = pixel.V(-gridSize, 0)
	right = pixel.V(gridSize, 0)

	// diagonal
	upright   = pixel.V(gridSize, gridSize)
	upleft    = pixel.V(-gridSize, gridSize)
	downright = pixel.V(gridSize, -gridSize)
	downleft  = pixel.V(-gridSize, -gridSize)
)

type square struct {
	hasBomb bool
	pos     pixel.Rect
	covered bool
	flagged bool
	// count is the number of surrounding squares have bombs
	count int
}

func (s *square) Draw(imd *imdraw.IMDraw, overlay *pixelgl.Canvas) {
	txt := text.New(winBounds.Center(), atlas)
	fmt.Fprint(txt, s.count)
	txt.Draw(overlay, pixel.IM.Moved(s.pos.Center()))
	if s.covered {
		imd.Color = coveredColour
	} else if s.hasBomb {
		imd.Color = lostColour
	} else {
		imd.Color = uncoveredColour
	}

	unit := pixel.V(borderWidth, borderWidth)
	imd.Push(s.pos.Min.Add(unit), s.pos.Max.Sub(unit))
	imd.Rectangle(0)
}

func (s *square) shift(v pixel.Vec) pixel.Rect {
	return s.pos.Moved(v)
}

// neighbours returns a list of square which surround this square.  This does
// not include diagonals
func (s *square) neighbours() []*square {
	var ns []*square

	if _, ok := grid[s.shift(up)]; ok {
		ns = append(ns, grid[s.shift(up)])
	}
	if _, ok := grid[s.shift(down)]; ok {
		ns = append(ns, grid[s.shift(down)])
	}
	if _, ok := grid[s.shift(right)]; ok {
		ns = append(ns, grid[s.shift(right)])
	}
	if _, ok := grid[s.shift(left)]; ok {
		ns = append(ns, grid[s.shift(left)])
	}

	return ns
}

func (s *square) allNeighbours() []*square {
	ns := s.neighbours()

	if _, ok := grid[s.shift(upright)]; ok {
		ns = append(ns, grid[s.shift(upright)])
	}
	if _, ok := grid[s.shift(upleft)]; ok {
		ns = append(ns, grid[s.shift(upleft)])
	}
	if _, ok := grid[s.shift(downright)]; ok {
		ns = append(ns, grid[s.shift(downright)])
	}
	if _, ok := grid[s.shift(downleft)]; ok {
		ns = append(ns, grid[s.shift(downleft)])
	}

	return ns
}

// click returns if the click has lost the game
func (s *square) click() bool {
	s.covered = false

	if s.hasBomb {
		return true
	}

	// uncover neighbours
	s.reveal()

	return false
}

func (s *square) reveal() {
	for _, n := range s.neighbours() {
		if !n.hasBomb {
			n.covered = false
		}

		if n.count == 0 {
			n.reveal()
		}
	}
}

func (s *square) flag() {
	s.flagged = true
}

// generateGrid creates and allocates the grid
func generateGrid() {
	m := make(map[pixel.Rect]*square)
	var keys []pixel.Rect

	for x := 0.; x < winBounds.Max.X; x += gridSize {
		for y := 0.; y < winBounds.Max.Y; y += gridSize {
			r := pixel.R(x, y, x+gridSize, y+gridSize)
			m[r] = &square{
				pos:     r,
				covered: true,
			}
			keys = append(keys, r)
		}
	}

	grid = m

	// Allocate bombs
	for i := 0; i < numBombs; i++ {
		k := r.Intn(len(keys))
		m[keys[k]].hasBomb = true
	}

	// Allocate counts
	for _, s := range m {
		for _, n := range s.allNeighbours() {
			if n.hasBomb {
				s.count++
			}
		}
	}
}
