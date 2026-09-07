package achievement

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var BVKLogged = AchievementManager.Register(&Achievement{
	Name: "BVK_LOGGED",
	Title: Text{
		{Lang: "en", Text: "What he came for"},
		{Lang: "hu", Text: "Amiért jött"},
	},
	Text: Text{
		{Lang: "en", Text: "Paul leaves the BVK with what he came for."},
		{Lang: "hu", Text: "Paul azzal jön ki a BVK-ból, amiért jött."},
	},
	When: inkwell.Flag(inc.FlagBVKLogged),
})
