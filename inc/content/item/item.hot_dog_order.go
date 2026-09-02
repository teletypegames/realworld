package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemHotDogOrder,
	Description: inkwell.Text{
		{Lang: "en", Text: "tray of hot dogs"},
		{Lang: "hu", Text: "tálca hotdog"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Six with everything. Getting cold. I'm not eating the cover story."},
			{Lang: "hu", Text: "Hat mindennel. Hűl. Nem eszem meg a fedősztorit."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Deliver it upstairs before it stops being a reason to be there."},
			{Lang: "hu", Text: "Vidd fel, mielőtt már nem indok arra, hogy ott legyél."},
		}),
	),
})
