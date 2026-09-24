package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"

	"realworld/inc/util"
)

var scanlineShade = util.RGBA(0x000000, 0x33)

type Scanlines struct {
	Name string
	When inkwell.Condition
}

func (lines *Scanlines) GetName() string                { return lines.Name }
func (lines *Scanlines) VisibleWhen() inkwell.Condition { return lines.When }
func (lines *Scanlines) Layer() inkwell.Layer           { return inkwell.LayerCurtain }
func (lines *Scanlines) Tick(context *inkwell.UICtx)    {}

func (lines *Scanlines) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	for y := 0.0; y < ScreenHeight; y += 4 {
		fill(destination, inkwell.Rect(0, y, ScreenWidth, 2), scanlineShade)
	}
}
