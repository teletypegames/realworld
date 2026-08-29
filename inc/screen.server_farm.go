package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenServerFarm,
		Title:      "SERVER FARM",
		Background: BgServerFarm,
		OnSelector: true,
	})
}
