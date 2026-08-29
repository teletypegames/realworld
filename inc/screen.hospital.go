package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenHospital,
		Title:      "HOSPITAL — PSYCHIATRIC WING",
		Background: BgHospital,
		OnSelector: true,
	})
}
