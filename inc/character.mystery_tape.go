package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// mysteryTape is the tape found in the alley — in truth Norman's Personal
// Tape, but the player only learns that in the finale.
//
// Its voice is a colder, greyer amber than Dex's. That pays off in reverse:
// it never once spoke in the same voice as the friend.
func characterMysteryTape() inkwell.Character {
	return inkwell.Character{
		Name:        TapeMystery,
		SpeechColor: RGB(0xB9A98A),
	}
}
