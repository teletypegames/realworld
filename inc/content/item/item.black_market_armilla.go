package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ItemManager.Register(Item{
	Name: inc.ItemBlackArmilla,
	Description: inkwell.Text{
		{Lang: "en", Text: "black-market Armilla"},
		{Lang: "hu", Text: "feketepiaci Armilla"},
	},
	OnUseSelf: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Jailbroken. Pushes ads, but it was cheap."},
			{Lang: "hu", Text: "Feltört. Reklámokat tol, de olcsó volt."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Two slots, and both of them lie about the temperature. Don't trade me in for it."},
			{Lang: "hu", Text: "Két foglalat, és mindkettő hazudik a hőmérsékletről. Ne cserélj be érte."},
		}),
	),
	OnUseWith: map[string]inkwell.Action{
		inc.ItemNoodleLetter: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "A letter and a bracelet. Brilliant."},
				{Lang: "hu", Text: "Egy levél és egy karperec. Zseniális."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Nothing. Which is what I'd charge for it."},
				{Lang: "hu", Text: "Semmi. Ennyit is kérnék érte."},
			}),
		),
	},
})
