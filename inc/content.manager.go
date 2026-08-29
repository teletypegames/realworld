package inc

// Content — the composition root for everything the game declares:
// assets, characters, items, dialogues, scripts and scenes.
//
// Each registered entity lives in its own file, named after its category:
// adding a screen means adding screen.<name>.go and background.<name>.go, plus
// one line in each category's manager list — nothing else moves.

// registerContent adds every entity to the game. Order matters only in that
// screens reference backgrounds and characters, which Game.Validate cross-checks.
func registerContent(w *World) {
	registerBackground(w)
	registerCharacter(w)
	registerItem(w)
	registerDialog(w)
	registerScript(w)
	registerScreen(w)
}
