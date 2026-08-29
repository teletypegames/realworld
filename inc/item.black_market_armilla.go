package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// blackMarketArmilla is flavour, not a quest trigger: the gap between
// legitimised technology and street reality.
func itemBlackMarketArmilla() inkwell.Item {
	return inkwell.Item{
		Name:        ItemBlackArmilla,
		Description: "black-market Armilla",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(Paul, "Jailbroken. Pushes ads, but it was cheap."),
			TapeSay(Dex, "Two slots, and both of them lie about the temperature. Don't trade me in for it."),
		),
		OnUseWith: map[string]inkwell.Action{
			// An AUTHORED failure. Nothing is consumed, the character reacts,
			// the tape remarks. Never a generic error line.
			ItemNoodleLetter: inkwell.Seq(
				inkwell.Say(Paul, "A letter and a bracelet. Brilliant."),
				TapeSay(Dex, "Nothing. Which is what I'd charge for it."),
			),
		},
	}
}
