package widget

import (
	"image/color"

	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/inc"
	"realworld/inc/content/tape"
)

func init() {
	WidgetManager.Register(&tapeChannel{
		Name:   "channel",
		Bounds: inkwell.Rect(DividerX+2, HUDTop+2, ScreenWidth-DividerX-4, ScreenHeight-HUDTop-4),
	})
}

type tapeChannel struct {
	Name   string
	Bounds inkwell.Rectangle
}

func (channel *tapeChannel) GetName() string                { return channel.Name }
func (channel *tapeChannel) VisibleWhen() inkwell.Condition { return whenPlaying }
func (channel *tapeChannel) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (channel *tapeChannel) Tick(context *inkwell.UICtx) {}

func (channel *tapeChannel) BlocksClickAt(point inkwell.Point) bool {
	return channel.Bounds.Contains(point)
}

func (channel *tapeChannel) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game
	theme := game.Theme()
	bounds := channel.Bounds
	if bounds.W <= 0 || bounds.H <= 0 {
		return
	}

	vector.DrawFilledRect(destination, float32(bounds.X), float32(bounds.Y),
		float32(bounds.W), float32(bounds.H), theme.ChatLogBG, false)

	speaker, speakerColor := channel.lastSpeaker(game, theme)
	game.DrawText(destination, speaker, int(bounds.X)+Padding, int(bounds.Y)+Padding, speakerColor)
	ruleY := float32(bounds.Y+Padding+LineHeight) + 2
	vector.StrokeLine(destination, float32(bounds.X)+Padding, ruleY,
		float32(bounds.X+bounds.W)-Padding, ruleY, 1, theme.CharacterPanelBorder, false)

	type line struct {
		text  string
		color color.Color
	}
	wrapWidth := int(bounds.W) - Padding*2
	var rendered []line
	for _, message := range game.Messages() {
		var lineColor color.Color
		var text string
		switch message.Kind {
		case inkwell.LogAction:
			lineColor, text = theme.ChatLogPrompt, message.Text
		case inkwell.LogResponse:
			if !tape.Is(message.Speaker) {
				continue
			}
			lineColor, text = channel.speechColor(game, message.Speaker, theme.ChatLogResponse), message.Text
		default:
			lineColor, text = theme.ChatLogSystem, message.Text
		}
		for _, wrapped := range inkwell.WrapText(text, wrapWidth) {
			rendered = append(rendered, line{
				text:  wrapped,
				color: lineColor,
			})
		}
	}

	top := int(bounds.Y) + Padding + LineHeight + 6
	maxLines := (int(bounds.Y+bounds.H) - Padding - top) / LineHeight
	if maxLines <= 0 {
		return
	}
	if start := len(rendered) - maxLines; start > 0 {
		rendered = rendered[start:]
	}
	for index, rendering := range rendered {
		game.DrawText(destination, rendering.text, int(bounds.X)+Padding, top+index*LineHeight, rendering.color)
	}
}

func (channel *tapeChannel) lastSpeaker(game *inkwell.Game, theme inkwell.Theme) (string, color.Color) {
	messages := game.Messages()
	for index := len(messages) - 1; index >= 0; index-- {
		message := messages[index]
		if message.Kind == inkwell.LogResponse && tape.Is(message.Speaker) {
			return game.CharacterLabel(message.Speaker).String(),
				channel.speechColor(game, message.Speaker, theme.ChatLogResponse)
		}
	}
	return game.CharacterLabel(inc.Dex).String(), channel.speechColor(game, inc.Dex, theme.ChatLogResponse)
}

func (channel *tapeChannel) speechColor(game *inkwell.Game, name string, fallback color.Color) color.Color {
	if character, found := game.CharacterManager.Get(name); found && character.SpeechColor != nil {
		return character.SpeechColor
	}
	return fallback
}
