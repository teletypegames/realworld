package inc

import (
	"math/rand"

	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	WidgetManager.Register(&useWithGuard{
		Name: "usewith_guard",
	})
}

type useWithGuard struct {
	Name string
}

func (u *useWithGuard) GetName() string                            { return u.Name }
func (u *useWithGuard) Layer() inkwell.Layer                       { return inkwell.LayerScene }
func (u *useWithGuard) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

func (u *useWithGuard) Tick(ctx *inkwell.UICtx) {
	g := ctx.Game
	if !World.HUDVisible() || !g.Input.LeftClicked() {
		return
	}
	sel := g.Inventory.Selected()
	if sel == "" {
		return
	}
	h := g.HotspotAt(g.Input.Point())
	if h == nil || hasAuthoredPair(g, h, sel) {
		return
	}

	g.Input.ConsumeLeft()
	g.Inventory.Select("")

	label := h.Label
	if label == "" {
		label = h.Name
	}
	g.LogAction("> Use " + itemLabel(g, sel) + " on: " + label)

	key := "fail." + sel + "." + h.Name
	g.State.NoteTalked(key)
	n := g.State.Talked(key)

	World.Do(inkwell.Seq(
		inkwell.Say(Paul, pick(paulFails, n)),
		TapeSay(Dex, dexFail(n)),
	))
}

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
