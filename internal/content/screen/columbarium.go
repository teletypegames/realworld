package screen

import "git.teletypegames.org/games/realworld/internal/names"

// columbarium is a 💡 suggestion: where Dex is buried. A quiet place, away
// from the noise of the terminals — the loss-and-letting-go thread, which is
// the same thread as Norman and The AI
// (wiki: entities#potenciális-további-helyszínek).
func columbarium() screen {
	return screen{
		name:       names.ScreenColumbarium,
		title:      "COLUMBARIUM — DEX'S MEMORIAL",
		background: names.BgColumbarium,
		onSelector: true,
	}
}
