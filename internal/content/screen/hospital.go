package screen

import "git.teletypegames.org/games/realworld/internal/names"

// hospital is the clinical dead end: Dr. Maren looked for a medical reason for
// Norman's state and did not find one (wiki: entities, timeline 21).
func hospital() screen {
	return screen{
		name:       names.ScreenHospital,
		title:      "HOSPITAL — PSYCHIATRIC WING",
		background: names.BgHospital,
		onSelector: true,
	}
}
