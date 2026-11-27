package world

import "testing"

func TestPositionString(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		expected string
	}{
		{"(0,0)", Position{X: 0, Y: 0}, "(0,0)"},
		{"(3,5)", Position{X: 3, Y: 5}, "(3,5)"},
		{"(-3,5)", Position{X: -3, Y: 5}, "(-3,5)"},
		{"(3,-5)", Position{X: 3, Y: -5}, "(3,-5)"},
		{"(-3,-5)", Position{X: -3, Y: -5}, "(-3,-5)"},
		{"(100,250)", Position{X: 100, Y: 250}, "(100,250)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.position.String()

			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestPositionOffset(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		dx       int
		dy       int
		expected Position
	}{
		{"(0,0)->(0,0)", Position{X: 0, Y: 0}, 0, 0, Position{X: 0, Y: 0}},
		{"(2,3)->(1,0)", Position{X: 2, Y: 3}, 1, 0, Position{X: 3, Y: 3}},
		{"(2,3)->(0,1)", Position{X: 2, Y: 3}, 0, 1, Position{X: 2, Y: 4}},
		{"(2,3)->(-1,0)", Position{X: 2, Y: 3}, -1, 0, Position{X: 1, Y: 3}},
		{"(2,3)->(0,-1)", Position{X: 2, Y: 3}, 0, -1, Position{X: 2, Y: 2}},
		{"(0,0)->(-1,-1)", Position{X: 0, Y: 0}, -1, -1, Position{X: -1, Y: -1}},
		{"(4,5)->(3,-2)", Position{X: 4, Y: 5}, 3, -2, Position{X: 7, Y: 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := tt.position // Store the original position to check for immutability
			actual := tt.position.Offset(tt.dx, tt.dy)

			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}

			if tt.position != original {
				t.Errorf("Offset modified the original position: expected %v, got %v", original, tt.position)
			}
		})
	}
}

func TestPositionManhattanDistance(t *testing.T) {
	tests := []struct {
		name      string
		position1 Position
		position2 Position
		expected  int
	}{
		{"(0,0)->(0,0)", Position{X: 0, Y: 0}, Position{X: 0, Y: 0}, 0},
		{"(0,0)->(5,0)", Position{X: 0, Y: 0}, Position{X: 5, Y: 0}, 5},
		{"(0,0)->(0,5)", Position{X: 0, Y: 0}, Position{X: 0, Y: 5}, 5},
		{"(1,2)->(4,6)", Position{X: 1, Y: 2}, Position{X: 4, Y: 6}, 7},
		{"(0,0)->(-3,-4)", Position{X: 0, Y: 0}, Position{X: -3, Y: -4}, 7},
		{"(-2,-3)->(2,3)", Position{X: -2, Y: -3}, Position{X: 2, Y: 3}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.position1.ManhattanDistance(tt.position2)
			reversed := tt.position2.ManhattanDistance(tt.position1)

			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}

			if reversed != tt.expected {
				t.Errorf("expected %d, got %d (reversed)", tt.expected, reversed)
			}
		})
	}
}
