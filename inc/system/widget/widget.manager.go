package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system/world"
)

type Widget = inkwell.Widget

var Manager = inkwell.NewManager[Widget]()

const (
	ScreenW = 640
	ScreenH = 380

	LineH = inkwell.GlyphH + 2
	Pad   = 6

	TopBarH  = LineH + 2
	DividerX = 394

	HUDH   = Pad + LineH + 4 + invSlot*2 + invGap + Pad
	HUDTop = ScreenH - HUDH

	statusY = HUDTop + Pad
	slotsY  = HUDTop + Pad + LineH
	slotsH  = LineH*2 + Pad*2
	invY    = slotsY + 4
	invSlot = 40
	invGap  = 4
)

var (
	whenPlaying  = inkwell.VarEq(inc.VarMode, string(world.ModePlay))
	whenCutscene = inkwell.VarEq(inc.VarMode, string(world.ModeCutscene))
	whenUnpaused = inkwell.And(whenPlaying, inkwell.Not(inkwell.VarEq(inc.VarNote, inc.NotePaused)))
)
