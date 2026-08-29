package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenHackerspace,
		Title:      "HACKERSPACE",
		Background: BgHackerspace,
		OnSelector: true,
	})
}
