package screen

import "git.teletypegames.org/games/realworld/internal/names"

// policeStation is where Noodle is held for dealing in cracked Armillas, and
// where a clerk quietly tips Paul off about the hiding place
// (wiki: entities, Paul's timeline 4 and 15).
//
// The wiki's graph splits it three ways — utcafront → recepció → kihallgató —
// but the deck paints it as one screen, so the reception desk, the waiting
// room and the interview window are hotspots inside this one, not screens of
// their own.
func policeStation() screen {
	return screen{
		name:       names.ScreenPoliceStation,
		title:      "SFPD — STATION AND HOLDING",
		background: names.BgPoliceStation,
		onSelector: true,
	}
}
