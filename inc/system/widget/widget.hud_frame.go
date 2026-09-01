package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var _ = WidgetManager.Register(&hudFrame{
	Name: "hudframe",
})

type hudFrame struct {
	Name string
}

func (frame *hudFrame) GetName() string                { return frame.Name }
func (frame *hudFrame) VisibleWhen() inkwell.Condition { return WhenPlaying }
func (frame *hudFrame) Layer() inkwell.Layer           { return inkwell.LayerPanel }
func (frame *hudFrame) Tick(context *inkwell.UICtx)    {}

func (frame *hudFrame) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	theme := context.Game.Theme()
	vector.DrawFilledRect(destination, 0, HUDTop+1, ScreenWidth, ScreenHeight-HUDTop-1, theme.PanelBG, false)

	vector.StrokeLine(destination, 0, HUDTop, ScreenWidth, HUDTop, 1, theme.CharacterPanelBorder, false)

	vector.StrokeLine(destination, DividerX, HUDTop+1, DividerX, ScreenHeight, 1, theme.CharacterPanelBorder, false)
}

func (frame *hudFrame) BlocksClickAt(point inkwell.Point) bool { return point.Y >= HUDTop }
