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

	statusY           = HUDTop + Padding
	slotsY            = HUDTop + Padding + LineHeight
	slotsHeight       = LineHeight*2 + Padding*2
	inventoryY        = slotsY + 4
	inventorySlotSize = 40
	inventoryGap      = 4
)

var (
	whenPlaying  = inkwell.VarEq(inc.VarMode, string(system.ModePlay))
	whenCutscene = inkwell.VarEq(inc.VarMode, string(system.ModeCutscene))
	whenDebug    = inkwell.Flag(inc.FlagDebug)
)
