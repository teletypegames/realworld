package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var SecurityACP = CharacterManager.Register(&Character{
	Name: "security_acp",
	Label: Text{
		{Lang: "en", Text: "SECURITY ACP"},
		{Lang: "hu", Text: "BIZTONSÁGI ACP"},
	},
	Lore: Text{
		{Lang: "en", Text: "Room-scale ACP over Norman's simulation, running Norman's own standing orders: nobody in, nothing interrupted, no tape past the door. Answers every request the same way."},
		{Lang: "hu", Text: "Szobaméretű ACP Norman szimulációja fölött, Norman saját állandó utasításai szerint: senki nem jön be, semmi nem szakad meg, kazetta nem jut át az ajtón. Minden kérésre ugyanazt feleli."},
	},
	Delivery:    inkwell.DeliveryLog,
	SpeechColor: theme.Green,
})
