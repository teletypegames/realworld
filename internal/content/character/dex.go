package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/theme"
)

// dex is the personality imprint of a dead friend, permanently in slot 1 of
// Paul's Armilla. Constant commentary, hint system, dialogue partner.
func dex() inkwell.Character {
	return inkwell.Character{
		Name:        names.Dex,
		SpeechColor: theme.Amber,
	}
}
