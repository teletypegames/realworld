package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type screenNav struct {
	Name string
}

func (n *screenNav) GetName() string                            { return n.Name }
func (n *screenNav) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

func (n *screenNav) Tick(ctx *inkwell.UICtx) {
	if !World.HUDVisible() || World.Paused() {
		return
	}
	step := 0

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight):
		step = 1
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		step = -1
	default:
		return
	}
	if next := n.neighbour(ctx.Game, step); next != "" {
		World.Do(inkwell.GoTo(next))
	}
}

func (n *screenNav) neighbour(g *inkwell.Game, step int) string {
	deck := g.SceneManager.Names()
	if len(deck) < 2 {
		return ""
	}
	for i, name := range deck {
		if name == World.Scene() {
			return deck[((i+step)%len(deck)+len(deck))%len(deck)]
		}
	}
	return deck[0]
}
