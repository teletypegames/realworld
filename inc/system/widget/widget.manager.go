package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
)

type Widget = inkwell.Widget

var WidgetManager = inkwell.NewManager[Widget]()

const (
	ScreenWidth  = 640
	ScreenHeight = 380

	LineHeight = inkwell.GlyphH + 2
	Padding    = 6

	TopBarHeight = LineHeight + 2
	DividerX     = 394

	HUDHeight = Padding + LineHeight + 4 + inventorySlotSize*2 + inventoryGap + Padding
	HUDTop    = ScreenHeight - HUDHeight

	SlotsY      = HUDTop + Padding + LineHeight
	SlotsHeight = LineHeight*2 + Padding*2

	statusY           = HUDTop + Padding
	inventoryY        = SlotsY + 4
	inventorySlotSize = 40
	inventoryGap      = 4
)

var (
	onStage = inkwell.OnStage(inc.StageGame)

	WhenPlaying = inkwell.And(onStage, inkwell.Not(inCutscene))

	whenCutscene = inkwell.And(onStage, inCutscene)
	whenDebug    = inkwell.And(onStage, inkwell.Flag(inc.FlagDebug))

	inCutscene = inkwell.VarEq(inc.VarMode, string(system.ModeCutscene))
)
