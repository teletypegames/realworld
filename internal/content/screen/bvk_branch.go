package screen

import "git.teletypegames.org/games/realworld/internal/names"

// bvkBranch is a 💡 suggestion: the local certification office, imprimatur
// stamps and licence checks and numbered waiting rooms — everything a man with
// a jailbroken two-slot Armilla should stay away from
// (wiki: entities#potenciális-további-helyszínek).
func bvkBranch() screen {
	return screen{
		name:       names.ScreenBVKBranch,
		title:      "BVK — SAN FRANCISCO BRANCH",
		background: names.BgBVKBranch,
		onSelector: true,
	}
}
