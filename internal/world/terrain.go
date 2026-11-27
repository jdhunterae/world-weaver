package world

type Terrain uint8

const (
	Unknown Terrain = iota
	Grassland
	Forest
	Mountain
	Water
	Barren
)

func (t Terrain) String() string {
	switch t {
	case Grassland:
		return "Grassland"
	case Forest:
		return "Forest"
	case Mountain:
		return "Mountain"
	case Water:
		return "Water"
	case Barren:
		return "Barren"
	case Unknown:
		return "Unknown"
	default:
		return "Invalid"
	}
}

func (t Terrain) Symbol() rune {
	switch t {
	case Grassland:
		return '.'
	case Forest:
		return 'T'
	case Mountain:
		return '^'
	case Water:
		return '~'
	case Barren:
		return ':'
	case Unknown:
		return '?'
	default:
		return '!'
	}
}

func (t Terrain) IsValid() bool {
	switch t {
	case Grassland, Forest, Mountain, Water, Barren:
		return true
	default:
		return false
	}
}
