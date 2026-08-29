package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenSamizdatPress,
		Title:      "SAMIZDAT PRINTING HOUSE",
		Background: BgSamizdatPress,
		OnSelector: true,
	})
}
