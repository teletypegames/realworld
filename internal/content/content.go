// Package content is the composition root for everything the game declares:
// assets, characters, items, dialogues, scripts and scenes.
//
// Each registered entity lives in its own file, under a subpackage named after
// its kind. Adding a scene means adding scene/<name>.go and one line in
// scene/scene.go — nothing else moves.
package content

import (
	"git.teletypegames.org/games/realworld/internal/content/asset"
	"git.teletypegames.org/games/realworld/internal/content/character"
	"git.teletypegames.org/games/realworld/internal/content/dialog"
	"git.teletypegames.org/games/realworld/internal/content/item"
	"git.teletypegames.org/games/realworld/internal/content/scene"
	"git.teletypegames.org/games/realworld/internal/content/script"
	"git.teletypegames.org/games/realworld/internal/world"
)

// Register adds every entity to the game. Order matters only in that scenes
// reference assets and characters, which Game.Validate cross-checks.
func Register(w *world.World) {
	asset.Register(w)
	character.Register(w)
	item.Register(w)
	dialog.Register(w)
	script.Register(w)
	scene.Register(w)
}
