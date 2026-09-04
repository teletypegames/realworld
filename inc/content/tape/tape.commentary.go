package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var Commentary = TapeManager.Register(&Tape{
	Name:    "tape_commentary",
	Item:    item.CommentaryTape,
	Speaker: speaker.TapeCommentary,
	Dialog:  dialog.CommentaryTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(speaker.TapeCommentary, inkwell.Text{
			{Lang: "en", Text: "Recording. If you're hearing this, Paul, you were faster than they were. Good."},
			{Lang: "hu", Text: "Felvétel. Ha ezt hallod, Paul, gyorsabb voltál náluk. Helyes."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Noodle's voice. He recorded this for you. Before."},
			{Lang: "hu", Text: "Noodle hangja. Neked vette fel. Még előtte."},
		}),
	),
})
