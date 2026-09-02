package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptSecretRoomEmpty,
	Actions: inkwell.Seq(
		inkwell.SetFlag(inc.FlagRoomRaided),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Stripped. Shelves bare, drawers out, evidence tags on the floor."},
			{Lang: "hu", Text: "Kipucolva. Üres polcok, kihúzott fiókok, bűnjelcédulák a padlón."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "They came back for the room while you were learning the code. Whatever he left you here is in a box on Valencia."},
			{Lang: "hu", Text: "Visszajöttek a szobáért, amíg te a kódot tanultad. Amit itt hagyott neked, az egy dobozban van a Valencia utcán."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Next time I'll know the code before I need it."},
			{Lang: "hu", Text: "Legközelebb előbb tudom a kódot, mint hogy kellene."},
		}),
	),
})
