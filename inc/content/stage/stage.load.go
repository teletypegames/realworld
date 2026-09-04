package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

var Load = StageManager.Register(&Stage{
	Name: "load",
	Title: Text{
		{Lang: "en", Text: "LOAD"},
		{Lang: "hu", Text: "BETÖLTÉS"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "load_menu",
	When:   inkwell.OnStage(Load),
	Sticky: true,
	Bounds: window(360, 180),
	Root: MenuItem{
		Label: Text{
			{Lang: "en", Text: "LOAD"},
			{Lang: "hu", Text: "BETÖLTÉS"},
		},
		Expand: loadExpand,
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("load_escape", Load))

func loadExpand(context *inkwell.Ctx) []MenuItem {
	var rows []MenuItem
	for _, info := range context.Game.Slots(inc.SaveSlots) {
		rows = append(rows, MenuItem{
			Label:   inkwell.Format(slotName, info.Slot),
			Value:   inkwell.FixedValue(slotValue(info)),
			Enabled: fixed(info.Exists),
			Do: inkwell.Seq(
				inkwell.LoadSlot(info.Slot),
				inkwell.GoToStage(Game),
			),
		})
	}
	return append(rows, backRow())
}
