package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

var Save = StageManager.Register(&Stage{
	Name: "save",
	Title: Text{
		{Lang: "en", Text: "SAVE"},
		{Lang: "hu", Text: "MENTÉS"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "save_menu",
	When:   inkwell.OnStage(Save),
	Sticky: true,
	Bounds: window(360, 180),
	Root: MenuItem{
		Label: Text{
			{Lang: "en", Text: "SAVE"},
			{Lang: "hu", Text: "MENTÉS"},
		},
		Expand: saveExpand,
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("save_escape", Save))

func saveExpand(context *inkwell.Ctx) []MenuItem {
	var rows []MenuItem
	for _, info := range context.Game.Slots(inc.SaveSlots) {
		rows = append(rows, MenuItem{
			Label: inkwell.Format(slotName, info.Slot),
			Value: inkwell.FixedValue(slotValue(info)),
			Do: inkwell.Seq(
				inkwell.SaveSlot(info.Slot),
				inkwell.GoToStage(MainMenu),
			),
		})
	}
	return append(rows, backRow())
}
