package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

type FrameSplit struct {
	Name string
	When inkwell.Condition

	top float64
}

func (frame *FrameSplit) GetName() string                { return frame.Name }
func (frame *FrameSplit) VisibleWhen() inkwell.Condition { return frame.When }
func (frame *FrameSplit) Layer() inkwell.Layer           { return inkwell.LayerPanel }

func (frame *FrameSplit) Tick(context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	frame.top = uiTop(game)
	if game.Input.LeftClicked() && foldButton(game).Contains(game.Input.Point()) {
		game.Input.ConsumeLeft()
		game.Do(FoldToggle())
	}
}

func (frame *FrameSplit) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	theme := game.Theme()
	top := uiTop(game)
	frame.top = top
	divider := paneWidth(game)

	fill(destination, inkwell.Rect(0, 0, ScreenWidth, top), theme.SceneBackdrop)
	fill(destination, inkwell.Rect(0, top, ScreenWidth, ScreenHeight-top), theme.PanelBG)
	fill(destination, inkwell.Rect(divider, top, ScreenWidth-divider, ScreenHeight-top), theme.ChatLogBG)

	fill(destination, inkwell.Rect(0, top, ScreenWidth, 1), theme.CharacterPanelBorder)
	fill(destination, inkwell.Rect(divider, top, 1, ScreenHeight-top), theme.TopBarAccent)

	button := foldButton(game)
	border, ink := theme.CharacterPanelBorder, theme.TopBarText
	if button.Contains(game.Input.Point()) {
		border, ink = theme.StatusText, theme.StatusText
	}
	stroke(destination, button, border)
	arrow := "↓"
	if folded(game) {
		arrow = "↑"
	}
	centered(destination, game, arrow, button, int(button.Y)-1, ink)
}

func (frame *FrameSplit) BlocksClickAt(point inkwell.Point) bool {
	if frame.top == 0 {
		return point.Y >= UITop
	}
	return point.Y >= frame.top
}
