package screen

import "git.teletypegames.org/games/realworld/internal/names"

// serverFarm is the cooled, humming corridors where The AI's physical
// infrastructure can be seen for what it is. Getting in needs the access tape
// from the hackerspace or from Noodle (wiki: entities, timeline 24).
func serverFarm() screen {
	return screen{
		name:       names.ScreenServerFarm,
		title:      "SERVER FARM",
		background: names.BgServerFarm,
		onSelector: true,
	}
}
