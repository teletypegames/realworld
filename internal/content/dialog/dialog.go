// Package dialog registers the game's dialogue trees.
//
// Talking to a tape carries the same weight as talking to a human NPC — the
// wiki states this explicitly. Tapes are dialogue partners, not menus.
//
// A note on DialogueChoice.Once: inkwell HIDES a spent choice. The wireframe
// deck wanted it struck through but still visible, so the list becomes a
// memory of what you already tried. That is still open — see README, "Known
// gaps".
package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/world"
)

// Register adds every dialogue to the game.
func Register(w *world.World) {
	for _, d := range []inkwell.Dialogue{
		dexTalk(),
		mysteryTapeSilent(),
	} {
		w.G.DialogueManager.Register(d)
	}
}
