package screen

import "git.teletypegames.org/games/realworld/internal/names"

// curatorShop is a 💡 suggestion: the last workshop of the declining curators'
// guild — quartz tuning, tape repair, craft handed down for generations. The
// man who built Paul's second slot (wiki: entities#potenciális-további-helyszínek).
func curatorShop() screen {
	return screen{
		name:       names.ScreenCuratorShop,
		title:      "CURATOR SERVICE — THE MASTER'S WORKSHOP",
		background: names.BgCuratorShop,
		onSelector: true,
	}
}
