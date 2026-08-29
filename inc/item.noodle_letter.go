package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// noodleLetter starts the story: it is what brings Paul to the city.
func itemNoodleLetter() inkwell.Item {
	return inkwell.Item{
		Name:        ItemNoodleLetter,
		Description: "Noodle's letter",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(Paul, "“Come out, Paul. Not a phone conversation.” That's all it says."),
			TapeSay(Dex, "And now you're here. And he isn't."),
		),
	}
}
