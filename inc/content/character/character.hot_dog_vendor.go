package character

import (
	"realworld/inc/system/theme"
)

var HotDogVendor = CharacterManager.Register(&Character{
	Name: "hot_dog_vendor",
	Label: Text{
		{Lang: "en", Text: "hot dog vendor"},
		{Lang: "hu", Text: "hotdogárus"},
	},
	Lore: Text{
		{Lang: "en", Text: "Six with everything, on the street below Norman's block. Stamps a loyalty card, remembers where his regulars work, and doesn't mind a madman on the pavement as long as he buys."},
		{Lang: "hu", Text: "Hat mindennel, Norman háza alatt az utcán. Pontgyűjtő kártyát bélyegez, emlékszik, hol dolgoznak a törzsvendégei, és nem zavarja egy őrült a járdán, amíg vásárol."},
	},
	Speed:       80,
	W:           30,
	H:           66,
	SpeechColor: theme.AmberPale,
})
