package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	ItemManager.Register(Item{
		Name:        ItemNoodleLetter,
		Description: "Noodle's letter",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(Paul, "“Come out, Paul. Not a phone conversation.” That's all it says."),
			TapeSay(Dex, "And now you're here. And he isn't."),
		),
	})
}
