package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenPublicBBS,
		Title:      "PUBLIC BBS TERMINAL",
		Background: BgPublicBBS,
		OnSelector: true,
	})
}
