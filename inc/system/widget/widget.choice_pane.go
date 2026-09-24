package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

type Choice struct {
	Label  Text
	Dim    bool
	Struck bool
}

type ChoicePane struct {
	Name    string
	When    inkwell.Condition
	Header  Text
	Hint    Text
	Rows    []Choice
	Speaker Text
	Body    Text
	Footer  Text

	rowCursor
}

func (pane *ChoicePane) GetName() string                { return pane.Name }
func (pane *ChoicePane) VisibleWhen() inkwell.Condition { return pane.When }
func (pane *ChoicePane) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (pane *ChoicePane) Tick(context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	pane.tick(game, pane.areas(game))
}

func (pane *ChoicePane) areas(game *inkwell.Game) []inkwell.Rectangle {
	if folded(game) {
		return nil
	}
	width := paneWidth(game)
	top := uiTop(game) + Padding + inkwell.GlyphH + 6
	areas := make([]inkwell.Rectangle, len(pane.Rows))
	for index := range pane.Rows {
		areas[index] = inkwell.Rect(
			PaneInset,
			top+float64(index*(RowHeight+RowGap)),
			width-PaneInset*2,
			RowHeight,
		)
	}
	return areas
}

func (pane *ChoicePane) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	theme := game.Theme()
	width := paneWidth(game)
	paneHeader(destination, game, pane.Header, pane.Hint)

	if folded(game) {
		if len(pane.Body) > 0 {
			platedCenter(destination, game, pane.Body.String(),
				int(uiTop(game))-inkwell.GlyphH-14, theme.StatusText)
		}
		return
	}

	for index, area := range pane.areas(game) {
		row(destination, game, area, pane.Rows[index].Label, pane.state(index))
	}

	if len(pane.Body) > 0 {
		x := PaneInset
		top := int(uiTop(game)) + Padding + inkwell.GlyphH + 14
		if len(pane.Speaker) > 0 {
			label(destination, game, pane.Speaker, x, top, theme.ChatLogSystem)
			x += inkwell.GlyphW * 6
		}
		for offset, line := range inkwell.WrapText(pane.Body.String(), int(width)-x-PaneInset) {
			game.DrawText(destination, line, x, top+offset*LineHeight, theme.StatusText)
		}
	}

	if len(pane.Footer) > 0 {
		game.DrawText(destination, pane.Footer.String(), PaneInset,
			ScreenHeight-Padding-inkwell.GlyphH, theme.ChatLogSystem)
	}
}

func (pane *ChoicePane) state(index int) rowState {
	switch {
	case pane.Rows[index].Struck:
		return rowStruck
	case index == pane.selected:
		return rowSelected
	case pane.Rows[index].Dim:
		return rowDim
	}
	return rowNormal
}
