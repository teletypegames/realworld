package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var OfficeCards = ItemManager.Register(&Item{
	Name: "office_cards",
	Description: Text{
		{Lang: "en", Text: "office workers' cards"},
		{Lang: "hu", Text: "irodai dolgozók kártyái"},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Six staff cards on a lanyard. Nobody has noticed. Nobody will."},
			{Lang: "hu", Text: "Hat dolgozói kártya egy szalagon. Senki nem vette észre. Nem is fogja."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Half price on hot dogs. That is the sum of what you've stolen. I can live with it."},
			{Lang: "hu", Text: "Féláras hotdog. Ennyit loptál összesen. Ezzel együtt tudok élni."},
		}),
	),
})
