package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenBBSTerminal,
		Title:      "TERMINAL — BBS ACCESS POINT",
		Background: BgBBSTerminal,
		Exits: []exit{
			{
				label: "step back from the terminal",
				side:  sideNear,
			},
		},
	})
}
