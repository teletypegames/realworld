package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

type Widget = inkwell.Widget

const (
	ScreenW = 640
	ScreenH = 380

	LineH = glyphH + 2
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

type gate struct {
	Name  string
	Inner Widget
}

func gated(name string, inner Widget) Widget {
	return &gate{
		Name:  name,
		Inner: inner,
	}
}

func (n *gate) GetName() string { return n.Name }

func (n *gate) Layer() inkwell.Layer { return inkwell.LayerOf(n.Inner) }

func (n *gate) Tick(ctx *inkwell.UICtx) {
	if World.HUDVisible() {
		n.Inner.Tick(ctx)
	}
}

func (n *gate) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if World.HUDVisible() {
		n.Inner.Draw(dst, ctx)
	}
}

func (n *gate) BlocksClickAt(p inkwell.Point) bool {
	if !World.HUDVisible() {
		return false
	}
	b, ok := n.Inner.(interface{ BlocksClickAt(inkwell.Point) bool })
	return ok && b.BlocksClickAt(p)
}
