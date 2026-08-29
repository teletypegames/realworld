package inc

// Items — the game's inventory, one file each.
//
// Personal Tapes are inventory items and quest triggers at the same time: they
// go into slot 2 of Paul's Armilla, and loading one activates its effect.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// registerItem adds every item to the game.
func registerItem(w *World) {
	for _, i := range []inkwell.Item{
		itemNoodleLetter(),
		itemBlackMarketArmilla(),
		itemMysteryTape(),
	} {
		w.G.ItemManager.Register(i)
	}
}
