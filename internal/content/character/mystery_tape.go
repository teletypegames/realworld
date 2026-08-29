package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/names"
	"realworld/internal/theme"
)

// mysteryTape is the tape found in the alley — in truth Norman's Personal
// Tape, but the player only learns that in the finale.
//
// Its voice is a colder, greyer amber than Dex's. That pays off in reverse:
// it never once spoke in the same voice as the friend.
func mysteryTape() inkwell.Character {
	return inkwell.Character{
		Name:        names.TapeMystery,
		SpeechColor: theme.RGB(0xB9A98A),
	}
}
