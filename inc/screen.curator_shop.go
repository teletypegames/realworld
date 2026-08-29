package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenCuratorShop,
		Title:      "CURATOR SERVICE — THE MASTER'S WORKSHOP",
		Background: BgCuratorShop,
		OnSelector: true,
	})
}
