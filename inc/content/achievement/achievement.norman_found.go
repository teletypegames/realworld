package achievement

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var NormanFound = AchievementManager.Register(&Achievement{
	Name: "NORMAN_FOUND",
	Title: Text{
		{Lang: "en", Text: "Not finding him. Getting him out."},
		{Lang: "hu", Text: "Nem megtalálni. Kihozni."},
	},
	Text: Text{
		{Lang: "en", Text: "Paul finds Norman, and the job changes."},
		{Lang: "hu", Text: "Paul megtalálja Normant, és a feladat megváltozik."},
	},
	When: inkwell.Flag(inc.FlagNormanFound),
})
