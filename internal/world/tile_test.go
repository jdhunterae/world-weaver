package world

import "testing"

func TestTileTerrain(t *testing.T) {
	var tile Tile
	tile.terrain = Grassland
	terrain := tile.Terrain()

	if terrain != Grassland {
		t.Errorf("expected %q, got %q", terrain, Grassland)
	}
}

func TestTileZeroValue(t *testing.T) {
	var tile Tile
	terrain := tile.Terrain()

	if terrain != Unknown {
		t.Errorf("expected tile to contain Unknown terrain, got %d", terrain)
	}

	if terrain.IsValid() {
		t.Error("uninitialized tile should not contain valid terrain")
	}
}
