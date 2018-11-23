package main

import (
	"fmt"
	"image/color"
	"math"
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
	// maxReveal is the maximum tiles revealed in one go
	maxReveal = 15
)

var (
	numBombs = int(math.Floor((winBounds.Max.X * winBounds.Max.Y * 2) / (math.Pow(gridSize, 2) * 10)))

	borderColour    = color.RGBA{0xd2, 0xd5, 0xb6, 0xff}
	coveredColour   = color.RGBA{0x94, 0x94, 0x94, 0xff}
	uncoveredColour = color.RGBA{0xc6, 0xc9, 0xac, 0xff}
	lostColour      = color.RGBA{0xb3, 0x00, 0x00, 0xff}
	flaggedColour   = color.RGBA{0x00, 0x00, 0x00, 0xff}

	src = rand.NewSource(time.Now().Unix())
	r   = rand.New(src)

	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Direction vectors
	// lateral
	up    = pixel.V(0, 1)
	down  = pixel.V(0, -1)
	left  = pixel.V(-1, 0)
	right = pixel.V(1, 0)

	// diagonal
	upright   = pixel.V(1, 1)
	upleft    = pixel.V(-1, 1)
	downright = pixel.V(1, -1)
	downleft  = pixel.V(-1, -1)
)

func vecToRowCol(pos pixel.Vec) pixel.Vec {
	intX := int(pos.X) / int(gridSize)
	intY := int(pos.Y) / int(gridSize)

	return pixel.V(
		float64(intX),
		float64(intY),
	)
}

type square struct {
	hasBomb bool
	gridPos pixel.Vec
	pos     pixel.Rect
	covered bool
	flagged bool
	// count is the number of surrounding squares have bombs
	count int
}

func (s *square) Draw(imd *imdraw.IMDraw, overlay *pixelgl.Canvas) {
	if !s.covered && s.count > 0 {
		txt := text.New(pixel.ZV, atlas)
		fmt.Fprint(txt, s.count)
		txt.Draw(overlay, pixel.IM.Moved(s.pos.Center()))
	}

	if s.flagged {
		imd.Color = flaggedColour
	} else if s.covered {
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

func (s *square) shift(v pixel.Vec) pixel.Vec {
	return s.gridPos.Add(v)
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
	// Don't try uncover if already uncovered or flagged
	if !s.covered || s.flagged {
		return false
	}

	s.covered = false

	if s.hasBomb {
		return true
	}

	// uncover neighbours
	s.reveal(0)

	return false
}

func (s *square) reveal(calls int) {
	for _, n := range s.neighbours() {
		calls++
		if calls > maxReveal {
			return
		}

		if !n.hasBomb {
			n.covered = false
		}

		if n.count == 0 {
			n.reveal(calls)
		}
	}
}

func (s *square) flag() bool {
	s.flagged = !s.flagged

	for _, n := range grid {
		if n.hasBomb && !n.flagged {
			return false
		}
	}
	return true
}

// generateGrid creates and allocates the grid
func generateGrid() {
	m := make(map[pixel.Vec]*square)
	var keys []pixel.Vec

	for x := 0.; x < winBounds.Max.X/gridSize; x++ {
		for y := 0.; y < winBounds.Max.Y/gridSize; y++ {
			v := pixel.V(x, y)
			m[v] = &square{
				pos:     pixel.R(x*gridSize, y*gridSize, (x*gridSize)+gridSize, (y*gridSize)+gridSize),
				gridPos: v,
				covered: true,
			}
			keys = append(keys, v)
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
