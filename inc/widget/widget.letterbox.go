package widget

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
	"git.teletypegames.org/games/realworld/inc/theme"
	"git.teletypegames.org/games/realworld/inc/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func init() {
	Manager.Register(&letterbox{
		Name: "letterbox",
		Bar:  36,
	})
}

type letterbox struct {
	Name string
	Bar  float64
}

func (l *letterbox) GetName() string                { return l.Name }
func (l *letterbox) VisibleWhen() inkwell.Condition { return whenCutscene }
func (l *letterbox) Layer() inkwell.Layer           { return inkwell.LayerScene }

func (l *letterbox) Tick(ctx *inkwell.UICtx) {
	if !world.TapeWaiting() {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		world.TakeTapeOffer()
	}
}

func (l *letterbox) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
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

	if world.TapeWaiting() {
		msg := "SPACE — " + g.CharacterLabel(inc.Dex) + " has something to say"
		g.DrawText(dst, msg, g.Width-inkwell.TextWidth(msg)-4, g.Height-int(bar)+4, th.ChatLogResponse)
	}
}
