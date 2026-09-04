package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var ServerFarmTape = ItemManager.Register(&Item{
	Name: "server_farm_tape",
	Description: Text{
		{Lang: "en", Text: "datacenter access tape"},
		{Lang: "hu", Text: "adatközponti belépőkazetta"},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Somebody's badge, cloned onto a tape. No name on it."},
			{Lang: "hu", Text: "Valakinek a kitűzője, kazettára klónozva. Név nincs rajta."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "It's a key, not a person. Put it in slot two when a door asks for it, and take it out after."},
			{Lang: "hu", Text: "Ez kulcs, nem személy. Tedd a kettes foglalatba, amikor egy ajtó kéri, aztán vedd ki."},
		}),
	),
})
