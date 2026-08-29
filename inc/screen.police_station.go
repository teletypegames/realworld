package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenPoliceStation,
		Title:      "SFPD — STATION AND HOLDING",
		Background: BgPoliceStation,
		OnSelector: true,
	})
}
