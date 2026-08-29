package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/names"
	"realworld/internal/theme"
)

// supportTape answers everything with basic information, insufferably. Dex
// hates it. Acquired in Chinatown (beat 5) — not reachable yet.
func supportTape() inkwell.Character {
	return inkwell.Character{
		Name:        names.TapeSupport,
		SpeechColor: theme.RGB(0xE8C86A),
	}
}
