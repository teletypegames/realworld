package inc

func screenPoliceStation() screen {
	return screen{
		name:       ScreenPoliceStation,
		title:      "SFPD — STATION AND HOLDING",
		background: BgPoliceStation,
		onSelector: true,
	}
}
