package achievement

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var TapeSpeaks = AchievementManager.Register(&Achievement{
	Name: "TAPE_SPEAKS",
	Title: Text{
		{Lang: "en", Text: "The tape speaks"},
		{Lang: "hu", Text: "A kazetta megszólal"},
	},
	Text: Text{
		{Lang: "en", Text: "The mystery tape says its first line."},
		{Lang: "hu", Text: "A rejtélyes kazetta kimondja az első sorát."},
	},
	When: inkwell.Flag(inc.FlagTapeSpeaks),
})
