package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenSecretLab,
		Title:      "SECRET RESEARCH LABORATORY",
		Background: BgSecretLab,
		OnSelector: true,
	})
}
