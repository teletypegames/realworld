package inc

func screenBbsTerminal() screen {
	return screen{
		name:       ScreenBBSTerminal,
		title:      "TERMINAL — BBS ACCESS POINT",
		background: BgBBSTerminal,
		exits: []exit{
			{label: "step back from the terminal", side: sideNear},
		},
	}
}
