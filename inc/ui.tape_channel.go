package inc

// TapeChannel — the tape channel down the right-hand side.
//
// Not assistant UI: this is the voice of CHARACTERS. The wiki is explicit that
// ACPs and Personal Tapes are not props but characters with their own voice.
// That is why we do not use the built-in ChatLog — it paints every LogResponse
// in one Theme.ChatLogResponse colour, whereas here three or four tapes can
// speak into the same strip (Dex, the mystery tape, the tech support tape) and
// the colour comes per speaker from Character.SpeechColor.
//
// Paul and the NPCs never appear here: they speak in the scene, through the
// SpeechBubble. The channel is what the tapes said, plus — quietly — what you
// did to prompt it. That gives the "joke book, not a scoreboard" log for free.

import (
	"image/color"

	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type TapeChannel struct {
	Name   string
	W      *World
	Bounds inkwell.Rectangle
}

func (t *TapeChannel) GetName() string { return t.Name }

func (t *TapeChannel) Tick(ctx *inkwell.UICtx) {}

// BlocksClickAt — the channel is not interactive, but it is a solid input
// zone: the verb coin must not open over it.
func (t *TapeChannel) BlocksClickAt(p inkwell.Point) bool {
	return t.W.HUDVisible() && t.Bounds.Contains(p)
}

func (t *TapeChannel) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if !t.W.HUDVisible() {
		return
	}
	g := ctx.Game
	th := g.Theme()
	b := t.Bounds
	if b.W <= 0 || b.H <= 0 {
		return
	}
	const lh, pad = LineH, Pad

	vector.DrawFilledRect(dst, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), th.ChatLogBG, false)

	// Header: the name of whichever tape spoke last — not "console".
	speaker, speakerCol := t.lastSpeaker(g, th)
	drawTextC(dst, speaker, int(b.X)+pad, int(b.Y)+pad, speakerCol)
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
			// Your own actions — quiet, but visible: that is what makes the
			// joke book readable.
			col, txt = th.ChatLogPrompt, m.Text
		case inkwell.LogResponse:
			if !IsTape(m.Speaker) {
				continue // Paul and the NPCs speak in the scene, not here
			}
			col, txt = speechColor(g, m.Speaker, th.ChatLogResponse), m.Text
		default:
			col, txt = th.ChatLogSystem, m.Text
		}
		for _, wl := range wrap(txt, wrapW) {
			rendered = append(rendered, line{text: wl, col: col})
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
		drawTextC(dst, ln.text, int(b.X)+pad, top+i*lh, ln.col)
	}
}

// lastSpeaker returns the display name and colour of whichever tape spoke
// last, defaulting to the occupant of slot 1.
func (t *TapeChannel) lastSpeaker(g *inkwell.Game, th inkwell.Theme) (string, color.Color) {
	msgs := g.Messages()
	for i := len(msgs) - 1; i >= 0; i-- {
		m := msgs[i]
		if m.Kind == inkwell.LogResponse && IsTape(m.Speaker) {
			return TapeDisplayName(m.Speaker), speechColor(g, m.Speaker, th.ChatLogResponse)
		}
	}
	return TapeDisplayName(Dex), speechColor(g, Dex, th.ChatLogResponse)
}

// speechColor is the character's voice colour, falling back to the theme's.
func speechColor(g *inkwell.Game, name string, fallback color.Color) color.Color {
	if c, ok := g.CharacterManager.Get(name); ok && c.SpeechColor != nil {
		return c.SpeechColor
	}
	return fallback
}
