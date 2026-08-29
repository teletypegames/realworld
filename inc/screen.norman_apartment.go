package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenNormanApartment,
		Title:      "NORMAN'S APARTMENT",
		Background: BgNormanApartment,
		OnSelector: true,
		Exits: []exit{
			{
				to:    ScreenRooftopHideout,
				label: "the stairs up to the roof",
				side:  sideBack,
			},
			{
				to:    ScreenBBSTerminal,
				label: "Norman's terminal",
				side:  sideRight,
			},
		},
	})
}
