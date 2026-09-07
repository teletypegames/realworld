package achievement

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var NoodleReleased = AchievementManager.Register(&Achievement{
	Name: "NOODLE_RELEASED",
	Title: Text{
		{Lang: "en", Text: "Out of the box"},
		{Lang: "hu", Text: "Kint a dobozból"},
	},
	Text: Text{
		{Lang: "en", Text: "Noodle walks out of the police station."},
		{Lang: "hu", Text: "Noodle kisétál a rendőrségről."},
	},
	When: inkwell.Flag(inc.FlagNoodleReleased),
})
