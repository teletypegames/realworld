package screen

import "git.teletypegames.org/games/realworld/internal/names"

// bbsTerminal is the screen you get when Paul sits down at a terminal: the
// catalogue calls it a "Tárgy/portál", a physical object in the world that
// opens the SPOKE — BBS surface (wiki: entities#helyszinek).
//
// It is the same screen wherever the terminal stands — the club's ECHO box,
// Norman's machine at the flat, the old one on the roof — which is why its way
// out is "back", and not a place.
func bbsTerminal() screen {
	return screen{
		name:       names.ScreenBBSTerminal,
		title:      "TERMINAL — BBS ACCESS POINT",
		background: names.BgBBSTerminal,
		exits: []exit{
			{label: "step back from the terminal", side: sideNear},
		},
	}
}
