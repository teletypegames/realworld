package achievement

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var OfficerFed = AchievementManager.Register(&Achievement{
	Name: "OFFICER_FED",
	Title: Text{
		{Lang: "en", Text: "Breakfast for the law"},
		{Lang: "hu", Text: "Reggeli a törvénynek"},
	},
	Text: Text{
		{Lang: "en", Text: "Paul hands the officer at Noodle's door his breakfast."},
		{Lang: "hu", Text: "Paul odaadja a reggelijét a Noodle ajtaja előtt álló rendőrnek."},
	},
	When: inkwell.Flag(inc.FlagOfficerFed),
})
