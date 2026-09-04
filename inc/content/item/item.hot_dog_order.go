package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var HotDogOrder = ItemManager.Register(&Item{
	Name: "hot_dog_order",
	Description: Text{
		{Lang: "en", Text: "tray of hot dogs"},
		{Lang: "hu", Text: "tálca hotdog"},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Six with everything. Getting cold. I'm not eating the cover story."},
			{Lang: "hu", Text: "Hat mindennel. Hűl. Nem eszem meg a fedősztorit."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Deliver it upstairs before it stops being a reason to be there."},
			{Lang: "hu", Text: "Vidd fel, mielőtt már nem indok arra, hogy ott legyél."},
		}),
	),
})
