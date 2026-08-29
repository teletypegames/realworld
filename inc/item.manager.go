package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func registerItem(w *World) {
	for _, i := range []inkwell.Item{
		itemNoodleLetter(),
		itemBlackMarketArmilla(),
		itemMysteryTape(),
	} {
		w.G.ItemManager.Register(i)
	}
}
