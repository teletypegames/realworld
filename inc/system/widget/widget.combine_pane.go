package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	combineCardWidth  = 86
	combineCardHeight = 62
	combineColumnGap  = 14
)

type TriedRow struct {
	Pair Text
	Quip Text
}

type CombinePane struct {
	Name     string
	When     inkwell.Condition
	Header   Text
	Hint     Text
	Cards    []Text
	Operator Text
	Dimmed   bool
	Caption  Text
	Rows     []Choice
	TriedTop Text
	Tried    []TriedRow
	Footer   Text

	rowCursor
}

func (pane *CombinePane) GetName() string                { return pane.Name }
func (pane *CombinePane) VisibleWhen() inkwell.Condition { return pane.When }
func (pane *CombinePane) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (pane *CombinePane) Tick(context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	pane.tick(game, pane.areas(game))
}

func (pane *CombinePane) rowsLeft() float64 {
	return PaneInset + float64(len(pane.Cards)*(combineCardWidth+combineColumnGap)) + combineColumnGap
}

func (pane *CombinePane) areas(game *inkwell.Game) []inkwell.Rectangle {
	if folded(game) {
		return nil
	}
	width := paneWidth(game)
	top := uiTop(game) + Padding + inkwell.GlyphH + 6
	if len(pane.Caption) > 0 {
		top += inkwell.GlyphH + 2
	}
	left := pane.rowsLeft()
	areas := make([]inkwell.Rectangle, len(pane.Rows))
	for index := range pane.Rows {
		areas[index] = inkwell.Rect(
			left,
			top+float64(index*(RowHeight+RowGap)),
			width-left-PaneInset,
			RowHeight,
		)
	}
	return areas
}

func (pane *CombinePane) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	theme := game.Theme()
	width := paneWidth(game)
	top := paneHeader(destination, game, pane.Header, pane.Hint)
	if folded(game) {
		return
	}

	border, ink := theme.StatusText, theme.StatusText
	if pane.Dimmed {
		border, ink = theme.HotspotOutline, theme.ChatLogSystem
	}
	for index, name := range pane.Cards {
		card := inkwell.Rect(
			PaneInset+float64(index*(combineCardWidth+combineColumnGap)),
			top,
			combineCardWidth,
			combineCardHeight,
		)
		if pane.Dimmed {
			strokeDashed(destination, card, border)
		} else {
			stroke(destination, card, border)
		}
		picture := inkwell.Rect(card.X+10, card.Y+8, combineCardWidth-20, 26)
		strokeDashed(destination, picture, theme.ChatLogSystem)
		centered(destination, game, inkwell.ClipText(name.String(), combineCardWidth-6), card,
			int(card.Y+card.H)-inkwell.GlyphH-6, ink)
		if index+1 < len(pane.Cards) && len(pane.Operator) > 0 {
			game.DrawText(destination, pane.Operator.String(),
				int(card.X+card.W)+combineColumnGap/2-inkwell.GlyphW/2,
				int(card.Y+card.H/2)-inkwell.GlyphH/2, theme.TopBarText)
		}
	}

	if len(pane.Caption) > 0 {
		game.DrawText(destination, pane.Caption.String(), int(pane.rowsLeft()), int(top), theme.TopBarText)
	}
	for index, area := range pane.areas(game) {
		state := rowNormal
		switch {
		case pane.Rows[index].Struck:
			state = rowStruck
		case index == pane.selected:
			state = rowSelected
		case pane.Rows[index].Dim:
			state = rowDim
		}
		row(destination, game, area, pane.Rows[index].Label, state)
	}

	triedY := int(top+combineCardHeight) + 8
	if len(pane.TriedTop) > 0 {
		game.DrawText(destination, pane.TriedTop.String(), PaneInset, triedY, theme.TopBarText)
		triedY += inkwell.GlyphH + 2
	}
	for _, tried := range pane.Tried {
		game.DrawText(destination, tried.Pair.String(), PaneInset, triedY, theme.TopBarText)
		quipX := PaneInset + 150
		game.DrawText(destination, inkwell.ClipText(tried.Quip.String(), int(width)-quipX-PaneInset),
			quipX, triedY, theme.ChatLogSystem)
		triedY += LineHeight
	}

	if len(pane.Footer) > 0 {
		game.DrawText(destination, pane.Footer.String(), PaneInset,
			ScreenHeight-Padding-inkwell.GlyphH, theme.ChatLogSystem)
	}
}
