package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	inventoryColumnWidth = 110
	inventoryRowHeight   = 18
	inventoryArtWidth    = 78
	inventoryArtHeight   = 46
)

type InventoryPane struct {
	Name        string
	When        inkwell.Condition
	Header      Text
	Hint        Text
	Items       []Choice
	ArtLabel    Text
	Description Text
	Verbs       []Choice

	rowCursor
}

func (pane *InventoryPane) GetName() string                { return pane.Name }
func (pane *InventoryPane) VisibleWhen() inkwell.Condition { return pane.When }
func (pane *InventoryPane) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (pane *InventoryPane) Tick(context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	pane.tick(game, pane.areas(game))
}

func (pane *InventoryPane) areas(game *inkwell.Game) []inkwell.Rectangle {
	if folded(game) {
		return nil
	}
	top := uiTop(game) + Padding + inkwell.GlyphH + 6
	areas := make([]inkwell.Rectangle, len(pane.Items))
	for index := range pane.Items {
		areas[index] = inkwell.Rect(
			PaneInset,
			top+float64(index*(inventoryRowHeight+2)),
			inventoryColumnWidth,
			inventoryRowHeight,
		)
	}
	return areas
}

func (pane *InventoryPane) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	theme := game.Theme()
	width := paneWidth(game)
	top := paneHeader(destination, game, pane.Header, pane.Hint)
	if folded(game) {
		return
	}

	for index, area := range pane.areas(game) {
		state := rowNormal
		switch {
		case index == pane.selected:
			state = rowSelected
		case pane.Items[index].Dim:
			state = rowDim
		}
		row(destination, game, area, pane.Items[index].Label, state)
	}

	detailX := PaneInset + inventoryColumnWidth + 12
	picture := inkwell.Rect(float64(detailX), top, inventoryArtWidth, inventoryArtHeight)
	strokeDashed(destination, picture, theme.TopBarText)
	if len(pane.ArtLabel) > 0 {
		centered(destination, game, pane.ArtLabel.String(), picture,
			int(picture.Y+picture.H/2)-inkwell.GlyphH/2, theme.ChatLogSystem)
	}

	textX := detailX + inventoryArtWidth + 10
	for index, line := range inkwell.WrapText(pane.Description.String(), int(width)-textX-PaneInset) {
		game.DrawText(destination, line, textX, int(top)+index*LineHeight, theme.TopBarText)
	}

	chipX := float64(detailX)
	chipY := top + inventoryArtHeight + 10
	for _, verb := range pane.Verbs {
		chip := inkwell.Rect(chipX, chipY,
			float64(inkwell.TextWidth(verb.Label.String())+Padding*2), RowHeight)
		if chip.X+chip.W > width-PaneInset {
			chipX, chipY = float64(detailX), chipY+RowHeight+RowGap
			chip.X, chip.Y = chipX, chipY
		}
		row(destination, game, chip, verb.Label, rowNormal)
		chipX += chip.W + RowGap
	}
}
