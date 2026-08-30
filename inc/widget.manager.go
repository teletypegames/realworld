package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type Widget = inkwell.Widget

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
	whenPlaying  = inkwell.VarEq(VarMode, string(ModePlay))
	whenCutscene = inkwell.VarEq(VarMode, string(ModeCutscene))
	whenUnpaused = inkwell.And(whenPlaying, inkwell.Not(inkwell.VarEq(VarNote, NotePaused)))
)
