package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	ItemManager.Register(Item{
		Name:        ItemBlackArmilla,
		Description: "black-market Armilla",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(Paul, "Jailbroken. Pushes ads, but it was cheap."),
			TapeSay(Dex, "Two slots, and both of them lie about the temperature. Don't trade me in for it."),
		),
		OnUseWith: map[string]inkwell.Action{
			ItemNoodleLetter: inkwell.Seq(
				inkwell.Say(Paul, "A letter and a bracelet. Brilliant."),
				TapeSay(Dex, "Nothing. Which is what I'd charge for it."),
			),
		},
	})
}
