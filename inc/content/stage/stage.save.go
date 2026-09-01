package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

type saveScreen struct{}

var _ = StageManager.Register(Stage{
	Name: inc.StageSave,
	Title: inkwell.Text{
		{Lang: "en", Text: "SAVE"},
		{Lang: "hu", Text: "MENTÉS"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "save_menu",
	When:   inkwell.OnStage(inc.StageSave),
	Sticky: true,
	Bounds: window(360, 180),
	Root: inkwell.MenuItem{
		Label: inkwell.Text{
			{Lang: "en", Text: "SAVE"},
			{Lang: "hu", Text: "MENTÉS"},
		},
		Expand: saveScreen{}.rows,
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("save_escape", inc.StageSave))

func (saveScreen) rows(context *inkwell.Ctx) []inkwell.MenuItem {
	var rows []inkwell.MenuItem
	for _, info := range context.Game.Slots(inc.SaveSlots) {
		rows = append(rows, inkwell.MenuItem{
			Label: inkwell.Format(slotName, info.Slot),
			Value: inkwell.FixedValue(slotValue(info)),
			Do: inkwell.Seq(
				inkwell.SaveSlot(info.Slot),
				inkwell.GoToStage(inc.StageMainMenu),
			),
		})
	}
	return append(rows, backRow())
}
