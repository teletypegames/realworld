// Package content is the composition root for everything the game declares:
// assets, characters, items, dialogues, scripts and scenes.
//
// Each registered entity lives in its own file, under a subpackage named after
// its kind. Adding a screen means adding screen/<name>.go, background/<name>.go
// and one line in each package's list — nothing else moves.
package content

import (
	"git.teletypegames.org/games/realworld/internal/content/background"
	"git.teletypegames.org/games/realworld/internal/content/character"
	"git.teletypegames.org/games/realworld/internal/content/dialog"
	"git.teletypegames.org/games/realworld/internal/content/item"
	"git.teletypegames.org/games/realworld/internal/content/screen"
	"git.teletypegames.org/games/realworld/internal/content/script"
	"git.teletypegames.org/games/realworld/internal/world"
)

// Register adds every entity to the game. Order matters only in that screens
// reference backgrounds and characters, which Game.Validate cross-checks.
func Register(w *world.World) {
	background.Register(w)
	character.Register(w)
	item.Register(w)
	dialog.Register(w)
	script.Register(w)
	screen.Register(w)
}
