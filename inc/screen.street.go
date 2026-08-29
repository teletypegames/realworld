package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenStreet,
		Title:      "STREET",
		Background: BgStreet,
		OnSelector: true,
	})
}
