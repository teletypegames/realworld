package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/world"
)

// noodleLetter starts the story: it is what brings Paul to the city.
func noodleLetter() inkwell.Item {
	return inkwell.Item{
		Name:        names.ItemNoodleLetter,
		Description: "Noodle's letter",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(names.Paul, "“Come out, Paul. Not a phone conversation.” That's all it says."),
			world.TapeSay(names.Dex, "And now you're here. And he isn't."),
		),
	}
}
