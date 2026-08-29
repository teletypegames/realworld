package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/world"
)

// blackMarketArmilla is flavour, not a quest trigger: the gap between
// legitimised technology and street reality.
func blackMarketArmilla() inkwell.Item {
	return inkwell.Item{
		Name:        names.ItemBlackArmilla,
		Description: "black-market Armilla",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(names.Paul, "Jailbroken. Pushes ads, but it was cheap."),
			world.TapeSay(names.Dex, "Two slots, and both of them lie about the temperature. Don't trade me in for it."),
		),
		OnUseWith: map[string]inkwell.Action{
			// An AUTHORED failure. Nothing is consumed, the character reacts,
			// the tape remarks. Never a generic error line.
			names.ItemNoodleLetter: inkwell.Seq(
				inkwell.Say(names.Paul, "A letter and a bracelet. Brilliant."),
				world.TapeSay(names.Dex, "Nothing. Which is what I'd charge for it."),
			),
		},
	}
}
