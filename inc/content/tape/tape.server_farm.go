package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var ServerFarm = TapeManager.Register(&Tape{
	Name:    "tape_server_farm",
	Item:    item.ServerFarmTape,
	Speaker: speaker.TapeServerFarm,
	Dialog:  dialog.ServerFarmTape,
	OnInsert: inkwell.Seq(
		inkwell.Say(speaker.TapeServerFarm, inkwell.Text{
			{Lang: "en", Text: "CREDENTIAL LOADED. DOOR GROUP: DATACENTER-WEST."},
			{Lang: "hu", Text: "HITELESÍTŐ BETÖLTVE. AJTÓCSOPORT: ADATKÖZPONT-NYUGAT."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "No personality at all. I'm almost jealous."},
			{Lang: "hu", Text: "Semmi személyiség. Majdnem irigylem."},
		}),
	),
})
