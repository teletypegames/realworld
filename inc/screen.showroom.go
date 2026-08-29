package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenShowroom,
		Title:      "NEUMATRONIC SHOWROOM",
		Background: BgShowroom,
		OnSelector: true,
	})
}
