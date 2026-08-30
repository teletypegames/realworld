package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	WidgetManager.Register(&actionPump{
		Name: "pump",
	})
}

type actionPump struct {
	Name string
}

func (p *actionPump) GetName() string                            { return p.Name }
func (p *actionPump) Layer() inkwell.Layer                       { return inkwell.LayerScene }
func (p *actionPump) Tick(ctx *inkwell.UICtx)                    { World.PumpTick(ctx.DT) }
func (p *actionPump) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}
