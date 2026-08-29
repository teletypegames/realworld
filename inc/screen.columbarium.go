package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenColumbarium,
		Title:      "COLUMBARIUM — DEX'S MEMORIAL",
		Background: BgColumbarium,
		OnSelector: true,
	})
}
