package ui

// Letterbox — the cutscene frame.
//
// The one exception to the fixed layout. In cutscene mode the HUD strip and
// the TopBar are hidden, the scene gets a black bar top and bottom, and a line
// at the foot of the screen says a tape would like to speak — but it does not
// speak on its own. The floor is the player's: a tape signals, never interrupts.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"git.teletypegames.org/games/realworld/internal/names"
	"git.teletypegames.org/games/realworld/internal/theme"
	"git.teletypegames.org/games/realworld/internal/world"
)

type Letterbox struct {
	Name string
	W    *world.World
	Bar  float64 // height of the black bar, top and bottom
}

func (l *Letterbox) GetName() string { return l.Name }

func (l *Letterbox) Tick(ctx *inkwell.UICtx) {
	if l.W.Mode() != world.ModeCutscene || !l.W.TapeWaiting() {
		return
	}
	// inkwell's Input only covers the mouse; read the key directly.
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		l.W.TakeTapeOffer()
	}
}

func (l *Letterbox) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if l.W.Mode() != world.ModeCutscene {
		return
	}
	g := ctx.Game
	th := g.Theme()
	bar := l.Bar
	if bar <= 0 {
		bar = 18
	}
	w, h := float32(g.Width), float32(g.Height)
	black := theme.RGB(0x000000)
	vector.DrawFilledRect(dst, 0, 0, w, float32(bar), black, false)
	vector.DrawFilledRect(dst, 0, h-float32(bar), w, float32(bar), black, false)

	if l.W.TapeWaiting() {
		msg := "SPACE — " + names.TapeDisplayName(names.Dex) + " has something to say"
		drawTextC(dst, msg, g.Width-textW(msg)-4, g.Height-int(bar)+4, th.ChatLogResponse)
	}
}
