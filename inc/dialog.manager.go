package inc

// Dialogues — the game's dialogue trees, one file each.
//
// Talking to a tape carries the same weight as talking to a human NPC — the
// wiki states this explicitly. Tapes are dialogue partners, not menus.
//
// A note on DialogueChoice.Once: inkwell HIDES a spent choice. The wireframe
// deck wanted it struck through but still visible, so the list becomes a
// memory of what you already tried. That is still open — see README, "Known
// gaps".

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// registerDialog adds every dialogue to the game.
func registerDialog(w *World) {
	for _, d := range []inkwell.Dialogue{
		dialogDexTalk(),
		dialogMysteryTapeSilent(),
	} {
		w.G.DialogueManager.Register(d)
	}
}
