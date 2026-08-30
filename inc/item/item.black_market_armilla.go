package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Item{
		Name:        inc.ItemBlackArmilla,
		Description: "black-market Armilla",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(inc.Paul, "Jailbroken. Pushes ads, but it was cheap."),
			inkwell.Say(inc.Dex, "Two slots, and both of them lie about the temperature. Don't trade me in for it."),
		),
		OnUseWith: map[string]inkwell.Action{
			inc.ItemNoodleLetter: inkwell.Seq(
				inkwell.Say(inc.Paul, "A letter and a bracelet. Brilliant."),
				inkwell.Say(inc.Dex, "Nothing. Which is what I'd charge for it."),
			),
		},
	})
}
