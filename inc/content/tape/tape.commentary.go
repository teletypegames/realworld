package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeCommentary,
	Item: inc.ItemCommentaryTape,
	Label: inkwell.Text{
		{Lang: "en", Text: "DEV COMMENTARY"},
		{Lang: "hu", Text: "FEJLESZTŐI KOMMENTÁR"},
	},
	Color:  theme.Amber,
	Dialog: inc.DialogCommentaryTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(inc.TapeCommentary, inkwell.Text{
			{Lang: "en", Text: "Recording. If you're hearing this, Paul, you were faster than they were. Good."},
			{Lang: "hu", Text: "Felvétel. Ha ezt hallod, Paul, gyorsabb voltál náluk. Helyes."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Noodle's voice. He recorded this for you. Before."},
			{Lang: "hu", Text: "Noodle hangja. Neked vette fel. Még előtte."},
		}),
	),
})
