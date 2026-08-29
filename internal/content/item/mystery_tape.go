package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/world"
)

// mysteryTape is the engine of the investigation. Password-locked; it only
// starts speaking after the club trigger (beat 6).
func mysteryTape() inkwell.Item {
	return inkwell.Item{
		Name:        names.ItemMysteryTape,
		Description: "unmarked Personal Tape",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(names.Paul, "Password-locked. Of course."),
			world.TapeSay(names.Dex, "Put it in the second slot if you care that much. I did warn you."),
		),
	}
}
