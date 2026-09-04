package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var SumphorCard = ItemManager.Register(&Item{
	Name: "sumphor_card",
	Description: Text{
		{Lang: "en", Text: "Sumphor's access card"},
		{Lang: "hu", Text: "Sumphor belépőkártyája"},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Staff card, Tower B. The photo is of a man who still slept at night."},
			{Lang: "hu", Text: "Dolgozói kártya, B torony. A fényképen egy ember, aki még aludt éjszakánként."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Four-one-nine-one. Don't make me repeat it at the lift."},
			{Lang: "hu", Text: "Négy-egy-kilenc-egy. Ne kelljen a liftnél megismételnem."},
		}),
	),
})
