package speaker

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/system/theme"
)

var Dex = SpeakerManager.Register(&Speaker{
	Name: "dex",
	Label: Text{
		{Lang: "en", Text: "DEX"},
		{Lang: "hu", Text: "DEX"},
	},
	Lore: Text{
		{Lang: "en", Text: "Paul's Personal Tape, slot one, for years. Dry, loyal, and against everything Paul is about to do. The only friend he carries around."},
		{Lang: "hu", Text: "Paul Personal Tape-je, egyes foglalat, évek óta. Száraz, hű, és minden ellen, amit Paul épp tenni készül. Az egyetlen barát, akit magával hord."},
	},
	Delivery: inkwell.DeliveryLog,
	Color:    theme.Amber,
})
