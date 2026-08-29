package inc

// scrapMarket is a 💡 suggestion: the black market in condemned ACPs and bare,
// unstamped tape. Canon says the leak of scrap is not an accident — "something
// wants to live on" (wiki: entities#potenciális-további-helyszínek).
func screenScrapMarket() screen {
	return screen{
		name:       ScreenScrapMarket,
		title:      "SCRAP MARKET",
		background: BgScrapMarket,
		onSelector: true,
	}
}
