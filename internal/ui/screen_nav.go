package ui

// screenNav — walking the deck with the arrow keys.
//
// The screens are connected — each one records where you can get to from it
// (content/screen/exit.go) — but walking the map is the player's job, and a
// screen whose beat is not written yet has nothing in it to walk towards. LEFT
// and RIGHT step to the previous and next screen so the deck can be reviewed as
// a deck: in concept-art order, wrapping at both ends. It is a tool for looking
// at the game, not a way through it.
//
// It is deliberately inert whenever the game is actually doing something: no
// stepping out of a cutscene, a menu, or a stopped world during dialogue. So
// once a beat owns a screen, this cannot cut across it.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"git.teletypegames.org/games/realworld/internal/world"
)

type screenNav struct {
	Name string
	W    *world.World
}

func (n *screenNav) GetName() string                            { return n.Name }
func (n *screenNav) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

func (n *screenNav) Tick(ctx *inkwell.UICtx) {
	if !n.W.HUDVisible() || n.W.Paused() {
		return
	}
	step := 0
	// inkwell's Input only covers the mouse; read the keys directly.
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight):
		step = 1
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		step = -1
	default:
		return
	}
	if next := n.neighbour(ctx.Game, step); next != "" {
		n.W.Do(inkwell.GoTo(next))
	}
}

// neighbour returns the screen step places from the current one, wrapping.
// Registration order is the deck order — see content/screen/screen.go.
func (n *screenNav) neighbour(g *inkwell.Game, step int) string {
	deck := g.SceneManager.Names()
	if len(deck) < 2 {
		return ""
	}
	for i, name := range deck {
		if name == n.W.Scene() {
			return deck[((i+step)%len(deck)+len(deck))%len(deck)]
		}
	}
	return deck[0]
}
