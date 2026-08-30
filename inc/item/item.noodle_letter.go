package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Item{
		Name:        inc.ItemNoodleLetter,
		Description: "Noodle's letter",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(inc.Paul, "“Come out, Paul. Not a phone conversation.” That's all it says."),
			inkwell.Say(inc.Dex, "And now you're here. And he isn't."),
		),
	})
}
