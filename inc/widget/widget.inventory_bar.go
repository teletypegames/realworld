package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	Manager.Register(&inkwell.InventoryBar{
		Name: "inventory",
		When: whenPlaying,
		Origin: inkwell.Point{
			X: Pad + 2,
			Y: invY,
		},
		Slots:    8,
		Cols:     4,
		SlotSize: invSlot,
		Gap:      invGap,
	})
}
