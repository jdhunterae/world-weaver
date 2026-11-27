package world

import (
	"fmt"
	"math"
)

type World struct {
	width  int
	height int
	tiles  []Tile
}

func NewWorld(width, height int) (*World, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("invalid world dimensions: %dx%d", width, height)
	}
	if width > math.MaxInt/height {
		return nil, fmt.Errorf("world dimensions too large: %dx%d", width, height)
	}

	tiles := make([]Tile, width*height)
	return &World{
		width:  width,
		height: height,
		tiles:  tiles,
	}, nil
}

func (w *World) Width() int {
	return w.width
}

func (w *World) Height() int {
	return w.height
}
