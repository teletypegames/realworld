// Package item registers the game's inventory items.
//
// Personal Tapes are inventory items and quest triggers at the same time: they
// go into slot 2 of Paul's Armilla, and loading one activates its effect.
package item

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"git.teletypegames.org/games/realworld/internal/world"
)

// Register adds every item to the game.
func Register(w *world.World) {
	for _, i := range []inkwell.Item{
		noodleLetter(),
		blackMarketArmilla(),
		mysteryTape(),
	} {
		w.G.ItemManager.Register(i)
	}
}
