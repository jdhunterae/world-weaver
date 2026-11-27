package world

type Tile struct {
	terrain Terrain
}

func (t Tile) Terrain() Terrain {
	return t.terrain
}
