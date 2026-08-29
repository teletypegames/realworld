package screen

import "git.teletypegames.org/games/realworld/internal/names"

// hackerspace: the community workshop at the edge of the Mission, printers and
// stripped servers and paranoid amateur cryptographers. Where the server-farm
// access tape can be had (wiki: entities#helyszinek, timeline 24).
//
// The wiki's graph gives it no physical neighbour, so it hangs off the
// selector like every other main location.
func hackerspace() screen {
	return screen{
		name:       names.ScreenHackerspace,
		title:      "HACKERSPACE",
		background: names.BgHackerspace,
		onSelector: true,
	}
}
