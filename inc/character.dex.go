package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// dex is the personality imprint of a dead friend, permanently in slot 1 of
// Paul's Armilla. Constant commentary, hint system, dialogue partner.
func characterDex() inkwell.Character {
	return inkwell.Character{
		Name:        Dex,
		SpeechColor: Amber,
	}
}
