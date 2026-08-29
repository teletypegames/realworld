package screen

import "git.teletypegames.org/games/realworld/internal/names"

// showroom is a 💡 suggestion: the legitimate, glittering side — Armilla
// cabinets, licensed operators, matte-black Nokia-punk under showroom lights.
// The contrast with Paul's edge of the world
// (wiki: entities#potenciális-további-helyszínek).
func showroom() screen {
	return screen{
		name:       names.ScreenShowroom,
		title:      "NEUMATRONIC SHOWROOM",
		background: names.BgShowroom,
		onSelector: true,
	}
}
