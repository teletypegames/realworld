package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptClubAdmission,
	Actions: inkwell.Seq(
		inkwell.Say(inc.Shopkeeper, inkwell.Text{
			{Lang: "en", Text: "Behind the curtain. Down the step. Knock twice, or don't, nobody there to hear it."},
			{Lang: "hu", Text: "A függöny mögött. Le a lépcsőn. Kopogjon kétszer, vagy ne, nincs ott senki, aki hallaná."},
		}),
		inkwell.SetFlag(inc.FlagClubEntry),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "A shop with a back room is never just a shop. Write that down somewhere I can't see it."},
			{Lang: "hu", Text: "Egy bolt hátsó szobával sosem csak egy bolt. Írd fel valahová, ahol nem látom."},
		}),
	),
})
