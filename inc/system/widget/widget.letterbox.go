package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/util"
)

var _ = WidgetManager.Register(&inkwell.Letterbox{
	Name:  "letterbox",
	When:  whenCutscene,
	Bar:   36,
	Color: util.RGB(0x000000),
})
