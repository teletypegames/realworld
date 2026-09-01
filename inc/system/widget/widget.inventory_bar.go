package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(&inkwell.InventoryBar{
		Name: "inventory",
		When: WhenPlaying,
		Origin: inkwell.Point{
			X: Padding + 2,
			Y: inventoryY,
		},
		Slots:    8,
		Cols:     4,
		SlotSize: inventorySlotSize,
		Gap:      inventoryGap,
	})
}
