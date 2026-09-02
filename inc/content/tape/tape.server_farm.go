package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/theme"
)

var _ = TapeManager.Register(Tape{
	Name: inc.TapeServerFarm,
	Item: inc.ItemServerFarmTape,
	Label: inkwell.Text{
		{Lang: "en", Text: "ACCESS TAPE"},
		{Lang: "hu", Text: "BELÉPŐKAZETTA"},
	},
	Color:  theme.AmberLow,
	Dialog: inc.DialogServerFarmTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(inc.TapeServerFarm, inkwell.Text{
			{Lang: "en", Text: "CREDENTIAL LOADED. DOOR GROUP: DATACENTER-WEST."},
			{Lang: "hu", Text: "HITELESÍTŐ BETÖLTVE. AJTÓCSOPORT: ADATKÖZPONT-NYUGAT."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "No personality at all. I'm almost jealous."},
			{Lang: "hu", Text: "Semmi személyiség. Majdnem irigylem."},
		}),
	),
})
