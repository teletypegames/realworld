package widget

import (
	"image/color"
	"math"

	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/inc"
)

type (
	Widget = inkwell.Widget
	Text   = inkwell.Text
)

var WidgetManager = inkwell.NewManager[Widget]()

const (
	ScreenWidth  = 640
	ScreenHeight = 380

	LineHeight = inkwell.GlyphH + 2
	Padding    = 6

	TopBarHeight = LineHeight + 2

	ArtHeight = ScreenHeight * 3 / 5
	UITop     = ArtHeight
	UIHeight  = ScreenHeight - UITop

	ConsoleWidth = 224
	RailWidth    = 32
	DividerX     = ScreenWidth - ConsoleWidth
	RailX        = ScreenWidth - RailWidth

	RowHeight = inkwell.GlyphH + 6
	RowGap    = 3
	PaneInset = 10

	UICollapsedHeight = ScreenHeight / 10
	UIFoldedTop       = ScreenHeight - UICollapsedHeight

	foldButtonSize = 14

	HUDTop    = UITop
	HUDHeight = UIHeight

	SlotsY      = HUDTop + Padding + LineHeight
	SlotsHeight = LineHeight*2 + Padding*2

	statusY           = HUDTop + Padding
	inventoryY        = SlotsY + 4
	inventorySlotSize = 40
	inventoryGap      = 4
)

var (
	onWorld = inkwell.OnWorld()

	WhenPlaying = onWorld
	whenDebug   = inkwell.And(onWorld, inkwell.Flag(inc.FlagDebug))
)

func FoldToggle() inkwell.Action {
	return inkwell.If(
		inkwell.Flag(inc.FlagUICollapsed),
		inkwell.ClearFlag(inc.FlagUICollapsed),
		inkwell.SetFlag(inc.FlagUICollapsed),
	)
}

func folded(game *inkwell.Game) bool { return game.State.Flag(inc.FlagUICollapsed) }

func uiTop(game *inkwell.Game) float64 {
	if folded(game) {
		return UIFoldedTop
	}
	return UITop
}

func foldButton(game *inkwell.Game) inkwell.Rectangle {
	return inkwell.Rect(
		paneWidth(game)-PaneInset-foldButtonSize,
		uiTop(game)+5,
		foldButtonSize,
		foldButtonSize,
	)
}

func paneWidth(game *inkwell.Game) float64 {
	if game.State.Flag(inc.FlagConsoleCollapsed) {
		return RailX
	}
	return DividerX
}

func fill(destination *ebiten.Image, area inkwell.Rectangle, shade color.Color) {
	vector.DrawFilledRect(destination, float32(area.X), float32(area.Y), float32(area.W), float32(area.H), shade, false)
}

func stroke(destination *ebiten.Image, area inkwell.Rectangle, shade color.Color) {
	vector.StrokeRect(destination, float32(area.X), float32(area.Y), float32(area.W), float32(area.H), 1, shade, false)
}

func strokeDashed(destination *ebiten.Image, area inkwell.Rectangle, shade color.Color) {
	for offset := 0.0; offset < area.W; offset += 6 {
		length := math.Min(3, area.W-offset)
		fill(destination, inkwell.Rect(area.X+offset, area.Y, length, 1), shade)
		fill(destination, inkwell.Rect(area.X+offset, area.Y+area.H-1, length, 1), shade)
	}
	for offset := 0.0; offset < area.H; offset += 6 {
		length := math.Min(3, area.H-offset)
		fill(destination, inkwell.Rect(area.X, area.Y+offset, 1, length), shade)
		fill(destination, inkwell.Rect(area.X+area.W-1, area.Y+offset, 1, length), shade)
	}
}

func label(destination *ebiten.Image, game *inkwell.Game, text Text, x, y int, shade color.Color) int {
	body := text.String()
	game.DrawText(destination, body, x, y, shade)
	return x + inkwell.TextWidth(body)
}

func centered(destination *ebiten.Image, game *inkwell.Game, body string, area inkwell.Rectangle, y int, shade color.Color) {
	x := int(area.X) + (int(area.W)-inkwell.TextWidth(body))/2
	game.DrawText(destination, body, x, y, shade)
}

type rowState int

const (
	rowNormal rowState = iota
	rowSelected
	rowDim
	rowStruck
)

func row(destination *ebiten.Image, game *inkwell.Game, area inkwell.Rectangle, text Text, state rowState) {
	theme := game.Theme()
	border, ink := theme.ChatLogSystem, theme.TopBarText
	switch state {
	case rowSelected:
		fill(destination, area, theme.VerbButtonBG)
		border, ink = theme.StatusText, theme.StatusText
	case rowDim, rowStruck:
		border, ink = theme.HotspotOutline, theme.ChatLogSystem
	}
	if state == rowStruck {
		strokeDashed(destination, area, border)
	} else {
		stroke(destination, area, border)
	}
	body := inkwell.ClipText(text.String(), int(area.W)-Padding*2)
	baseline := int(area.Y) + (int(area.H)-inkwell.GlyphH)/2
	game.DrawText(destination, body, int(area.X)+Padding, baseline, ink)
	if state == rowStruck {
		fill(destination, inkwell.Rect(
			area.X+Padding,
			float64(baseline+inkwell.GlyphH/2),
			float64(inkwell.TextWidth(body)),
			1,
		), ink)
	}
}

type rowCursor struct {
	selected int
}

func (cursor *rowCursor) tick(game *inkwell.Game, areas []inkwell.Rectangle) {
	if len(areas) == 0 {
		return
	}
	if cursor.selected >= len(areas) {
		cursor.selected = len(areas) - 1
	}
	if game.Input.ClaimKey(ebiten.KeyArrowDown) {
		cursor.selected = (cursor.selected + 1) % len(areas)
	}
	if game.Input.ClaimKey(ebiten.KeyArrowUp) {
		cursor.selected = (cursor.selected + len(areas) - 1) % len(areas)
	}
	point := game.Input.Point()
	for index, area := range areas {
		if area.Contains(point) {
			cursor.selected = index
		}
	}
}

func paneHeader(destination *ebiten.Image, game *inkwell.Game, header, hint Text) float64 {
	theme := game.Theme()
	top := uiTop(game) + Padding
	x := label(destination, game, header, PaneInset, int(top), theme.TopBarText)
	if len(hint) > 0 {
		room := int(foldButton(game).X) - x - inkwell.GlyphW*2 - 4
		game.DrawText(destination, inkwell.ClipText(hint.String(), room),
			x+inkwell.GlyphW*2, int(top), theme.ChatLogSystem)
	}
	return top + inkwell.GlyphH + 6
}

func plate(destination *ebiten.Image, game *inkwell.Game, body string, x, y int, ink color.Color) {
	theme := game.Theme()
	fill(destination, inkwell.Rect(
		float64(x-4),
		float64(y-3),
		float64(inkwell.TextWidth(body)+8),
		inkwell.GlyphH+6,
	), theme.PanelBG)
	game.DrawText(destination, body, x, y, ink)
}

func platedCenter(destination *ebiten.Image, game *inkwell.Game, body string, y int, ink color.Color) {
	plate(destination, game, body, (ScreenWidth-inkwell.TextWidth(body))/2, y, ink)
}
