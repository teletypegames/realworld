package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	WidgetManager.Register(gated("topbar", &titleBar{
		TopBar: &inkwell.TopBar{
			Name:    "topbar",
			Height:  TopBarH,
			TimeVar: VarArmillaStrip,
		},
	}))
}

type titleBar struct {
	*inkwell.TopBar
}

func (t *titleBar) Tick(ctx *inkwell.UICtx) {
	t.LeftText = World.SceneTitle()
	if World.Paused() {
		t.LeftText += " — paused"
	}
	t.TopBar.Tick(ctx)
}
