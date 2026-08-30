package widget

import (
	"image/color"

	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/tape"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func init() {
	Manager.Register(&tapeChannel{
		Name:   "channel",
		Bounds: inkwell.Rect(DividerX+2, HUDTop+2, ScreenW-DividerX-4, ScreenH-HUDTop-4),
	})
}

type tapeChannel struct {
	Name   string
	Bounds inkwell.Rectangle
}

func (t *tapeChannel) GetName() string                { return t.Name }
func (t *tapeChannel) VisibleWhen() inkwell.Condition { return whenPlaying }
func (t *tapeChannel) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (t *tapeChannel) Tick(ctx *inkwell.UICtx) {}

func (t *tapeChannel) BlocksClickAt(p inkwell.Point) bool { return t.Bounds.Contains(p) }

func (t *tapeChannel) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	g := ctx.Game
	th := g.Theme()
	b := t.Bounds
	if b.W <= 0 || b.H <= 0 {
		return
	}
	const lh, pad = LineH, Pad

	vector.DrawFilledRect(dst, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), th.ChatLogBG, false)

	speaker, speakerCol := t.lastSpeaker(g, th)
	g.DrawText(dst, speaker, int(b.X)+pad, int(b.Y)+pad, speakerCol)
	hy := float32(b.Y+pad+lh) + 2
	vector.StrokeLine(dst, float32(b.X)+pad, hy,
		float32(b.X+b.W)-pad, hy, 1, th.CharacterPanelBorder, false)

	type line struct {
		text string
		col  color.Color
	}
	wrapW := int(b.W) - pad*2
	var rendered []line
	for _, m := range g.Messages() {
		var col color.Color
		var txt string
		switch m.Kind {
		case inkwell.LogAction:

			col, txt = th.ChatLogPrompt, m.Text
		case inkwell.LogResponse:
			if !tape.Is(m.Speaker) {
				continue
			}
			col, txt = speechColor(g, m.Speaker, th.ChatLogResponse), m.Text
		default:
			col, txt = th.ChatLogSystem, m.Text
		}
		for _, wl := range inkwell.WrapText(txt, wrapW) {
			rendered = append(rendered, line{
				text: wl,
				col:  col,
			})
		}
	}

	top := int(b.Y) + pad + lh + 6
	maxLines := (int(b.Y+b.H) - pad - top) / lh
	if maxLines <= 0 {
		return
	}
	if start := len(rendered) - maxLines; start > 0 {
		rendered = rendered[start:]
	}
	for i, ln := range rendered {
		g.DrawText(dst, ln.text, int(b.X)+pad, top+i*lh, ln.col)
	}
}

func (t *tapeChannel) lastSpeaker(g *inkwell.Game, th inkwell.Theme) (string, color.Color) {
	msgs := g.Messages()
	for i := len(msgs) - 1; i >= 0; i-- {
		m := msgs[i]
		if m.Kind == inkwell.LogResponse && tape.Is(m.Speaker) {
			return g.CharacterLabel(m.Speaker), speechColor(g, m.Speaker, th.ChatLogResponse)
		}
	}
	return g.CharacterLabel(inc.Dex), speechColor(g, inc.Dex, th.ChatLogResponse)
}

func speechColor(g *inkwell.Game, name string, fallback color.Color) color.Color {
	if c, ok := g.CharacterManager.Get(name); ok && c.SpeechColor != nil {
		return c.SpeechColor
	}
	return fallback
}
