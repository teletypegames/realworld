package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Letterbox struct {
	Name string
	W    *World
	Bar  float64
}

func (l *Letterbox) GetName() string { return l.Name }

func (l *Letterbox) Tick(ctx *inkwell.UICtx) {
	if l.W.Mode() != ModeCutscene || !l.W.TapeWaiting() {
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		l.W.TakeTapeOffer()
	}
}

func (l *Letterbox) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if l.W.Mode() != ModeCutscene {
		return
	}
	g := ctx.Game
	th := g.Theme()
	bar := l.Bar
	if bar <= 0 {
		bar = 18
	}
	w, h := float32(g.Width), float32(g.Height)
	black := RGB(0x000000)
	vector.DrawFilledRect(dst, 0, 0, w, float32(bar), black, false)
	vector.DrawFilledRect(dst, 0, h-float32(bar), w, float32(bar), black, false)

	if l.W.TapeWaiting() {
		msg := "SPACE — " + TapeDisplayName(Dex) + " has something to say"
		drawTextC(dst, msg, g.Width-textW(msg)-4, g.Height-int(bar)+4, th.ChatLogResponse)
	}
}
