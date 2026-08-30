package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	WidgetManager.Register(&windowSizer{
		Name:  "window",
		Scale: 2,
	})
}

type windowSizer struct {
	Name  string
	Scale int
	done  bool
}

func (s *windowSizer) GetName() string                            { return s.Name }
func (s *windowSizer) Layer() inkwell.Layer                       { return inkwell.LayerScene }
func (s *windowSizer) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

func (s *windowSizer) Tick(ctx *inkwell.UICtx) {
	if s.done {
		return
	}
	s.done = true
	scale := s.Scale
	if scale <= 0 {
		scale = 2
	}
	ebiten.SetWindowSize(ctx.Game.Width*scale, ctx.Game.Height*scale)
}
