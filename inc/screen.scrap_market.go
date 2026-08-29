package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenScrapMarket,
		Title:      "SCRAP MARKET",
		Background: BgScrapMarket,
		OnSelector: true,
	})
}
