package inc

// UseWithGuard — enforcing authored failure.
//
// The rule: a failed interaction fails IN CHARACTER, never with an error
// message. But inkwell hardcodes a "Nem ehhez." flash in core.engine.go when an
// item/hotspot pair has no OnUseWith handler, and that string cannot be
// replaced from the domain.
//
// The fix, without touching the engine: widgets tick BEFORE the engine's
// handleSceneInput and may consume the click. So this widget catches the
// "selected item + hotspot with no authored pair" case and fails in Paul's
// voice with a remark from Dex — costing nothing, since nothing is consumed and
// nothing breaks.
//
// IMPORTANT: it must be registered FIRST so that it ticks LAST among the
// widgets. That way every HUD widget (inventory, tape slots, dialogue) sees the
// click first, and this only ever catches clicks that land on the scene.

import (
	"math/rand"

	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

type UseWithGuard struct {
	Name string
	W    *World
}

func (u *UseWithGuard) GetName() string                            { return u.Name }
func (u *UseWithGuard) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

func (u *UseWithGuard) Tick(ctx *inkwell.UICtx) {
	g := ctx.Game
	if !u.W.HUDVisible() || !g.Input.LeftClicked() {
		return
	}
	sel := g.Inventory.Selected()
	if sel == "" {
		return
	}
	h := g.HotspotAt(g.Input.Point())
	if h == nil || hasAuthoredPair(g, h, sel) {
		return // leave authored pairs to the engine
	}

	g.Input.ConsumeLeft()
	g.Inventory.Select("")

	label := h.Label
	if label == "" {
		label = h.Name
	}
	g.LogAction("> Use " + itemLabel(g, sel) + " on: " + label)

	// The tone escalates on repeats: dry, then teasing, then a real nudge.
	key := "fail." + sel + "." + h.Name
	g.State.NoteTalked(key)
	n := g.State.Talked(key)

	u.W.Do(inkwell.Seq(
		inkwell.Say(Paul, pick(paulFails, n)),
		TapeSay(Dex, dexFail(n)),
	))
}

// hasAuthoredPair mirrors the engine's lookup order in invokeHotspotVerb.
func hasAuthoredPair(g *inkwell.Game, h *inkwell.Hotspot, item string) bool {
	if h.OnUseWith != nil {
		if a, ok := h.OnUseWith[item]; ok && a != nil {
			return true
		}
	}
	if it, ok := g.ItemManager.Get(item); ok && it.OnUseWith != nil {
		if a, ok := it.OnUseWith[h.Name]; ok && a != nil {
			return true
		}
	}
	return false
}

var paulFails = []string{
	"No. That's not going to work.",
	"Tried that. It didn't improve.",
	"All right, I know that one doesn't work.",
	"Now it's personal.",
}

var dexDry = []string{
	"No. Not like that.",
	"I heard it. Nothing happened.",
}

var dexTease = []string{
	"Twice the same. The second one rarely goes better.",
	"Do it a few more times, maybe physics reconsiders.",
}

// dexFail escalates: dry, then teasing, then an actual hint — so the joke
// never becomes the wall.
func dexFail(n int) string {
	switch {
	case n <= 1:
		return pick(dexDry, n)
	case n == 2:
		return pick(dexTease, n)
	default:
		return "Paul. Leave it. Look again at what you're carrying — that's where it is."
	}
}

func pick(s []string, n int) string {
	switch {
	case len(s) == 0:
		return ""
	case n <= 0:
		return s[0]
	case n-1 < len(s):
		return s[n-1]
	default:
		return s[rand.Intn(len(s))]
	}
}
