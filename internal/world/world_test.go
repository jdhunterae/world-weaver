package world

import (
	"math"
	"testing"
)

func TestWorldTileSize(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		expected int
	}{
		{"1x1", 1, 1, 1},
		{"5x5", 5, 5, 25},
		{"10x4", 10, 4, 40},
		{"4x10", 4, 10, 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			world, err := NewWorld(tt.width, tt.height)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if world == nil {
				t.Fatal("unexpected error: world is nil")
			}

			actual := len(world.tiles)

			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
			if world.Width() != tt.width {
				t.Errorf("expected width %v, got %v", tt.width, world.Width())
			}
			if world.Height() != tt.height {
				t.Errorf("expected height %v, got %v", tt.height, world.Height())
			}
		})
	}
}

func TestWorldInvalidSizes(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
	}{
		{"0x5", 0, 5},
		{"5x0", 5, 0},
		{"0x0", 0, 0},
		{"-1x5", -1, 5},
		{"5x-1", 5, -1},
		{"-1x-1", -1, -1},
		{"overflow_negative", math.MaxInt, 2},
		{"overflow_positive", math.MaxInt, math.MaxInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			world, err := NewWorld(tt.width, tt.height)

			if err == nil {
				t.Fatal("expected an error")
			}
			if world != nil {
				t.Fatalf("expected world to be nil, got %v", world)
			}
		})
	}
}

func TestWorldInitialTerrain(t *testing.T) {
	world, err := NewWorld(5, 5)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if world == nil {
		t.Fatal("unexpected error: world is nil")
	}

	for i, tile := range world.tiles {
		if tile.Terrain() != Unknown {
			t.Errorf("tile %d: expected %v, got %v", i, Unknown, tile.Terrain())
		}
	}
}
