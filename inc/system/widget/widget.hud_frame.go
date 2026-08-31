package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func init() {
	Manager.Register(&hudFrame{
		Name: "hudframe",
	})
}

type hudFrame struct {
	Name string
}

func (h *hudFrame) GetName() string                { return h.Name }
func (h *hudFrame) VisibleWhen() inkwell.Condition { return whenPlaying }
func (h *hudFrame) Layer() inkwell.Layer           { return inkwell.LayerPanel }
func (h *hudFrame) Tick(ctx *inkwell.UICtx)        {}

func (h *hudFrame) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	th := ctx.Game.Theme()
	vector.DrawFilledRect(dst, 0, HUDTop+1, ScreenW, ScreenH-HUDTop-1, th.PanelBG, false)

	vector.StrokeLine(dst, 0, HUDTop, ScreenW, HUDTop, 1, th.CharacterPanelBorder, false)

	vector.StrokeLine(dst, DividerX, HUDTop+1, DividerX, ScreenH, 1, th.CharacterPanelBorder, false)
}

func (h *hudFrame) BlocksClickAt(p inkwell.Point) bool { return p.Y >= HUDTop }
