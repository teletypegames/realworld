package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// mysteryTape is the engine of the investigation. Password-locked; it only
// starts speaking after the club trigger (beat 6).
func itemMysteryTape() inkwell.Item {
	return inkwell.Item{
		Name:        ItemMysteryTape,
		Description: "unmarked Personal Tape",
		OnUseSelf: inkwell.Seq(
			inkwell.Say(Paul, "Password-locked. Of course."),
			TapeSay(Dex, "Put it in the second slot if you care that much. I did warn you."),
		),
	}
}
