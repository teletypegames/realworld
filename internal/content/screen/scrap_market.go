package screen

import "git.teletypegames.org/games/realworld/internal/names"

// scrapMarket is a 💡 suggestion: the black market in condemned ACPs and bare,
// unstamped tape. Canon says the leak of scrap is not an accident — "something
// wants to live on" (wiki: entities#potenciális-további-helyszínek).
func scrapMarket() screen {
	return screen{
		name:       names.ScreenScrapMarket,
		title:      "SCRAP MARKET",
		background: names.BgScrapMarket,
		onSelector: true,
	}
}
