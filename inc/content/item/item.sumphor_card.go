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
	Lore: Text{
		{Lang: "en", Text: "Sumphor's staff card, Tower B, with the photo of a man who still slept at night. Opens the lift to the abandoned office and a door at the foot of the bridge."},
		{Lang: "hu", Text: "Sumphor dolgozói kártyája, B torony, egy ember fotójával, aki még aludt éjszakánként. Nyitja a liftet az elhagyatott irodához és egy ajtót a híd lábánál."},
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
