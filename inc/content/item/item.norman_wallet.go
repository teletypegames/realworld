package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var NormanWallet = ItemManager.Register(&Item{
	Name: "norman_wallet",
	Description: inkwell.Text{
		{Lang: "en", Text: "Norman's wallet"},
		{Lang: "hu", Text: "Norman tárcája"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Hot dog card, eight stamps. Behind it an access card for the complex up the hill. No cash."},
			{Lang: "hu", Text: "Hotdogkártya, nyolc pecsét. Mögötte egy belépőkártya a dombon lévő komplexumhoz. Készpénz semmi."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "The card with the stamps is the one that tells you where he ate. The other one tells you where he worked."},
			{Lang: "hu", Text: "A pecsétes kártya mondja meg, hol evett. A másik, hogy hol dolgozott."},
		}),
	),
	OnUseWith: map[*inkwell.Item]inkwell.Action{
		NoodleLetter: inkwell.Seq(
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Noodle's letter and Norman's wallet. Two men who wanted a beer together."},
				{Lang: "hu", Text: "Noodle levele és Norman tárcája. Két ember, akik együtt akartak sörözni."},
			}),
			inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "And neither of them is here. Put them away."},
				{Lang: "hu", Text: "És egyikük sincs itt. Tedd el őket."},
			}),
		),
	},
})
