package fsettings

import "encoding/json"

// This struct is used to generate server settings JSON files.
type FactorioServerSettings struct {
	Name                                 string             `json:"name"`
	Description                          string             `json:"description"`
	Tags                                 []string           `json:"tags"`
	MaxPlayers                           uint               `json:"max_players"`
	Visibility                           VisibilitySettings `json:"visibility"`
	Username                             string             `json:"username"`
	Password                             string             `json:"password"`
	Token                                string             `json:"token"`
	GamePassword                         string             `json:"game_password"`
	RequireUserVerification              bool               `json:"require_user_verification"`
	MaxUploadInKilobytesPerSecond        uint               `json:"max_upload_in_kilobytes_per_second"`
	MinimumLatencyInTicks                uint               `json:"minimum_latency_in_ticks"`
	IgnorePlayerLimitForReturningPlayers bool               `json:"ignore_player_limit_for_returning_players"`
	AllowCommands                        string             `json:"allow_commands"`
	AutoSaveInterval                     uint               `json:"autosave_interval"`
	AutoSaveSlots                        uint               `json:"autosave_slots"`
	AfkAutoKickInterval                  uint               `json:"afk_autokick_interval"`
	AutoPause                            bool               `json:"auto_pause"`
	OnlyAdminsCanPauseGame               bool               `json:"only_admins_can_pause_the_game"`
	AutoSaveOnlyOnServer                 bool               `json:"autosave_only_on_server"`
	Admins                               []string           `json:"admins,omitempty"`
}

func (fs FactorioServerSettings) String() string {
	j, _ := json.MarshalIndent(fs, "", "   ")
	return string(j)
}

type VisibilitySettings struct {
	Public bool `json:"public"`
	Lan    bool `json:"lan"`
}

// Return a new FactorioSettings struct with reasonable defaults.
// These do not match the example values in the Factorio distribution.
func NewFactorioServerSettings() FactorioServerSettings {
	return FactorioServerSettings{
		Name:                   "My Factorio Server",
		Description:            "I should really change my description",
		Visibility:             VisibilitySettings{Public: false, Lan: false},
		AllowCommands:          "admins-only",
		AutoSaveInterval:       uint(10),
		AutoSaveSlots:          uint(5),
		AutoPause:              true,
		OnlyAdminsCanPauseGame: true,
		AutoSaveOnlyOnServer:   true,
		Tags:                   []string{},
		Admins:                 []string{},
	}
}
