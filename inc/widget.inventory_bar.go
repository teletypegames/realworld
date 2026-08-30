package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(gated("inventory", &inkwell.InventoryBar{
		Name: "inventory",
		Origin: inkwell.Point{
			X: Pad + 2,
			Y: invY,
		},
		Slots:    8,
		Cols:     4,
		SlotSize: invSlot,
		Gap:      invGap,
	}))
}
