package character

import (
	"realworld/inc/system/theme"
)

var Receptionist = CharacterManager.Register(&Character{
	Name: "receptionist",
	Label: Text{
		{Lang: "en", Text: "receptionist"},
		{Lang: "hu", Text: "recepciós"},
	},
	Lore: Text{
		{Lang: "en", Text: "Thick glasses, buried in paperwork, bored rather than hostile. Everybody who wants into the back of the station has to give her a reason first, and most reasons aren't good enough."},
		{Lang: "hu", Text: "Vastag keretes szemüveg, papírokba temetkezve, unott inkább, mint ellenséges. Aki a kapitányság hátsó részébe akar bejutni, annak előbb indokot kell adnia, és a legtöbb indok nem elég jó."},
	},
	Speed:       80,
	W:           26,
	H:           60,
	SpeechColor: theme.Bone,
})
