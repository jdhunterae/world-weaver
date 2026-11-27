package world

import "fmt"

type Position struct {
	X, Y int
}

func (p Position) Offset(dx, dy int) Position {
	return Position{X: p.X + dx, Y: p.Y + dy}
}

func (p Position) ManhattanDistance(other Position) int {
	dx := p.X - other.X
	if dx < 0 {
		dx = -dx
	}

	dy := p.Y - other.Y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
