package screen

import "git.teletypegames.org/games/realworld/internal/names"

// secretLab is the lab the BVK's secret section runs, where Norman lies wired
// into the simulator: the two-thirds point where Paul finds him, and the
// screen the finale plays on (wiki: entities#helyszinek, timeline 22 and 25).
//
// The wiki's graph runs Híd → security szoba → Norman szimulációs szobája, and
// the story puts the whole thing inside the Golden Gate Bridge. The deck paints
// it as one screen, so the door at the foot of the bridge, the observation room
// and the glass wall are hotspots in here.
func secretLab() screen {
	return screen{
		name:       names.ScreenSecretLab,
		title:      "SECRET RESEARCH LABORATORY",
		background: names.BgSecretLab,
		onSelector: true,
	}
}
