package fsettings

import "encoding/json"

// This struct is used to generate map generation settings JSON files.
type MapGenSettings struct {
	TerrainSegmentation string            `json:"terrain_segmentation"`
	Water               string            `json:"water"`
	Width               uint              `json:"width"`
	Height              uint              `json:"height"`
	StartingArea        string            `json:"starting_area"`
	PeacefulMode        bool              `json:"peaceful_mode"`
	AutoplaceControls   AutoplaceControls `json:"autoplace_controls"`
}

func (fs MapGenSettings) String() string {
	j, _ := json.MarshalIndent(fs, "", "   ")
	return string(j)
}

type AutoplaceControls struct {
	Coal      PlacementControl `json:"coal"`
	CopperOre PlacementControl `json:"copper-ore"`
	CrudeOil  PlacementControl `json:"crude-oil"`
	EnemyBase PlacementControl `json:"enemy-base"`
	IronOre   PlacementControl `json:"iron-ore"`
	Stone     PlacementControl `json:"stone"`
}

type PlacementControl struct {
	Frequency string `json:"frequency"`
	Size      string `json:"size"`
	Richness  string `json:"richness"`
}

// Return a new MapGenSettings struct with reasonable defaults.
// These do not match the example values in the Factorio distribution.
func NewMapGenSettings() MapGenSettings {
	normalLevels := PlacementControl{Frequency: "normal", Size: "normal", Richness: "normal"}
	return MapGenSettings{
		TerrainSegmentation: "normal",
		Water:               "normal",
		Width:               uint(0),
		Height:              uint(0),
		StartingArea:        "normal",
		PeacefulMode:        false,
		AutoplaceControls: AutoplaceControls{
			Coal:      normalLevels,
			CopperOre: normalLevels,
			CrudeOil:  normalLevels,
			EnemyBase: normalLevels,
			IronOre:   normalLevels,
			Stone:     normalLevels,
		},
	}
}
