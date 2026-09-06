package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var BlackArmilla = ItemManager.Register(&Item{
	Name: "black_market_armilla",
	Description: Text{
		{Lang: "en", Text: "black-market Armilla"},
		{Lang: "hu", Text: "feketepiaci Armilla"},
	},
	Lore: Text{
		{Lang: "en", Text: "A jailbroken Armilla off the market. Two slots, both lying about the temperature, and it pushes ads. It was cheap."},
		{Lang: "hu", Text: "Egy feltört Armilla a piacról. Két foglalat, mindkettő hazudik a hőmérsékletről, és reklámot tol. Olcsó volt."},
	},
	OnUseSelf: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Jailbroken. Pushes ads, but it was cheap."},
			{Lang: "hu", Text: "Feltört. Reklámokat tol, de olcsó volt."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Two slots, and both of them lie about the temperature. Don't trade me in for it."},
			{Lang: "hu", Text: "Két foglalat, és mindkettő hazudik a hőmérsékletről. Ne cserélj be érte."},
		}),
	),
	OnUseWith: map[*inkwell.Item]Action{
		NoodleLetter: inkwell.Seq(
			character.Paul.Say(Text{
				{Lang: "en", Text: "A letter and a bracelet. Brilliant."},
				{Lang: "hu", Text: "Egy levél és egy karperec. Zseniális."},
			}),
			speaker.Dex.Say(Text{
				{Lang: "en", Text: "Nothing. Which is what I'd charge for it."},
				{Lang: "hu", Text: "Semmi. Ennyit is kérnék érte."},
			}),
		),
	},
})
