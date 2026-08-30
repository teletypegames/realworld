package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Item{
		Name:        inc.ItemMysteryTape,
		Description: "unmarked Personal Tape",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(inc.Paul, "Password-locked. Of course."),
			inkwell.Say(inc.Dex, "Put it in the second slot if you care that much. I did warn you."),
		),
	})
}
