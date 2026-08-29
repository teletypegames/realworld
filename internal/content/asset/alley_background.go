package asset

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
)

// alleyBackground is the alley behind Noodle's house (beat 3).
func alleyBackground() inkwell.Asset {
	return inkwell.Asset{
		Name: names.AssetAlleyBG,
		Path: "assets/bg/alley.png",
		Kind: inkwell.AssetImage,
	}
}
