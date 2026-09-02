package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemOfficeCards,
	Description: inkwell.Text{
		{Lang: "en", Text: "office workers' cards"},
		{Lang: "hu", Text: "irodai dolgozók kártyái"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Six staff cards on a lanyard. Nobody has noticed. Nobody will."},
			{Lang: "hu", Text: "Hat dolgozói kártya egy szalagon. Senki nem vette észre. Nem is fogja."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Half price on hot dogs. That is the sum of what you've stolen. I can live with it."},
			{Lang: "hu", Text: "Féláras hotdog. Ennyit loptál összesen. Ezzel együtt tudok élni."},
		}),
	),
})
