package world

import "testing"

func TestTerrainString(t *testing.T) {
	tests := []struct {
		name     string
		terrain  Terrain
		expected string
	}{
		{"Grassland", Grassland, "Grassland"},
		{"Forest", Forest, "Forest"},
		{"Mountain", Mountain, "Mountain"},
		{"Water", Water, "Water"},
		{"Barren", Barren, "Barren"},
		{"Unknown", Unknown, "Unknown"},
		{"Invalid_FirstUnsupported", Terrain(6), "Invalid"},
		{"Invalid_Arbitrary", Terrain(99), "Invalid"},
		{"Invalid_MaxUint8", Terrain(255), "Invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.terrain.String()

			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestTerrainSymbol(t *testing.T) {
	tests := []struct {
		name     string
		terrain  Terrain
		expected rune
	}{
		{"Grassland", Grassland, '.'},
		{"Forest", Forest, 'T'},
		{"Mountain", Mountain, '^'},
		{"Water", Water, '~'},
		{"Barren", Barren, ':'},
		{"Unknown", Unknown, '?'},
		{"Invalid_FirstUnsupported", Terrain(6), '!'},
		{"Invalid_Arbitrary", Terrain(99), '!'},
		{"Invalid_MaxUint8", Terrain(255), '!'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.terrain.Symbol()

			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestTerrainIsValid(t *testing.T) {
	tests := []struct {
		name     string
		terrain  Terrain
		expected bool
	}{
		{"Grassland", Grassland, true},
		{"Forest", Forest, true},
		{"Mountain", Mountain, true},
		{"Water", Water, true},
		{"Barren", Barren, true},
		{"Unknown", Unknown, false},
		{"Invalid_FirstUnsupported", Terrain(6), false},
		{"Invalid_Arbitrary", Terrain(99), false},
		{"Invalid_MaxUint8", Terrain(255), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.terrain.IsValid()

			if actual != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, actual)
			}
		})
	}
}

func TestTerrainZeroValue(t *testing.T) {
	var terrain Terrain

	if terrain != Unknown {
		t.Errorf("expected Unknown terrain, got %d", terrain)
	}

	if terrain.String() != "Unknown" {
		t.Errorf("expected Unknown, got %q", terrain.String())
	}

	if terrain.Symbol() != '?' {
		t.Errorf("expected '?', got %q", terrain.Symbol())
	}

	if terrain.IsValid() {
		t.Error("uninitialized terrain should not be valid")
	}
}
