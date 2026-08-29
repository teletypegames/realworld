package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenRooftopHideout,
		Title:      "NORMAN'S ROOFTOP HIDEOUT",
		Background: BgRooftopHideout,
		Exits: []exit{
			{
				to:    ScreenNormanApartment,
				label: "back down into the flat",
				side:  sideNear,
			},
			{
				to:    ScreenBBSTerminal,
				label: "the old terminal",
				side:  sideRight,
			},
		},
	})
}
