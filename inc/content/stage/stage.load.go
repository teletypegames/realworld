package stage

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/widget"
)

type loadScreen struct{}

var Load = StageManager.Register(&Stage{
	Name: "load",
	Title: inkwell.Text{
		{Lang: "en", Text: "LOAD"},
		{Lang: "hu", Text: "BETÖLTÉS"},
	},
})

var _ = widget.WidgetManager.Register(&inkwell.Menu{
	Name:   "load_menu",
	When:   inkwell.OnStage(Load),
	Sticky: true,
	Bounds: window(360, 180),
	Root: inkwell.MenuItem{
		Label: inkwell.Text{
			{Lang: "en", Text: "LOAD"},
			{Lang: "hu", Text: "BETÖLTÉS"},
		},
		Expand: loadScreen{}.expand,
	},
})

var _ = widget.WidgetManager.Register(escapeToMenu("load_escape", Load))

func (loadScreen) expand(context *inkwell.Ctx) []inkwell.MenuItem {
	var rows []inkwell.MenuItem
	for _, info := range context.Game.Slots(inc.SaveSlots) {
		rows = append(rows, inkwell.MenuItem{
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
