package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"

	"realworld/inc"
)

var consoleLabel = Text{
	{Lang: "en", Text: "DEX · ARMILLA"},
	{Lang: "hu", Text: "DEX · ARMILLA"},
}

var consoleTokens = Text{
	{Lang: "en", Text: "%d TOKENS"},
	{Lang: "hu", Text: "%d ZSETON"},
}

var consolePrompt = Text{
	{Lang: "en", Text: "talk to Dex…"},
	{Lang: "hu", Text: "szólj Dexhez…"},
}

var consoleRail = Text{
	{Lang: "en", Text: "DEX"},
	{Lang: "hu", Text: "DEX"},
}

var consoleTab = Text{
	{Lang: "en", Text: "TAB"},
	{Lang: "hu", Text: "TAB"},
}

const consoleTokenPips = 5

type ConsoleLine struct {
	Body Text
	Cost Text
}

type Console struct {
	Name   string
	When   inkwell.Condition
	Tokens int
	Lines  []ConsoleLine
}

func (console *Console) GetName() string                { return console.Name }
func (console *Console) VisibleWhen() inkwell.Condition { return console.When }
func (console *Console) Layer() inkwell.Layer           { return inkwell.LayerHUD }
func (console *Console) Tick(context *inkwell.UICtx)    {}

func (console *Console) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game.(*inkwell.Game)
	if game.State.Flag(inc.FlagConsoleCollapsed) {
		console.drawRail(destination, game)
		return
	}
	console.drawPanel(destination, game)
	if !folded(game) {
		console.drawPrompt(destination, game)
	}
}

func (console *Console) drawPanel(destination *ebiten.Image, game *inkwell.Game) {
	theme := game.Theme()
	left := DividerX + 8
	right := ScreenWidth - 8
	top := int(uiTop(game)) + Padding

	fill(destination, inkwell.Rect(float64(left), float64(top+5), 5, 5), theme.TopBarAccent)
	label(destination, game, consoleLabel, left+9, top, theme.TopBarAccent)

	chip := inkwell.Rect(float64(right-inkwell.TextWidth(consoleTab.String())-4), float64(top+2),
		float64(inkwell.TextWidth(consoleTab.String())+4), inkwell.GlyphH-2)
	stroke(destination, chip, theme.DialogBorder)
	game.DrawText(destination, consoleTab.String(), int(chip.X)+1, top, theme.DialogBorder)

	tokens := inkwell.Format(consoleTokens, console.Tokens).String()
	game.DrawText(destination, tokens, int(chip.X)-inkwell.TextWidth(tokens)-6, top, theme.ChatLogResponse)

	if folded(game) {
		return
	}

	y := top + inkwell.GlyphH + 8
	limit := ScreenHeight - Padding - 20 - 6
	for _, line := range console.Lines {
		wrapped := inkwell.WrapText(line.Body.String(), right-left)
		tail := 0
		for _, body := range wrapped {
			if y+inkwell.GlyphH > limit {
				return
			}
			game.DrawText(destination, body, left, y, theme.ChatLogResponse)
			tail = left + inkwell.TextWidth(body)
			y += LineHeight
		}
		if len(line.Cost) > 0 {
			cost := line.Cost.String()
			costY := y
			if tail+inkwell.TextWidth(cost)+6 <= right {
				costY = y - LineHeight
			} else {
				y += LineHeight
			}
			if costY+inkwell.GlyphH <= limit {
				game.DrawText(destination, cost, right-inkwell.TextWidth(cost), costY, theme.DialogBorder)
			}
		}
		y += 4
	}
}

func (console *Console) drawPrompt(destination *ebiten.Image, game *inkwell.Game) {
	theme := game.Theme()
	bar := inkwell.Rect(float64(DividerX+8), ScreenHeight-Padding-20, ConsoleWidth-16, 20)
	fill(destination, bar, theme.PanelBG)
	stroke(destination, bar, theme.DialogBorder)
	baseline := int(bar.Y) + (20-inkwell.GlyphH)/2
	game.DrawText(destination, ">", int(bar.X)+4, baseline, theme.TopBarAccent)
	game.DrawText(destination, consolePrompt.String(), int(bar.X)+14, baseline, theme.DialogBorder)
	fill(destination, inkwell.Rect(bar.X+bar.W-8, float64(baseline+2), 4, 11), theme.TopBarAccent)
}

func (console *Console) drawRail(destination *ebiten.Image, game *inkwell.Game) {
	theme := game.Theme()
	rail := inkwell.Rect(RailX, uiTop(game), RailWidth, ScreenHeight-uiTop(game))
	top := int(uiTop(game)) + Padding

	if folded(game) {
		fill(destination, inkwell.Rect(rail.X+RailWidth/2-4, float64(top+4), 8, 8), theme.TopBarAccent)
		return
	}

	centered(destination, game, "←", rail, top, theme.DialogBorder)
	for index, glyph := range []rune(consoleRail.String()) {
		centered(destination, game, string(glyph), rail, top+inkwell.GlyphH+2+index*inkwell.GlyphH, theme.TopBarAccent)
	}

	dotY := top + inkwell.GlyphH*4 + 6
	fill(destination, inkwell.Rect(rail.X+RailWidth/2-4, float64(dotY), 8, 8), theme.TopBarAccent)

	pipY := ScreenHeight - Padding - inkwell.GlyphH - 4 - consoleTokenPips*8
	for index := 0; index < consoleTokenPips; index++ {
		pip := inkwell.Rect(rail.X+6, float64(pipY+index*8), RailWidth-12, 5)
		if index < console.Tokens {
			fill(destination, pip, theme.TopBarAccent)
			continue
		}
		stroke(destination, pip, theme.DialogBorder)
	}
	centered(destination, game, consoleTab.String(), rail, ScreenHeight-Padding-inkwell.GlyphH, theme.DialogBorder)
}
